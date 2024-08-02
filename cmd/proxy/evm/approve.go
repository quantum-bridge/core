package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc1155"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc20"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc721"
	"math/big"
)

// Approve approves the given token for the bridge contract to transfer the tokens.
func (p *proxyEVM) Approve(tokenChain datashared.TokenChain, from string) (interface{}, error) {
	// Convert from address to common.Address type.
	fromAddress := common.HexToAddress(from)

	var tx *types.Transaction
	var err error

	// Switch on the token type to create the transaction object for the approval of the token to be transferred by the bridge contract.
	switch tokenChain.TokenType {
	case TokenERC20:
		tx, err = p.approveERC20(common.HexToAddress(tokenChain.TokenAddress), fromAddress)
	case TokenERC721:
		tx, err = p.approveERC721(common.HexToAddress(tokenChain.TokenAddress), fromAddress)
	case TokenERC1155:
		tx, err = p.approveERC1155(common.HexToAddress(tokenChain.TokenAddress), fromAddress)
	default:
		return nil, errors.New("unsupported type of token")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to approve token")
	}

	// If the transaction is nil (no approval needed), return nil.
	if tx == nil {
		return nil, nil
	}

	// Encode the transaction to be sent to the blockchain.
	encodedTx := encodeTransaction(tx, fromAddress, p.chainID, tokenChain.ChainID, nil)

	// Return the encoded transaction.
	return encodedTx, nil
}

// approveERC20 approves the given ERC20 token for the bridge contract.
func (p *proxyEVM) approveERC20(tokenAddress, fromAddress common.Address) (*types.Transaction, error) {
	// Create a new ERC20 contract instance.
	token, err := erc20.NewErc20(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ERC20 contract")
	}

	// Get the allowance of the token for the bridge contract.
	allowance, err := token.Allowance(&bind.CallOpts{}, fromAddress, p.bridgeContractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get allowance")
	}

	// If the allowance is greater than 0, return nil (no approval needed).
	if allowance.Cmp(big.NewInt(0)) == 1 {
		return nil, nil
	}

	// Create a new transaction to approve the token for the bridge contract with the maximum amount without executing the transaction in the blockchain.
	tx, err := token.Approve(buildTransactionOptions(fromAddress, nil), p.bridgeContractAddress, abi.MaxUint256)
	if err != nil {
		return nil, errors.Wrap(err, "failed to approve")
	}

	return tx, nil
}

// approveERC721 approves the given ERC721 token for the bridge contract.
func (p *proxyEVM) approveERC721(tokenAddress, from common.Address) (*types.Transaction, error) {
	// Create a new ERC721 contract instance.
	token, err := erc721.NewErc721(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ERC721 contract")
	}

	// Check if the token is already approved for the bridge contract.
	approved, err := token.IsApprovedForAll(nil, from, p.bridgeContractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check approval")
	}

	// If the token is already approved, return nil (no approval needed).
	if approved {
		return nil, nil
	}

	// Create a new transaction to approve the token for the bridge contract with the maximum amount without executing the transaction in the blockchain.
	tx, err := token.SetApprovalForAll(buildTransactionOptions(from, nil), p.bridgeContractAddress, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to approve")
	}

	return tx, nil
}

// approveERC1155 approves the given ERC1155 token for the bridge contract.
func (p *proxyEVM) approveERC1155(tokenAddress, from common.Address) (*types.Transaction, error) {
	// Create a new ERC1155 contract instance.
	token, err := erc1155.NewErc1155(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ERC1155 contract")
	}

	// Check if the token is already approved for the bridge contract.
	approved, err := token.IsApprovedForAll(nil, from, p.bridgeContractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check approval")
	}

	// If the token is already approved, return nil (no approval needed).
	if approved {
		return nil, nil
	}

	// Create a new transaction to approve the token for the bridge contract with the maximum amount without executing the transaction in the blockchain.
	tx, err := token.SetApprovalForAll(buildTransactionOptions(from, nil), p.bridgeContractAddress, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to approve")
	}

	return tx, nil
}
