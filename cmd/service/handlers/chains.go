package handlers

import (
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/requests"
	"github.com/quantum-bridge/core/cmd/service/responses"
	"net/http"
)

// GetChains is an HTTP handler that returns a list of chains and tokens based on the request.
// @Summary Get chains list
// @Description Get a list of chains and tokens based on the request.
// @ID getChains
// @Tags Chains
// @Accept json
// @Produce json
// @Param filter[chain_type] query array false "Filter by chain type. Items Value: [`'evm'`]"
// @Param include_tokens query bool false "Include tokens in the response"
// @Success 200 {object} shared.ChainListResponse "Successful operation"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /chains [get]
func GetChains(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get the filter type and include tokens.
	request, err := requests.NewGetChainsRequest(r)
	if err != nil {
		Log(r.Context()).Errorf("failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Get the chains from the RAM cache.
	chains := Chains(r.Context()).Select()
	if chains == nil {
		http.Error(w, "chains not found", http.StatusInternalServerError)

		return
	}

	// Filter the chains by the filter type.
	if len(request.FilterType) > 0 {
		filteredChains := make([]datashared.Chain, 0)

		for _, chain := range chains {
			for _, filter := range request.FilterType {
				if chain.Type == filter {
					filteredChains = append(filteredChains, chain)
				}
			}
		}

		chains = filteredChains
	}

	// Get the tokens if the request includes tokens.
	var tokens []datashared.Token
	if request.IncludeTokens {
		// Get the tokens from the RAM cache.
		tokens = Tokens(r.Context()).Select()

		// Filter tokens by the chains that are being used.
		tokenIDs := tokensID(chains)
		filteredTokens := make(map[string]*datashared.Token)
		for _, token := range tokens {
			for _, tokenChain := range token.Chains {
				for _, tokenID := range tokenIDs {
					if tokenID == tokenChain.TokenID {
						filteredTokens[token.ID] = &token
					}
				}
			}
		}

		// Convert the map back to a slice.
		tokens = make([]datashared.Token, 0, len(filteredTokens))
		for _, token := range filteredTokens {
			tokens = append(tokens, *token)
		}
	}

	// Create the response for the chains.
	response, err := responses.NewChainListResponse(chains, tokens)
	if err != nil {
		http.Error(w, "Failed to create response: "+err.Error(), http.StatusInternalServerError)

		return
	}

	// Return the chains as response.
	// Set the response content type to JSON.
	respond(w, response)
}

// tokensID returns a list of token IDs from a list of chains.
func tokensID(chains []datashared.Chain) []string {
	uniqueTokenIDs := make(map[string]bool)
	var tokenIDs []string

	for _, chain := range chains {
		for _, token := range chain.Tokens {
			if _, exists := uniqueTokenIDs[token.TokenID]; !exists {
				uniqueTokenIDs[token.TokenID] = true
				tokenIDs = append(tokenIDs, token.TokenID)
			}
		}
	}

	return tokenIDs
}
