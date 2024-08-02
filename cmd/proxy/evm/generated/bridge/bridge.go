// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMintable\",\"type\":\"bool\"}],\"name\":\"DepositedERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMintable\",\"type\":\"bool\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMintable\",\"type\":\"bool\"}],\"name\":\"DepositedERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oracles_\",\"type\":\"address[]\"}],\"name\":\"addOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"}],\"name\":\"checkHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"to_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"to_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"}],\"name\":\"getERC1155SignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"}],\"name\":\"getERC20SignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"}],\"name\":\"getERC721SignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"oracles_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"thresholdOracleSignatures_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oracles_\",\"type\":\"address[]\"}],\"name\":\"removeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thresholdOracleSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"thresholdOracleSignatures_\",\"type\":\"uint256\"}],\"name\":\"updateThresholdSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMintable_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bridge *BridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bridge *BridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bridge.Contract.UPGRADEINTERFACEVERSION(&_Bridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bridge *BridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bridge.Contract.UPGRADEINTERFACEVERSION(&_Bridge.CallOpts)
}

// CheckHash is a free data retrieval call binding the contract method 0xe9c22c88.
//
// Solidity: function checkHash(bytes32 txHash_, uint256 txNonce_) view returns(bool)
func (_Bridge *BridgeCaller) CheckHash(opts *bind.CallOpts, txHash_ [32]byte, txNonce_ *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "checkHash", txHash_, txNonce_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckHash is a free data retrieval call binding the contract method 0xe9c22c88.
//
// Solidity: function checkHash(bytes32 txHash_, uint256 txNonce_) view returns(bool)
func (_Bridge *BridgeSession) CheckHash(txHash_ [32]byte, txNonce_ *big.Int) (bool, error) {
	return _Bridge.Contract.CheckHash(&_Bridge.CallOpts, txHash_, txNonce_)
}

// CheckHash is a free data retrieval call binding the contract method 0xe9c22c88.
//
// Solidity: function checkHash(bytes32 txHash_, uint256 txNonce_) view returns(bool)
func (_Bridge *BridgeCallerSession) CheckHash(txHash_ [32]byte, txNonce_ *big.Int) (bool, error) {
	return _Bridge.Contract.CheckHash(&_Bridge.CallOpts, txHash_, txNonce_)
}

// GetERC1155SignHash is a free data retrieval call binding the contract method 0xb427d67c.
//
// Solidity: function getERC1155SignHash(address token_, uint256 tokenId_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetERC1155SignHash(opts *bind.CallOpts, token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isMintable_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getERC1155SignHash", token_, tokenId_, amount_, to_, txHash_, txNonce_, chainId_, tokenURI_, isMintable_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC1155SignHash is a free data retrieval call binding the contract method 0xb427d67c.
//
// Solidity: function getERC1155SignHash(address token_, uint256 tokenId_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetERC1155SignHash(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isMintable_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC1155SignHash(&_Bridge.CallOpts, token_, tokenId_, amount_, to_, txHash_, txNonce_, chainId_, tokenURI_, isMintable_)
}

// GetERC1155SignHash is a free data retrieval call binding the contract method 0xb427d67c.
//
// Solidity: function getERC1155SignHash(address token_, uint256 tokenId_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetERC1155SignHash(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isMintable_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC1155SignHash(&_Bridge.CallOpts, token_, tokenId_, amount_, to_, txHash_, txNonce_, chainId_, tokenURI_, isMintable_)
}

// GetERC20SignHash is a free data retrieval call binding the contract method 0xaaba091e.
//
// Solidity: function getERC20SignHash(address token_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetERC20SignHash(opts *bind.CallOpts, token_ common.Address, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, isMintable_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getERC20SignHash", token_, amount_, to_, txHash_, txNonce_, chainId_, isMintable_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC20SignHash is a free data retrieval call binding the contract method 0xaaba091e.
//
// Solidity: function getERC20SignHash(address token_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetERC20SignHash(token_ common.Address, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, isMintable_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC20SignHash(&_Bridge.CallOpts, token_, amount_, to_, txHash_, txNonce_, chainId_, isMintable_)
}

