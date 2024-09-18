package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature/logs"
	"math/big"
)

// WithdrawNonFungible withdraws the given non-fungible token to the given address.
func (p *proxyEVM) WithdrawNonFungible(tokenChain datashared.TokenChain, from, to, txHash, tokenID, tokenURI string) (interface{}, error) {
	// Check if the hash isn't already used.
	if err := p.checkHash(txHash); err != nil {
		return nil, errors.Wrapf(err, "failed to check hash %s", txHash)
	}

	// Get the threshold of the bridge for the signer signatures.
	threshold, err := p.getThresholdSignerSignatures()
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
	case TokenERC721:
		tx, err = p.withdrawERC721(tokenChain, fromAddress, txHash, to, tokenID, tokenURI)
	case TokenERC1155:
		tx, err = p.withdrawERC1155NonFungible(tokenChain, fromAddress, txHash, to, tokenID, tokenURI)
	default:
		return nil, errors.New("unsupported type of token")
	}

	if tx == nil {
		return nil, nil
	}

	signNumber := int64(1)

	// Check if the number of signatures is greater or equal to the threshold.
	confirmed := signNumber >= int64(threshold)

	// If autoSend is enabled and the threshold is reached, send the transaction to the blockchain.
	if tokenChain.AutoSend && confirmed {
		return p.sendTransaction(tx, tokenChain.ChainID)
	}

	// Encode the transaction to be sent to the blockchain.
	return encodeTransaction(tx, fromAddress, p.chainID, tokenChain.ChainID, &confirmed), nil
}

// withdrawERC721 withdraws the given ERC721 token to the given address.
func (p *proxyEVM) withdrawERC721(tokenChain datashared.TokenChain, fromAddress common.Address, txHash, to, tokenID string, tokenURI string) (*types.Transaction, error) {
	// Convert the transaction hash to common.Hash type.
	transactionHash := common.HexToHash(txHash)

	// Convert the token ID to big.Int type.
	tokenIDBigInt, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, errors.New("failed to convert token ID")
	}

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if the token is mintable")
	}

	// Create a new ERC721 log object.
	log := logs.ERC721Log{
		TokenAddress: tokenChain.TokenAddress,
		TokenID:      tokenIDBigInt,
		To:           to,
		TxHash:       transactionHash,
		EventIndex:   defaultEventIndex,
		ChainID:      p.chainID,
		TokenURI:     tokenURI,
		IsMintable:   isMintable,
	}

	// Sign the log object.
	signature, err := p.signer.Sign(log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	// Create the transaction object for the withdrawal of the ERC721 token to be transferred to the destination chain.
	return p.bridge.WithdrawERC721(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		tokenIDBigInt,
		common.HexToAddress(to),
		transactionHash,
		big.NewInt(int64(defaultEventIndex)),
		tokenURI,
		log.IsMintable,
		[][]byte{signature},
	)
}

// withdrawERC1155NonFungible withdraws the given ERC1155 token to the given address.
func (p *proxyEVM) withdrawERC1155NonFungible(tokenChain datashared.TokenChain, fromAddress common.Address, txHash, to, tokenID string, tokenURI string) (*types.Transaction, error) {
	// Convert the transaction hash to common.Hash type.
	transactionHash := common.HexToHash(txHash)

	// Convert the token ID to big.Int type.
	tokenIDBigInt, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, errors.New("failed to convert token ID")
	}

	// Set the amount to 1 for non-fungible tokens.
	amount := big.NewInt(1)

	// Check if the token is mintable/burnable.
	isMintable, err := p.isMintable(tokenChain.BridgeType)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if the token is mintable")
	}

	// Create a new ERC1155 log object.
	log := logs.ERC1155Log{
		TokenAddress: tokenChain.TokenAddress,
		TokenID:      tokenIDBigInt,
		Amount:       amount,
		To:           to,
		TxHash:       transactionHash,
		EventIndex:   defaultEventIndex,
		ChainID:      p.chainID,
		TokenURI:     tokenURI,
		IsMintable:   isMintable,
	}

	// Sign the log object.
	signature, err := p.signer.Sign(log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	// Create the transaction object for the withdrawal of the ERC1155 token to be transferred to the destination chain.
	return p.bridge.WithdrawERC1155(
		buildTransactionOptions(fromAddress, nil),
		common.HexToAddress(tokenChain.TokenAddress),
		tokenIDBigInt,
		amount,
		common.HexToAddress(to),
		transactionHash,
		big.NewInt(int64(defaultEventIndex)),
		tokenURI,
		log.IsMintable,
		[][]byte{signature},
	)
}
