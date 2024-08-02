package shared

// EntityType is the type of the entity.
type EntityType string

const (
	// BALANCE represents the balance entity type.
	BALANCE EntityType = "balance"
	// CHAIN represents the chain entity type.
	CHAIN EntityType = "chain"
	// EVM_TRANSACTION represents the EVM transaction entity type.
	EVM_TRANSACTION EntityType = "evm_transaction"
	// NFT represents the NFT entity type.
	NFT EntityType = "nft"
	// PROCESSED_TRANSACTION represents the processed transaction entity type.
	PROCESSED_TRANSACTION EntityType = "processed_transaction"
	// TOKEN represents the token entity type.
	TOKEN EntityType = "token"
)
