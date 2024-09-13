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

// CheckFungible checks the status of the given fungible token transfer.
func (p *proxyEVM) CheckFungible(txHash string, tokenChain datashared.TokenChain) (*datashared.FungibleLock, error) {
	// Get the transaction receipt for the given transaction hash.
	to, err := p.getTxReceipt(txHash)
	if err != nil {
		return &datashared.FungibleLock{}, err
	}

	// Switch on the token type to check the status of the given fungible token transfer.
	switch tokenChain.TokenType {
	case TokenERC20:
		return p.checkFungibleERC20(to, tokenChain)
	case TokenERC1155:
		return p.checkFungibleERC1155(to, tokenChain)
	default:
		return &datashared.FungibleLock{}, errors.New("unsupported type of token")
	}
}

// checkFungibleERC20 checks the status of the given ERC20 token transfer.
func (p *proxyEVM) checkFungibleERC20(to *types.Receipt, tokenChain datashared.TokenChain) (*datashared.FungibleLock, error) {
	// Create a new BridgeDepositedERC20 log object.
	log := bridge.BridgeDepositedERC20{}

	// Get the bridge event log for the given event index.
	err := p.getBridgeEvent(&log, depositERC20Event, to)
	if err != nil {
		return &datashared.FungibleLock{}, err
	}

	// Check if the token address matches the token address of the token chain.
	if !addressesEqual(log.Token, common.HexToAddress(tokenChain.TokenAddress)) {
		return &datashared.FungibleLock{}, errors.New("token address does not match")
	}

	// Check if the token is mintable and the bridge type matches the bridge type of the token chain.
	if log.IsMintable && tokenChain.BridgeType != data.BridgeTypeMintable {
		return &datashared.FungibleLock{}, bridgeErrors.ErrWrongLockEvent
	}

	// Get the decimals of the token to convert the amount with decimals.
	decimals, err := p.getDecimalsERC20(common.HexToAddress(tokenChain.TokenAddress))
	if err != nil {
		return &datashared.FungibleLock{}, errors.Wrap(err, "failed to get decimals for ERC20 token")
	}

	// Convert the amount with decimals.
	amount, err := p.ConvertAmountWithDecimals(log.Amount, int(decimals))
	if err != nil {
		return &datashared.FungibleLock{}, errors.Wrap(err, "failed to convert amount with decimals")
	}

	// Return the fungible lock object.
	return &datashared.FungibleLock{
		To:      log.To,
		Amount:  amount,
		Network: log.Network,
	}, nil
}

// checkFungibleERC1155 checks the status of the given ERC1155 token transfer.
func (p *proxyEVM) checkFungibleERC1155(to *types.Receipt, tokenChain datashared.TokenChain) (*datashared.FungibleLock, error) {
	// Create a new BridgeDepositedERC1155 log object to get the event log.
	log := bridge.BridgeDepositedERC1155{}

	// Get the bridge event log for the given event index.
	err := p.getBridgeEvent(&log, depositERC1155Event, to)
	if err != nil {
		return &datashared.FungibleLock{}, err
	}

	// Check if the token address matches the token address of the token chain.
	if !addressesEqual(log.Token, common.HexToAddress(tokenChain.TokenAddress)) {
		return &datashared.FungibleLock{}, errors.New("token address does not match")
	}

	// Check if the token is mintable and the bridge type matches the bridge type of the token chain.
	if log.IsMintable && tokenChain.BridgeType != data.BridgeTypeMintable {
		return &datashared.FungibleLock{}, bridgeErrors.ErrWrongLockEvent
	}

	// Check if the amount is not equal to 1.
	return &datashared.FungibleLock{
		To:      log.To,
		Amount:  log.Amount,
		Network: log.Network,
	}, nil
}
