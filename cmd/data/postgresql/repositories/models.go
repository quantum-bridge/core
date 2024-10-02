package repositories

// DepositsHistory is a struct that represents the deposits history table.
type DepositsHistory struct {
	ID                 int64  `db:"id" structs:"-"`
	SourceNetwork      string `db:"source_network" structs:"source_network"`
	TxHash             string `db:"tx_hash" structs:"tx_hash"`
	BlockNumber        uint64 `db:"block_number" structs:"block_number"`
	TokenAddress       string `db:"token_address" structs:"token_address"`
	TokenID            string `db:"token_id" structs:"token_id"`
	Amount             string `db:"amount" structs:"amount"`
	FromAddress        string `db:"from_address" structs:"from_address"`
	ToAddress          string `db:"to_address" structs:"to_address"`
	DestinationNetwork string `db:"destination_network" structs:"destination_network"`
	IsMintable         bool   `db:"is_mintable" structs:"is_mintable"`
}

// WithdrawalsHistory is a struct that represents the withdrawals history table.
type WithdrawalsHistory struct {
	ID                 int64  `db:"id" structs:"-"`
	SourceNetwork      string `db:"source_network" structs:"source_network"`
	TxHash             string `db:"tx_hash" structs:"tx_hash"`
	BlockNumber        uint64 `db:"block_number" structs:"block_number"`
	TokenAddress       string `db:"token_address" structs:"token_address"`
	TokenID            string `db:"token_id" structs:"token_id"`
	Amount             string `db:"amount" structs:"amount"`
	FromAddress        string `db:"from_address" structs:"from_address"`
	ToAddress          string `db:"to_address" structs:"to_address"`
	DestinationNetwork string `db:"destination_network" structs:"destination_network"`
	IsMintable         bool   `db:"is_mintable" structs:"is_mintable"`
}
