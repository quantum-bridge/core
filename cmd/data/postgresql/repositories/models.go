package repositories

// TransactionsHistory is a struct that represents the transactions history table.
type TransactionsHistory struct {
	ID           int64  `db:"id" structs:"-"`
	ChainID      int64  `db:"chain_id" structs:"chain_id"`
	TxHash       string `db:"tx_hash" structs:"tx_hash"`
	BlockNumber  uint64 `db:"block_number" structs:"block_number"`
	TokenAddress string `db:"token_address" structs:"token_address"`
	TokenID      string `db:"token_id" structs:"token_id"`
	Amount       string `db:"amount" structs:"amount"`
	FromAddress  string `db:"from_address" structs:"from_address"`
	ToAddress    string `db:"to_address" structs:"to_address"`
	ToNetwork    string `db:"to_network" structs:"to_network"`
	IsMintable   bool   `db:"is_mintable" structs:"is_mintable"`
}
