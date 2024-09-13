package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc1155"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc20"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc721"
	"math/big"
)

// BridgeBalance returns the balance of the given token for the bridge contract.
func (p *proxyEVM) BridgeBalance(tokenChain datashared.TokenChain, tokenId *string) (*big.Int, error) {
	return p.Balance(tokenChain, p.bridgeAddress.Hex(), tokenId)
}

// Balance returns the balance of the given token for the given address.
func (p *proxyEVM) Balance(tokenChain datashared.TokenChain, address string, tokenId *string) (*big.Int, error) {
	// Switch on the token type to get the balance of the token for the given address.
	switch tokenChain.TokenType {
	case TokenERC20:
		return p.balanceERC20(tokenChain, address)
	case TokenERC721:
		return p.balanceERC721(tokenChain, address, *tokenId)
	case TokenERC1155:
		return p.balanceERC1155(tokenChain, address, *tokenId)
	default:
		return &big.Int{}, errors.New("unsupported type of token")
	}
}

// balanceERC20 returns the balance of the given ERC20 token for the given address.
func (p *proxyEVM) balanceERC20(tokenChain datashared.TokenChain, address string) (*big.Int, error) {
	// Create a new ERC20 contract instance.
	token, err := erc20.NewErc20(common.HexToAddress(tokenChain.TokenAddress), p.client)
	if err != nil {
		return &big.Int{}, errors.Wrapf(err, "failed to create ERC20 contract for token %s", tokenChain.TokenAddress)
	}

	// Get the balance of the token for the given address.
	balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return &big.Int{}, errors.Wrapf(err, "failed to get balance for address %s", address)
	}

	// Get the decimals of the token.
	decimals, err := p.getDecimalsERC20(common.HexToAddress(tokenChain.TokenAddress))
	if err != nil {
		return &big.Int{}, errors.Wrap(err, "failed to get decimals for ERC20 token")
	}

	// Convert the balance to the amount with decimals.
	return p.ConvertAmountWithDecimals(balance, int(decimals))
}

// balanceERC721 returns the balance of the given ERC721 token for the given address.
func (p *proxyEVM) balanceERC721(tokenChain datashared.TokenChain, address string, tokenId string) (*big.Int, error) {
	// Create a new ERC721 contract instance.
	token, err := erc721.NewErc721(common.HexToAddress(tokenChain.TokenAddress), p.client)
	if err != nil {
		return &big.Int{}, errors.Wrapf(err, "failed to create ERC721 contract for token %s", tokenChain.TokenAddress)
	}

	// Convert the token ID to big int.
	tokenID, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return &big.Int{}, errors.New("failed to convert token ID to big int")
	}

	// Get the owner of the token with the given token ID.
	owner, err := token.OwnerOf(&bind.CallOpts{}, tokenID)
	if err != nil {
		return &big.Int{}, errors.Wrapf(err, "failed to get owner of token %s", tokenId)
	}

	// If the owner is the same as the given address, return 1 (balance is 1).
	if !addressesEqual(owner, common.HexToAddress(address)) {
		return big.NewInt(1), nil
	}

	return big.NewInt(0), nil
}

// balanceERC1155 returns the balance of the given ERC1155 token for the given address.
func (p *proxyEVM) balanceERC1155(tokenChain datashared.TokenChain, address string, tokenId string) (*big.Int, error) {
	// Create a new ERC1155 contract instance.
	token, err := erc1155.NewErc1155(common.HexToAddress(tokenChain.TokenAddress), p.client)
	if err != nil {
		return &big.Int{}, errors.Wrapf(err, "failed to create ERC1155 contract for token %s", tokenChain.TokenAddress)
	}

	// Convert the token ID to big int.
	tokenID, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return &big.Int{}, errors.New("failed to convert token ID to big int")
	}

	// Get the balance of the token for the given address and token ID.
	balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), tokenID)
	if err != nil {
		return &big.Int{}, errors.Wrapf(err, "failed to get balance for address %s and token %s", address, tokenId)
	}

	// If the balance is greater than 0, return 1 (balance is 1).
	if balance.Cmp(big.NewInt(0)) == 1 {
		return big.NewInt(1), nil
	}

	return balance, nil
}
