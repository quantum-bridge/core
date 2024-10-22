package handlers

import (
	"context"
	pgrepositories "github.com/quantum-bridge/core/cmd/data/postgresql/repositories"
	"github.com/quantum-bridge/core/cmd/data/repositories"
	"github.com/quantum-bridge/core/cmd/proxy"
	"go.uber.org/zap"
)

// contextKey is the type for the context keys used in the RAM cache.
type contextKey int

// All context keys used in the RAM cache for the service to store data in the context object.
const (
	// loggerCtx is the context key for the logger.
	loggerCtx contextKey = iota
	// tokensCtx is the context key for the tokens repository.
	tokensCtx
	// chainsCtx is the context key for the chains repository.
	chainsCtx
	// tokenChainsCtx is the context key for the token chains repository.
	tokenChainsCtx
	// proxyCtx is the context key for the proxy.
	proxyCtx
	// depositsHistoryCtx is the context key for the database.
	depositsHistoryCtx
)

// LogContextMiddleware is a middleware that adds the logger to the context.
func LogContextMiddleware(logger *zap.SugaredLogger) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, loggerCtx, logger)
	}
}

// TokensContextMiddleware is a middleware that adds the tokens repository to the context.
func TokensContextMiddleware(tokensRepository repositories.TokensRepository) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, tokensCtx, tokensRepository)
	}
}

// ChainsContextMiddleware is a middleware that adds the chains repository to the context.
func ChainsContextMiddleware(chainsRepository repositories.ChainsRepository) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, chainsCtx, chainsRepository)
	}
}

// TokenChainsContextMiddleware is a middleware that adds the token chains repository to the context.
func TokenChainsContextMiddleware(tokenChainsRepository repositories.TokenChainsRepository) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, tokenChainsCtx, tokenChainsRepository)
	}
}

// ProxyContextMiddleware is a middleware that adds the proxy to the context.
func ProxyContextMiddleware(proxy proxy.Proxy) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, proxyCtx, proxy)
	}
}

// DepositsHistoryContextMiddleware is a middleware that adds the database to the context.
func DepositsHistoryContextMiddleware(transactionsHistoryQuery pgrepositories.DepositsHistoryRepository) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, depositsHistoryCtx, transactionsHistoryQuery)
	}
}

// Log returns the logger from the RAM cache.
func Log(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(loggerCtx).(*zap.SugaredLogger)
}

// Tokens returns the tokens repository from the RAM cache.
func Tokens(ctx context.Context) repositories.TokensRepository {
	return ctx.Value(tokensCtx).(repositories.TokensRepository).New()
}

// Chains returns the chains repository from the RAM cache.
func Chains(ctx context.Context) repositories.ChainsRepository {
	return ctx.Value(chainsCtx).(repositories.ChainsRepository).New()
}

// TokenChains returns the token chains repository from the RAM cache.
func TokenChains(ctx context.Context) repositories.TokenChainsRepository {
	return ctx.Value(tokenChainsCtx).(repositories.TokenChainsRepository).New()
}

// Proxy returns the proxy from the RAM cache. Proxy is used to make requests to the blockchain.
func Proxy(ctx context.Context) proxy.Proxy {
	return ctx.Value(proxyCtx).(proxy.Proxy)
}

// DepositsHistoryQuery returns the database client that are being used in the bridge.
func DepositsHistoryQuery(ctx context.Context) pgrepositories.DepositsHistoryRepository {
	return ctx.Value(depositsHistoryCtx).(pgrepositories.DepositsHistoryRepository).New()
}
