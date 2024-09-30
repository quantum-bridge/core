package service

import (
	"github.com/quantum-bridge/core/cmd/config"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/ipfs"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature"
	txHistory "github.com/quantum-bridge/core/cmd/service/tx-history"
	"github.com/quantum-bridge/core/pkg/squirrelizer"
	"go.uber.org/zap"
	"net"
	"net/http"
)

// service is the main service that runs the HTTP server.
type service struct {
	logger      *zap.SugaredLogger
	listener    net.Listener
	db          *squirrelizer.DB
	tokens      []datashared.Token
	chains      []datashared.Chain
	tokenChains []datashared.TokenChain
	signer      signature.Signer
	ipfs        ipfs.IPFS
	services    *config.ServicesConfig
}

// run starts the service.
func (s *service) run() error {
	if s.services.TxHistory {
		go txHistory.NewTxHistory(s.logger, s.db, s.chains).Run()
		s.logger.Infof("transaction history service is started successfully")
	}

	s.logger.Infof("API service listening on %s", s.listener.Addr().String())

	return http.Serve(s.listener, s.router())
}

// newService creates a new service with the given configuration and logger.
func newService(cfg config.Config, logger *zap.SugaredLogger) *service {
	return &service{
		logger:      logger,
		listener:    cfg.Listener(),
		db:          cfg.DB(),
		tokens:      cfg.Token(),
		chains:      cfg.Chains(),
		tokenChains: cfg.TokenChains(),
		signer:      cfg.Signer(),
		ipfs:        cfg.IPFS(),
		services:    cfg.ServicesConfig(),
	}
}

// Run starts the service with the given configuration and logger.
func Run(cfg config.Config, logger *zap.SugaredLogger) {
	svc := newService(cfg, logger)
	if err := svc.run(); err != nil {
		panic(err)
	}
}
