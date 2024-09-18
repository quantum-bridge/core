package ipfs

import (
	"fmt"
	"strings"
)

const (
	// ipfsPrefix is the prefix for the IPFS url.
	ipfsPrefix = "ipfs://"
	// ipfsPrefixWithSlash is the prefix for the IPFS url with a slash.
	ipfsPrefixWithSuffix = "ipfs://ipfs/"
)

// IPFS is an interface for handling IPFS requests.
type IPFS interface {
	// CreateUrl creates a new url with IPFS endpoint.
	CreateUrl(url string) string
}

// ipfs is structure that stores the IPFS proxy url.
type ipfs struct {
	endpoint string
}

// New creates a new instance of the IPFS.
func New(endpoint string) IPFS {
	endpoint = strings.TrimSuffix(endpoint, "/")

	return &ipfs{
		endpoint: endpoint,
	}
}

// CreateUrl creates a new url with IPFS endpoint.
func (i *ipfs) CreateUrl(url string) string {
	if !i.isIPFSUrl(url) {
		return url
	}

	// Remove the ipfs://ipfs/ prefix.
	url = strings.Replace(url, ipfsPrefixWithSuffix, "", 1)
	// Remove the ipfs:// prefix if it is not ipfs://ipfs/.
	url = strings.Replace(url, ipfsPrefix, "", 1)

	// Wrap the url to the IPFS proxy url and return it (e.g. https://ipfs.io/ipfs/QmNt1NDTaHraeQoGaRidWTUPa2w4Fe8gkk2R3mJzCqCoPz).
	return fmt.Sprintf("%s/ipfs/%s", i.endpoint, url)
}

// isIPFSUrl checks if the url is an IPFS url (e.g. ipfs://ipfs/QmNt1NDTaHraeQoGaRidWTUPa2w4Fe8gkk2R3mJzCqCoPz).
func (i *ipfs) isIPFSUrl(url string) bool {
	return strings.HasPrefix(url, ipfsPrefix)
}
