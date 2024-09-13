package responses

import (
	"encoding/json"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/shared"
)

// NewTokenListResponse creates a new shared.TokenListResponse from the given tokens and chains.
func NewTokenListResponse(tokens []datashared.Token, chains []datashared.Chain) (shared.TokenListResponse, error) {
	response := shared.TokenListResponse{
		Data:     make([]shared.Token, len(tokens)),
		Included: make([]json.RawMessage, len(chains)),
	}

	// Create token models and add them to the data array of the response.
	for i, token := range tokens {
		response.Data[i] = shared.Token{
			Key: datashared.Key{
				ID:   token.ID,
				Type: datashared.TOKEN,
			},
			Attributes: shared.TokenAttributes{
				Icon:      &token.Icon,
				Name:      token.Name,
				Symbol:    token.Symbol,
				TokenType: token.Type,
			},
			Relationships: shared.TokenRelationships{
				Chains: shared.RelationCollection{
					Data: make([]datashared.Key, len(token.Chains)),
				},
			},
		}

		// Add the chain IDs to the token model.
		for j, chain := range token.Chains {
			response.Data[i].Relationships.Chains.Data[j] = datashared.Key{
				ID:   chain.ChainID,
				Type: datashared.CHAIN,
			}
		}
	}

	// Create chain models and add them to the included array of the response.
	for i, chain := range chains {
		chainModel := &shared.Chain{
			Key: datashared.Key{
				ID:   chain.ID,
				Type: datashared.CHAIN,
			},
			Attributes: shared.ChainAttributes{
				ChainParams: chain.ChainParams,
				ChainType:   string(chain.Type),
				Icon:        &chain.Icon,
				Name:        chain.Name,
			},
			Relationships: shared.ChainRelationships{
				Tokens: shared.RelationCollection{
					Data: make([]datashared.Key, 0),
				},
			},
		}

		// Add the token IDs to the chain model.
		for _, token := range chain.Tokens {
			chainModel.Relationships.Tokens.Data = append(chainModel.Relationships.Tokens.Data, datashared.Key{
				ID:   token.TokenID,
				Type: datashared.TOKEN,
			})
		}

		// Marshal the chain model to JSON and add it to the included array of the response.
		chainModelJSON, err := json.Marshal(chainModel)
		if err != nil {
			return shared.TokenListResponse{}, err
		}

		response.Included[i] = chainModelJSON
	}

	return response, nil
}
