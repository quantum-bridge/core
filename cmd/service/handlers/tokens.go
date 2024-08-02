package handlers

import (
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/requests"
	"github.com/quantum-bridge/core/cmd/service/responses"
	"net/http"
)

// GetTokens is an HTTP handler that returns a list of tokens based on the request.
func GetTokens(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get the filter type and include chains.
	request, err := requests.NewGetTokensRequest(r)
	if err != nil {
		Log(r.Context()).Errorf("failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Get the tokens from the RAM cache.
	tokens := Tokens(r.Context()).Select()

	// Filter the tokens based on the filter type from the request if it is not empty.
	if len(request.FilterType) > 0 {
		filteredTokens := make([]datashared.Token, 0)
		for _, token := range tokens {
			for _, filter := range request.FilterType {
				if token.Type == filter {
					filteredTokens = append(filteredTokens, token)
				}
			}
		}

		tokens = filteredTokens
	}

	// Include the chains in the response if the include chains flag is set.
	var includedChains []datashared.Chain
	if request.IncludeChains {
		chains := Chains(r.Context()).Select()
		if chains == nil {
			http.Error(w, "chains not found", http.StatusInternalServerError)

			return
		}

		// Get the chain IDs that are used by the tokens.
		chainIDs := chainsId(tokens)
		includedChains = make([]datashared.Chain, 0, len(chainIDs))
		for _, chain := range chains {
			for _, chainID := range chainIDs {
				if chainID == chain.ID {
					includedChains = append(includedChains, chain)
				}
			}
		}
	}

	// Create the response object.
	response, err := responses.NewTokenListResponse(tokens, includedChains)
	if err != nil {
		Log(r.Context()).Errorf("failed to create response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Return the tokens as response.
	// Set the response content type to JSON.
	respond(w, response)
}

// chainsId returns a list of unique chain IDs that are used by the given tokens.
func chainsId(tokens []datashared.Token) []string {
	uniqueChainIDs := make(map[string]bool)
	var chainIDs []string

	for _, token := range tokens {
		for _, chain := range token.Chains {
			if _, exists := uniqueChainIDs[chain.ChainID]; !exists {
				uniqueChainIDs[chain.ChainID] = true
				chainIDs = append(chainIDs, chain.ChainID)
			}
		}
	}

	return chainIDs
}
