package events

import (
	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"go.uber.org/zap"
	"sync"
)

// ProcessWithdrawnEvents is the interface for processing withdrawal events.
type ProcessWithdrawnEvents interface {
	// ProcessWithdrawnNativeEventWithBackoff processes the withdrawn native event with a backoff mechanism.
	ProcessWithdrawnNativeEventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessWithdrawnERC20EventWithBackoff processes the withdrawn ERC20 event with a backoff mechanism.
	ProcessWithdrawnERC20EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessWithdrawnERC721EventWithBackoff processes the withdrawn ERC721 event with a backoff mechanism.
	ProcessWithdrawnERC721EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessWithdrawnERC1155EventWithBackoff processes the withdrawn ERC1155 event with a backoff mechanism.
	ProcessWithdrawnERC1155EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
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

// ProcessWithdrawnNativeEventWithBackoff processes the withdrawn native event with a backoff mechanism.
func (p *processWithdrawnEvents) ProcessWithdrawnNativeEventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processWithdrawnNativeEvent, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn native event after retries")
	}
}

// ProcessWithdrawnERC20EventWithBackoff processes the withdrawn ERC20 event with a backoff mechanism.
func (p *processWithdrawnEvents) ProcessWithdrawnERC20EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processWithdrawnERC20Event, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn ERC20 event after retries")
	}
}

// ProcessWithdrawnERC721EventWithBackoff processes the withdrawn ERC721 event with a backoff mechanism.
func (p *processWithdrawnEvents) ProcessWithdrawnERC721EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processWithdrawnERC721Event, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn ERC721 event after retries")
	}
}

// ProcessWithdrawnERC1155EventWithBackoff processes the withdrawn ERC1155 event with a backoff mechanism.
func (p *processWithdrawnEvents) ProcessWithdrawnERC1155EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processWithdrawnERC1155Event, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn ERC1155 event after retries")
	}
}

// processWithdrawnNativeEvent processes the withdrawn native event.
func (p *processWithdrawnEvents) processWithdrawnNativeEvent(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	withdrawnNativeLogs, err := p.bridgeFilterer.FilterWithdrawnNative(filterQuery)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for withdrawn native event")

		return err
	}
	defer withdrawnNativeLogs.Close()

	for withdrawnNativeLogs.Next() {
		p.nativeChan <- withdrawnNativeLogs.Event
	}

	if err := withdrawnNativeLogs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn native event")

		return err
	}

	close(p.nativeChan)
	wg.Done()

	return nil
}

// processWithdrawnERC20Event processes the withdrawn ERC20 event.
func (p *processWithdrawnEvents) processWithdrawnERC20Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	withdrawnERC20Logs, err := p.bridgeFilterer.FilterWithdrawnERC20(filterQuery)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for withdrawn ERC20 event")

		return err
	}
	defer withdrawnERC20Logs.Close()

	for withdrawnERC20Logs.Next() {
		p.erc20Chan <- withdrawnERC20Logs.Event
	}

	if err := withdrawnERC20Logs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn ERC20 event")

		return err
	}

	close(p.erc20Chan)
	wg.Done()

	return nil
}

// processWithdrawnERC721Event processes the withdrawn ERC721 event.
func (p *processWithdrawnEvents) processWithdrawnERC721Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	withdrawnERC721Logs, err := p.bridgeFilterer.FilterWithdrawnERC721(filterQuery)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for withdrawn ERC721 event")

		return err
	}
	defer withdrawnERC721Logs.Close()

	for withdrawnERC721Logs.Next() {
		p.erc721Chan <- withdrawnERC721Logs.Event
	}

	if err := withdrawnERC721Logs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn ERC721 event")

		return err
	}

	close(p.erc721Chan)
	wg.Done()

	return nil
}

// processWithdrawnERC1155Event processes the withdrawn ERC1155 event.
func (p *processWithdrawnEvents) processWithdrawnERC1155Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	withdrawnERC1155Logs, err := p.bridgeFilterer.FilterWithdrawnERC1155(filterQuery)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for withdrawn ERC1155 event")

		return err
	}
	defer withdrawnERC1155Logs.Close()

	for withdrawnERC1155Logs.Next() {
		p.erc1155Chan <- withdrawnERC1155Logs.Event
	}

	if err := withdrawnERC1155Logs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process withdrawn ERC1155 event")

		return err
	}

	close(p.erc1155Chan)
	wg.Done()

	return nil
}

// withBackoff creates an operation for backoff.Retry
func (p *processWithdrawnEvents) withBackoff(operation func(*bind.FilterOpts, *sync.WaitGroup) error, filterQuery *bind.FilterOpts, wg *sync.WaitGroup) func() error {
	return func() error {
		return operation(filterQuery, wg)
	}
}
