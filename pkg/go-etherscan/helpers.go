package goetherscan

import (
	"strconv"
	"strings"
)

// HexToDecimal converts a hex string to a decimal int64.
func HexToDecimal(hex string) (int64, error) {
	// Remove the "0x" prefix if it exists in the hex string.
	hex = strings.TrimPrefix(hex, "0x")

	// Parse the hex string to an int64 decimal.
	dec, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}

	return dec, nil
}
