package proxy

import (
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/cmd/data"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/ipfs"
	proxybridge "github.com/quantum-bridge/core/cmd/proxy/bridge"
	"github.com/quantum-bridge/core/cmd/proxy/evm"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature"
)

// Proxy is the interface for the proxy of the chains that are being used in the bridge.
type Proxy interface {
	// Get returns the proxy for the chain by the given chain ID.
	Get(chainID string) proxybridge.Bridge
}

// proxy is the structure that holds the proxy data.
type proxy struct {
	// proxyChains is the map that holds the chain ID and the proxy instance.
	proxyChains map[string]proxybridge.Bridge
}

// NewProxy creates a new proxy instance with the given chains, signer and IPFS client.
func NewProxy(chains []datashared.Chain, signer signature.Signer, ipfsClient ipfs.IPFS) (Proxy, error) {
	// Create a new proxy instance.
	p := &proxy{
		proxyChains: make(map[string]proxybridge.Bridge),
	}

	// Iterate over the chains to create the proxy for each chain.
	for _, chain := range chains {
		// Switch on the chain type to create the proxy for the chain.
		switch chain.Type {
		// If the chain type is EVM, create a new EVM proxy.
		case data.EVM:
			// Create a new EVM proxy.
			evmProxy, err := evm.NewProxy(chain.RpcEndpoint, chain.BridgeAddress, signer, ipfsClient, chain.Confirmations)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to create EVM proxy for chain %s", chain.ID)
			}

			p.proxyChains[chain.ID] = evmProxy
		default:
			return nil, errors.Errorf("unsupported chain type: %s", chain.Type)
		}
	}

	return p, nil
}

// Get returns the proxy for the chain by the given chain ID.
func (p *proxy) Get(chainID string) proxybridge.Bridge {
	return p.proxyChains[chainID]
}
