package shared

// ContractCreationResponse is the response structure for the contract creation API.
type ContractCreationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		ContractAddress string `json:"contractAddress"`
		ContractCreator string `json:"contractCreator"`
		TxHash          string `json:"txHash"`
	} `json:"result"`
}

// TransactionByHashResponse is the response structure for the transaction by hash API.
type TransactionByHashResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash        string        `json:"blockHash"`
		BlockNumber      string        `json:"blockNumber"`
		From             string        `json:"from"`
		Gas              string        `json:"gas"`
		GasPrice         string        `json:"gasPrice"`
		Hash             string        `json:"hash"`
		Input            string        `json:"input"`
		Nonce            string        `json:"nonce"`
		To               *string       `json:"to"`
		TransactionIndex string        `json:"transactionIndex"`
		Value            string        `json:"value"`
		Type             string        `json:"type"`
		AccessList       []interface{} `json:"accessList"`
		ChainId          string        `json:"chainId"`
		V                string        `json:"v"`
		R                string        `json:"r"`
		S                string        `json:"s"`
		YParity          string        `json:"yParity"`
	} `json:"result"`
}
