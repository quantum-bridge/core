package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc20"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
	"math/big"
	"reflect"
	"strings"
)

const (
	// gasLimit represents the gas limit for the transaction.
	gasLimit = 300000
	// defaultPrecision represents the default precision for the token.
	defaultDecimals = 18
)

// buildTransactionOptions creates a new transaction options with the given value.
func buildTransactionOptions(fromAddress common.Address, value *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     fromAddress,
		Signer:   manualSigner,
		Value:    value,
		GasLimit: gasLimit,
		NoSend:   true,
	}
}

// manualSigner signs the transaction manually. This is used to bypass the automatic signer.
func manualSigner(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return tx, nil
}

// encodeTransaction encodes the transaction to the shared format for the API.
func encodeTransaction(tx *types.Transaction, from common.Address, chainID *big.Int, chain string, confirmed *bool) interface{} {
	return datashared.EVMTransaction{
		Key: datashared.Key{
			ID:   tx.Hash().Hex(),
			Type: datashared.EVM_TRANSACTION,
		},
		Attributes: datashared.EVMAttributes{
			Confirmed: confirmed,
			TxBody: datashared.EVMTxBody{
				ChainID: fmt.Sprintf("0x%x", chainID),
				From:    from.Hex(),
				To:      tx.To().Hex(),
				Value:   tx.Value().String(),
				Data:    hexutil.Encode(tx.Data()),
			},
		},
		Relationships: datashared.Relationships{
			Chain: datashared.ChainEntity{
				Data: datashared.DataEntity{
					ID:   chain,
					Type: datashared.CHAIN,
				},
			},
		},
	}
}

// encodeProcessedTransaction encodes the processed transaction to the shared format for the API.
func encodeProcessedTransaction(tx *types.Transaction, from, chain string, confirmed *bool) interface{} {
	return datashared.EVMTransaction{
		Key: datashared.Key{
			ID:   tx.Hash().Hex(),
			Type: datashared.PROCESSED_TRANSACTION,
		},
		Attributes: datashared.EVMAttributes{
			Confirmed: confirmed,
			TxBody: datashared.EVMTxBody{
				ChainID: tx.ChainId().String(),
				From:    from,
				To:      tx.To().Hex(),
				Value:   tx.Value().String(),
				Data:    string(tx.Data()),
			},
		},
		Relationships: datashared.Relationships{
			Chain: datashared.ChainEntity{
				Data: datashared.DataEntity{
					ID:   chain,
					Type: datashared.CHAIN,
				},
			},
		},
	}
}

// getDecimalsERC20 returns the decimals of the given ERC20 token.
func (p *proxyEVM) getDecimalsERC20(tokenAddress common.Address) (uint8, error) {
	token, err := erc20.NewErc20(tokenAddress, p.client)
	if err != nil {
		return 0, err
	}

	decimals, err := token.Decimals(nil)
	if err != nil {
		return 0, err
	}

	return decimals, nil
}

// ConvertAmountWithDecimals converts the given amount to the given precision.
func (p *proxyEVM) ConvertAmountWithDecimals(amount *big.Int, decimals int) (*big.Int, error) {
	if decimals == 0 {
		return big.NewInt(1), nil
	}

	if decimals > defaultDecimals {
		return nil, errors.New("precision is greater than the default precision")
	}

	precisionDiff := defaultDecimals - decimals

	return big.NewInt(0).Div(big.NewInt(0).Set(amount), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precisionDiff)), nil)), nil
}

// isMintable checks if the given bridging type is mintable.
func (p *proxyEVM) isMintable(bridgingType data.BridgeType) (bool, error) {
	switch bridgingType {
	case data.BridgeTypeMintable:
		return true, nil
	case data.BridgeTypeLP:
		return false, nil
	default:
		return false, errors.New("unsupported bridging type")
	}
}

// addressesEqual compares the given addresses and returns true if they are equal.
func addressesEqual(a, b common.Address) bool {
	return strings.ToLower(a.String()) == strings.ToLower(b.String())
}

