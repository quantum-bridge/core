// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chains": {
            "get": {
                "description": "Get a list of chains and tokens based on the request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chains"
                ],
                "summary": "Get chains list",
                "operationId": "getChains",
                "parameters": [
                    {
                        "type": "array",
                        "description": "Filter by chain type. Items Value: [` + "`" + `'evm'` + "`" + `]",
                        "name": "filter[chain_type]",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Include tokens in the response. Items Value: [` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `]",
                        "name": "include_tokens",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.ChainListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/tokens": {
            "get": {
                "description": "Get the list of tokens based on the filter type and include chains flag.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokens"
                ],
                "summary": "Get Tokens",
                "operationId": "getTokens",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Include chains in the response. Items Value: [` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `]",
                        "name": "include_chains",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "description": "Filter by chain type. Items Value: [` + "`" + `'chain'` + "`" + `]",
                        "name": "filter[token_type]",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.TokenListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/tokens/{token_id}/balance": {
            "get": {
                "description": "Get the balance of an account for a token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokens"
                ],
                "summary": "Get Balance",
                "operationId": "getBalance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "token_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Address of the account",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chain_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT ID",
                        "name": "nft",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.BalanceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/tokens/{token_id}/nfts/{nft_id}": {
            "get": {
                "description": "Get the metadata of a non-fungible token based on the token ID and NFT ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokens"
                ],
                "summary": "Get NFT metadata",
                "operationId": "getNFT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "token_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "NFT ID",
                        "name": "nft_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chain_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.NFTResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/transfers/approve": {
            "post": {
                "description": "Approve is an HTTP handler that creates an approval transaction for a spender.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "summary": "Approve",
                "operationId": "approve",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ApproveDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.TransactionsResponse"
                        }
                    },
                    "204": {
                        "description": "No content"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/transfers/lock": {
            "post": {
                "description": "Generates transaction that will lock a token in the source chain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "summary": "Lock Token",
                "operationId": "lock",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LockDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.TransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/transfers/withdraw": {
            "post": {
                "description": "Check if lock transaction is valid and withdraw the token from the bridge. Returns the transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transfers"
                ],
                "summary": "Withdraw",
                "operationId": "withdraw",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.WithdrawDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/shared.TransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "big.Int": {
            "type": "object"
        },
        "data.TokenType": {
            "type": "string",
            "enum": [
                "fungible",
                "non-fungible"
            ],
            "x-enum-varnames": [
                "FUNGIBLE",
                "NONFUNGIBLE"
            ]
        },
        "github_com_quantum-bridge_core_cmd_service_shared.Chain": {
            "type": "object",
            "required": [
                "attributes",
                "id",
                "relationships",
                "type"
            ],
            "properties": {
                "attributes": {
                    "description": "Attributes is the attributes of the chain entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.ChainAttributes"
                        }
                    ]
                },
                "id": {
                    "description": "ID is the identifier of the entity.",
                    "type": "string"
                },
                "relationships": {
                    "description": "Relationships is the relationships of the chain entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.ChainRelationships"
                        }
                    ]
                },
                "type": {
                    "description": "Type is the type of the entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.EntityType"
                        }
                    ]
                }
            }
        },
        "github_com_quantum-bridge_core_cmd_service_shared.NFTAttribute": {
            "type": "object",
            "required": [
                "trait_type",
                "value"
            ],
            "properties": {
                "trait_type": {
                    "description": "TraitType is the type of the trait.",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the value of the trait.",
                    "type": "string"
                }
            }
        },
        "github_com_quantum-bridge_core_cmd_service_shared.Token": {
            "type": "object",
            "required": [
                "attributes",
                "id",
                "relationships",
                "type"
            ],
            "properties": {
                "attributes": {
                    "description": "Attributes is the attributes of the token entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.TokenAttributes"
                        }
                    ]
                },
                "id": {
                    "description": "ID is the identifier of the entity.",
                    "type": "string"
                },
                "relationships": {
                    "description": "Relationships is the relationships of the token entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.TokenRelationships"
                        }
                    ]
                },
                "type": {
                    "description": "Type is the type of the entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.EntityType"
                        }
                    ]
                }
            }
        },
        "requests.ApproveDTO": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the data of the approval request.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/requests.ApproveRequest"
                        }
                    ]
                }
            }
        },
        "requests.ApproveRequest": {
            "type": "object",
            "required": [
                "address",
                "chain_id",
                "token_id"
            ],
            "properties": {
                "address": {
                    "description": "Address is the address of the spender.",
                    "type": "string"
                },
                "chain_id": {
                    "description": "ChainID is the ID of the chain.",
                    "type": "string"
                },
                "token_id": {
                    "description": "TokenID is the ID of the token.",
                    "type": "string"
                }
            }
        },
        "requests.LockDTO": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the data of the lock request.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/requests.LockRequest"
                        }
                    ]
                }
            }
        },
        "requests.LockRequest": {
            "type": "object",
            "required": [
                "chain_from",
                "chain_to",
                "from",
                "to",
                "token_id"
            ],
            "properties": {
                "amount": {
                    "description": "Amount is the amount of tokens to lock.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/big.Int"
                        }
                    ]
                },
                "chain_from": {
                    "description": "ChainFrom is the chain that the lock is from.",
                    "type": "string"
                },
                "chain_to": {
                    "description": "ChainTo is the chain that is receiving the amount of tokens.",
                    "type": "string"
                },
                "from": {
                    "description": "From is the sender address of the lock.",
                    "type": "string"
                },
                "nft_id": {
                    "description": "NFT is the ID of the NFT being locked in the chain.",
                    "type": "string"
                },
                "to": {
                    "description": "To is the receiver address of the lock.",
                    "type": "string"
                },
                "token_id": {
                    "description": "TokenID is the ID of the token being locked in the chain.",
                    "type": "string"
                }
            }
        },
        "requests.WithdrawDTO": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the data of the withdrawal request.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/requests.WithdrawRequest"
                        }
                    ]
                }
            }
        },
        "requests.WithdrawRequest": {
            "type": "object",
            "required": [
                "chain_from",
                "token_id",
                "tx_hash"
            ],
            "properties": {
                "chain_from": {
                    "description": "ChainFrom is the source chain ID.",
                    "type": "string"
                },
                "from": {
                    "description": "From is the address of the sender in the destination chain. Should be used only if the sender address is different with source chain.",
                    "type": "string"
                },
                "token_id": {
                    "description": "TokenID is the token ID of the token.",
                    "type": "string"
                },
                "tx_hash": {
                    "description": "TxHash is the hash of the transaction in the source chain that locked the token.",
                    "type": "string"
                }
            }
        },
        "shared.Balance": {
            "type": "object",
            "required": [
                "attributes",
                "id",
                "type"
            ],
            "properties": {
                "attributes": {
                    "description": "Attributes is the attributes of the balance entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.BalanceAttributes"
                        }
                    ]
                },
                "id": {
                    "description": "ID is the identifier of the entity.",
                    "type": "string"
                },
                "type": {
                    "description": "Type is the type of the entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.EntityType"
                        }
                    ]
                }
            }
        },
        "shared.BalanceAttributes": {
            "type": "object",
            "required": [
                "address",
                "amount",
                "token_address"
            ],
            "properties": {
                "address": {
                    "description": "Address is the address of the balance.",
                    "type": "string"
                },
                "amount": {
                    "description": "Amount is the amount of the balance.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/big.Int"
                        }
                    ]
                },
                "token_address": {
                    "description": "TokenAddress is the token address of the balance.",
                    "type": "string"
                }
            }
        },
        "shared.BalanceResponse": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the balance of the account.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.Balance"
                        }
                    ]
                },
                "included": {
                    "description": "Included is the included object in the response.",
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "shared.ChainAttributes": {
            "type": "object",
            "required": [
                "chain_params",
                "chain_type",
                "icon",
                "name"
            ],
            "properties": {
                "chain_params": {
                    "description": "ChainParams is the parameters of the chain."
                },
                "chain_type": {
                    "description": "ChainType is the type of the chain.",
                    "type": "string"
                },
                "icon": {
                    "description": "Icon is the icon of the chain.",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the chain.",
                    "type": "string"
                }
            }
        },
        "shared.ChainListResponse": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the list of chains.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_quantum-bridge_core_cmd_service_shared.Chain"
                    }
                },
                "included": {
                    "description": "Included is the included object in the response.",
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "shared.ChainRelationships": {
            "type": "object",
            "required": [
                "tokens"
            ],
            "properties": {
                "tokens": {
                    "description": "Tokens is the tokens that are used in the chain.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.RelationCollection"
                        }
                    ]
                }
            }
        },
        "shared.EntityType": {
            "type": "string",
            "enum": [
                "balance",
                "chain",
                "evm_transaction",
                "nft",
                "processed_transaction",
                "token"
            ],
            "x-enum-varnames": [
                "BALANCE",
                "CHAIN",
                "EVM_TRANSACTION",
                "NFT",
                "PROCESSED_TRANSACTION",
                "TOKEN"
            ]
        },
        "shared.Key": {
            "type": "object",
            "required": [
                "id",
                "type"
            ],
            "properties": {
                "id": {
                    "description": "ID is the identifier of the entity.",
                    "type": "string"
                },
                "type": {
                    "description": "Type is the type of the entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.EntityType"
                        }
                    ]
                }
            }
        },
        "shared.NFTAttributes": {
            "type": "object",
            "required": [
                "attributes",
                "image_url",
                "metadata_url",
                "name"
            ],
            "properties": {
                "animation_url": {
                    "description": "AnimationURL is the animation URL of the NFT.",
                    "type": "string"
                },
                "attributes": {
                    "description": "Attributes is the list of attributes of the NFT.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_quantum-bridge_core_cmd_service_shared.NFTAttribute"
                    }
                },
                "description": {
                    "description": "Description is the description of the NFT.",
                    "type": "string"
                },
                "external_url": {
                    "description": "ExternalURL is the external URL of the NFT.",
                    "type": "string"
                },
                "image_url": {
                    "description": "ImageURL is the image URL of the NFT.",
                    "type": "string"
                },
                "metadata_url": {
                    "description": "MetadataURL is the metadata URL of the NFT.",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the NFT.",
                    "type": "string"
                }
            }
        },
        "shared.NFTData": {
            "type": "object",
            "required": [
                "attributes",
                "key"
            ],
            "properties": {
                "attributes": {
                    "description": "Attributes is the attributes of the NFT entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.NFTAttributes"
                        }
                    ]
                },
                "key": {
                    "description": "Key is the key of the NFT entity.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.Key"
                        }
                    ]
                }
            }
        },
        "shared.NFTResponse": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the NFT.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.NFTData"
                        }
                    ]
                },
                "included": {
                    "description": "Includes is the included object in the response.",
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "shared.RelationCollection": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is list of Key objects.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.Key"
                    }
                }
            }
        },
        "shared.TokenAttributes": {
            "type": "object",
            "required": [
                "name",
                "symbol",
                "token_type"
            ],
            "properties": {
                "icon": {
                    "description": "Icon is the icon of the token.",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the token.",
                    "type": "string"
                },
                "symbol": {
                    "description": "Symbol is the symbol of the token.",
                    "type": "string"
                },
                "token_type": {
                    "description": "TokenType is the type of the token.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/data.TokenType"
                        }
                    ]
                }
            }
        },
        "shared.TokenListResponse": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the list of tokens.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_quantum-bridge_core_cmd_service_shared.Token"
                    }
                },
                "included": {
                    "description": "Included is the included object in the response.",
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "shared.TokenRelationships": {
            "type": "object",
            "required": [
                "chains"
            ],
            "properties": {
                "chains": {
                    "description": "Chains is the chains that are used by the token.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.RelationCollection"
                        }
                    ]
                }
            }
        },
        "shared.TransactionsResponse": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "Data is the body of transaction that is returned in the response."
                },
                "included": {
                    "description": "Included is the included chain data of the response.",
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{"http"},
	Title:            "Core bridge API",
	Description:      "Core bridge API is a service that responsible for the communication between blockchains.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
