package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data/postgresql/repositories"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"github.com/quantum-bridge/core/cmd/proxy/evm/subscriptions"
	"github.com/quantum-bridge/core/cmd/proxy/evm/subscriptions/events"
	pkgcommon "github.com/quantum-bridge/core/pkg/common"
	"go.uber.org/zap"
	"sync"
)

// EventHandlers is the interface for handling EVM events.
type EventHandlers interface {
	// HandleEVMEvents listens to the EVM events and stores them in the database.
	HandleEVMEvents(contractAddress common.Address)
	// HandleEVMMissedEvents handles the missed EVM events from the last processed block and stores them in the database.
	HandleEVMMissedEvents(ctx context.Context, contractAddress common.Address) error
}

// eventHandlers is the struct that holds the event handlers service.
type eventHandlers struct {
	logger              *zap.SugaredLogger
	evmClient           *ethclient.Client
	txHistoryRepository repositories.TransactionsHistoryRepository
	chainId             int64
}

// NewEventHandlers initializes and returns a new EventHandlers.
func NewEventHandlers(logger *zap.SugaredLogger, evmClient *ethclient.Client, transactionsHistoryRepository repositories.TransactionsHistoryRepository, chainId int64) EventHandlers {
	return &eventHandlers{
		logger:              logger,
		evmClient:           evmClient,
		txHistoryRepository: transactionsHistoryRepository,
		chainId:             chainId,
	}
}

// HandleEVMEvents listens to the EVM events and stores them in the database.
func (h *eventHandlers) HandleEVMEvents(contractAddress common.Address) {
	bridgeFilterer, err := bridge.NewBridgeFilterer(contractAddress, h.evmClient)
	if err != nil {
		h.logger.Error(err)

		return
	}

	subscription := subscriptions.NewSubscription(h.evmClient, *bridgeFilterer)

	depositedNativeEventChan, depositedNativeSubscription, err := subscription.DepositedNative()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer depositedNativeSubscription.Unsubscribe()

	depositedERC20EventChan, depositedERC20Subscription, err := subscription.DepositedERC20()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer depositedERC20Subscription.Unsubscribe()

	depositedERC721EventChan, depositedERC721Subscription, err := subscription.DepositedERC721()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer depositedERC721Subscription.Unsubscribe()

	depositedERC1155EventChan, depositedERC1155Subscription, err := subscription.DepositedERC1155()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer depositedERC1155Subscription.Unsubscribe()

	for {
		select {
		case event := <-depositedNativeEventChan:
			err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), pkgcommon.NativeTokenAddress, event.Amount.String(), "", event.Raw.Address.Hex(), event.To.String(), event.Network, false)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store Native Deposit event to the database"))

				return
			}
		case event := <-depositedERC20EventChan:
			err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), "", event.Raw.Address.Hex(), event.To, event.Network, event.IsMintable)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store ERC20 Deposit event to the database"))

				return
			}
		case event := <-depositedERC721EventChan:
			err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), "1", event.TokenId.String(), event.Raw.Address.Hex(), event.To.Hex(), event.Network, event.IsMintable)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store ERC721 Deposit event to the database"))

				return
			}
		case event := <-depositedERC1155EventChan:
			err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), event.TokenId.String(), event.Raw.Address.Hex(), event.To, event.Network, event.IsMintable)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store ERC1155 Deposit event to the database"))

				return
			}
		}
	}
}

// HandleEVMMissedEvents handles the missed EVM events from the last processed block and stores them in the database.
func (h *eventHandlers) HandleEVMMissedEvents(ctx context.Context, contractAddress common.Address) error {
	lastTxHistory, err := h.txHistoryRepository.
		Where("chain_id", fmt.Sprintf("%d", h.chainId)).
		OrderBy("block_number", "DESC").
		Limit(1).
		Get()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// TODO: Set lastTxHistory.BlockNumber to the deployed contract block number in the chain.
		lastTxHistory.BlockNumber = 0
	}

	filterQuery := &bind.FilterOpts{
		Start:   lastTxHistory.BlockNumber + 1, // Increment by 1 to start from the next block from the last processed block.
		End:     nil,
		Context: ctx,
	}

	nativeChan := make(chan *bridge.BridgeDepositedNative)
	erc20Chan := make(chan *bridge.BridgeDepositedERC20)
	erc721Chan := make(chan *bridge.BridgeDepositedERC721)
	erc1155Chan := make(chan *bridge.BridgeDepositedERC1155)
	errorChan := make(chan error)

	bridgeFilterer, err := bridge.NewBridgeFilterer(contractAddress, h.evmClient)
	if err != nil {
		return err
	}

	processEvents := events.NewProcessEvents(h.logger, *bridgeFilterer, nativeChan, erc20Chan, erc721Chan, erc1155Chan, errorChan)

	var wg sync.WaitGroup
	wg.Add(4) // 4 events to process (native, erc20, erc721, erc1155)

	go processEvents.ProcessDepositedNativeEvent(filterQuery, &wg)
	go processEvents.ProcessDepositedERC20Event(filterQuery, &wg)
	go processEvents.ProcessDepositedERC721Event(filterQuery, &wg)
	go processEvents.ProcessDepositedERC1155Event(filterQuery, &wg)

	go h.processMissedEvents(nativeChan, erc20Chan, erc721Chan, erc1155Chan, errorChan)

	wg.Wait()

	return nil
}

// processMissedEvents processes the missed events and stores them in the database.
func (h *eventHandlers) processMissedEvents(
	nativeChan chan *bridge.BridgeDepositedNative,
	erc20Chan chan *bridge.BridgeDepositedERC20,
	erc721Chan chan *bridge.BridgeDepositedERC721,
	erc1155Chan chan *bridge.BridgeDepositedERC1155,
	errorChan chan error,
) {
	for {
		select {
		case event := <-nativeChan:
			if event != nil {
				err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), pkgcommon.NativeTokenAddress, event.Amount.String(), "", event.Raw.Address.Hex(), event.To.String(), event.Network, false)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store Native Deposit event to the database"))
				}
			}
		case event := <-erc20Chan:
			if event != nil {
				err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), "", event.Raw.Address.Hex(), event.To, event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC20 Deposit event to the database"))
				}
			}
		case event := <-erc721Chan:
			if event != nil {
				err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), "1", event.TokenId.String(), event.Raw.Address.Hex(), event.To.Hex(), event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC721 Deposit event to the database"))
				}
			}
		case event := <-erc1155Chan:
			if event != nil {
				err := h.storeTxHistory(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), event.TokenId.String(), event.Raw.Address.Hex(), event.To, event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC1155 Deposit event to the database"))
				}
			}
		case err := <-errorChan:
			if err != nil {
				h.logger.Error(err)
			}
		}
	}
}

// storeTxHistory stores the transaction history in the database.
func (h *eventHandlers) storeTxHistory(blockNumber uint64, txHash, tokenAddress, amount, tokenId, fromAddress, toAddress, toNetwork string, isMintable bool) error {
	return h.txHistoryRepository.Insert(&repositories.TransactionsHistory{
		TxHash:       txHash,
		ChainID:      h.chainId,
		BlockNumber:  blockNumber,
		TokenAddress: tokenAddress,
		TokenID:      tokenId,
		Amount:       amount,
		FromAddress:  fromAddress,
		ToAddress:    toAddress,
		ToNetwork:    toNetwork,
		IsMintable:   isMintable,
	})
}
