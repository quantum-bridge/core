package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"math/big"
)

// LockFungible locks the given fungible token for the bridge contract to transfer the tokens.
func (p *proxyEVM) LockFungible(tokenChain datashared.TokenChain, from, to, network, amount string) (interface{}, error) {
	// Convert from address to common.Address type.
	fromAddress := common.HexToAddress(from)

	var tx *types.Transaction
	var err error

	// Switch on the token type to create the transaction object for the lock of the token to be transferred to the destination chain.
	switch tokenChain.TokenType {
	case TokenNative:
		tx, err = p.lockNative(fromAddress, to, network, amount)
	case TokenERC20:
		tx, err = p.lockERC20(tokenChain, fromAddress, to, network, amount)
	case TokenERC1155:
		tx, err = p.lockERC1155Fungible(tokenChain, fromAddress, to, network, amount)
	default:
		return nil, errors.New("unsupported type of token")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to lock token")
	}

	// If the transaction is nil (no lock needed), return nil.
	if tx == nil {
		return nil, nil
	}

	// Encode the transaction to be sent to the blockchain.
	encodedTx := encodeTransaction(tx, fromAddress, p.chainID, network, nil)

	return encodedTx, nil
}

// lockNative locks the native token to be transferred to the destination chain.
func (p *proxyEVM) lockNative(fromAddress common.Address, to, network, amount string) (*types.Transaction, error) {
	// Convert the amount to big.Int type.
	amountBigInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("failed to convert amount to big.Int")
	}

	// Convert the destination address to common.Address type.
	toAddress := common.HexToAddress(to)

	// Create the transaction object for the lock of the native token to be transferred to the destination chain.
	tx, err := p.bridge.DepositNative(
		buildTransactionOptions(fromAddress, amountBigInt),
		network,
		toAddress,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deposit native token")
	}

	return tx, nil
}

// lockERC20 locks the given ERC20 token to be transferred to the destination chain.
func (p *proxyEVM) lockERC20(tokenChain datashared.TokenChain, fromAddress common.Address, to, network, amount string) (*types.Transaction, error) {
	// Get the decimals of the ERC20 token.
	decimals, err := p.getDecimalsERC20(common.HexToAddress(tokenChain.TokenAddress))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get decimals of ERC20 token %s", tokenChain.TokenAddress)
	}

	// Convert the amount to big.Int type.
	amountBigInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("failed to convert amount to big.Int")
	}

	// Convert the amount to the amount with decimals.
	amountWithDecimals, err := p.ConvertAmountWithDecimals(amountBigInt, int(decimals))
	if err != nil {
		return nil, err
	}

	// Check if the token is mintable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if token is mintable")
	}

	// Create the transaction object for the lock of the ERC20 token to be transferred to the destination chain.
	tx, err := p.bridge.DepositERC20(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		amountWithDecimals,
		to,
		network,
		isMintable,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deposit ERC20 token")
	}

	return tx, nil
}

// lockERC1155Fungible locks the given ERC1155 token to be transferred to the destination chain.
func (p *proxyEVM) lockERC1155Fungible(tokenChain datashared.TokenChain, fromAddress common.Address, to, network, amount string) (*types.Transaction, error) {
	// Convert the amount to big.Int type.
	amountBigInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("failed to convert amount to big.Int")
	}

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)

	// Convert the token ID to big.Int type.
	tokenID, ok := new(big.Int).SetString(tokenChain.TokenID, 10)
	if !ok {
		return nil, errors.New("failed to convert tokenID to big.Int")
	}

	// Create the transaction object for the lock of the ERC1155 token to be transferred to the destination chain.
	tx, err := p.bridge.DepositERC1155(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		tokenID,
		amountBigInt,
		to,
		network,
		isMintable,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deposit ERC1155 token")
	}

	return tx, nil
}
