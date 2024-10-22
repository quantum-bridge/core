package shared

import "encoding/json"

// EntityType is the type of the entity.
type EntityType string

const (
	// BALANCE represents the balance entity type.
	BALANCE EntityType = "balance"
	// CHAIN represents the chain entity type.
	CHAIN EntityType = "chain"
	// EVM_TRANSACTION represents the EVM transaction entity type.
	EVM_TRANSACTION EntityType = "evm_transaction"
	// NFT represents the NFT entity type.
	NFT EntityType = "nft"
	// PROCESSED_TRANSACTION represents the processed transaction entity type.
	PROCESSED_TRANSACTION EntityType = "processed_transaction"
	// TOKEN represents the token entity type.
	TOKEN EntityType = "token"
)

// GetEtherscanParams returns the Etherscan API URL and API key from the given chain parameters (json.RawMessage).
func GetEtherscanParams(params json.RawMessage) (string, string, error) {
	var etherscanParams struct {
		ApiUrl string `json:"api_url"`
		ApiKey string `json:"api_key"`
	}

	if err := json.Unmarshal(params, &etherscanParams); err != nil {
		return "", "", err
	}

	return etherscanParams.ApiUrl, etherscanParams.ApiKey, nil
}
