package configure

import (
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/env"
	"net"
)

// Listener returns the HTTP listener that are being used in the bridge.
func (c *config) Listener() net.Listener {
	// Load the listener.
	c.onceListener.Do(func() {
		// Create a struct to hold the listener.
		cfg := struct {
			Addr string `config:"port,required"`
		}{}

		// Get the data map from the getter.
		dataMap, err := c.getter.GetStringMap("listener")
		if err != nil {
			panic(errors.Wrap(err, "unable to get listener config"))
		}

		// Load the configuration from the data map into the struct.
		err = env.NewConfiguration(&cfg).From(dataMap).Load()
		if err != nil {
			panic(errors.Wrap(err, "unable to load listener config"))
		}

		// Create the listener that will be used in the bridge.
		listener, err := net.Listen("tcp", cfg.Addr)
		if err != nil {
			panic(errors.Wrap(err, "unable to create listener"))
		}

		// Set the listener in the config.
		c.listener = listener
	})

	return c.listener
}
