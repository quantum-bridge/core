package responses

// HistoryEntry represents a single history entry that can be either deposit or withdrawal
type HistoryEntry struct {
	Type               string `json:"type"` // "deposit" or "withdrawal"
	TxHash             string `json:"tx_hash,omitempty"`
	BlockNumber        uint64 `json:"block_number"`
	TokenAddress       string `json:"token_address"`
	TokenID            string `json:"token_id,omitempty"`
	Amount             string `json:"amount"`
	FromAddress        string `json:"from_address,omitempty"`
	ToAddress          string `json:"to_address"`
	SourceNetwork      string `json:"source_network,omitempty"`
	DestinationNetwork string `json:"destination_network"`
	IsMintable         bool   `json:"is_mintable"`
	WithdrawalTxHash   string `json:"withdrawal_tx_hash,omitempty"`
	DepositTxHash      string `json:"deposit_tx_hash,omitempty"`
}

// HistoryResponse represents the response for history endpoint
type HistoryResponse struct {
	Data       []HistoryEntry `json:"data"`
	Pagination struct {
		CurrentPage int   `json:"current_page"`
		PageSize    int   `json:"page_size"`
		TotalItems  int64 `json:"total_items"`
		TotalPages  int   `json:"total_pages"`
	} `json:"pagination"`
}
