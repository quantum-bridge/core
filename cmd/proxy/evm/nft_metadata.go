package evm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	datashared "github.com/quantum-bridge/core/cmd/data/shared"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc1155"
	"github.com/quantum-bridge/core/cmd/proxy/evm/generated/erc721"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
	"io"
	"math/big"
	"net/http"
	"strings"
)

const (
	// OpenSeaAPIURL is the URL of the OpenSea API.
	OpenSeaAPIURL = "https://api.opensea.io"
)

// GetNFTMetadata returns the metadata for the given non-fungible token.
func (p *proxyEVM) GetNFTMetadata(tokenChain datashared.TokenChain, tokenId string) (*datashared.NFTMetadata, error) {
	metadataURI, err := p.GetNFTMetadataURI(tokenChain, tokenId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get metadata URI for token %s", tokenId)
	}

	// Get the metadata for the given metadata URI and token ID.
	return p.getMetadata(metadataURI, tokenId)
}

// GetNFTMetadataURI returns the metadata URI for the given non-fungible token.
func (p *proxyEVM) GetNFTMetadataURI(tokenChain datashared.TokenChain, tokenId string) (string, error) {
	// Parse the token ID to big.Int type.
	tokenID, ok := big.NewInt(0).SetString(tokenId, 10)
	if !ok {
		return "", errors.Errorf("failed to parse token ID %s", tokenId)
	}

	var metadataURI string
	var err error

	// Switch on the token type to get the metadata URI for the given non-fungible token.
	switch tokenChain.TokenType {
	case TokenERC721:
		metadataURI, err = p.getERC721MetadataURI(tokenChain, tokenID)
	case TokenERC1155:
		metadataURI, err = p.getERC1155MetadataURI(tokenChain, tokenID)
	default:
		return "", errors.New("unsupported type of token")
	}
	if err != nil {
		return "", errors.Wrap(err, "failed to get metadata URI")
	}

	return metadataURI, nil
}

// getERC721MetadataURI returns the metadata URI for the given ERC721 token.
func (p *proxyEVM) getERC721MetadataURI(tokenChain datashared.TokenChain, tokenId *big.Int) (string, error) {
	// Create a new ERC721 contract instance.
	token, err := erc721.NewErc721(common.HexToAddress(tokenChain.TokenAddress), p.client)
	if err != nil {
		return "", errors.Wrap(err, "failed to get ERC721 contract")
	}

	// Get the token URI for the given token ID.
	uri, err := token.TokenURI(nil, tokenId)
	if err != nil {
		return "", errors.Wrap(err, "failed to get token URI")
	}

	return uri, nil
}

// getERC1155MetadataURI returns the metadata URI for the given ERC1155 token.
func (p *proxyEVM) getERC1155MetadataURI(tokenChain datashared.TokenChain, tokenId *big.Int) (string, error) {
	// Create a new ERC1155 contract instance.
	token, err := erc1155.NewErc1155(common.HexToAddress(tokenChain.TokenAddress), p.client)
	if err != nil {
		return "", errors.Wrap(err, "failed to get ERC1155 contract")
	}

	// Get the token URI for the given token ID.
	uri, err := token.Uri(nil, tokenId)
	if err != nil {
		return "", errors.Wrap(err, "failed to get token URI")
	}

	return uri, nil
}

// getMetadata returns the metadata for the given URL and token ID.
func (p *proxyEVM) getMetadata(url string, tokenId string) (*datashared.NFTMetadata, error) {
	// Create the token URL with the given token ID and URL.
	tokenUrl := p.ipfs.CreateUrl(url)
	// Construct the OpenSea API URL with the given token ID.
	tokenUrl = constructOpenSeaAPIURL(tokenUrl, tokenId)
	// Replace the "{id}" placeholder in the URL with the given token ID.
	tokenUrl = replaceTokenIDInURL(tokenUrl, tokenId)

	// Get the metadata from the given URL.
	metadata, err := getMetadata(tokenUrl)
	if err != nil {
		return nil, err
	}

	// Set the metadata URL, icon URL.
	metadata.MetadataURL = url
	metadata.Image = p.ipfs.CreateUrl(metadata.Image)

	// Wrap the animation URL and external URL with the IPFS URL.
	metadata.AnimationURL = p.createIPFSURL(metadata.AnimationURL)
	metadata.ExternalURL = p.createIPFSURL(metadata.ExternalURL)

	return metadata, nil
}

// createIPFSURL creates an IPFS URL with the given URL if it is not nil.
func (p *proxyEVM) createIPFSURL(url *string) *string {
	// Check if the URL is nil.
	if url == nil {
		return nil
	}

	// Create the IPFS URL with the given URL.
	wrapped := p.ipfs.CreateUrl(*url)

	return &wrapped
}

// getMetadata returns the metadata for the given URL.
func getMetadata(url string) (*datashared.NFTMetadata, error) {
	// Get the metadata from the given URL with an HTTP GET request.
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get metadata from %s", url)
	}
	defer response.Body.Close()

	// Read the body of the response.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read metadata from %s", url)
	}

	// Check if the status code of the response is not 200 OK - return not found error.
	if response.StatusCode != http.StatusOK {
		return nil, bridgeErrors.ErrNotFound
	}

	// Check if the status code of the response is not 200 OK - return not found error.
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.Errorf("failed to get metadata status code: %d, body: %s", response.StatusCode, body)
	}

	// Unmarshal the metadata from the body of the response to the NFT metadata structure.
	metadata := &datashared.NFTMetadata{}
	if err := json.Unmarshal(body, metadata); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal metadata from %s", url)
	}

	return metadata, nil
}

// constructOpenSeaAPIURL constructs the OpenSea API URL with the given token ID.
func constructOpenSeaAPIURL(baseURL, tokenID string) string {
	// Check if the URL starts with the OpenSea API URL.
	if !strings.HasPrefix(baseURL, OpenSeaAPIURL) {
		return baseURL
	}

	// Replace the token ID placeholder with the given token ID.
	return strings.Replace(baseURL, "0x{id}", tokenID, 1)
}

// replaceTokenIDInURL replaces the "{id}" placeholder in the given URL with the provided token ID.
func replaceTokenIDInURL(url, tokenId string) string {
	return strings.Replace(url, "{id}", tokenId, 1)
}
