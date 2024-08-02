package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/ipfs"
	proxybridge "github.com/quantum-bridge/core/cmd/proxy/bridge"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/bridge"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature"
	"math/big"
)

const (
	// TokenERC20 is the name of the ERC20 token type.
	TokenERC20 = "erc20"
	// TokenERC721 is the name of the ERC721 token type.
	TokenERC721 = "erc721"
	// TokenERC1155 is the name of the ERC1155 token type.
	TokenERC1155 = "erc1155"
)

// proxyEVM is the EVM proxy implementation.
type proxyEVM struct {
	client                *ethclient.Client
	signer                signature.Signer
	chainID               *big.Int
	bridgeContractAddress common.Address
	bridge                *bridge.Bridge
	ipfs                  ipfs.IPFS
	confirmations         int64
}

// NewProxy creates a new EVM proxy with the given RPC URL, signer, bridge contract address, IPFS client and number of confirmations.
func NewProxy(rpc string, signer signature.Signer, bridgeContractAddress string, ipfs ipfs.IPFS, confirmations int64) (proxybridge.Bridge, error) {
	// Dial the Ethereum client with the given RPC URL.
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial ethereum client")
	}

	// Get the chain ID of the Ethereum client.
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	// Create a new bridge contract instance.
	bridgeInstance, err := bridge.NewBridge(common.HexToAddress(bridgeContractAddress), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bridge contract")
	}

	// Return the new EVM proxy.
	return &proxyEVM{
		client:                client,
		signer:                signer,
		chainID:               chainID,
		bridgeContractAddress: common.HexToAddress(bridgeContractAddress),
		bridge:                bridgeInstance,
		ipfs:                  ipfs,
		confirmations:         confirmations,
	}, nil
}
