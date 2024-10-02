package tx_history

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	pgrepositories "github.com/quantum-bridge/core/cmd/data/postgresql/repositories"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/tx-history/handlers"
	goetherscan "github.com/quantum-bridge/core/pkg/go-etherscan"
)

// initializeEVM initializes the EVM event handlers and returns them.
func (s *txHistoryService) initializeEVM(client *ethclient.Client, chain datashared.Chain) (handlers.EventHandlers, error) {
	depositsHistoryRepository := pgrepositories.NewDepositsHistory(s.db)
	withdrawalsRepository := pgrepositories.NewWithdrawalsHistory(s.db)

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain ID")
	}

	eventHandlers := handlers.NewEventHandlers(s.logger, client, depositsHistoryRepository, withdrawalsRepository, chain, chainId.Int64())

	return eventHandlers, nil
}

// storeEVMDepositedEvents stores the EVM events in the database for the given chain in the real-time.
func (s *txHistoryService) storeEVMDepositedEvents(client *ethclient.Client, chain datashared.Chain) {
	eventHandlers, err := s.initializeEVM(client, chain)
	if err != nil {
		panic(err)
	}

	go eventHandlers.HandleEVMDepositedEvents(common.HexToAddress(chain.BridgeAddress))

}

// storeMissedEVMDepositedEvents stores the missed EVM events in the database for the given chain.
func (s *txHistoryService) storeMissedEVMDepositedEvents(client *ethclient.Client, chain datashared.Chain) {
	eventHandlers, err := s.initializeEVM(client, chain)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize EVM event handlers"))
	}

	// Get api_url and api_key from chain params (json.RawMessage).
	apiUrl, apiKey, err := datashared.GetEtherscanParams(chain.ChainParams)
	if err != nil {
		panic(errors.Wrap(err, "failed to get Etherscan API URL and API key"))
	}

	etherscan := goetherscan.New(apiUrl, apiKey)

	// Handle missed deposit events.
	err = eventHandlers.HandleEVMMissedDepositEvents(context.Background(), common.HexToAddress(chain.BridgeAddress), etherscan)
	if err != nil {
		s.logger.Errorf("failed to save missed events to the database: %s", err)
	}
}

// storeEVMWithdrawnEvents stores the EVM events in the database for the given chain in the real-time.
func (s *txHistoryService) storeEVMWithdrawnEvents(client *ethclient.Client, chain datashared.Chain) {
	eventHandlers, err := s.initializeEVM(client, chain)
	if err != nil {
		panic(err)
	}

	go eventHandlers.HandleEVMWithdrawnEvents(common.HexToAddress(chain.BridgeAddress))
}

// storeMissedEVMWithdrawnEvents stores the missed EVM events in the database for the given chain.
func (s *txHistoryService) storeMissedEVMWithdrawnEvents(client *ethclient.Client, chain datashared.Chain) {
	eventHandlers, err := s.initializeEVM(client, chain)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize EVM event handlers"))
	}

	// Get api_url and api_key from chain params (json.RawMessage).
	apiUrl, apiKey, err := datashared.GetEtherscanParams(chain.ChainParams)
	if err != nil {
		panic(errors.Wrap(err, "failed to get Etherscan API URL and API key"))
	}

	etherscan := goetherscan.New(apiUrl, apiKey)

	// Handle missed withdraw events.
	err = eventHandlers.HandleEVMMissedWithdrawalEvents(context.Background(), common.HexToAddress(chain.BridgeAddress), etherscan)
	if err != nil {
		s.logger.Errorf("failed to save missed events to the database: %s", err)
	}
}
