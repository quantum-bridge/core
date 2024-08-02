package signature

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature/logs"
	"math/big"
)

// Signer represents the interface for the signer.
type Signer interface {
	// Sign signs the given data and returns the signature.
	Sign(logs.SHA3) ([]byte, error)
	// SignTransaction signs the given transaction and returns the signed transaction.
	SignTransaction(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	// Address returns the address of the signer.
	Address() common.Address
	// PublicKey returns the public key of the signer.
	PublicKey() *ecdsa.PublicKey
}

// signer represents the signer.
type signer struct {
	privKey *ecdsa.PrivateKey
	pubKey  *ecdsa.PublicKey
	address common.Address
}

// NewSigner creates a new signer with the given private key.
func NewSigner(privKey *ecdsa.PrivateKey) Signer {
	// Convert the private key to the public key.
	pubKey, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("cannot convert private key to public key for signer")
	}

	// Create a new signer with the private key, public key, and address.
	return &signer{
		privKey: privKey,
		pubKey:  pubKey,
		address: crypto.PubkeyToAddress(*pubKey),
	}
}

// Sign signs the given data and returns the signature.
func (s *signer) Sign(data logs.SHA3) ([]byte, error) {
	// Sign the hash of the data and return the signature.
	signature, err := crypto.Sign(signHash(data.Hash()), s.privKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign data")
	}

	// Add 27 to the last byte of the signature to make it compatible with EVM signatures.
	signature[64] += 27

	return signature, nil
}

// SignTransaction signs the given transaction and returns the signed transaction.
func (s *signer) SignTransaction(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	// Create a new Cancun signer with the chain ID.
	txSigner := types.NewCancunSigner(chainID)

	// Sign the transaction and return the signed transaction object.
	signedTx, err := types.SignTx(tx, txSigner, s.privKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction")
	}

	return signedTx, nil
}

// Address returns the address of the signer.
func (s *signer) Address() common.Address {
	return s.address
}

// PublicKey returns the public key of the signer.
func (s *signer) PublicKey() *ecdsa.PublicKey {
	return s.pubKey
}

// signHash signs the hash with the Ethereum prefix and returns the signed hash.
func signHash(hash []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)

	return crypto.Keccak256([]byte(msg))
}
