package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature/logs"
	"math/big"
)

// WithdrawFungible withdraws the given amount of fungible token to the given address.
func (p *proxyEVM) WithdrawFungible(tokenChain datashared.TokenChain, from, to, txHash string, amount *big.Int) (interface{}, error) {
	// Check if the hash isn't already used.
	if err := p.checkHash(txHash); err != nil {
		return nil, errors.Wrapf(err, "failed to check hash %s", txHash)
	}

	// Get the threshold of the bridge for the oracle signatures.
	threshold, err := p.getThresholdOracleSignatures()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get threshold")
	}

	// Convert the from address to common.Address type.
	fromAddress := common.HexToAddress(from)

	// If autoSend is enabled, use the signer address as the sender.
	if tokenChain.AutoSend {
		fromAddress = p.signer.Address()
	}

	var tx *types.Transaction

	// Switch on the token type to create the transaction object for the withdrawal of the token to be transferred to the destination chain.
	switch tokenChain.TokenType {
	case TokenERC20:
		tx, err = p.withdrawERC20(tokenChain, fromAddress, txHash, to, amount)
	case TokenERC1155:
		tx, err = p.withdrawERC1155Fungible(tokenChain, fromAddress, txHash, to, amount)
	default:
		return nil, errors.New("unsupported type of token")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to withdraw token")
	}

	if tx == nil {
		return nil, nil
	}

	// Get the number of signatures for the transaction.
	signNumber, err := p.getSignatureNumber(tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get signature number")
	}

	// Check if the threshold is reached (signNumber is greater or equal to the threshold).
	confirmed := signNumber >= int64(threshold)

	// Send the transaction if the auto send is enabled and the threshold is reached.
	if tokenChain.AutoSend && confirmed {
		// Send the transaction to the network.
		return p.sendTransaction(tx, tokenChain.ChainID)
	}

	// Encode the transaction to be sent to the blockchain.
	return encodeTransaction(tx, fromAddress, p.chainID, tokenChain.ChainID, &confirmed), nil
}

// withdrawERC20 withdraws the given amount of ERC20 token to the given address.
func (p *proxyEVM) withdrawERC20(tokenChain datashared.TokenChain, fromAddress common.Address, txHash, to string, amount *big.Int) (*types.Transaction, error) {
	// Get the decimals of the ERC20 token.
	decimals, err := p.getDecimalsERC20(common.HexToAddress(tokenChain.TokenAddress))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get decimals of ERC20 token %s", tokenChain.TokenAddress)
	}

	// Convert the amount to the amount with decimals.
	amountWithDecimals, err := p.ConvertAmountWithDecimals(amount, int(decimals))
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert amount with decimals")
	}

	// Convert the transaction hash to common.Hash type.
	transactionHash := common.HexToHash(txHash)

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if token is mintable")
	}

	// Create the log object for the ERC20 token withdrawal.
	log := logs.ERC20Log{
		TokenAddress: tokenChain.TokenAddress,
		Amount:       amountWithDecimals,
		To:           to,
		TxHash:       transactionHash,
		EventIndex:   defaultEventIndex,
		ChainID:      p.chainID,
		IsMintable:   isMintable,
	}

	// Sign the log object to get the signature.
	signature, err := p.signer.Sign(log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	// Create the transaction object for the withdrawal of the ERC20 token to be transferred to the destination chain.
	return p.bridge.WithdrawERC20(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		amountWithDecimals,
		common.HexToAddress(to),
		transactionHash,
		big.NewInt(int64(defaultEventIndex)),
		log.IsMintable,
		[][]byte{signature},
	)
}

// withdrawERC1155Fungible withdraws the given amount of ERC1155 token to the given address.
func (p *proxyEVM) withdrawERC1155Fungible(tokenChain datashared.TokenChain, fromAddress common.Address, txHash, to string, amount *big.Int) (*types.Transaction, error) {
	// Convert the transaction hash to common.Hash type.
	transactionHash := common.HexToHash(txHash)

	// Convert the token ID to big.Int type.
	tokenID, ok := new(big.Int).SetString(tokenChain.TokenID, 10)
	if !ok {
		return nil, errors.New("failed to convert tokenID to big.Int")
	}

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if the token is mintable")
	}

	// Create the log object for the ERC1155 token withdrawal.
	log := logs.ERC1155Log{
		TokenAddress: tokenChain.TokenAddress,
		TokenID:      tokenID,
		Amount:       amount,
		To:           to,
		TxHash:       transactionHash,
		EventIndex:   defaultEventIndex,
		ChainID:      p.chainID,
		IsMintable:   isMintable,
	}

	// Sign the log object to get the signature.
	signature, err := p.signer.Sign(log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	// Create the transaction object for the withdrawal of the ERC1155 token to be transferred to the destination chain.
	return p.bridge.WithdrawERC1155(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		tokenID,
		amount,
		common.HexToAddress(to),
		transactionHash,
		big.NewInt(int64(defaultEventIndex)),
		"", // tokenUri is not used for fungible tokens
		log.IsMintable,
		[][]byte{signature},
	)
}
