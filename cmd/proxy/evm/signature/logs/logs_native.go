package logs

import (
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

// NativeLog represents the log of the native token.
type NativeLog struct {
	Amount     *big.Int
	To         string
	TxHash     common.Hash
	EventIndex int
	ChainID    *big.Int
}

// Hash returns the hash of the native log object.
func (log NativeLog) Hash() []byte {
	return solsha3.SoliditySHA3(
		solsha3.Uint256(log.Amount),
		solsha3.Address(common.HexToAddress(log.To)),
		solsha3.Bytes32(log.TxHash.String()),
		solsha3.Uint256(big.NewInt(int64(log.EventIndex))),
		solsha3.Uint256(log.ChainID),
	)
}
