package events

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"go.uber.org/zap"
	"sync"
)

// ProcessWithdrawnEvents is the interface for processing withdrawal events.
type ProcessWithdrawnEvents interface {
	// ProcessWithdrawnNativeEvent processes the withdrawn native event.
	ProcessWithdrawnNativeEvent(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessWithdrawnERC20Event processes the withdrawn ERC20 event.
	ProcessWithdrawnERC20Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessWithdrawnERC721Event processes the withdrawn ERC721 event.
	ProcessWithdrawnERC721Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessWithdrawnERC1155Event processes the withdrawn ERC1155 event.
	ProcessWithdrawnERC1155Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
}

// processWithdrawEvents is the struct that holds the process withdrawn events service.
type processWithdrawnEvents struct {
	logger         *zap.SugaredLogger
	bridgeFilterer bridge.BridgeFilterer
	nativeChan     chan *bridge.BridgeWithdrawnNative
	erc20Chan      chan *bridge.BridgeWithdrawnERC20
	erc721Chan     chan *bridge.BridgeWithdrawnERC721
	erc1155Chan    chan *bridge.BridgeWithdrawnERC1155
	errorChan      chan error
}

// NewProcessWithdrawnEvents creates a new process events service with the given bridge filterer and channels.
func NewProcessWithdrawnEvents(
	logger *zap.SugaredLogger,
	bridgeFilterer bridge.BridgeFilterer,
	nativeChan chan *bridge.BridgeWithdrawnNative,
	erc20Chan chan *bridge.BridgeWithdrawnERC20,
	erc721Chan chan *bridge.BridgeWithdrawnERC721,
	erc1155Chan chan *bridge.BridgeWithdrawnERC1155,
	errorChan chan error,
) ProcessWithdrawnEvents {
	return &processWithdrawnEvents{
		logger:         logger,
		bridgeFilterer: bridgeFilterer,
		nativeChan:     nativeChan,
		erc20Chan:      erc20Chan,
		erc721Chan:     erc721Chan,
		erc1155Chan:    erc1155Chan,
		errorChan:      errorChan,
	}
}

// ProcessWithdrawnNativeEvent processes the withdrawn native event.
func (p *processWithdrawnEvents) ProcessWithdrawnNativeEvent(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	withdrawnNativeLogs, err := p.bridgeFilterer.FilterWithdrawnNative(filterQuery)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer withdrawnNativeLogs.Close()

	for withdrawnNativeLogs.Next() {
		p.nativeChan <- withdrawnNativeLogs.Event
	}

	if err := withdrawnNativeLogs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.nativeChan)
	wg.Done()
}

// ProcessWithdrawnERC20Event processes the withdrawn ERC20 event.
func (p *processWithdrawnEvents) ProcessWithdrawnERC20Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	withdrawnERC20Logs, err := p.bridgeFilterer.FilterWithdrawnERC20(filterQuery)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer withdrawnERC20Logs.Close()

	for withdrawnERC20Logs.Next() {
		p.erc20Chan <- withdrawnERC20Logs.Event
	}

	if err := withdrawnERC20Logs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.erc20Chan)
	wg.Done()
}

// ProcessWithdrawnERC721Event processes the withdrawn ERC721 event.
func (p *processWithdrawnEvents) ProcessWithdrawnERC721Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	withdrawnERC721Logs, err := p.bridgeFilterer.FilterWithdrawnERC721(filterQuery)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer withdrawnERC721Logs.Close()

	for withdrawnERC721Logs.Next() {
		p.erc721Chan <- withdrawnERC721Logs.Event
	}

	if err := withdrawnERC721Logs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.erc721Chan)
	wg.Done()
}

// ProcessWithdrawnERC1155Event processes the withdrawn ERC1155 event.
func (p *processWithdrawnEvents) ProcessWithdrawnERC1155Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	withdrawnERC1155Logs, err := p.bridgeFilterer.FilterWithdrawnERC1155(filterQuery)
	if err != nil {
		p.errorChan <- err

		return
	}
	defer withdrawnERC1155Logs.Close()

	for withdrawnERC1155Logs.Next() {
		p.erc1155Chan <- withdrawnERC1155Logs.Event
	}

	if err := withdrawnERC1155Logs.Error(); err != nil {
		p.errorChan <- err

		return
	}

	close(p.erc1155Chan)
	wg.Done()
}
