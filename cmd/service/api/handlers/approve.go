package handlers

import (
	"github.com/quantum-bridge/core/cmd/service/api/requests"
	"github.com/quantum-bridge/core/cmd/service/api/responses"
	"net/http"
)

// Approve is an HTTP handler that approves a token for a spender.
// @Summary Approve
// @Description Approve is an HTTP handler that creates an approval transaction for a spender.
// @ID approve
// @Tags Transfers
// @Accept json
// @Produce json
// @Param _ body requests.ApproveDTO true "Request body"
// @Success 200 {object} shared.TransactionsResponse "Successful operation"
// @Success 204 "No content"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /transfers/approve [post]
func Approve(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get the token ID, chain ID and holder address.
	request, err := requests.NewApproveRequest(r)
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

	// Get approval transaction for the given token chain and holder address.
	tx, err := Proxy(r.Context()).Get(tokenChain.ChainID).Approve(*tokenChain, request.Address)
	if err != nil {
		Log(r.Context()).Errorf("failed to approve: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// If the transaction is nil (no approval needed), return no content.
	if tx == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	// Get the chain for the given chain ID from the RAM cache.
	chain := Chains(r.Context()).FilterByChainID(request.ChainID).Get()
	if chain == nil {
		Log(r.Context()).Error("chain not found")
		http.Error(w, "chain not found", http.StatusNotFound)

		return
	}

	// Create the response for the transaction.
	response, err := responses.NewTransactionResponse(tx, *chain)
	if err != nil {
		Log(r.Context()).Errorf("failed to create response: %v", err)
		http.Error(w, "failed to create response", http.StatusInternalServerError)

		return
	}

	// Return the transaction as response.
	// Set the response content type to JSON.
	respond(w, response)
}
