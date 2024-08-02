package repositories

import (
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/pkg/common"
)

// TokenChainsRepository is the interface that defines the token chains repository.
type TokenChainsRepository interface {
	// New creates a new token chains repository instance.
	New() TokenChainsRepository
	// Select returns the token chains that match the given filters.
	Select() []datashared.TokenChain
	// Get returns the first token chain that matches the given filters.
	Get() *datashared.TokenChain
	// FilterByTokenID filters the token chains by the given token IDs.
	FilterByTokenID(IDs ...string) TokenChainsRepository
	// FilterByChainID filters the token chains by the given chain IDs.
	FilterByChainID(IDs ...string) TokenChainsRepository
	// FilterByTokenType filters the token chains by the given token types.
	FilterByTokenType(types ...string) TokenChainsRepository
	// FilterByBridgeType filters the token chains by the given bridge types.
	FilterByBridgeType(types ...string) TokenChainsRepository
}

// tokenChainFilter is a function that filters the token chain by the given filters.
type tokenChainFilter func(tokenChain datashared.TokenChain) bool

// tokenChainsRepository is the structure that holds the token chains repository data and filters.
type tokenChainsRepository struct {
	tokenChains       []datashared.TokenChain
	tokenChainFilters []tokenChainFilter
}

// NewTokenChains creates a new token chains repository instance.
func NewTokenChains(tokenChains []datashared.TokenChain) TokenChainsRepository {
	return &tokenChainsRepository{
		tokenChains:       tokenChains,
		tokenChainFilters: make([]tokenChainFilter, 0),
	}
}

// New creates a new token chains repository instance.
func (r *tokenChainsRepository) New() TokenChainsRepository {
	return NewTokenChains(r.tokenChains)
}

// Select returns the token chains that match the given filters.
func (r *tokenChainsRepository) Select() []datashared.TokenChain {
	var tokenChains []datashared.TokenChain
	for _, tokenChain := range r.tokenChains {
		if r.tokenChainFilter(tokenChain) {
			tokenChains = append(tokenChains, tokenChain)
		}
	}

	return tokenChains
}

// Get returns the first token chain that matches the given filters.
func (r *tokenChainsRepository) Get() *datashared.TokenChain {
	for _, tokenChain := range r.tokenChains {
		if r.tokenChainFilter(tokenChain) {
			return &tokenChain
		}
	}

	return &datashared.TokenChain{}
}

// FilterByTokenID filters the token chains by the given token IDs.
func (r *tokenChainsRepository) FilterByTokenID(tokenIDs ...string) TokenChainsRepository {
	r.tokenChainFilters = append(r.tokenChainFilters, func(tokenChain datashared.TokenChain) bool {
		return common.Contains(tokenIDs, tokenChain.TokenID)
	})

	return r
}

// FilterByChainID filters the token chains by the given chain IDs.
func (r *tokenChainsRepository) FilterByChainID(chainIDs ...string) TokenChainsRepository {
	r.tokenChainFilters = append(r.tokenChainFilters, func(tokenChain datashared.TokenChain) bool {
		return common.Contains(chainIDs, tokenChain.ChainID)
	})

	return r
}

// FilterByTokenType filters the token chains by the given token types.
func (r *tokenChainsRepository) FilterByTokenType(tokenTypes ...string) TokenChainsRepository {
	r.tokenChainFilters = append(r.tokenChainFilters, func(tokenChain datashared.TokenChain) bool {
		return common.Contains(tokenTypes, tokenChain.TokenType)
	})

	return r
}

// FilterByBridgeType filters the token chains by the given bridge types.
func (r *tokenChainsRepository) FilterByBridgeType(types ...string) TokenChainsRepository {
	r.tokenChainFilters = append(r.tokenChainFilters, func(tokenChain datashared.TokenChain) bool {
		return common.Contains(types, tokenChain.BridgeType.String())
	})

	return r
}

// tokenChainFilter filters the token chain by the given filters.
func (r *tokenChainsRepository) tokenChainFilter(tokenChain datashared.TokenChain) bool {
	for _, filter := range r.tokenChainFilters {
		if !filter(tokenChain) {
			return false
		}
	}

	return true
}
