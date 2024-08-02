package configure

import (
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/env"
	"github.com/quantum-bridge/core/cmd/ipfs"
	"github.com/quantum-bridge/core/cmd/proxy/evm/signature"
	"net"
	"sync"
)

// Config is the interface for the configuration of the bridge.
type Config interface {
	// Chains returns the chains that are being used in the bridge.
	Chains() []datashared.Chain
	// Token returns the tokens that are being used in the bridge.
	Token() []datashared.Token
	// TokenChains returns the token chains that are being used in the bridge.
	TokenChains() []datashared.TokenChain
	// Listener returns the HTTP listener that are being used in the bridge.
	Listener() net.Listener
	// Signer returns the signer that are being used in the bridge.
	Signer() signature.Signer
	// IPFS returns the IPFS client that are being used in the bridge.
	IPFS() ipfs.IPFS
}

// config is struct that holds the configuration of the bridge.
type config struct {
	getter          env.Getter
	onceTokenChains sync.Once
	onceListener    sync.Once
	onceIPFS        sync.Once
	onceSigner      sync.Once

	chains      []datashared.Chain
	tokens      []datashared.Token
	tokenChains []datashared.TokenChain

	listener net.Listener
}

// NewConfig creates a new configuration instance with the given getter.
func NewConfig(getter env.Getter) Config {
	return &config{
		getter: getter,
	}
}
