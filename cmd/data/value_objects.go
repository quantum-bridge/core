package data

// ChainType is the type of the chain.
type ChainType string

// BridgeType is the type of the bridge that is used in the token chain.
type BridgeType string

// TokenType is the type of the token that is used in the bridge.
type TokenType string

const (
	// EVM represents the Ethereum Virtual Machine chain type (EVM) that is used in the bridge.
	EVM ChainType = "evm"

	// BridgeTypeLP represents the liquidity pool bridge type that is used in the bridge.
	BridgeTypeLP BridgeType = "liquidity_pool"
	// BridgeTypeMintable represents the mintable/burnable bridge type that is used in the bridge.
	BridgeTypeMintable BridgeType = "mintable"

	// FUNGIBLE represents the fungible token type.
	FUNGIBLE TokenType = "fungible"
	// NONFUNGIBLE represents the non-fungible token type.
	NONFUNGIBLE TokenType = "non-fungible"
)

func (t *ChainType) String() string {
	return string(*t)
}

// String returns the string representation of the bridge type.
func (t *BridgeType) String() string {
	return string(*t)
}

// String returns the string representation of the token type.
func (t *TokenType) String() string {
	return string(*t)
}
