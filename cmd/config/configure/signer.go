package configure

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/env"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature"
)

// Signer returns the signer that are being used in the bridge.
func (c *config) Signer() signature.Signer {
	// Create a struct to hold the signer.
	cfg := struct {
		PrivateKey string `config:"evm_signer,required"`
	}{}

	// Load the signer.
	c.onceSigner.Do(func() {
		// Get the data map from the getter.
		dataMap, err := c.getter.GetStringMap("signer")
		if err != nil {
			panic(errors.Wrap(err, "unable to get signature config"))
		}

		// Load the configuration from the data map into the struct.
		err = env.NewConfiguration(&cfg).From(dataMap).Load()
		if err != nil {
			panic(errors.Wrap(err, "unable to load signature config"))
		}
	})

	// Create the signer that will be used in the bridge.
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		panic(errors.Wrap(err, "unable to create private key"))
	}

	return signature.NewSigner(privateKey)
}
