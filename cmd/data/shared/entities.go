package shared

import "math/big"

// Key represents the key of the entity.
type Key struct {
	// ID is the identifier of the entity.
	ID string `json:"id"`
	// Type is the type of the entity.
	Type EntityType `json:"type"`
}

// GetKey returns the key of the entity.
func (k *Key) GetKey() Key {
	return *k
}

// EVMTxBody represents the body of an EVM transaction.
type EVMTxBody struct {
	// ChainID is the chain ID of the transaction.
	ChainID string `json:"chain_id"`
	// From is the address of the sender.
	From string `json:"from"`
	// To is the address of the receiver.
	To string `json:"to"`
	// Value is the amount of the transaction.
	Value string `json:"value"`
	// Data is the data of the transaction.
	Data string `json:"data"`
}

// EVMAttributes represents the attributes of the entity for the EVM chain.
type EVMAttributes struct {
	// Confirmed is the flag that indicates whether the transaction is confirmed.
	Confirmed *bool `json:"confirmed"`
	// TxBody is the body of the transaction.
	TxBody EVMTxBody `json:"tx_body"`
}

// DataEntity represents the data of the entity.
type DataEntity struct {
	// ID is the identifier of the entity.
	ID string `json:"id"`
	// Type is the type of the entity.
	Type EntityType `json:"type"`
}

// ChainEntity represents the chain entity.
type ChainEntity struct {
	// Data is the data of the entity.
	Data DataEntity `json:"data"`
}

// Relationships represents the relationships of the entity.
type Relationships struct {
	// Chain is the chain of the entity.
	Chain ChainEntity `json:"chain"`
}

// EVMTransaction represents the EVM transaction entity.
type EVMTransaction struct {
	// Key is the key of the entity.
	Key Key `json:"key"`
	// Attributes are the attributes of the entity.
	Attributes EVMAttributes `json:"attributes"`
	// Relationships are the relationships of the entity.
	Relationships Relationships `json:"relationships"`
}

// FungibleLock represents the lock of a fungible token.
type FungibleLock struct {
	// To is the address of the receiver.
	To string `json:"to"`
	// Amount is the amount of the token.
	Amount *big.Int `json:"amount"`
	// Network is the network of the token.
	Network string `json:"network"`
}

// NonFungibleLock represents the lock of a non-fungible token.
type NonFungibleLock struct {
	// To is the address of the receiver.
	To string `json:"to"`
	// TokenID is the identifier of the token.
	TokenID string `json:"token_id"`
	// Network is the network of the token.
	Network string `json:"network"`
}

// NFTAttribute represents the attribute of an NFT.
type NFTAttribute struct {
	// TraitType is the type of the trait.
	TraitType string `json:"trait_type"`
	// Value is the value of the trait.
	Value string `json:"value"`
}

// NFTMetadata represents the metadata of an NFT.
type NFTMetadata struct {
	// MetadataURL is the URL of the metadata.
	MetadataURL string `json:"metadata_url"`
	// Name is the name of the NFT.
	Name string `json:"name"`
	// Image is the URL of the icon.
	Image string `json:"image"`
	// Description is the description of the NFT.
	Description *string `json:"description"`
	// AnimationURL is the URL of the animation.
	AnimationURL *string `json:"animation_url"`
	// ExternalURL is the URL of the external resource.
	ExternalURL *string `json:"external_url"`
	// Attributes are the attributes of the NFT.
	Attributes []NFTAttribute `json:"attributes"`
}
