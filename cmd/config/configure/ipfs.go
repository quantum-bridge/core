package configure

import (
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/env"
	"github.com/quantum-bridge/core/cmd/ipfs"
)

// IPFS returns the IPFS client that are being used in the bridge.
func (c *config) IPFS() ipfs.IPFS {
	// Create a struct to hold the IPFS client.
	cfg := struct {
		Endpoint string `config:"endpoint,required"`
	}{}

	// Load the IPFS client.
	c.onceIPFS.Do(func() {
		// Get the data map from the getter.
		dataMap, err := c.getter.GetStringMap("ipfs")
		if err != nil {
			panic(errors.Wrap(err, "unable to get ipfs config"))
		}

		// Load the configuration from the data map into the struct.
		err = env.NewConfiguration(&cfg).From(dataMap).Load()
		if err != nil {
			panic(errors.Wrap(err, "unable to load ipfs config"))
		}
	})

	return ipfs.New(cfg.Endpoint)
}
