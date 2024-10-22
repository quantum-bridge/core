package config

// Services is the interface for the services configuration.
type Services interface {
	ServicesConfig() *ServicesConfig
	SetServicesConfig(config *ServicesConfig)
}

// ServicesConfig is the configuration for the services of the bridge.
type ServicesConfig struct {
	// TxHistory is the configuration for the transaction history service of the bridge.
	TxHistory bool
}

// services is the struct that holds the configuration for the services of the bridge.
type services struct {
	config *ServicesConfig
}

// NewServices creates a new services instance with the given configuration for the services.
func NewServices(config *ServicesConfig) Services {
	return &services{
		config: config,
	}
}

// ServicesConfig returns the configuration for the services.
func (s *services) ServicesConfig() *ServicesConfig {
	return s.config
}

// SetServicesConfig sets the configuration for the services.
func (s *services) SetServicesConfig(config *ServicesConfig) {
	s.config = config
}
