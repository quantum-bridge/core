package responses

import (
	"fmt"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/api/shared"
	"math/big"
)

// NewBalanceResponse creates a new balance response.
func NewBalanceResponse(tokenID, tokenAddress, address string, NFT *string, balance *big.Int) (shared.BalanceResponse, error) {
	// Check if the NFT is nil.
	var id string
	if *NFT == "" {
		id = fmt.Sprintf("%s-%s", tokenID, address)
	} else {
		id = fmt.Sprintf("%s-%s-%s", tokenID, address, *NFT)
	}

	// Create the balance response.
	return shared.BalanceResponse{
		Data: shared.Balance{
			Key: datashared.Key{
				ID:   id,
				Type: datashared.BALANCE,
			},
			Attributes: shared.BalanceAttributes{
				Amount:       balance,
				Address:      address,
				TokenAddress: tokenAddress,
			},
		},
	}, nil
}
