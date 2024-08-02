package requests

import (
	"encoding/json"
	"github.com/quantum-bridge/core/cmd/data"
	"net/http"
)

// GetTokensDTO is the data transfer object for the get tokens request.
type GetTokensDTO struct {
	FilterType    []data.TokenType
	IncludeChains bool
}

// NewGetTokensRequest creates a new get tokens request from the given HTTP request.
func NewGetTokensRequest(r *http.Request) (GetTokensDTO, error) {
	// Create a new request object.
	request := GetTokensDTO{}
	// Get the query parameters from the URL.
	params := r.URL.Query()

	// Get 'token_type' query parameter (optional).
	if chainTypes, ok := params["filter[token_type]"]; ok {
		// The query parameter is a JSON array, so unmarshal it.
		err := json.Unmarshal([]byte(chainTypes[0]), &request.FilterType)
		if err != nil {
			return request, err
		}
	}

	// Get 'include_tokens' query parameter (optional).
	includeChains := params.Get("include_chains")
	if includeChains == "true" {
		request.IncludeChains = true
	}

	return request, nil
}
