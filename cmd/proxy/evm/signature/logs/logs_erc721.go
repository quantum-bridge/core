package logs

import (
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

// ERC721Log represents the log of the ERC721 token.
type ERC721Log struct {
	TokenAddress string
	TokenID      *big.Int
	To           string
	TxHash       common.Hash
	EventIndex   int
	ChainID      *big.Int
	TokenURI     string
	IsMintable   bool
}

// Hash returns the hash of the ERC721 log object.
func (l ERC721Log) Hash() []byte {
	return solsha3.SoliditySHA3(
		solsha3.Address(l.TokenAddress),
		solsha3.Uint256(l.TokenID),
		solsha3.Address(l.To),
		solsha3.Bytes32(l.TxHash.String()),
		solsha3.Uint256(big.NewInt(int64(l.EventIndex))),
		solsha3.Uint256(l.ChainID),
		solsha3.String(l.TokenURI),
		solsha3.Bool(l.IsMintable),
	)
}
