package handlers

import (
	"context"
	"database/sql"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data/postgresql/repositories"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"github.com/quantum-bridge/core/cmd/proxy/evm/subscriptions"
	"github.com/quantum-bridge/core/cmd/proxy/evm/subscriptions/events"
	pkgcommon "github.com/quantum-bridge/core/pkg/common"
	goetherscan "github.com/quantum-bridge/core/pkg/go-etherscan"
	"go.uber.org/zap"
	"sync"
)

// EventHandlers is the interface for handling EVM events.
type EventHandlers interface {
	// HandleEVMDepositedEvents listens to the EVM events and stores them in the database.
	HandleEVMDepositedEvents(contractAddress common.Address)
	// HandleEVMWithdrawnEvents listens to the EVM withdrawal events and updates deposit status in the database accordingly.
	HandleEVMWithdrawnEvents(contractAddress common.Address)
	// HandleEVMMissedDepositEvents handles the missed EVM deposit events from the last processed block and stores them in the database.
	HandleEVMMissedDepositEvents(ctx context.Context, contractAddress common.Address, etherscan goetherscan.Client) error
	// HandleEVMMissedWithdrawalEvents listens to the EVM withdrawal events and updates deposit status in the database accordingly.
	HandleEVMMissedWithdrawalEvents(ctx context.Context, contractAddress common.Address, etherscan goetherscan.Client) error
}

// eventHandlers is the struct that holds the event handlers service.
type eventHandlers struct {
	logger                       *zap.SugaredLogger
	evmClient                    *ethclient.Client
	depositsHistoryRepository    repositories.DepositsHistoryRepository
	withdrawalsHistoryRepository repositories.WithdrawalsHistoryRepository
	chain                        datashared.Chain
	chainId                      int64
}

// NewEventHandlers initializes and returns a new EventHandlers.
func NewEventHandlers(logger *zap.SugaredLogger, evmClient *ethclient.Client, depositsHistoryRepository repositories.DepositsHistoryRepository, withdrawalsRepository repositories.WithdrawalsHistoryRepository, chain datashared.Chain, chainId int64) EventHandlers {
	return &eventHandlers{
		logger:                       logger,
		evmClient:                    evmClient,
		depositsHistoryRepository:    depositsHistoryRepository,
		withdrawalsHistoryRepository: withdrawalsRepository,
		chain:                        chain,
		chainId:                      chainId,
	}
}

// HandleEVMDepositedEvents listens to the EVM events and stores them in the database.
func (h *eventHandlers) HandleEVMDepositedEvents(contractAddress common.Address) {
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
			if event != nil {
				err = h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), pkgcommon.NativeTokenAddress, event.Amount.String(), "", event.From.Hex(), event.To.String(), event.Network, false)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store Native Deposit event to the database"))

					return
				}
			}
		case event := <-depositedERC20EventChan:
			if event != nil {
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), "", event.From.Hex(), event.To, event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC20 Deposit event to the database"))

					return
				}
			}
		case event := <-depositedERC721EventChan:
			if event != nil {
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), "1", event.TokenId.String(), event.From.Hex(), event.To.Hex(), event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC721 Deposit event to the database"))

					return
				}
			}
		case event := <-depositedERC1155EventChan:
			if event != nil {
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), event.TokenId.String(), event.From.Hex(), event.To, event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC1155 Deposit event to the database"))

					return
				}
			}
		}
	}
}

// HandleEVMWithdrawnEvents listens to the EVM withdrawal events and updates deposit status in the database accordingly.
func (h *eventHandlers) HandleEVMWithdrawnEvents(contractAddress common.Address) {
	bridgeFilterer, err := bridge.NewBridgeFilterer(contractAddress, h.evmClient)
	if err != nil {
		h.logger.Error(err)

		return
	}

	subscription := subscriptions.NewSubscription(h.evmClient, *bridgeFilterer)

	withdrawnNativeEventChan, withdrawnNativeSubscription, err := subscription.WithdrawnNative()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer withdrawnNativeSubscription.Unsubscribe()

	withdrawnERC20EventChan, withdrawnERC20Subscription, err := subscription.WithdrawnERC20()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer withdrawnERC20Subscription.Unsubscribe()

	withdrawnERC721EventChan, withdrawnERC721Subscription, err := subscription.WithdrawnERC721()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer withdrawnERC721Subscription.Unsubscribe()

	withdrawnERC1155EventChan, withdrawnERC1155Subscription, err := subscription.WithdrawnERC1155()
	if err != nil {
		h.logger.Error(err)

		return
	}
	defer withdrawnERC1155Subscription.Unsubscribe()

	for {
		select {
		case event := <-withdrawnNativeEventChan:
			err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, pkgcommon.NativeTokenAddress, event.Amount.String(), "", event.To.Hex(), false)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store Native Withdrawn event to the database"))

				return
			}
		case event := <-withdrawnERC20EventChan:
			err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, event.Token.Hex(), event.Amount.String(), "", event.To.Hex(), event.IsMintable)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store ERC20 Withdrawn event to the database"))

				return
			}
		case event := <-withdrawnERC721EventChan:
			err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, event.Token.Hex(), "1", event.TokenId.String(), event.To.Hex(), event.IsMintable)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store ERC721 Withdrawn event to the database"))

				return
			}
		case event := <-withdrawnERC1155EventChan:
			err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, event.Token.Hex(), event.Amount.String(), event.TokenId.String(), event.To.Hex(), event.IsMintable)
			if err != nil {
				h.logger.Error(errors.Wrap(err, "failed to store ERC1155 Withdrawn event to the database"))

				return
			}
		}
	}
}

