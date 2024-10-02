package goetherscan

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/pkg/go-etherscan/shared"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// defaultTimeout is the default timeout for the HTTP client in seconds.
	defaultTimeout = 30 * time.Second // 30 seconds.
)

// Client is the interface for the etherscan client.
type Client interface {
	// GetContractCreation gets the contract creation information for the provided contract addresses.
	GetContractCreation(contractAddresses ...string) (*shared.ContractCreationResponse, error)
	// GetTransactionByHash gets the transaction information for the provided transaction hash.
	GetTransactionByHash(txHash string) (*shared.TransactionByHashResponse, error)
}

// client is the structure that holds the HTTP client and the API key.
type client struct {
	client *http.Client
	apiKey string
	apiUrl string
}

// New creates a new etherscan client with the provided API URL and API key.
func New(ApiUrl, ApiKey string) Client {
	return &client{
		client: &http.Client{Timeout: defaultTimeout},
		apiKey: ApiKey,
		apiUrl: ApiUrl,
	}
}

// SetTimeout sets the timeout for the HTTP client.
func (c *client) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
}

// get makes a get request to the provided URL and decodes the JSON response.
func (c *client) get(u *url.URL, result interface{}) error {
	response, err := c.client.Get(u.String())
	if err != nil {
		return errors.Wrap(err, "failed to make HTTP request")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.Wrap(err, "failed to make HTTP request")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	if err := json.Unmarshal(body, result); err != nil {
		return errors.Wrap(err, "failed to unmarshal JSON response")
	}

	return nil
}

// GetContractCreation gets the contract creation information for the provided contract addresses.
func (c *client) GetContractCreation(contractAddresses ...string) (*shared.ContractCreationResponse, error) {
	parsedUrl, err := url.Parse(c.apiUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse API URL")
	}

	query := parsedUrl.Query()
	query.Set("module", "contract")
	query.Set("action", "getcontractcreation")
	query.Set("contractaddresses", strings.Join(contractAddresses, ","))
	query.Set("apikey", c.apiKey)
	parsedUrl.RawQuery = query.Encode()

	var contractCreationResponse shared.ContractCreationResponse
	if err := c.get(parsedUrl, &contractCreationResponse); err != nil {
		return nil, err
	}

	return &contractCreationResponse, nil
}

// GetTransactionByHash gets the transaction information for the provided transaction hash.
func (c *client) GetTransactionByHash(txHash string) (*shared.TransactionByHashResponse, error) {
	parsedUrl, err := url.Parse(c.apiUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse API URL")
	}

	query := parsedUrl.Query()
	query.Set("module", "proxy")
	query.Set("action", "eth_getTransactionByHash")
	query.Set("txhash", txHash)
	query.Set("apikey", c.apiKey)
	parsedUrl.RawQuery = query.Encode()

	var txResponse shared.TransactionByHashResponse
	if err := c.get(parsedUrl, &txResponse); err != nil {
		return nil, err
	}

	return &txResponse, nil
}
