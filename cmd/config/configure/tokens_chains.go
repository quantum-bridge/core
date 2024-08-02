package configure

import (
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/env"
)

func (c *config) Chains() []datashared.Chain {
	c.loadTokensAndChains()

	return c.chains
}

func (c *config) Token() []datashared.Token {
	c.loadTokensAndChains()

	return c.tokens
}

func (c *config) TokenChains() []datashared.TokenChain {
	c.loadTokensAndChains()

	return c.tokenChains
}

func (c *config) loadTokensAndChains() {
	c.onceTokenChains.Do(func() {
		// Create a struct to hold the tokens and chains.
		cfg := struct {
			Tokens []datashared.Token `config:"tokens,required"`
			Chains []datashared.Chain `config:"chains,required"`
		}{}

		// Get the data map from the getter.
		dataMap, err := c.getter.GetStringMap("data")
		if err != nil {
			panic(err)
		}

		// Load the configuration from the data map into the struct.
		err = env.NewConfiguration(&cfg).From(dataMap).Load()
		if err != nil {
			panic(errors.Wrap(err, "unable to load config"))
		}

		// Initialize the Tokens and Chains in the struct.
		for i, chain := range cfg.Chains {
			chain.Tokens = make([]datashared.TokenChain, 0)
			cfg.Chains[i] = chain
		}

		// Set the TokenID for each TokenChain and add the TokenChain to the Tokens and Chains.
		tokenChains := make([]datashared.TokenChain, 0)
		for _, token := range cfg.Tokens {
			for i, tokenChain := range token.Chains {
				tokenChain.TokenID = token.ID
				token.Chains[i] = tokenChain
				tokenChains = append(tokenChains, tokenChain)
				for j, chain := range cfg.Chains {
					if chain.ID == tokenChain.ChainID {
						chain.Tokens = append(chain.Tokens, tokenChain)
						cfg.Chains[j] = chain
					}
				}
			}
		}

		// Assign the Tokens and Chains to c.tokens and c.chains
		c.tokens = cfg.Tokens
		c.chains = cfg.Chains
		c.tokenChains = tokenChains
	})
}
