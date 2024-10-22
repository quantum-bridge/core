package handlers

import (
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data"
	"github.com/quantum-bridge/core/cmd/service/api/requests"
	"github.com/quantum-bridge/core/cmd/service/api/responses"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
	"net/http"
)

// GetNFT is an HTTP handler that returns the metadata of a non-fungible token based on the request.
// @Summary Get NFT metadata
// @Description Get the metadata of a non-fungible token based on the token ID and NFT ID.
// @ID getNFT
// @Tags Tokens
// @Accept json
// @Produce json
// @Param token_id path string true "Token ID"
// @Param nft_id path string true "NFT ID"
// @Param chain_id query string true "Chain ID"
// @Success 200 {object} shared.NFTResponse "Successful operation"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /tokens/{token_id}/nfts/{nft_id} [get]
func GetNFT(w http.ResponseWriter, r *http.Request) {
	// Parse the request to get the token ID, NFT ID, and chain ID.
	request, err := requests.NewGetNFTRequest(r)
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

	// Check if the token is non-fungible or not.
	if token.Type != data.NONFUNGIBLE {
		Log(r.Context()).Error("token is not non-fungible")
		http.Error(w, "token is not non-fungible", http.StatusBadRequest)

		return
	}

	// Get the token chain from the RAM cache based on the token ID and chain ID from the request.
	tokenChain := TokenChains(r.Context()).FilterByTokenID(request.TokenID).FilterByChainID(request.ChainID).Get()
	if tokenChain == nil {
		Log(r.Context()).Error("token chain not found")
		http.Error(w, "token chain not found", http.StatusNotFound)

		return
	}

	// Get the nft metadata from the bridge.
	metadata, err := Proxy(r.Context()).Get(tokenChain.ChainID).GetNFTMetadata(*tokenChain, request.NFTID)
	if err != nil {
		// Check if the error is a not found error and return a 404 status code.
		if errors.Is(err, bridgeErrors.ErrNotFound) {
			Log(r.Context()).Error("nft not found")
			http.Error(w, "nft not found", http.StatusNotFound)

			return
		}

		// Log the error and return a 500 status code.
		Log(r.Context()).Errorf("failed to get nft metadata: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Create a new nft response.
	response := responses.NewNFTResponse(request.NFTID, *metadata)

	// Return the nft metadata as response.
	// Set the response content type to JSON.
	respond(w, response)
}
