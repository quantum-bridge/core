package repositories

// TransactionsHistory is a struct that represents the transactions history table.
type TransactionsHistory struct {
	ID           int64  `db:"id" structs:"-"`
	TxHash       string `db:"tx_hash" structs:"tx_hash"`
	TokenAddress string `db:"token_address" structs:"token_address"`
	TokenID      string `db:"token_id" structs:"token_id"`
	Amount       string `db:"amount" structs:"amount"`
	FromAddress  string `db:"from_address" structs:"from_address"`
	ToAddress    string `db:"to_address" structs:"to_address"`
	NetworkFrom  string `db:"network_from" structs:"network_from"`
	NetworkTo    string `db:"network_to" structs:"network_to"`
}
