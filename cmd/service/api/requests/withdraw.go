package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

// WithdrawRequest is the request to withdraw tokens from the bridge.
type WithdrawRequest struct {
	// ChainFrom is the source chain ID.
	ChainFrom string `json:"chain_from" binding:"required"`
	// From is the address of the sender in the destination chain. Should be used only if the sender address is different with source chain.
	From *string `json:"from,omitempty"`
	// TokenID is the token ID of the token.
	TokenID string `json:"token_id" binding:"required"`
	// TxHash is the hash of the transaction in the source chain that locked the token.
	TxHash string `json:"tx_hash" binding:"required"`
}

// WithdrawDTO is the data transfer object for the withdrawal request.
type WithdrawDTO struct {
	// Data is the data of the withdrawal request.
	Data WithdrawRequest `json:"data" binding:"required"`
}

// NewWithdrawRequest creates a new withdrawal request from the given HTTP request.
func NewWithdrawRequest(r *http.Request) (WithdrawRequest, error) {
	request := WithdrawDTO{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return WithdrawRequest{}, errors.Wrap(err, "failed to decode request")
	}

	return request.Data, nil
}
