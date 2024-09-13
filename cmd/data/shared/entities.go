package shared

import "math/big"

// Key represents the key of the entity.
type Key struct {
	// ID is the identifier of the entity.
	ID string `json:"id" binding:"required"`
	// Type is the type of the entity.
	Type EntityType `json:"type" binding:"required"`
}

// GetKey returns the key of the entity.
func (k *Key) GetKey() Key {
	return *k
}

// EVMTxBody represents the body of an EVM transaction.
type EVMTxBody struct {
	// ChainID is the chain ID of the transaction.
	ChainID string `json:"chain_id" binding:"required"`
	// From is the address of the sender.
	From string `json:"from" binding:"required"`
	// To is the address of the receiver.
	To string `json:"to" binding:"required"`
	// Value is the amount of the transaction.
	Value string `json:"value" binding:"required"`
	// Data is the data of the transaction.
	Data string `json:"data" binding:"required"`
}

// EVMAttributes represents the attributes of the entity for the EVM chain.
type EVMAttributes struct {
	// Confirmed is the flag that indicates whether the transaction is confirmed.
	Confirmed *bool `json:"confirmed" binding:"required"`
	// TxBody is the body of the transaction.
	TxBody EVMTxBody `json:"tx_body" binding:"required"`
}

// DataEntity represents the data of the entity.
type DataEntity struct {
	// ID is the identifier of the entity.
	ID string `json:"id" binding:"required"`
	// Type is the type of the entity.
	Type EntityType `json:"type" binding:"required"`
}

// ChainEntity represents the chain entity.
type ChainEntity struct {
	// Data is the data of the entity.
	Data DataEntity `json:"data" binding:"required"`
}

// Relationships represents the relationships of the entity.
type Relationships struct {
	// Chain is the chain of the entity.
	Chain ChainEntity `json:"chain" binding:"required"`
}

// EVMTransaction represents the EVM transaction entity.
type EVMTransaction struct {
	// Key is the key of the entity.
	Key Key `json:"key" binding:"required"`
	// Attributes are the attributes of the entity.
	Attributes EVMAttributes `json:"attributes" binding:"required"`
	// Relationships are the relationships of the entity.
	Relationships Relationships `json:"relationships" binding:"required"`
}

// FungibleLock represents the lock of a fungible token.
type FungibleLock struct {
	// To is the address of the receiver.
	To string `json:"to" binding:"required"`
	// Amount is the amount of the token.
	Amount *big.Int `json:"amount" binding:"required"`
	// Network is the network of the token.
	Network string `json:"network" binding:"required"`
}

// NonFungibleLock represents the lock of a non-fungible token.
type NonFungibleLock struct {
	// To is the address of the receiver.
	To string `json:"to" binding:"required"`
	// TokenID is the identifier of the token.
	TokenID string `json:"token_id" binding:"required"`
	// Network is the network of the token.
	Network string `json:"network" binding:"required"`
}

// NFTAttribute represents the attribute of an NFT.
type NFTAttribute struct {
	// TraitType is the type of the trait.
	TraitType string `json:"trait_type" binding:"required"`
	// Value is the value of the trait.
	Value string `json:"value" binding:"required"`
}

// NFTMetadata represents the metadata of an NFT.
type NFTMetadata struct {
	// MetadataURL is the URL of the metadata.
	MetadataURL string `json:"metadata_url" binding:"required"`
	// Name is the name of the NFT.
	Name string `json:"name" binding:"required"`
	// Image is the URL of the icon.
	Image string `json:"image" binding:"required"`
	// Description is the description of the NFT.
	Description *string `json:"description"`
	// AnimationURL is the URL of the animation.
	AnimationURL *string `json:"animation_url"`
	// ExternalURL is the URL of the external resource.
	ExternalURL *string `json:"external_url"`
	// Attributes are the attributes of the NFT.
	Attributes []NFTAttribute `json:"attributes" binding:"required"`
}
