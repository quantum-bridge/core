package repositories

import (
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/pkg/common"
)

// TokensRepository is the interface that defines logic for the tokens repository.
type TokensRepository interface {
	// New creates a new tokens repository instance.
	New() TokensRepository
	// Select returns the tokens that match the given filters.
	Select() []datashared.Token
	// Get returns the first token that matches the given filters.
	Get() *datashared.Token
	// FilterByType filters the tokens by the given types.
	FilterByType(tokenTypes ...data.TokenType) TokensRepository
	// FilterByTokenID filters the tokens by the given token IDs.
	FilterByTokenID(tokenIDs ...string) TokensRepository
}

// tokenFilter is a function that filters the token by the given filters.
type tokenFilter func(token datashared.Token) bool

// tokensRepository is the structure that holds the tokens repository data and filters.
type tokensRepository struct {
	tokens       []datashared.Token
	tokenFilters []tokenFilter
}

// NewTokens creates a new tokens repository instance.
func NewTokens(tokens []datashared.Token) TokensRepository {
	return &tokensRepository{
		tokens:       tokens,
		tokenFilters: make([]tokenFilter, 0),
	}
}

// New creates a new tokens repository instance.
func (r *tokensRepository) New() TokensRepository {
	return NewTokens(r.tokens)
}

// Select returns the tokens that match the given filters.
func (r *tokensRepository) Select() []datashared.Token {
	var tokens []datashared.Token
	for _, token := range r.tokens {
		if r.tokenFilter(token) {
			tokens = append(tokens, token)
		}
	}

	return tokens
}

// Get returns the first token that matches the given filters.
func (r *tokensRepository) Get() *datashared.Token {
	for _, token := range r.tokens {
		if r.tokenFilter(token) {
			return &token
		}
	}

	return &datashared.Token{}
}

// FilterByType filters the tokens by the given types.
func (r *tokensRepository) FilterByType(tokenTypes ...data.TokenType) TokensRepository {
	r.tokenFilters = append(r.tokenFilters, func(token datashared.Token) bool {
		return common.Contains(tokenTypes, token.Type)
	})

	return r
}

// FilterByTokenID filters the tokens by the given token IDs.
func (r *tokensRepository) FilterByTokenID(tokenIDs ...string) TokensRepository {
	r.tokenFilters = append(r.tokenFilters, func(token datashared.Token) bool {
		return common.Contains(tokenIDs, token.ID)
	})

	return r
}

// tokenFilter filters the token by the given filters.
func (r *tokensRepository) tokenFilter(token datashared.Token) bool {
	for _, filter := range r.tokenFilters {
		if !filter(token) {
			return false
		}
	}

	return true
}
