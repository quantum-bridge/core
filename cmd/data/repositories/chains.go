package repositories

import (
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/pkg/common"
)

// ChainsRepository is the interface for the chains repository.
type ChainsRepository interface {
	// New creates a new chains repository instance.
	New() ChainsRepository
	// Select returns the chains that match the given filters.
	Select() []datashared.Chain
	// Get returns the first chain that matches the given filters.
	Get() *datashared.Chain
	// FilterByChainID filters the chains by the given chain IDs.
	FilterByChainID(chainIDs ...string) ChainsRepository
	// FilterByType filters the chains by the given types.
	FilterByType(types ...data.ChainType) ChainsRepository
}

// chainFilter is a function that filters the chain by the given filters.
type chainFilter func(chain datashared.Chain) bool

// chainsRepository is the structure that holds the chains repository data and filters.
type chainsRepository struct {
	chains       []datashared.Chain
	chainFilters []chainFilter
}

// NewChains creates a new chains repository instance.
func NewChains(chains []datashared.Chain) ChainsRepository {
	return &chainsRepository{
		chains:       chains,
		chainFilters: make([]chainFilter, 0),
	}
}

// New creates a new chains repository instance.
func (r *chainsRepository) New() ChainsRepository {
	return NewChains(r.chains)
}

// Select returns the chains that match the given filters.
func (r *chainsRepository) Select() []datashared.Chain {
	var chains []datashared.Chain
	for _, chain := range r.chains {
		if r.chainFilter(chain) {
			chains = append(chains, chain)
		}
	}

	return chains
}

// Get returns the first chain that matches the given filters.
func (r *chainsRepository) Get() *datashared.Chain {
	for _, chain := range r.chains {
		if r.chainFilter(chain) {
			return &chain
		}
	}

	return &datashared.Chain{}
}

// FilterByChainID filters the chains by the given chain IDs.
func (r *chainsRepository) FilterByChainID(chainIDs ...string) ChainsRepository {
	r.chainFilters = append(r.chainFilters, func(chain datashared.Chain) bool {
		return common.Contains(chainIDs, chain.ID)
	})

	return r
}

// FilterByType filters the chains by the given types.
func (r *chainsRepository) FilterByType(types ...data.ChainType) ChainsRepository {
	r.chainFilters = append(r.chainFilters, func(chain datashared.Chain) bool {
		return common.Contains(types, chain.Type)
	})

	return r
}

// chainFilter filters the chain by the given filters.
func (r *chainsRepository) chainFilter(chain datashared.Chain) bool {
	for _, filter := range r.chainFilters {
		if !filter(chain) {
			return false
		}
	}
	return true
}
