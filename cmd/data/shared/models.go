package shared

import (
	"encoding/json"
	"github.com/quantum-bridge/core/cmd/data"
)

// TokenChain is the structure that holds the token chain data.
type TokenChain struct {
	TokenID      string          `json:"token_id"`
	ChainID      string          `config:"chain_id,required" json:"chain_id"`
	TokenAddress string          `config:"token_address" json:"token_address"`
	TokenType    string          `config:"token_type,required" json:"token_type"`
	BridgeType   data.BridgeType `config:"bridge_type,required" json:"bridge_type"`
	AutoSend     bool            `config:"auto_send" json:"auto_send"`
}

// Chain is the structure that holds the chain data.
type Chain struct {
	ID            string          `config:"id,required" json:"id"`
	Name          string          `config:"name,required" json:"name"`
	Type          data.ChainType  `config:"type,required" json:"type"`
	ChainParams   json.RawMessage `config:"chain_params,required" json:"chain_params"`
	BridgeAddress string          `config:"bridge_address,required" json:"bridge_address"`
	Icon          string          `config:"icon" json:"icon"`
	RpcEndpoint   string          `config:"rpc_endpoint,required" json:"rpc_endpoint"`
	Confirmations int64           `config:"confirmations,required" json:"confirmations"`
	Tokens        []TokenChain    `json:"tokens"`
}

// Token is the structure that holds the token data.
type Token struct {
	ID     string         `config:"id,required" json:"id"`
	Name   string         `config:"name,required" json:"name"`
	Symbol string         `config:"symbol,required" json:"symbol"`
	Icon   string         `config:"icon" json:"icon"`
	Type   data.TokenType `config:"type,required" json:"type"`
	Chains []TokenChain   `config:"chains,required" json:"chains"`
}
