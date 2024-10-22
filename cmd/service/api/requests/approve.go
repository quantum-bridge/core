package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

// ApproveRequest is the request to approve a token.
type ApproveRequest struct {
	// Address is the address of the spender.
	Address string `json:"address" binding:"required"`
	// ChainID is the ID of the chain.
	ChainID string `json:"chain_id" binding:"required"`
	// TokenID is the ID of the token.
	TokenID string `json:"token_id" binding:"required"`
}

// ApproveDTO is the data transfer object for the approval request.
type ApproveDTO struct {
	// Data is the data of the approval request.
	Data ApproveRequest `json:"data" binding:"required"`
}

// NewApproveRequest creates a new approve request from the given HTTP request.
func NewApproveRequest(r *http.Request) (ApproveRequest, error) {
	request := ApproveDTO{}

	// Decode the request body to get the data.
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to decode request")
	}

	return request.Data, nil
}
