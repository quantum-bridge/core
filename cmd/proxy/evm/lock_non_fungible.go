package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"math/big"
)

// LockNonFungible locks the given non-fungible token for the bridge contract to transfer the tokens.
func (p *proxyEVM) LockNonFungible(tokenChain datashared.TokenChain, from, to, network, tokenId string) (interface{}, error) {
	// Convert from address to common.Address type.
	fromAddress := common.HexToAddress(from)

	var tx *types.Transaction
	var err error

	// Switch on the token type to create the transaction object for the lock of the token to be transferred to the destination chain.
	switch tokenChain.TokenType {
	case TokenERC721:
		tx, err = p.lockERC721(tokenChain, fromAddress, common.HexToAddress(to), network, tokenId)
	case TokenERC1155:
		tx, err = p.lockERC1155NonFungible(tokenChain, fromAddress, common.HexToAddress(to), network, tokenId)
	default:
		return nil, errors.New("unsupported type of token")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to lock non-fungible token")
	}

	if tx == nil {
		return nil, nil
	}

	// Encode the transaction to be sent to the blockchain.
	encodedTx := encodeTransaction(tx, fromAddress, p.chainID, tokenChain.ChainID, nil)

	return encodedTx, nil
}

// lockERC721 locks the given ERC721 token to be transferred to the destination chain.
func (p *proxyEVM) lockERC721(tokenChain datashared.TokenChain, fromAddress, toAddress common.Address, network, tokenId string) (*types.Transaction, error) {
	// Convert the token ID to big.Int type.
	tokenIDBigInt, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return nil, errors.New("failed to convert token ID to big int")
	}

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if token is mintable")
	}

	// Create a new transaction object for the lock of the ERC721 token to be transferred to the destination chain.
	tx, err := p.bridge.DepositERC721(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		tokenIDBigInt,
		toAddress,
		network,
		isMintable,
	)

	return tx, nil
}

// lockERC1155NonFungible locks the given ERC1155 token to be transferred to the destination chain.
func (p *proxyEVM) lockERC1155NonFungible(tokenChain datashared.TokenChain, fromAddress, toAddress common.Address, network, tokenId string) (*types.Transaction, error) {
	// Convert the token ID to big.Int type.
	tokenIDBigInt, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return nil, errors.New("failed to convert token ID to big int")
	}

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if token is mintable")
	}

	// Create a new transaction object for the lock of the ERC1155 token to be transferred to the destination chain.
	tx, err := p.bridge.DepositERC1155(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		tokenIDBigInt,
		big.NewInt(1),
		toAddress.Hex(),
		network,
		isMintable,
	)

	return tx, nil
}
