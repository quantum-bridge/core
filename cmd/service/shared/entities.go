package shared

import (
	"encoding/json"
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"math/big"
)

// Chain is entity that represents a blockchain network in the system.
type Chain struct {
	// Key is the key of the chain entity.
	datashared.Key
	// Attributes is the attributes of the chain entity.
	Attributes ChainAttributes `json:"attributes"`
	// Relationships is the relationships of the chain entity.
	Relationships ChainRelationships `json:"relationships"`
}

// ChainAttributes is the attributes of the chain entity.
type ChainAttributes struct {
	// ChainParams is the parameters of the chain.
	ChainParams interface{} `json:"chain_params"`
	// ChainType is the type of the chain.
	ChainType string `json:"chain_type"`
	// Icon is the icon of the chain.
	Icon *string `json:"icon,omitempty"`
	// Name is the name of the chain.
	Name string `json:"name"`
}

// ChainRelationships is the relationships of the chain entity.
type ChainRelationships struct {
	// Tokens is the tokens that are used in the chain.
	Tokens RelationCollection `json:"tokens"`
}

// RelationCollection is a collection of keys.
type RelationCollection struct {
	// Data is list of Key objects.
	Data []datashared.Key `json:"data"`
}

// ChainListResponse is the response object for the list of chains.
type ChainListResponse struct {
	// Data is the list of chains.
	Data []Chain `json:"data"`
	// Included is the included object in the response.
	Included []json.RawMessage `json:"included"`
}

// Included is the included object in the response.
type Included struct {
	// Includes is the list of included objects in the response.
	Includes json.RawMessage `json:"includes"`
}

// Token is the entity that represents a token in the system.
type Token struct {
	// Key is the key of the token entity.
	datashared.Key
	// Attributes is the attributes of the token entity.
	Attributes TokenAttributes `json:"attributes"`
	// Relationships is the relationships of the token entity.
	Relationships TokenRelationships `json:"relationships"`
}

// TokenAttributes is the attributes of the token entity.
type TokenAttributes struct {
	// Icon is the icon of the token.
	Icon *string `json:"icon,omitempty"`
	// Name is the name of the token.
	Name string `json:"name"`
	// Symbol is the symbol of the token.
	Symbol string `json:"symbol"`
	// TokenType is the type of the token.
	TokenType data.TokenType `json:"token_type"`
}

// TokenRelationships is the relationships of the token entity.
type TokenRelationships struct {
	// Chains is the chains that are used by the token.
	Chains RelationCollection `json:"chains"`
}

// TokenListResponse is the response object for the list of tokens.
type TokenListResponse struct {
	// Data is the list of tokens.
	Data []Token `json:"data"`
	// Included is the included object in the response.
	Included []json.RawMessage `json:"included"`
}

// BalanceAttributes is the attributes of the balance entity.
type BalanceAttributes struct {
	// Amount is the amount of the balance.
	Amount *big.Int `json:"amount"`
	// Address is the address of the balance.
	Address string `json:"address"`
	// TokenAddress is the token address of the balance.
	TokenAddress string `json:"token_address"`
}

// Balance is the entity that represents the balance of an account.
type Balance struct {
	// Key is the key of the balance entity.
	datashared.Key
	// Attributes is the attributes of the balance entity.
	Attributes BalanceAttributes `json:"attributes"`
}

// BalanceResponse is the response object for the balance of an account.
type BalanceResponse struct {
	// Data is the balance of the account.
	Data Balance `json:"data"`
	// Included is the included object in the response.
	Included []json.RawMessage `json:"included"`
}

// NFTAttribute is the attribute of an NFT.
type NFTAttribute struct {
	// TraitType is the type of the trait.
	TraitType string `json:"trait_type"`
	// Value is the value of the trait.
	Value string `json:"value"`
}

// NFTAttributes is the attributes of an NFT.
type NFTAttributes struct {
	// Name is the name of the NFT.
	Name string `json:"name"`
	// Description is the description of the NFT.
	Description *string `json:"description"`
	// MetadataURL is the metadata URL of the NFT.
	MetadataURL string `json:"metadata_url"`
	// ExternalURL is the external URL of the NFT.
	ExternalURL *string `json:"external_url"`
	// ImageURL is the image URL of the NFT.
	ImageURL string `json:"image_url"`
	// AnimationURL is the animation URL of the NFT.
	AnimationURL *string `json:"animation_url"`
	// Attributes is the list of attributes of the NFT.
	Attributes []NFTAttribute `json:"attributes"`
}

// NFTData is the entity that represents an NFT.
type NFTData struct {
	// Key is the key of the NFT entity.
	Key datashared.Key `json:"key"`
	// Attributes is the attributes of the NFT entity.
	Attributes NFTAttributes `json:"attributes"`
}

// NFTResponse is the response object for an NFT.
type NFTResponse struct {
	// Data is the NFT.
	Data NFTData `json:"data"`
	// Includes is the included object in the response.
	Includes []json.RawMessage `json:"included"`
}
