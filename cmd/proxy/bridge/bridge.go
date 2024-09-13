package bridge

import (
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"math/big"
)

// Bridge is the interface for the proxy of the bridge contract that realizes methods for cross-chain token transfers, withdrawals, approvals, and balance checks.
type Bridge interface {
	// Approve is a method that approves the token to be transferred by the bridge contract.
	Approve(tokenChain datashared.TokenChain, from string) (interface{}, error)
	// LockFungible is a method that locks the given amount of tokens to be transferred to the destination chain.
	LockFungible(tokenChain datashared.TokenChain, from, to, network, amount string) (interface{}, error)
	// LockNonFungible is a method that locks the given non-fungible token to be transferred to the destination chain.
	LockNonFungible(tokenChain datashared.TokenChain, from, to, network, tokenID string) (interface{}, error)
	// CheckFungible is a method that checks the status of the given fungible token transfer.
	CheckFungible(txHash string, tokenChain datashared.TokenChain) (*datashared.FungibleLock, error)
	// CheckNonFungible is a method that checks the status of the given non-fungible token transfer.
	CheckNonFungible(txHash string, tokenChain datashared.TokenChain) (*datashared.NonFungibleLock, error)
	// Balance is a method that returns the balance of the given token for the given address.
	Balance(tokenChain datashared.TokenChain, address string, tokenID *string) (*big.Int, error)
	// BridgeBalance is a method that returns the balance of the given token for the bridge contract.
	BridgeBalance(tokenChain datashared.TokenChain, tokenID *string) (*big.Int, error)
	// WithdrawFungible is a method that withdraws the given amount of tokens from the bridge contract.
	WithdrawFungible(tokenChain datashared.TokenChain, from, to, txHash string, amount *big.Int) (interface{}, error)
	// WithdrawNonFungible is a method that withdraws the given non-fungible token from the bridge contract.
	WithdrawNonFungible(tokenChain datashared.TokenChain, from, to, txHash, tokenID, tokenURI string) (interface{}, error)
	// GetNFTMetadataURI is a method that returns the metadata URI of the given non-fungible token.
	GetNFTMetadataURI(tokenChain datashared.TokenChain, tokenID string) (string, error)
	// GetNFTMetadata is a method that returns the metadata of the given non-fungible token.
	GetNFTMetadata(tokenChain datashared.TokenChain, tokenID string) (*datashared.NFTMetadata, error)
}
