package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
	"math/big"
)

const (
	// depositERC20Event is the name of the ERC20 deposit event.
	depositERC20Event = "DepositedERC20"
	// depositERC721Event is the name of the ERC721 deposit event.
	depositERC721Event = "DepositedERC721"
	// depositERC1155Event is the name of the ERC1155 deposit event.
	depositERC1155Event = "DepositedERC1155"
)

// getTxReceipt returns the transaction receipt for the given transaction hash.
func (p *proxyEVM) getTxReceipt(txHash string) (*types.Receipt, error) {
	// Get the transaction receipt for the given transaction hash.
	receipt, err := p.client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		if err.Error() == "not found" {
			return &types.Receipt{}, bridgeErrors.ErrTxNotFound
		}
	}

	// Check if the transaction receipt status is successful.
	if receipt.Status != types.ReceiptStatusSuccessful {
		return receipt, bridgeErrors.ErrTxFailed
	}

	// Get the current block height.
	blockHeight, err := p.client.BlockNumber(context.Background())
	if err != nil {
		return receipt, errors.Wrap(err, "failed to get block number")
	}

	// Check if the transaction is confirmed – the block number of the transaction + the number of confirmations is less than the current block height.
	if (receipt.BlockNumber.Uint64() + uint64(p.confirmations-1)) > blockHeight {
		return nil, bridgeErrors.ErrTxNotConfirmed
	}

	return receipt, nil
}

// getBridgeEvent returns the bridge event from the given receipt by the given log name and event index.
func (p *proxyEVM) getBridgeEvent(dest interface{}, logName string, eventIndex int, receipt *types.Receipt) error {
	// Get the bridge ABI to unpack the log.
	abi, err := bridge.BridgeMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "failed to get bridge ABI")
	}

	// Create a new bound contract with the bridge ABI.
	contract := bind.NewBoundContract(common.Address{}, *abi, nil, nil, nil)

	// iterate over the logs in the receipt
	index := 0
	for _, log := range receipt.Logs {
		if log == nil {
			continue
		}

		// Unpack the log with the given log name.
		err = contract.UnpackLog(dest, logName, *log)
		if err != nil {
			// If the log cannot be unpacked, continue with the next log.
			if index == eventIndex {
				// If the event index is reached, return nil.
				return nil
			}

			index++
		}
	}

	// If the event index is not found, return an error.
	return bridgeErrors.ErrEventNotFound
}

// checkHash checks if the given transaction hash contains the given event index.
func (p *proxyEVM) checkHash(txHash string, eventIndex int) error {
	// Check if the transaction hash contains the given event index.
	containsHash, err := p.bridge.CheckHash(&bind.CallOpts{}, common.HexToHash(txHash), big.NewInt(int64(eventIndex)))
	if err != nil {
		return errors.Wrap(err, "failed to check hash")
	}

	// If the transaction hash contains the given event index, return an error.
	if containsHash {
		return bridgeErrors.ErrAlreadyWithdrawn
	}

	return nil
}
