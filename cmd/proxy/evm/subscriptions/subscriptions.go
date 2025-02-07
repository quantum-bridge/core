package subscriptions

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
)

// Subscription is the interface for the subscription service.
type Subscription interface {
	// DepositedNative returns a channel for the deposited native events and the subscription.
	DepositedNative() (chan *bridge.BridgeDepositedNative, event.Subscription, error)
	// DepositedERC20 returns a channel for the deposited ERC20 events and the subscription.
	DepositedERC20() (chan *bridge.BridgeDepositedERC20, event.Subscription, error)
	// DepositedERC721 returns a channel for the deposited ERC721 events and the subscription.
	DepositedERC721() (chan *bridge.BridgeDepositedERC721, event.Subscription, error)
	// DepositedERC1155 returns a channel for the deposited ERC1155 events and the subscription.
	DepositedERC1155() (chan *bridge.BridgeDepositedERC1155, event.Subscription, error)
	// WithdrawnNative returns a channel for the withdrawn native events and the subscription.
	WithdrawnNative() (chan *bridge.BridgeWithdrawnNative, event.Subscription, error)
	// WithdrawnERC20 returns a channel for the withdrawn ERC20 events and the subscription.
	WithdrawnERC20() (chan *bridge.BridgeWithdrawnERC20, event.Subscription, error)
	// WithdrawnERC721 returns a channel for the withdrawn ERC721 events and the subscription.
	WithdrawnERC721() (chan *bridge.BridgeWithdrawnERC721, event.Subscription, error)
	// WithdrawnERC1155 returns a channel for the withdrawn ERC1155 events and the subscription.
	WithdrawnERC1155() (chan *bridge.BridgeWithdrawnERC1155, event.Subscription, error)
}

// subscription is the struct that holds the subscription service.
type subscription struct {
	client         *ethclient.Client
	bridgeFilterer bridge.BridgeFilterer
}

// NewSubscription creates a new subscription service with the given client and bridge filterer.
func NewSubscription(client *ethclient.Client, bridgeFilterer bridge.BridgeFilterer) Subscription {
	return &subscription{
		client:         client,
		bridgeFilterer: bridgeFilterer,
	}
}

// DepositedNative returns a channel for the deposited native events and the subscription.
func (s *subscription) DepositedNative() (chan *bridge.BridgeDepositedNative, event.Subscription, error) {
	events := make(chan *bridge.BridgeDepositedNative)
	sub, err := s.bridgeFilterer.WatchDepositedNative(nil, events)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// DepositedERC20 returns a channel for the deposited ERC20 events and the subscription.
func (s *subscription) DepositedERC20() (chan *bridge.BridgeDepositedERC20, event.Subscription, error) {
	events := make(chan *bridge.BridgeDepositedERC20)
	sub, err := s.bridgeFilterer.WatchDepositedERC20(nil, events, nil)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// DepositedERC721 returns a channel for the deposited ERC721 events and the subscription.
func (s *subscription) DepositedERC721() (chan *bridge.BridgeDepositedERC721, event.Subscription, error) {
	events := make(chan *bridge.BridgeDepositedERC721)
	sub, err := s.bridgeFilterer.WatchDepositedERC721(nil, events, nil)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// DepositedERC1155 returns a channel for the deposited ERC1155 events and the subscription.
func (s *subscription) DepositedERC1155() (chan *bridge.BridgeDepositedERC1155, event.Subscription, error) {
	events := make(chan *bridge.BridgeDepositedERC1155)
	sub, err := s.bridgeFilterer.WatchDepositedERC1155(nil, events)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// WithdrawnNative returns a channel for the withdrawn native events and the subscription.
func (s *subscription) WithdrawnNative() (chan *bridge.BridgeWithdrawnNative, event.Subscription, error) {
	events := make(chan *bridge.BridgeWithdrawnNative)
	sub, err := s.bridgeFilterer.WatchWithdrawnNative(nil, events)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// WithdrawnERC20 returns a channel for the withdrawn ERC20 events and the subscription.
func (s *subscription) WithdrawnERC20() (chan *bridge.BridgeWithdrawnERC20, event.Subscription, error) {
	events := make(chan *bridge.BridgeWithdrawnERC20)
	sub, err := s.bridgeFilterer.WatchWithdrawnERC20(nil, events)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// WithdrawnERC721 returns a channel for the withdrawn ERC721 events and the subscription.
func (s *subscription) WithdrawnERC721() (chan *bridge.BridgeWithdrawnERC721, event.Subscription, error) {
	events := make(chan *bridge.BridgeWithdrawnERC721)
	sub, err := s.bridgeFilterer.WatchWithdrawnERC721(nil, events)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}

// WithdrawnERC1155 returns a channel for the withdrawn ERC1155 events and the subscription.
func (s *subscription) WithdrawnERC1155() (chan *bridge.BridgeWithdrawnERC1155, event.Subscription, error) {
	events := make(chan *bridge.BridgeWithdrawnERC1155)
	sub, err := s.bridgeFilterer.WatchWithdrawnERC1155(nil, events)
	if err != nil {
		return nil, nil, err
	}

	return events, sub, nil
}
