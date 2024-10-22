package tx_history

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/pkg/squirrelizer"
	"go.uber.org/zap"
)

// TxHistoryService defines the interface for the transaction history service.
type TxHistoryService interface {
	Run()
}

type txHistoryService struct {
	logger *zap.SugaredLogger
	db     *squirrelizer.DB
	chains []datashared.Chain
}

// NewTxHistory initializes and returns a new TxHistoryService.
func NewTxHistory(logger *zap.SugaredLogger, db *squirrelizer.DB, chains []datashared.Chain) TxHistoryService {
	return &txHistoryService{
		logger: logger,
		db:     db,
		chains: chains,
	}
}

// Run runs the transaction history service with the given configuration.
func (s *txHistoryService) Run() {
	for _, chain := range s.chains {
		switch chain.Type {
		case data.EVM:
			client, err := ethclient.Dial(chain.RpcEndpoint)
			if err != nil {
				panic(errors.Wrap(err, "failed to connect to the EVM client"))
			}

			// Store missed events in the database before listening to new events.
			s.storeMissedEVMDepositedEvents(client, chain)

			// Listen to new events from the chain and store them in the database.
			go s.storeEVMDepositedEvents(client, chain)

			// Store missed events in the database before listening to new events.
			s.storeMissedEVMWithdrawnEvents(client, chain)

			// Listen to new events from the chain and store them in the database.
			go s.storeEVMWithdrawnEvents(client, chain)

		default:
			s.logger.Errorf("unsupported chain type: %s", chain.Type)
		}
	}
}
