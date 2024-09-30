package requests

import (
	"encoding/json"
	"github.com/quantum-bridge/core/cmd/data"
	"net/http"
)

// GetChainsDTO is the data transfer object for the GetChains request.
type GetChainsDTO struct {
	// FilterType is the type of chain to filter by.
	FilterType []data.ChainType
	// IncludeTokens is a flag to include tokens in the response.
	IncludeTokens bool
}

// NewGetChainsRequest creates a new GetChainsDTO from an HTTP request.
func NewGetChainsRequest(r *http.Request) (GetChainsDTO, error) {
	// Create a new request object.
	request := GetChainsDTO{}
	// Get the query parameters from the URL.
	params := r.URL.Query()

	// Get 'chain_type' query parameter (optional).
	if chainTypes, ok := params["filter[chain_type]"]; ok {
		// The query parameter is a JSON array, so unmarshal it
		err := json.Unmarshal([]byte(chainTypes[0]), &request.FilterType)
		if err != nil {
			return request, err
		}
	}

	// Get 'token' query parameter (optional).
	includeTokens := params.Get("include_tokens")
	if includeTokens == "true" {
		request.IncludeTokens = true
	}

	return request, nil
}