// GetERC20SignHash is a free data retrieval call binding the contract method 0xaaba091e.
//
// Solidity: function getERC20SignHash(address token_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetERC20SignHash(token_ common.Address, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, isMintable_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC20SignHash(&_Bridge.CallOpts, token_, amount_, to_, txHash_, txNonce_, chainId_, isMintable_)
}

// GetERC721SignHash is a free data retrieval call binding the contract method 0xaf94570d.
//
// Solidity: function getERC721SignHash(address token_, uint256 tokenId_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetERC721SignHash(opts *bind.CallOpts, token_ common.Address, tokenId_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isMintable_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getERC721SignHash", token_, tokenId_, to_, txHash_, txNonce_, chainId_, tokenURI_, isMintable_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC721SignHash is a free data retrieval call binding the contract method 0xaf94570d.
//
// Solidity: function getERC721SignHash(address token_, uint256 tokenId_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetERC721SignHash(token_ common.Address, tokenId_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isMintable_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC721SignHash(&_Bridge.CallOpts, token_, tokenId_, to_, txHash_, txNonce_, chainId_, tokenURI_, isMintable_)
}

// GetERC721SignHash is a free data retrieval call binding the contract method 0xaf94570d.
//
// Solidity: function getERC721SignHash(address token_, uint256 tokenId_, address to_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isMintable_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetERC721SignHash(token_ common.Address, tokenId_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isMintable_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC721SignHash(&_Bridge.CallOpts, token_, tokenId_, to_, txHash_, txNonce_, chainId_, tokenURI_, isMintable_)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Bridge *BridgeCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Bridge *BridgeSession) GetOracles() ([]common.Address, error) {
	return _Bridge.Contract.GetOracles(&_Bridge.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Bridge *BridgeCallerSession) GetOracles() ([]common.Address, error) {
	return _Bridge.Contract.GetOracles(&_Bridge.CallOpts)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Hashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Hashes(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Hashes(&_Bridge.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Hashes(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Hashes(&_Bridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeSession) ProxiableUUID() ([32]byte, error) {
	return _Bridge.Contract.ProxiableUUID(&_Bridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bridge.Contract.ProxiableUUID(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// ThresholdOracleSignatures is a free data retrieval call binding the contract method 0x1d28dc86.
//
// Solidity: function thresholdOracleSignatures() view returns(uint256)
func (_Bridge *BridgeCaller) ThresholdOracleSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "thresholdOracleSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThresholdOracleSignatures is a free data retrieval call binding the contract method 0x1d28dc86.
//
// Solidity: function thresholdOracleSignatures() view returns(uint256)
func (_Bridge *BridgeSession) ThresholdOracleSignatures() (*big.Int, error) {
	return _Bridge.Contract.ThresholdOracleSignatures(&_Bridge.CallOpts)
}

// ThresholdOracleSignatures is a free data retrieval call binding the contract method 0x1d28dc86.
//
// Solidity: function thresholdOracleSignatures() view returns(uint256)
func (_Bridge *BridgeCallerSession) ThresholdOracleSignatures() (*big.Int, error) {
	return _Bridge.Contract.ThresholdOracleSignatures(&_Bridge.CallOpts)
}

// AddOracles is a paid mutator transaction binding the contract method 0x205b931e.
//
// Solidity: function addOracles(address[] oracles_) returns()
func (_Bridge *BridgeTransactor) AddOracles(opts *bind.TransactOpts, oracles_ []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addOracles", oracles_)
}

// AddOracles is a paid mutator transaction binding the contract method 0x205b931e.
//
// Solidity: function addOracles(address[] oracles_) returns()
func (_Bridge *BridgeSession) AddOracles(oracles_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddOracles(&_Bridge.TransactOpts, oracles_)
}

// AddOracles is a paid mutator transaction binding the contract method 0x205b931e.
//
// Solidity: function addOracles(address[] oracles_) returns()
func (_Bridge *BridgeTransactorSession) AddOracles(oracles_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddOracles(&_Bridge.TransactOpts, oracles_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x1270ce5a.
//
// Solidity: function depositERC1155(address token_, uint256 tokenId_, uint256 amount_, string to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeTransactor) DepositERC1155(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ string, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC1155", token_, tokenId_, amount_, to_, network_, isMintable_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x1270ce5a.
//
// Solidity: function depositERC1155(address token_, uint256 tokenId_, uint256 amount_, string to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeSession) DepositERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ string, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, to_, network_, isMintable_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x1270ce5a.
//
// Solidity: function depositERC1155(address token_, uint256 tokenId_, uint256 amount_, string to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ string, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, to_, network_, isMintable_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xfc0a9870.
//
// Solidity: function depositERC20(address token_, uint256 amount_, string to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeTransactor) DepositERC20(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, to_ string, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20", token_, amount_, to_, network_, isMintable_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xfc0a9870.
//
// Solidity: function depositERC20(address token_, uint256 amount_, string to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeSession) DepositERC20(token_ common.Address, amount_ *big.Int, to_ string, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, token_, amount_, to_, network_, isMintable_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xfc0a9870.
//
// Solidity: function depositERC20(address token_, uint256 amount_, string to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC20(token_ common.Address, amount_ *big.Int, to_ string, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, token_, amount_, to_, network_, isMintable_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xc0f5e247.
//
// Solidity: function depositERC721(address token_, uint256 tokenId_, address to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeTransactor) DepositERC721(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, to_ common.Address, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC721", token_, tokenId_, to_, network_, isMintable_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xc0f5e247.
//
// Solidity: function depositERC721(address token_, uint256 tokenId_, address to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeSession) DepositERC721(token_ common.Address, tokenId_ *big.Int, to_ common.Address, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, token_, tokenId_, to_, network_, isMintable_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xc0f5e247.
//
// Solidity: function depositERC721(address token_, uint256 tokenId_, address to_, string network_, bool isMintable_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC721(token_ common.Address, tokenId_ *big.Int, to_ common.Address, network_ string, isMintable_ bool) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, token_, tokenId_, to_, network_, isMintable_)
}

// Initialize is a paid mutator transaction binding the contract method 0x3ede50c6.
//
// Solidity: function initialize(address initialOwner_, address[] oracles_, uint256 thresholdOracleSignatures_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, initialOwner_ common.Address, oracles_ []common.Address, thresholdOracleSignatures_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", initialOwner_, oracles_, thresholdOracleSignatures_)
}

// Initialize is a paid mutator transaction binding the contract method 0x3ede50c6.
//
// Solidity: function initialize(address initialOwner_, address[] oracles_, uint256 thresholdOracleSignatures_) returns()
func (_Bridge *BridgeSession) Initialize(initialOwner_ common.Address, oracles_ []common.Address, thresholdOracleSignatures_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, initialOwner_, oracles_, thresholdOracleSignatures_)
}

// Initialize is a paid mutator transaction binding the contract method 0x3ede50c6.
//
// Solidity: function initialize(address initialOwner_, address[] oracles_, uint256 thresholdOracleSignatures_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(initialOwner_ common.Address, oracles_ []common.Address, thresholdOracleSignatures_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, initialOwner_, oracles_, thresholdOracleSignatures_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0x45644fd6.
//
// Solidity: function removeOracles(address[] oracles_) returns()
func (_Bridge *BridgeTransactor) RemoveOracles(opts *bind.TransactOpts, oracles_ []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeOracles", oracles_)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0x45644fd6.
//
// Solidity: function removeOracles(address[] oracles_) returns()
func (_Bridge *BridgeSession) RemoveOracles(oracles_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveOracles(&_Bridge.TransactOpts, oracles_)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0x45644fd6.
//
// Solidity: function removeOracles(address[] oracles_) returns()
func (_Bridge *BridgeTransactorSession) RemoveOracles(oracles_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveOracles(&_Bridge.TransactOpts, oracles_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// UpdateThresholdSignatures is a paid mutator transaction binding the contract method 0x60dbac34.
//
// Solidity: function updateThresholdSignatures(uint256 thresholdOracleSignatures_) returns()
func (_Bridge *BridgeTransactor) UpdateThresholdSignatures(opts *bind.TransactOpts, thresholdOracleSignatures_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateThresholdSignatures", thresholdOracleSignatures_)
}

// UpdateThresholdSignatures is a paid mutator transaction binding the contract method 0x60dbac34.
//
// Solidity: function updateThresholdSignatures(uint256 thresholdOracleSignatures_) returns()
func (_Bridge *BridgeSession) UpdateThresholdSignatures(thresholdOracleSignatures_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateThresholdSignatures(&_Bridge.TransactOpts, thresholdOracleSignatures_)
}

// UpdateThresholdSignatures is a paid mutator transaction binding the contract method 0x60dbac34.
//
// Solidity: function updateThresholdSignatures(uint256 thresholdOracleSignatures_) returns()
func (_Bridge *BridgeTransactorSession) UpdateThresholdSignatures(thresholdOracleSignatures_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateThresholdSignatures(&_Bridge.TransactOpts, thresholdOracleSignatures_)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToAndCall(&_Bridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToAndCall(&_Bridge.TransactOpts, newImplementation, data)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xb3953d44.
//
// Solidity: function withdrawERC1155(address token_, uint256 tokenId_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC1155(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC1155", token_, tokenId_, amount_, to_, txHash_, txNonce_, tokenURI_, isMintable_, signatures_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xb3953d44.
//
// Solidity: function withdrawERC1155(address token_, uint256 tokenId_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, to_, txHash_, txNonce_, tokenURI_, isMintable_, signatures_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xb3953d44.
//
// Solidity: function withdrawERC1155(address token_, uint256 tokenId_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, to_, txHash_, txNonce_, tokenURI_, isMintable_, signatures_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x94995fc4.
//
// Solidity: function withdrawERC20(address token_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC20(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20", token_, amount_, to_, txHash_, txNonce_, isMintable_, signatures_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x94995fc4.
//
// Solidity: function withdrawERC20(address token_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC20(token_ common.Address, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, token_, amount_, to_, txHash_, txNonce_, isMintable_, signatures_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x94995fc4.
//
// Solidity: function withdrawERC20(address token_, uint256 amount_, address to_, bytes32 txHash_, uint256 txNonce_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC20(token_ common.Address, amount_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, token_, amount_, to_, txHash_, txNonce_, isMintable_, signatures_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x7eb9d447.
//
// Solidity: function withdrawERC721(address token_, uint256 tokenId_, address to_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC721(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC721", token_, tokenId_, to_, txHash_, txNonce_, tokenURI_, isMintable_, signatures_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x7eb9d447.
//
// Solidity: function withdrawERC721(address token_, uint256 tokenId_, address to_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC721(token_ common.Address, tokenId_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, token_, tokenId_, to_, txHash_, txNonce_, tokenURI_, isMintable_, signatures_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x7eb9d447.
//
// Solidity: function withdrawERC721(address token_, uint256 tokenId_, address to_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isMintable_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC721(token_ common.Address, tokenId_ *big.Int, to_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isMintable_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, token_, tokenId_, to_, txHash_, txNonce_, tokenURI_, isMintable_, signatures_)
}

// BridgeDepositedERC1155Iterator is returned from FilterDepositedERC1155 and is used to iterate over the raw logs and unpacked data for DepositedERC1155 events raised by the Bridge contract.
type BridgeDepositedERC1155Iterator struct {
	Event *BridgeDepositedERC1155 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeDepositedERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC1155)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeDepositedERC1155)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeDepositedERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC1155 represents a DepositedERC1155 event raised by the Bridge contract.
type BridgeDepositedERC1155 struct {
	Token      common.Address
	TokenId    *big.Int
	Amount     *big.Int
	To         string
	Network    string
	IsMintable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC1155 is a free log retrieval operation binding the contract event 0xc4bbd2c4c89eca35a4569ba4be47a2ecd453b39abc66de1c93920b8977224c87.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, string to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) FilterDepositedERC1155(opts *bind.FilterOpts) (*BridgeDepositedERC1155Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC1155Iterator{contract: _Bridge.contract, event: "DepositedERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositedERC1155 is a free log subscription operation binding the contract event 0xc4bbd2c4c89eca35a4569ba4be47a2ecd453b39abc66de1c93920b8977224c87.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, string to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) WatchDepositedERC1155(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC1155) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC1155)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedERC1155 is a log parse operation binding the contract event 0xc4bbd2c4c89eca35a4569ba4be47a2ecd453b39abc66de1c93920b8977224c87.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, string to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) ParseDepositedERC1155(log types.Log) (*BridgeDepositedERC1155, error) {
	event := new(BridgeDepositedERC1155)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the Bridge contract.
type BridgeDepositedERC20Iterator struct {
	Event *BridgeDepositedERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeDepositedERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC20 represents a DepositedERC20 event raised by the Bridge contract.
type BridgeDepositedERC20 struct {
	Token      common.Address
	Amount     *big.Int
	To         string
	Network    string
	IsMintable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC20 is a free log retrieval operation binding the contract event 0xe45ffec1054683dac57438f7878191986caad7d21a6090f93d132031dce00be0.
//
// Solidity: event DepositedERC20(address indexed token, uint256 amount, string to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) FilterDepositedERC20(opts *bind.FilterOpts, token []common.Address) (*BridgeDepositedERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC20", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC20Iterator{contract: _Bridge.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0xe45ffec1054683dac57438f7878191986caad7d21a6090f93d132031dce00be0.
//
// Solidity: event DepositedERC20(address indexed token, uint256 amount, string to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC20, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC20", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC20)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedERC20 is a log parse operation binding the contract event 0xe45ffec1054683dac57438f7878191986caad7d21a6090f93d132031dce00be0.
//
// Solidity: event DepositedERC20(address indexed token, uint256 amount, string to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) ParseDepositedERC20(log types.Log) (*BridgeDepositedERC20, error) {
	event := new(BridgeDepositedERC20)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC721Iterator is returned from FilterDepositedERC721 and is used to iterate over the raw logs and unpacked data for DepositedERC721 events raised by the Bridge contract.
type BridgeDepositedERC721Iterator struct {
	Event *BridgeDepositedERC721 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeDepositedERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC721)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeDepositedERC721)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeDepositedERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC721 represents a DepositedERC721 event raised by the Bridge contract.
type BridgeDepositedERC721 struct {
	Token      common.Address
	TokenId    *big.Int
	To         common.Address
	Network    string
	IsMintable bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC721 is a free log retrieval operation binding the contract event 0x9a60b2a4bebc9d9514117778b942585428fcffcf1eecd58c27056e0844b83aab.
//
// Solidity: event DepositedERC721(address indexed token, uint256 tokenId, address to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) FilterDepositedERC721(opts *bind.FilterOpts, token []common.Address) (*BridgeDepositedERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC721", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC721Iterator{contract: _Bridge.contract, event: "DepositedERC721", logs: logs, sub: sub}, nil
}

// WatchDepositedERC721 is a free log subscription operation binding the contract event 0x9a60b2a4bebc9d9514117778b942585428fcffcf1eecd58c27056e0844b83aab.
//
// Solidity: event DepositedERC721(address indexed token, uint256 tokenId, address to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) WatchDepositedERC721(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC721, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC721", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC721)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedERC721 is a log parse operation binding the contract event 0x9a60b2a4bebc9d9514117778b942585428fcffcf1eecd58c27056e0844b83aab.
//
// Solidity: event DepositedERC721(address indexed token, uint256 tokenId, address to, string network, bool isMintable)
func (_Bridge *BridgeFilterer) ParseDepositedERC721(log types.Log) (*BridgeDepositedERC721, error) {
	event := new(BridgeDepositedERC721)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bridge contract.
type BridgeUpgradedIterator struct {
	Event *BridgeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUpgraded represents a Upgraded event raised by the Bridge contract.
type BridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUpgradedIterator{contract: _Bridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUpgraded)
				if err := _Bridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) ParseUpgraded(log types.Log) (*BridgeUpgraded, error) {
	event := new(BridgeUpgraded)
	if err := _Bridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