// sendTransaction sends the given transaction to the network.
func (p *proxyEVM) sendTransaction(tx *types.Transaction, chain string) (interface{}, error) {
	// Sign the transaction. If the transaction is already signed, the signature will be overwritten.
	tx, err := p.signer.SignTransaction(tx, p.chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction")
	}

	// Send the transaction to the network and get the transaction receipt.
	err = p.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send transaction")
	}

	// Wait for the transaction to be mined and get the transaction receipt.
	receipt, err := p.waitTransactionConfirmation(context.Background(), p.client, tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to wait for transaction")
	}

	// Check if the transaction is confirmed.
	confirmed := receipt.Status == types.ReceiptStatusSuccessful

	// Encode the processed transaction to be sent to the blockchain.
	return encodeProcessedTransaction(tx, p.signer.Address().Hex(), chain, &confirmed), nil
}

// waitTransactionConfirmation waits for the given transaction to be mined.
func (p *proxyEVM) waitTransactionConfirmation(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	// Wait for the transaction to be mined and get the transaction receipt.
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to wait for transaction")
	}

	// Check if the transaction failed.
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, errors.New("transaction failed")
	}

	return receipt, nil
}

// addSignature adds the given signature to the transaction.
func (p *proxyEVM) addSignature(transactionOptions *bind.TransactOpts, transaction *types.Transaction, rawTransactionData []byte) (*types.Transaction, int64, error) {
	// Get the bridge ABI.
	bridgeABI, err := bridge.BridgeMetaData.GetAbi()
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get bridge ABI")
	}

	// Decode the transaction parameters.
	transactionParams, transactionMethod, err := decodeTransactionParams(*bridgeABI, transaction.Data())
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to decode transaction params")
	}

	// Decode the old transaction parameters.
	oldTransactionParams, oldTransactionMethod, err := decodeTransactionParams(*bridgeABI, rawTransactionData)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to decode transaction params")
	}

	// Check if the number of parameters isn't equal to the number of old parameters - return an error.
	if len(oldTransactionParams) != len(transactionParams) {
		return nil, 0, bridgeErrors.ErrWrongSignedTx
	}

	// Check if the method name isn't equal to the old method name - return an error.
	if transactionMethod.Name != oldTransactionMethod.Name {
		return nil, 0, bridgeErrors.ErrWrongSignedTx
	}

	// Check if params except signature is equal.
	if !reflect.DeepEqual(oldTransactionParams[:len(oldTransactionParams)-1], transactionParams[:len(transactionParams)-1]) {
		return nil, 0, bridgeErrors.ErrWrongSignedTx
	}

	existingSignatures := oldTransactionParams[len(oldTransactionParams)-1].([][]byte)
	newSignature := transactionParams[len(transactionParams)-1].([][]byte)[0]

	// Check if the signature is already added to the transaction.
	for _, existingSignature := range existingSignatures {
		if reflect.DeepEqual(existingSignature, newSignature) {
			return nil, int64(len(existingSignatures)), errors.New("double signature")
		}
	}

	// Add the new signature to the existing signatures.
	transactionParams[len(transactionParams)-1] = append(existingSignatures, newSignature)

	// Check if signatures are added to the transaction parameters. If not, return an error.
	if signatures, ok := transactionParams[len(transactionParams)-1].([][]byte); ok {
		// Get the bridge contract address.
		contract := bind.NewBoundContract(p.bridgeContractAddress, *bridgeABI, nil, p.client, nil)
		// Create a new transaction with the transaction options and the transaction method name and parameters.
		newTransaction, err := contract.Transact(transactionOptions, transactionMethod.Name, transactionParams...)

		// Return the new transaction, the number of signatures and an error.
		return newTransaction, int64(len(signatures)), err
	}

	return nil, 0, errors.New("failed to add signature")
}

// getThresholdOracleSignatures returns the threshold of the bridge.
func (p *proxyEVM) getThresholdOracleSignatures() (uint64, error) {
	// Get the threshold of the bridge for the oracle signatures.
	threshold, err := p.bridge.ThresholdOracleSignatures(&bind.CallOpts{})
	if err != nil {
		return 0, errors.Wrap(err, "failed to get threshold")
	}

	// Return the threshold as an uint64.
	return threshold.Uint64(), nil
}

// decodeTransactionParams decodes the transaction parameters using the given ABI.
func decodeTransactionParams(abi abi.ABI, data []byte) ([]interface{}, *abi.Method, error) {
	// Get the method by the first 4 bytes of the data to get the method name.
	method, err := abi.MethodById(data[:4])
	if err != nil {
		return nil, nil, err
	}

	// Unpack the values of the method using the data without the first 4 bytes to get the parameters.
	result, err := method.Inputs.UnpackValues(data[4:])
	if err != nil {
		return nil, nil, err
	}

	return result, method, nil
}
