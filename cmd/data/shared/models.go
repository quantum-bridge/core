package shared

import (
	"encoding/json"
	"github.com/quantum-bridge/core/cmd/data"
)

// Chain is the structure that holds the chain data.
type Chain struct {
	ID                    string          `config:"id,required"`
	Name                  string          `config:"name,required"`
	Type                  data.ChainType  `config:"type,required"`
	ChainParams           json.RawMessage `config:"chain_params,required"`
	BridgeContractAddress string          `config:"bridge_contract_address,required"`
	RpcEndpoint           string          `config:"rpc_endpoint,required"`
	Confirmations         int64           `config:"confirmations,required"`
	Tokens                []TokenChain
}

// Token is the structure that holds the token data.
type Token struct {
	ID     string         `config:"id,required"`
	Name   string         `config:"name,required"`
	Symbol string         `config:"symbol,required"`
	Type   data.TokenType `config:"type,required"`
	Chains []TokenChain   `config:"chains,required"`
}

// TokenChain is the structure that holds the token chain data.
type TokenChain struct {
	TokenID      string
	ChainID      string          `config:"chain_id,required"`
	TokenAddress string          `config:"token_address,required"`
	TokenType    string          `config:"token_type,required"`
	BridgeType   data.BridgeType `config:"bridge_type,required"`
	AutoSend     bool            `config:"auto_send"`
}
