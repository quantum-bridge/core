package service

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	pgrepositories "github.com/quantum-bridge/core/cmd/data/postgresql/repositories"
	"github.com/quantum-bridge/core/cmd/data/repositories"
	"github.com/quantum-bridge/core/cmd/proxy"
	"github.com/quantum-bridge/core/cmd/service/api/handlers"
	"net/http"
)

// router is a method that returns that creates a new router with the given middlewares, routes and handlers.
func (s *service) router() chi.Router {
	router := chi.NewRouter()

	proxies, err := proxy.NewProxy(s.chains, s.signer, s.ipfs)
	if err != nil {
		s.logger.Fatalf("failed to create proxy: %v", err)

		return nil
	}

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		contextMiddleware(
			handlers.LogContextMiddleware(s.logger),
			handlers.TokensContextMiddleware(repositories.NewTokens(s.tokens)),
			handlers.ChainsContextMiddleware(repositories.NewChains(s.chains)),
			handlers.TokenChainsContextMiddleware(repositories.NewTokenChains(s.tokenChains)),
			handlers.ProxyContextMiddleware(proxies),
			handlers.DepositsHistoryContextMiddleware(pgrepositories.NewDepositsHistory(s.db)),
			handlers.WithdrawalsHistoryContextMiddleware(pgrepositories.NewWithdrawalsHistory(s.db)),
		),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/chains", handlers.GetChains)
		r.Route("/tokens", func(r chi.Router) {
			r.Get("/", handlers.GetTokens)
			r.Route("/{tokenID}", func(r chi.Router) {
				r.Get("/balance", handlers.GetBalance)
				r.Get("/nfts/{nftID}", handlers.GetNFT)
			})
		})
		r.Route("/transfers", func(r chi.Router) {
			r.Post("/approve", handlers.Approve)
			r.Post("/lock", handlers.Lock)
			r.Post("/withdraw", handlers.Withdraw)
		})
		r.Get("/history", handlers.GetHistory)
	})

	return router
}

// contextMiddleware is a middleware that adds the given context middlewares to the request context.
func contextMiddleware(middlewares ...func(context.Context) context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for _, m := range middlewares {
				ctx = m(ctx)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
