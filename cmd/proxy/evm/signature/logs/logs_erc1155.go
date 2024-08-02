package logs

import (
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

// ERC1155Log represents the log of the ERC1155 token.
type ERC1155Log struct {
	TokenAddress string
	TokenID      *big.Int
	Amount       *big.Int
	To           string
	TxHash       common.Hash
	EventIndex   int
	ChainID      *big.Int
	TokenURI     string
	IsMintable   bool
}

// Hash returns the hash of the ERC1155 log object.
func (l ERC1155Log) Hash() []byte {
	return solsha3.SoliditySHA3(
		solsha3.Address(l.TokenAddress),
		solsha3.Uint256(l.TokenID),
		solsha3.Uint256(l.Amount),
		solsha3.Address(l.To),
		solsha3.Bytes32(l.TxHash.String()),
		solsha3.Uint256(big.NewInt(int64(l.EventIndex))),
		solsha3.Uint256(l.ChainID),
		solsha3.String(l.TokenURI),
		solsha3.Bool(l.IsMintable),
	)
}
