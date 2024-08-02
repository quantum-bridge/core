package ipfs

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

// IPFS is an interface for handling IPFS requests.
type IPFS interface {
	// Get gets data from IPFS by url.
	Get(result interface{}, url string) error
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

// Get gets data from IPFS by url.
func (i *ipfs) Get(result interface{}, url string) error {
	resp, err := http.Get(i.CreateUrl(url))
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response")
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return errors.New(fmt.Sprintf("failed to get data from ipfs status %d, body %s",
			resp.StatusCode, string(body)))
	}

	return json.Unmarshal(body, result)
}

// CreateUrl creates a new url with IPFS endpoint.
func (i *ipfs) CreateUrl(url string) string {
	if !i.isIPFSUrl(url) {
		return url
	}

	// Remove the ipfs://ipfs/ prefix.
	url = strings.Replace(url, "ipfs://ipfs/", "", 1)
	// Remove the ipfs:// prefix if it is not ipfs://ipfs/.
	url = strings.Replace(url, "ipfs://", "", 1)

	// Wrap the url to the IPFS proxy url and return it (e.g. https://ipfs.io/ipfs/QmNt1NDTaHraeQoGaRidWTUPa2w4Fe8gkk2R3mJzCqCoPz).
	return fmt.Sprintf("%s/ipfs/%s", i.endpoint, url)
}

// isIPFSUrl checks if the url is an IPFS url (e.g. ipfs://ipfs/QmNt1NDTaHraeQoGaRidWTUPa2w4Fe8gkk2R3mJzCqCoPz).
func (i *ipfs) isIPFSUrl(url string) bool {
	return strings.HasPrefix(url, "ipfs://")
}
