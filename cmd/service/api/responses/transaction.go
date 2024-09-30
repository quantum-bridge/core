package responses

import (
	"encoding/json"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/api/shared"
)

// NewTransactionResponse creates a new transactions response from the given transaction and chain.
func NewTransactionResponse(tx interface{}, chain datashared.Chain) (shared.TransactionsResponse, error) {
	// Create a new response with the transaction data.
	response := shared.TransactionsResponse{
		Data: tx,
	}

	// Marshal the chain data to JSON.
	chainData, err := json.Marshal(chain)
	if err != nil {
		return shared.TransactionsResponse{}, err
	}

	// Append the chain data to the included data of the
	response.Included = append(response.Included, chainData)

	return response, nil
}
