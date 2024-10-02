package events

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"go.uber.org/zap"
	"sync"
)

// ProcessDepositedEvents is the interface for processing deposit events.
type ProcessDepositedEvents interface {
	// ProcessDepositedNativeEvent processes the deposited native event.
	ProcessDepositedNativeEvent(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessDepositedERC20Event processes the deposited ERC20 event.
	ProcessDepositedERC20Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessDepositedERC721Event processes the deposited ERC721 event.
	ProcessDepositedERC721Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessDepositedERC1155Event processes the deposited ERC1155 event.
	ProcessDepositedERC1155Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
}

// processDepositedEvents is the struct that holds the process deposited events service.
type processDepositedEvents struct {
	logger         *zap.SugaredLogger
	bridgeFilterer bridge.BridgeFilterer
	nativeChan     chan *bridge.BridgeDepositedNative
	erc20Chan      chan *bridge.BridgeDepositedERC20
	erc721Chan     chan *bridge.BridgeDepositedERC721
	erc1155Chan    chan *bridge.BridgeDepositedERC1155
	errorChan      chan error
}

// NewProcessDepositedEvents creates a new process events service with the given bridge filterer and channels.
func NewProcessDepositedEvents(
	logger *zap.SugaredLogger,
	bridgeFilterer bridge.BridgeFilterer,
	nativeChan chan *bridge.BridgeDepositedNative,
	erc20Chan chan *bridge.BridgeDepositedERC20,
	erc721Chan chan *bridge.BridgeDepositedERC721,
	erc1155Chan chan *bridge.BridgeDepositedERC1155,
	errorChan chan error,
) ProcessDepositedEvents {
	return &processDepositedEvents{
		logger:         logger,
		bridgeFilterer: bridgeFilterer,
		nativeChan:     nativeChan,
		erc20Chan:      erc20Chan,
		erc721Chan:     erc721Chan,
		erc1155Chan:    erc1155Chan,
		errorChan:      errorChan,
	}
}

// ProcessDepositedNativeEvent processes the deposited native event.
func (p *processDepositedEvents) ProcessDepositedNativeEvent(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	depositedNativeLogs, err := p.bridgeFilterer.FilterDepositedNative(filterQuery)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer depositedNativeLogs.Close()

	for depositedNativeLogs.Next() {
		p.nativeChan <- depositedNativeLogs.Event
	}

	if err := depositedNativeLogs.Error(); err != nil {
		p.errorChan <- err
	}

	close(p.nativeChan)
	wg.Done()
}

// ProcessDepositedERC20Event processes the deposited ERC20 event.
func (p *processDepositedEvents) ProcessDepositedERC20Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	depositedERC20Logs, err := p.bridgeFilterer.FilterDepositedERC20(filterQuery, nil)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer depositedERC20Logs.Close()

	for depositedERC20Logs.Next() {
		p.erc20Chan <- depositedERC20Logs.Event
	}

	if err := depositedERC20Logs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.erc20Chan)
	wg.Done()
}

// ProcessDepositedERC721Event processes the deposited ERC721 event.
func (p *processDepositedEvents) ProcessDepositedERC721Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	depositedERC721Logs, err := p.bridgeFilterer.FilterDepositedERC721(filterQuery, nil)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer depositedERC721Logs.Close()

	for depositedERC721Logs.Next() {
		p.erc721Chan <- depositedERC721Logs.Event
	}

	if err := depositedERC721Logs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.erc721Chan)
	wg.Done()
}

// ProcessDepositedERC1155Event processes the deposited ERC1155 event.
func (p *processDepositedEvents) ProcessDepositedERC1155Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	depositedERC1155Logs, err := p.bridgeFilterer.FilterDepositedERC1155(filterQuery)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer depositedERC1155Logs.Close()

	for depositedERC1155Logs.Next() {
		p.erc1155Chan <- depositedERC1155Logs.Event
	}

	if err := depositedERC1155Logs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.erc1155Chan)
	wg.Done()
}
