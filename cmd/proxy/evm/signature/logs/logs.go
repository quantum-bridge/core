package logs

// SHA3 represents the Solidity SHA3 hash interface for the logs.
type SHA3 interface {
	// Hash returns the hash of the data.
	Hash() []byte
}
