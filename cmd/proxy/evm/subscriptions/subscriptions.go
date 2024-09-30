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
