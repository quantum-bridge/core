package tx_history

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	pgrepositories "github.com/quantum-bridge/core/cmd/data/postgresql/repositories"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/tx-history/handlers"
)

// initializeEVM initializes the EVM event handlers and returns them.
func (s *txHistoryService) initializeEVM(client *ethclient.Client) (handlers.EventHandlers, error) {
	transactionsHistoryRepository := pgrepositories.NewTransactionsHistory(s.db)

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain ID")
	}

	eventHandlers := handlers.NewEventHandlers(s.logger, client, transactionsHistoryRepository, chainId.Int64())

	return eventHandlers, nil
}

// storeEVMEvents stores the EVM events in the database for the given chain in the real-time.
func (s *txHistoryService) storeEVMEvents(client *ethclient.Client, chain datashared.Chain) {
	eventHandlers, err := s.initializeEVM(client)
	if err != nil {
		panic(err)
	}

	go eventHandlers.HandleEVMEvents(common.HexToAddress(chain.BridgeAddress))
}

// storeMissedEVMEvents stores the missed EVM events in the database for the given chain.
func (s *txHistoryService) storeMissedEVMEvents(client *ethclient.Client, chain datashared.Chain) {
	eventHandlers, err := s.initializeEVM(client)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize EVM event handlers"))
	}

	err = eventHandlers.HandleEVMMissedEvents(context.Background(), common.HexToAddress(chain.BridgeAddress))
	if err != nil {
		s.logger.Errorf("failed to save missed events to the database: %s", err)
	}
}