// HandleEVMMissedDepositEvents handles the missed EVM deposit events from the last processed block and stores them in the database.
func (h *eventHandlers) HandleEVMMissedDepositEvents(ctx context.Context, contractAddress common.Address, etherscan goetherscan.Client) error {
	lastDepositedTxHistory, err := h.depositsHistoryRepository.
		Where("source_network", h.chain.ID).
		OrderBy("block_number", "DESC").
		Limit(1).
		Get()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// If there are no deposits in the database, get the contract creation information from Etherscan.
		contractCreationInfo, err := etherscan.GetContractCreation(contractAddress.Hex())
		if err != nil {
			return err
		}

		// Get the transaction hash from the contract creation information.
		txHash := contractCreationInfo.Result[0].TxHash

		// Get the transaction information from the transaction hash.
		transactionInfo, err := etherscan.GetTransactionByHash(txHash)
		if err != nil {
			return err
		}

		// Get the block number from the transaction information and convert it to decimal format.
		blockNumber, err := goetherscan.HexToDecimal(transactionInfo.Result.BlockNumber)
		if err != nil {
			return err
		}

		lastDepositedTxHistory.BlockNumber = uint64(blockNumber)
	}

	filterQuery := &bind.FilterOpts{
		Start:   lastDepositedTxHistory.BlockNumber + 1, // Increment by 1 to start from the next block from the last processed block.
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

	processEvents := events.NewProcessDepositedEvents(h.logger, *bridgeFilterer, nativeChan, erc20Chan, erc721Chan, erc1155Chan, errorChan)

	var wg sync.WaitGroup
	wg.Add(4) // 4 events to process (native, erc20, erc721, erc1155)

	go processEvents.ProcessDepositedNativeEvent(filterQuery, &wg)
	go processEvents.ProcessDepositedERC20Event(filterQuery, &wg)
	go processEvents.ProcessDepositedERC721Event(filterQuery, &wg)
	go processEvents.ProcessDepositedERC1155Event(filterQuery, &wg)

	go h.processMissedDepositedEvents(nativeChan, erc20Chan, erc721Chan, erc1155Chan, errorChan)

	wg.Wait()

	return nil
}

// HandleEVMMissedWithdrawalEvents listens to the EVM withdrawal events and updates deposit status in the database accordingly.
func (h *eventHandlers) HandleEVMMissedWithdrawalEvents(ctx context.Context, contractAddress common.Address, etherscan goetherscan.Client) error {
	lastWithdrawnTxHistory, err := h.withdrawalsHistoryRepository.
		Where("destination_network", h.chain.ID).
		OrderBy("block_number", "DESC").
		Limit(1).
		Get()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// If there are no deposits in the database, get the contract creation information from Etherscan.
		contractCreationInfo, err := etherscan.GetContractCreation(contractAddress.Hex())
		if err != nil {
			return err
		}

		// Get the transaction hash from the contract creation information.
		txHash := contractCreationInfo.Result[0].TxHash

		// Get the transaction information from the transaction hash.
		transactionInfo, err := etherscan.GetTransactionByHash(txHash)
		if err != nil {
			return err
		}

		// Get the block number from the transaction information and convert it to decimal format.
		blockNumber, err := goetherscan.HexToDecimal(transactionInfo.Result.BlockNumber)
		if err != nil {
			return err
		}

		lastWithdrawnTxHistory.BlockNumber = uint64(blockNumber)
	}

	// If there are no deposits in the database
	if lastWithdrawnTxHistory == nil {
		return nil
	}

	filterQuery := &bind.FilterOpts{
		Start:   lastWithdrawnTxHistory.BlockNumber + 1,
		End:     nil,
		Context: ctx,
	}

	nativeChan := make(chan *bridge.BridgeWithdrawnNative)
	erc20Chan := make(chan *bridge.BridgeWithdrawnERC20)
	erc721Chan := make(chan *bridge.BridgeWithdrawnERC721)
	erc1155Chan := make(chan *bridge.BridgeWithdrawnERC1155)
	errorChan := make(chan error)

	bridgeFilterer, err := bridge.NewBridgeFilterer(contractAddress, h.evmClient)
	if err != nil {
		return err
	}

	processEvents := events.NewProcessWithdrawnEvents(h.logger, *bridgeFilterer, nativeChan, erc20Chan, erc721Chan, erc1155Chan, errorChan)

	var wg sync.WaitGroup
	wg.Add(4) // 4 events to process (native, erc20, erc721, erc1155)

	go processEvents.ProcessWithdrawnNativeEvent(filterQuery, &wg)
	go processEvents.ProcessWithdrawnERC20Event(filterQuery, &wg)
	go processEvents.ProcessWithdrawnERC721Event(filterQuery, &wg)
	go processEvents.ProcessWithdrawnERC1155Event(filterQuery, &wg)

	go h.processMissedWithdrawnEvents(nativeChan, erc20Chan, erc721Chan, erc1155Chan, errorChan)

	wg.Wait()

	return nil
}

