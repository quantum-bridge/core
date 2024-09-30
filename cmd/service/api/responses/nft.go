package responses

import (
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/service/api/shared"
)

// NewNFTResponse creates a new shared.NFTResponse from the given token ID and metadata.
func NewNFTResponse(tokenID string, metadata datashared.NFTMetadata) shared.NFTResponse {
	response := shared.NFTResponse{
		Data: shared.NFTData{
			Key: datashared.Key{
				ID:   tokenID,
				Type: datashared.NFT,
			},
			Attributes: shared.NFTAttributes{
				AnimationURL: metadata.AnimationURL,
				Description:  metadata.Description,
				ExternalURL:  metadata.ExternalURL,
				ImageURL:     metadata.Image,
				MetadataURL:  metadata.MetadataURL,
				Name:         metadata.Name,
				Attributes:   make([]shared.NFTAttribute, len(metadata.Attributes)),
			},
		},
	}

	// Add the attributes from the metadata to the response.
	for _, attribute := range metadata.Attributes {
		response.Data.Attributes.Attributes = append(response.Data.Attributes.Attributes, shared.NFTAttribute{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}

	return response
}
