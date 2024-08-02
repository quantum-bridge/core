package requests

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/service/shared"
	"net/http"
)

// GetBalanceDTO is the data transfer object for the get balance request.
type GetBalanceDTO struct {
	// TokenID is the ID of the token.
	TokenID string
	// Address is the address of the account.
	Address string
	// ChainID is the ID of the chain.
	ChainID string
	// NFT is the ID of the non-fungible token.
	NFT *string
}

// NewGetBalanceRequest creates a new get balance request from the given HTTP request.
func NewGetBalanceRequest(r *http.Request) (GetBalanceDTO, error) {
	// Create a new request object.
	request := GetBalanceDTO{}

	// Get the query parameters from the URL.
	params := r.URL.Query()

	// Get the token ID from the URL.
	request.TokenID = chi.URLParam(r, "tokenID")

	// Get the address from params.
	address := params.Get("address")

	// Check if the address is a valid address.
	if !shared.IsValidEthereumAddress(address) {
		return request, errors.New("invalid address")
	}
	request.Address = address

	// Get the chain ID from params.
	chain := params.Get("chain")
	if chain == "" {
		return request, errors.New("chain is required")
	}
	request.ChainID = chain

	// Get the NFT from params if it exists.
	nft := params.Get("nft")
	request.NFT = &nft

	return request, nil
}
