package responses

import (
	"encoding/json"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/api/shared"
)

// NewChainListResponse creates a new shared.ChainListResponse from the given chains and tokens.
func NewChainListResponse(chains []datashared.Chain, tokens []datashared.Token) (shared.ChainListResponse, error) {
	response := shared.ChainListResponse{
		Data:     make([]shared.Chain, len(chains)),
		Included: make([]json.RawMessage, len(tokens)),
	}

	for i, chain := range chains {
		response.Data[i] = shared.Chain{
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
					Data: make([]datashared.Key, len(chain.Tokens)),
				},
			},
		}

		for j, token := range chain.Tokens {
			response.Data[i].Relationships.Tokens.Data[j] = datashared.Key{
				ID:   token.TokenID,
				Type: datashared.TOKEN,
			}
		}
	}

	// Create token models and add them to the included array of the response.
	for i, token := range tokens {
		tokenModel := &shared.Token{
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
					Data: make([]datashared.Key, 0),
				},
			},
		}

		for _, chain := range token.Chains {
			tokenModel.Relationships.Chains.Data = append(tokenModel.Relationships.Chains.Data, datashared.Key{
				ID:   chain.ChainID,
				Type: datashared.CHAIN,
			})
		}

		// Marshal the token model to JSON and add it to the included array of the response.
		tokenModelJSON, err := json.Marshal(tokenModel)
		if err != nil {
			return shared.ChainListResponse{}, errors.Wrap(err, "failed to marshal token model")
		}

		response.Included[i] = tokenModelJSON
	}

	return response, nil
}
