package requests

import (
	"encoding/json"
	"math/big"
	"net/http"
)

// LockRequest is the data transfer object for the Lock request.
type LockRequest struct {
	//Amount is the amount of tokens to lock.
	Amount *big.Int `json:"amount,omitempty"`
	// ChainFrom is the chain that the lock is from.
	ChainFrom string `json:"chain_from"`
	// ChainTo is the chain that is receiving the amount of tokens.
	ChainTo string `json:"chain_to"`
	// From is the sender address of the lock.
	From string `json:"from"`
	// To is the receiver address of the lock.
	To string `json:"to"`
	// TokenID is the ID of the token being locked in the chain.
	TokenID string `json:"token_id"`
	// NFTID is the ID of the NFT being locked in the chain.
	NFTID *string `json:"nft_id,omitempty"`
}

// LockDTO is the data transfer object for the Lock request.
type LockDTO struct {
	Data LockRequest `json:"data"`
}

// NewLockRequest creates a new LockDTO from an HTTP request.
func NewLockRequest(r *http.Request) (LockRequest, error) {
	// Create a new request object.
	request := LockDTO{}

	// Decode the request body into the DTO object.
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return LockRequest{}, err
	}

	return request.Data, nil
}
