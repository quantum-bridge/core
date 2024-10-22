package shared

import "strings"

// IsValidEthereumAddress checks if the given string is a valid Ethereum address.
func IsValidEthereumAddress(address string) bool {
	// Ethereum addresses start with '0x' and are 42 characters long
	return len(address) == 42 && strings.HasPrefix(address, "0x")
}
