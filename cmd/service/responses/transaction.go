package responses

import (
	"encoding/json"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
)

// TransactionsResponse is the response for the transactions.
type TransactionsResponse struct {
	// Data is the data of the response.
	Data interface{} `json:"data"`
	// Included is the included data of the response.
	Included []json.RawMessage `json:"included"`
}

// NewTransactionResponse creates a new transactions response from the given transaction and chain.
func NewTransactionResponse(tx interface{}, chain datashared.Chain) (TransactionsResponse, error) {
	// Create a new response with the transaction data.
	response := TransactionsResponse{
		Data: tx,
	}

	// Marshal the chain data to JSON.
	chainData, err := json.Marshal(chain)
	if err != nil {
		return TransactionsResponse{}, err
	}

	// Append the chain data to the included data of the
	response.Included = append(response.Included, chainData)

	return response, nil
}
