package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
)

// CheckNonFungible checks the status of the given non-fungible token transfer.
func (p *proxyEVM) CheckNonFungible(txHash string, tokenChain datashared.TokenChain) (*datashared.NonFungibleLock, error) {
	// Get the transaction receipt for the given transaction hash.
	to, err := p.getTxReceipt(txHash)
	if err != nil {
		return &datashared.NonFungibleLock{}, err
	}

	// Switch on the token type to check the status of the given non-fungible token transfer.
	switch tokenChain.TokenType {
	case TokenERC721:
		return p.checkNonFungibleERC721(to, tokenChain)
	case TokenERC1155:
		return p.checkNonFungibleERC1155(to, tokenChain)
	default:
		return &datashared.NonFungibleLock{}, errors.New("unsupported type of token")
	}
}

// checkNonFungibleERC721 checks the status of the given ERC721 token transfer.
func (p *proxyEVM) checkNonFungibleERC721(to *types.Receipt, tokenChain datashared.TokenChain) (*datashared.NonFungibleLock, error) {
	// Create a new BridgeDepositedERC721 log object.
	log := bridge.BridgeDepositedERC721{}

	// Get the bridge event log for the given event index.
	err := p.getBridgeEvent(&log, depositERC721Event, to)
	if err != nil {
		return &datashared.NonFungibleLock{}, err
	}

	// Check if the token address matches the token address of the token chain.
	if !addressesEqual(log.Token, common.HexToAddress(tokenChain.TokenAddress)) {
		return &datashared.NonFungibleLock{}, errors.New("token address does not match")
	}

	// Check if the token is mintable and the bridge type matches the bridge type of the token chain.
	if log.IsMintable && tokenChain.BridgeType != data.BridgeTypeMintable {
		return &datashared.NonFungibleLock{}, bridgeErrors.ErrWrongLockEvent
	}

	// Return the non-fungible lock object.
	return &datashared.NonFungibleLock{
		To:      log.To.Hex(),
		TokenID: log.TokenId.String(),
		Network: log.Network,
	}, nil
}

// checkNonFungibleERC1155 checks the status of the given ERC1155 token transfer.
func (p *proxyEVM) checkNonFungibleERC1155(to *types.Receipt, tokenChain datashared.TokenChain) (*datashared.NonFungibleLock, error) {
	// Create a new BridgeDepositedERC1155 log object.
	log := bridge.BridgeDepositedERC1155{}

	// Get the bridge event log for the given event index.
	err := p.getBridgeEvent(&log, depositERC1155Event, to)
	if err != nil {
		return &datashared.NonFungibleLock{}, err
	}

	// Check if the token address matches the token address of the token chain.
	if !addressesEqual(log.Token, common.HexToAddress(tokenChain.TokenAddress)) {
		return &datashared.NonFungibleLock{}, errors.New("token address does not match")
	}

	// Check if the token is mintable and the bridge type matches the bridge type of the token chain.
	if log.IsMintable && tokenChain.BridgeType != data.BridgeTypeMintable {
		return &datashared.NonFungibleLock{}, bridgeErrors.ErrWrongLockEvent
	}

	// Check if the amount is 1. If not, return an error. ERC1155 non-fungible tokens are always locked with the amount of 1.
	if log.Amount.Uint64() != 1 {
		return &datashared.NonFungibleLock{}, bridgeErrors.ErrWrongLockEvent
	}

	// Return the non-fungible lock object.
	return &datashared.NonFungibleLock{
		To:      log.To,
		TokenID: log.TokenId.String(),
		Network: log.Network,
	}, nil
}
