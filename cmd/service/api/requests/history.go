package requests

import (
	"net/http"
	"strconv"
)

// HistoryRequest represents the request parameters for fetching history
type HistoryRequest struct {
	FromAddress        *string `json:"from_address"`
	ToAddress          *string `json:"to_address"`
	SourceNetwork      *string `json:"source_network"`
	DestinationNetwork *string `json:"destination_network"`
	TokenAddress       *string `json:"token_address"`
	TransactionType    *string `json:"transaction_type"` // "deposits", "withdrawals", or nil for both
	FromBlock          *uint64 `json:"from_block"`
	ToBlock            *uint64 `json:"to_block"`
	Page               int     `json:"page"`
	PageSize           int     `json:"page_size"`
	SortBy             string  `json:"sort_by"`
	SortOrder          string  `json:"sort_order"` // "asc" or "desc"
}

// NewHistoryRequest creates a new history request from HTTP request
func NewHistoryRequest(r *http.Request) (*HistoryRequest, error) {
	query := r.URL.Query()

	req := &HistoryRequest{
		Page:      1,
		PageSize:  10,
		SortBy:    "block_number",
		SortOrder: "desc",
	}

	if page := query.Get("page"); page != "" {
		if pageNum, err := strconv.Atoi(page); err == nil && pageNum > 0 {
			req.Page = pageNum
		}
	}

	if pageSize := query.Get("page_size"); pageSize != "" {
		if size, err := strconv.Atoi(pageSize); err == nil && size > 0 {
			req.PageSize = size
		}
	}

	if fromAddr := query.Get("from_address"); fromAddr != "" {
		req.FromAddress = &fromAddr
	}

	if toAddr := query.Get("to_address"); toAddr != "" {
		req.ToAddress = &toAddr
	}

	if srcNet := query.Get("source_network"); srcNet != "" {
		req.SourceNetwork = &srcNet
	}

	if destNet := query.Get("destination_network"); destNet != "" {
		req.DestinationNetwork = &destNet
	}

	if tokenAddr := query.Get("token_address"); tokenAddr != "" {
		req.TokenAddress = &tokenAddr
	}

	if txType := query.Get("transaction_type"); txType != "" {
		if txType == "deposits" || txType == "withdrawals" {
			req.TransactionType = &txType
		}
	}

	if fromBlock := query.Get("from_block"); fromBlock != "" {
		if block, err := strconv.ParseUint(fromBlock, 10, 64); err == nil {
			req.FromBlock = &block
		}
	}

	if toBlock := query.Get("to_block"); toBlock != "" {
		if block, err := strconv.ParseUint(toBlock, 10, 64); err == nil {
			req.ToBlock = &block
		}
	}

	if sortBy := query.Get("sort_by"); sortBy != "" {
		req.SortBy = sortBy
	}

	if sortOrder := query.Get("sort_order"); sortOrder != "" {
		if sortOrder == "asc" || sortOrder == "desc" {
			req.SortOrder = sortOrder
		}
	}

	return req, nil
}
