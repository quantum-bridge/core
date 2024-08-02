package requests

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
)

// GetNFTDTO is the data transfer object for the GetNFT request.
type GetNFTDTO struct {
	// TokenID is the ID of the token.
	TokenID string
	// NFTID is the ID of the non-fungible token.
	NFTID string
	// ChainID is the ID of the chain.
	ChainID string
}

// NewGetNFTRequest creates a new GetNFTDTO from an HTTP request.
func NewGetNFTRequest(r *http.Request) (GetNFTDTO, error) {
	// Create a new request object.
	var request GetNFTDTO

	// Get the token ID and NFT ID from the URL.
	request.TokenID = chi.URLParam(r, "tokenID")
	request.NFTID = chi.URLParam(r, "nftID")

	// Get the chain ID from the query parameters.
	params := r.URL.Query()

	// Get 'chain_id' query parameter (required).
	chain := params.Get("chain_id")
	if chain == "" {
		return GetNFTDTO{}, errors.New("chain is required")
	}
	request.ChainID = chain

	return request, nil
}
