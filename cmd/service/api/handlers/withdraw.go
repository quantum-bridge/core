package handlers

import (
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data"
	"github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/api/requests"
	"github.com/quantum-bridge/core/cmd/service/api/responses"
	"net/http"
)

// Withdraw is an HTTP handler that withdraws the tokens from the bridge.
// @Summary Withdraw
// @Description Check if lock transaction is valid and withdraw the token from the bridge. Returns the transaction
// to send to the destination chain to withdraw tokens. If service configuration is set to auto withdraw,
// the transaction will be sent to the destination chain automatically and the response will be the transaction hash.
// @ID withdraw
// @Tags Transfers
// @Accept json
// @Produce json
// @Param _ body requests.WithdrawDTO true "Request body"
// @Success 200 {object} shared.TransactionsResponse "Successful operation"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /transfers/withdraw [post]
func Withdraw(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewWithdrawRequest(r)
	if err != nil {
		Log(r.Context()).Errorf("failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Get the token from the RAM cache based on the token ID from the request.
	token := Tokens(r.Context()).FilterByTokenID(request.TokenID).Get()
	if token == nil {
		Log(r.Context()).Error("token not found")
		http.Error(w, "token not found", http.StatusNotFound)

		return
	}

	// Get the token chain from the RAM cache based on the token ID and chain ID from the request.
	tokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(request.ChainFrom).Get()
	if tokenChain == nil {
		Log(r.Context()).Error("token chain not found")
		http.Error(w, "token chain not found", http.StatusNotFound)

		return
	}

	// Check the token type and withdraw the token from the bridge.
	switch token.Type {
	case data.FUNGIBLE:
		// Withdraw the fungible token from the bridge.
		withdrawFungible(w, r, request, tokenChain)
	case data.NONFUNGIBLE:
		// Withdraw the non-fungible token from the bridge.
		withdrawNonFungible(w, r, request, tokenChain)
	default:
		Log(r.Context()).Error("token type not supported")
		http.Error(w, "token type not supported", http.StatusBadRequest)

		return
	}
}

// withdrawFungible withdraws the fungible token from the bridge.
func withdrawFungible(w http.ResponseWriter, r *http.Request, request requests.WithdrawRequest, tokenChain *shared.TokenChain) {
	// Check if the event is valid.
	event, err := Proxy(r.Context()).Get(tokenChain.ChainID).CheckFungible(request.TxHash, *tokenChain)
	if err != nil {
		Log(r.Context()).Errorf("failed to withdraw fungible: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Get the destination token chain.
	destinationTokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(event.Network).Get()
	if destinationTokenChain == nil {
		Log(r.Context()).Error("destination token chain not found")
		http.Error(w, "token does not connected to the destination chain", http.StatusNotFound)

		return
	}

	// Set the from address to the event to address if it is nil.
	if request.From == nil {
		request.From = &event.To
	}

	// Withdraw the fungible token from the bridge and get the transaction.
	tx, err := Proxy(r.Context()).Get(destinationTokenChain.ChainID).WithdrawFungible(
		*destinationTokenChain,
		*request.From,
		event.To,
		request.TxHash,
		event.Amount,
	)
	if err != nil {
		Log(r.Context()).Errorf("failed to withdraw fungible: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Get the chain for the given chain ID from the RAM cache.
	chain := Chains(r.Context()).FilterByChainID(destinationTokenChain.ChainID).Get()

	// Create the response for the transaction.
	response, err := responses.NewTransactionResponse(tx, *chain)
	if err != nil {
		Log(r.Context()).Errorf("failed to create response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	respond(w, response)
}

// withdrawNonFungible withdraws the non-fungible token from the bridge.
func withdrawNonFungible(w http.ResponseWriter, r *http.Request, request requests.WithdrawRequest, tokenChain *shared.TokenChain) {
	// Check if the event is valid.
	event, err := Proxy(r.Context()).Get(tokenChain.ChainID).CheckNonFungible(request.TxHash, *tokenChain)
	if err != nil {
		Log(r.Context()).Errorf("failed to withdraw fungible: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Get the destination token chain.
	destinationTokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(event.Network).Get()
	if destinationTokenChain == nil {
		Log(r.Context()).Error("destination token chain not found")
		http.Error(w, "token does not connected to the destination chain", http.StatusNotFound)

		return
	}

	// Get the original URI of the NFT from the bridge.
	uri, err := getOriginalURI(r, request.TokenID, event.TokenID)
	if err != nil {
		Log(r.Context()).Errorf("failed to get original URI: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Set the from address to the event to address if it is nil.
	if request.From == nil {
		request.From = &event.To
	}

	// Withdraw the non-fungible token from the bridge and get the transaction.
	tx, err := Proxy(r.Context()).Get(destinationTokenChain.ChainID).WithdrawNonFungible(
		*destinationTokenChain,
		*request.From,
		event.To,
		request.TxHash,
		event.TokenID,
		uri,
	)
	if err != nil {
		Log(r.Context()).Errorf("failed to withdraw non-fungible: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Get the chain for the given chain ID from the RAM cache.
	chain := Chains(r.Context()).FilterByChainID(destinationTokenChain.ChainID).Get()

	// Create the response for the transaction.
	response, err := responses.NewTransactionResponse(tx, *chain)
	if err != nil {
		Log(r.Context()).Errorf("failed to create response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	respond(w, response)
}

// getOriginalURI gets the original URI of the NFT.
func getOriginalURI(r *http.Request, tokenID, nftID string) (string, error) {
	// Get the token chain from the RAM cache based on the token ID and bridge type.
	tokenChain := TokenChains(r.Context()).FilterByTokenID(tokenID).FilterByBridgeType(data.BridgeTypeLP.String()).Get()
	if tokenChain == nil {
		return "", errors.New("token chain not found")
	}

	// Get the NFT metadata URI from the bridge.
	uri, err := Proxy(r.Context()).Get(tokenChain.ChainID).GetNFTMetadataURI(*tokenChain, nftID)
	if err != nil {
		return "", errors.Wrap(err, "failed to get NFT metadata URI")
	}

	return uri, nil
}
