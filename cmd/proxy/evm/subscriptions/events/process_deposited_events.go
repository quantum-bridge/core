package events

import (
	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"go.uber.org/zap"
	"sync"
)

// ProcessDepositedEvents is the interface for processing deposit events.
type ProcessDepositedEvents interface {
	// ProcessDepositedNativeEventWithBackoff processes the deposited native event with a backoff mechanism.
	ProcessDepositedNativeEventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessDepositedERC20EventWithBackoff processes the deposited ERC20 event with a backoff mechanism.
	ProcessDepositedERC20EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessDepositedERC721EventWithBackoff processes the deposited ERC721 event with a backoff mechanism.
	ProcessDepositedERC721EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
	// ProcessDepositedERC1155EventWithBackoff processes the deposited ERC1155 event with a backoff mechanism.
	ProcessDepositedERC1155EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup)
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

// ProcessDepositedNativeEventWithBackoff processes the deposited native event.
func (p *processDepositedEvents) ProcessDepositedNativeEventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processDepositedNativeEvent, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited native event after retries")
	}
}

// ProcessDepositedERC20EventWithBackoff processes the deposited ERC20 event.
func (p *processDepositedEvents) ProcessDepositedERC20EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processDepositedERC20Event, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited ERC20 event after retries")
	}
}

// ProcessDepositedERC721EventWithBackoff processes the deposited ERC721 event.
func (p *processDepositedEvents) ProcessDepositedERC721EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processDepositedERC721Event, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited ERC721 event after retries")
	}
}

// ProcessDepositedERC1155EventWithBackoff processes the deposited ERC1155 event.
func (p *processDepositedEvents) ProcessDepositedERC1155EventWithBackoff(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) {
	err := backoff.Retry(p.withBackoff(p.processDepositedERC1155Event, filterQuery, wg), backoff.NewExponentialBackOff())
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited ERC1155 event after retries")
	}
}

// processDepositedNativeEvent processes the deposited native event.
func (p *processDepositedEvents) processDepositedNativeEvent(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	depositedNativeLogs, err := p.bridgeFilterer.FilterDepositedNative(filterQuery)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for deposited native event")

		return err
	}
	defer depositedNativeLogs.Close()

	for depositedNativeLogs.Next() {
		p.nativeChan <- depositedNativeLogs.Event
	}

	if err := depositedNativeLogs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited native event")

		return err
	}

	close(p.nativeChan)
	wg.Done()

	return nil
}

// processDepositedERC20Event processes the deposited ERC20 event.
func (p *processDepositedEvents) processDepositedERC20Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	depositedERC20Logs, err := p.bridgeFilterer.FilterDepositedERC20(filterQuery, nil)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for deposited ERC20 event")

		return err
	}
	defer depositedERC20Logs.Close()

	for depositedERC20Logs.Next() {
		p.erc20Chan <- depositedERC20Logs.Event
	}

	if err := depositedERC20Logs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited ERC20 event")

		return err
	}

	close(p.erc20Chan)
	wg.Done()

	return nil
}

// processDepositedERC721Event processes the deposited ERC721 event.
func (p *processDepositedEvents) processDepositedERC721Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	depositedERC721Logs, err := p.bridgeFilterer.FilterDepositedERC721(filterQuery, nil)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for deposited ERC721 event")

		return err
	}
	defer depositedERC721Logs.Close()

	for depositedERC721Logs.Next() {
		p.erc721Chan <- depositedERC721Logs.Event
	}

	if err := depositedERC721Logs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited ERC721 event")

		return err
	}

	close(p.erc721Chan)
	wg.Done()

	return nil
}

// processDepositedERC1155Event processes the deposited ERC1155 event.
func (p *processDepositedEvents) processDepositedERC1155Event(filterQuery *bind.FilterOpts, wg *sync.WaitGroup) error {
	depositedERC1155Logs, err := p.bridgeFilterer.FilterDepositedERC1155(filterQuery)
	if err != nil {
		p.errorChan <- errors.Wrap(err, "failed to create filter for deposited ERC1155 event")

		return err
	}
	defer depositedERC1155Logs.Close()

	for depositedERC1155Logs.Next() {
		p.erc1155Chan <- depositedERC1155Logs.Event
	}

	if err := depositedERC1155Logs.Error(); err != nil {
		p.errorChan <- errors.Wrap(err, "failed to process deposited ERC1155 event")

		return err
	}

	close(p.erc1155Chan)
	wg.Done()

	return nil
}

// withBackoff creates an operation for backoff.Retry
func (p *processDepositedEvents) withBackoff(operation func(*bind.FilterOpts, *sync.WaitGroup) error, filterQuery *bind.FilterOpts, wg *sync.WaitGroup) func() error {
	return func() error {
		return operation(filterQuery, wg)
	}
}