// processMissedDepositedEvents processes the missed events and stores them in the database.
func (h *eventHandlers) processMissedDepositedEvents(
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
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), pkgcommon.NativeTokenAddress, event.Amount.String(), "", event.From.Hex(), event.To.String(), event.Network, false)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store Native Deposit event to the database"))
				}
			}
		case event := <-erc20Chan:
			if event != nil {
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), "", event.From.Hex(), event.To, event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC20 Deposit event to the database"))
				}
			}
		case event := <-erc721Chan:
			if event != nil {
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), "1", event.TokenId.String(), event.From.Hex(), event.To.Hex(), event.Network, event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC721 Deposit event to the database"))
				}
			}
		case event := <-erc1155Chan:
			if event != nil {
				err := h.storeDepositedEvent(event.Raw.BlockNumber, event.Raw.TxHash.Hex(), event.Token.Hex(), event.Amount.String(), event.TokenId.String(), event.From.Hex(), event.To, event.Network, event.IsMintable)
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

// processMissedWithdrawnEvents processes the missed withdrawn events and updates the withdrawn status in the database.
func (h *eventHandlers) processMissedWithdrawnEvents(
	nativeChan chan *bridge.BridgeWithdrawnNative,
	erc20Chan chan *bridge.BridgeWithdrawnERC20,
	erc721Chan chan *bridge.BridgeWithdrawnERC721,
	erc1155Chan chan *bridge.BridgeWithdrawnERC1155,
	errorChan chan error,
) {
	for {
		select {
		case event := <-nativeChan:
			if event != nil {
				err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, pkgcommon.NativeTokenAddress, event.Amount.String(), "", event.To.Hex(), false)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store Native Withdrawn event to the database"))
				}
			}
		case event := <-erc20Chan:
			if event != nil {
				err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, event.Token.Hex(), event.Amount.String(), "", event.To.Hex(), event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC20 Withdrawn event to the database"))
				}
			}
		case event := <-erc721Chan:
			if event != nil {
				err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, event.Token.Hex(), "1", event.TokenId.String(), event.To.Hex(), event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC721 Withdrawn event to the database"))
				}
			}
		case event := <-erc1155Chan:
			if event != nil {
				err := h.storeWithdrawnEvent(event.Raw.BlockNumber, event.TxHash, event.Token.Hex(), event.Amount.String(), event.TokenId.String(), event.To.Hex(), event.IsMintable)
				if err != nil {
					h.logger.Error(errors.Wrap(err, "failed to store ERC1155 Withdrawn event to the database"))
				}
			}
		case err := <-errorChan:
			if err != nil {
				h.logger.Error(err)
			}
		}
	}
}

// storeDepositedEvent stores the deposit event in the database.
func (h *eventHandlers) storeDepositedEvent(blockNumber uint64, txHash, tokenAddress, amount, tokenId, fromAddress, toAddress, destinationNetwork string, isMintable bool) error {
	return h.depositsHistoryRepository.Insert(&repositories.DepositsHistory{
		SourceNetwork:      h.chain.ID,
		TxHash:             txHash,
		BlockNumber:        blockNumber,
		TokenAddress:       tokenAddress,
		TokenID:            tokenId,
		Amount:             amount,
		FromAddress:        fromAddress,
		ToAddress:          toAddress,
		DestinationNetwork: destinationNetwork,
		IsMintable:         isMintable,
	})
}

// storeWithdrawnEvent stores the withdrawn event in the database.
func (h *eventHandlers) storeWithdrawnEvent(blockNumber uint64, txHash [32]byte, tokenAddress, amount, tokenId, toAddress string, isMintable bool) error {
	// Convert the transaction hash to hexadecimal format.
	transactionHash := common.HexToHash(common.Bytes2Hex(txHash[:])).Hex()

	// Get the deposited event with the transaction hash.
	depositedEvent, err := h.depositsHistoryRepository.
		New(). // Create a new query builder instance to get the deposited event with the transaction hash.
		Where("destination_network", h.chain.ID).
		Where("tx_hash", transactionHash).
		Get()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// TODO: fix sql error no rows in case when the deposited event is not found, but missed deposited event is already processed and stored in the database.
		h.logger.Warnf("failed to get deposited event with tx hash %s", transactionHash)

		return nil
	}

	return h.withdrawalsHistoryRepository.Insert(&repositories.WithdrawalsHistory{
		SourceNetwork:      depositedEvent.SourceNetwork,
		TxHash:             transactionHash,
		BlockNumber:        blockNumber,
		TokenAddress:       tokenAddress,
		TokenID:            tokenId,
		Amount:             amount,
		FromAddress:        depositedEvent.FromAddress,
		ToAddress:          toAddress,
		DestinationNetwork: h.chain.ID,
		IsMintable:         isMintable,
	})
}
