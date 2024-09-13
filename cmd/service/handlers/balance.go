package handlers

import (
	"github.com/quantum-bridge/core/cmd/data"
	"github.com/quantum-bridge/core/cmd/service/requests"
	"github.com/quantum-bridge/core/cmd/service/responses"
	"net/http"
)

// GetBalance is an HTTP handler that returns the balance of an account.
// @Summary Get Balance
// @Description Get the balance of an account for a token.
// @ID getBalance
// @Tags Tokens
// @Accept json
// @Produce json
// @Param token_id path string true "Token ID"
// @Param address query string true "Address of the account"
// @Param chain query string true "Chain ID"
// @Param nft query string false "NFT ID"
// @Success 200 {object} shared.BalanceResponse "Successful operation"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /tokens/{token_id}/balance [get]
func GetBalance(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get data of the request.
	request, err := requests.NewGetBalanceRequest(r)
	if err != nil {
		Log(r.Context()).Errorf("failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Get the token chain for the given token ID and chain ID from the RAM cache.
	tokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(request.ChainID).Get()
	if tokenChain == nil {
		Log(r.Context()).Error("token chain not found")
		http.Error(w, "token chain not found", http.StatusNotFound)

		return
	}

	// Get the token for the given token ID from the RAM cache.
	token := Tokens(r.Context()).FilterByTokenID(tokenChain.TokenID).Get()
	if token == nil {
		Log(r.Context()).Error("token not found")
		http.Error(w, "token not found", http.StatusNotFound)

		return
	}

	// Check if the token is non-fungible and the NFT is nil.
	if token.Type == data.NONFUNGIBLE && request.NFT == nil {
		Log(r.Context()).Error("nft is required for non-fungible tokens")
		http.Error(w, "nft is required for non-fungible tokens", http.StatusBadRequest)

		return
	}

	// Get the balance of the account.
	balance, err := Proxy(r.Context()).Get(tokenChain.ChainID).Balance(*tokenChain, request.Address, request.NFT)
	if err != nil {
		Log(r.Context()).Errorf("failed to get balance: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Create the balance response.
	response, err := responses.NewBalanceResponse(tokenChain.TokenID, tokenChain.TokenAddress, request.Address, request.NFT, balance)
	if err != nil {
		Log(r.Context()).Errorf("failed to create balance response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Return the balance response.
	// Set the response content type to JSON.
	respond(w, response)
}
