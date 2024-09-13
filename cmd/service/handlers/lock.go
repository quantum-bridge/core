package handlers

import (
	"github.com/quantum-bridge/core/cmd/data"
	"github.com/quantum-bridge/core/cmd/service/requests"
	"github.com/quantum-bridge/core/cmd/service/responses"
	"math/big"
	"net/http"
)

// One is a big integer with the value of 1 * 10^18 (EVMs wei).
var One = big.NewInt(1000000000000000000)

// Lock is an HTTP handler that locks a token from one chain to another.
// @Summary Lock Token
// @Description Generates transaction that will lock a token in the source chain.
// @ID lock
// @Tags Transfers
// @Accept json
// @Produce json
// @Param _ body requests.LockDTO true "Request body"
// @Success 200 {object} shared.TransactionsResponse "Successful operation"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /transfers/lock [post]
func Lock(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get the token ID.
	request, err := requests.NewLockRequest(r)
	if err != nil {
		Log(r.Context()).Errorf("failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// Get the token from the RAM cache based on the token ID from the request.
	token := Tokens(r.Context()).FilterByTokenID(request.TokenID).Get()
	if token == nil {
		Log(r.Context()).Error("token not found")
		http.Error(w, "token not found or does not exists", http.StatusBadRequest)

		return
	}

	// Get the token chain for the given token ID and chain ID from the RAM cache.
	tokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(request.ChainFrom).Get()
	if tokenChain == nil {
		Log(r.Context()).Error("token chain not found")
		http.Error(w, "token chain not found", http.StatusBadRequest)

		return
	}

	// Get the destination token chain for the given token ID and chain ID from the RAM cache.
	destinationTokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(request.ChainTo).Get()
	if destinationTokenChain == nil {
		Log(r.Context()).Error("destination token chain not found")
		http.Error(w, "destination token chain not found", http.StatusBadRequest)

		return
	}

	// Check if the destination token chain is a liquidity pool.
	if destinationTokenChain.BridgeType == data.BridgeTypeLP {
		balance, err := Proxy(r.Context()).Get(destinationTokenChain.ChainID).BridgeBalance(*destinationTokenChain, request.NFT)
		if err != nil {
			Log(r.Context()).Errorf("failed to get balance: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		// Set the amount to the request amount if it is not nil.
		amount := new(big.Int).Set(One)
		if request.Amount != nil {
			amount = request.Amount
		}

		// Check if the balance is sufficient for the lock.
		if balance.Cmp(amount) == -1 {
			Log(r.Context()).Error("insufficient balance in destination chain")
			http.Error(w, "insufficient balance", http.StatusBadRequest)

			return
		}
	}

	// Lock the token based on the token type and create the transaction object.
	var tx interface{}
	switch token.Type {
	case data.FUNGIBLE:
		if request.Amount == nil {
			Log(r.Context()).Error("amount is required for fungible tokens")
			http.Error(w, "amount is required for fungible tokens", http.StatusBadRequest)

			return
		}

		tx, err = Proxy(r.Context()).Get(tokenChain.ChainID).LockFungible(*tokenChain, request.From, request.To, request.ChainTo, request.Amount.String())
	case data.NONFUNGIBLE:
		if request.NFT == nil {
			Log(r.Context()).Error("nft id is required for non-fungible tokens")
			http.Error(w, "nft id is required for non-fungible tokens", http.StatusBadRequest)

			return
		}

		tx, err = Proxy(r.Context()).Get(tokenChain.ChainID).LockNonFungible(*tokenChain, request.From, request.To, request.ChainTo, *request.NFT)
	default:
		Log(r.Context()).Errorf("unsupported token type: %v, token ID: %v", token.Type, token.ID)
		http.Error(w, "unsupported token type", http.StatusBadRequest)

		return
	}
	if err != nil {
		Log(r.Context()).Errorf("failed to lock token: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Get the chain for the given chain ID from the RAM cache.
	chain := Chains(r.Context()).FilterByChainID(request.ChainFrom).Get()
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

	respond(w, response)
}
