// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm

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
)

// ZetaNonEthMetaData contains all meta data concerning the ZetaNonEth contract.
var ZetaNonEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tssAddressUpdater_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotConnector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZetaTransferError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burnee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceTssAddressUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tssAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"connectorAddress_\",\"type\":\"address\"}],\"name\":\"updateTssAndConnectorAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZetaNonEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaNonEthMetaData.ABI instead.
var ZetaNonEthABI = ZetaNonEthMetaData.ABI

// ZetaNonEth is an auto generated Go binding around an Ethereum contract.
type ZetaNonEth struct {
	ZetaNonEthCaller     // Read-only binding to the contract
	ZetaNonEthTransactor // Write-only binding to the contract
	ZetaNonEthFilterer   // Log filterer for contract events
}

// ZetaNonEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaNonEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaNonEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaNonEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaNonEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaNonEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaNonEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaNonEthSession struct {
	Contract     *ZetaNonEth       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaNonEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaNonEthCallerSession struct {
	Contract *ZetaNonEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZetaNonEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaNonEthTransactorSession struct {
	Contract     *ZetaNonEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZetaNonEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaNonEthRaw struct {
	Contract *ZetaNonEth // Generic contract binding to access the raw methods on
}

// ZetaNonEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaNonEthCallerRaw struct {
	Contract *ZetaNonEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaNonEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaNonEthTransactorRaw struct {
	Contract *ZetaNonEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaNonEth creates a new instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEth(address common.Address, backend bind.ContractBackend) (*ZetaNonEth, error) {
	contract, err := bindZetaNonEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEth{ZetaNonEthCaller: ZetaNonEthCaller{contract: contract}, ZetaNonEthTransactor: ZetaNonEthTransactor{contract: contract}, ZetaNonEthFilterer: ZetaNonEthFilterer{contract: contract}}, nil
}

// NewZetaNonEthCaller creates a new read-only instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaNonEthCaller, error) {
	contract, err := bindZetaNonEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthCaller{contract: contract}, nil
}

// NewZetaNonEthTransactor creates a new write-only instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaNonEthTransactor, error) {
	contract, err := bindZetaNonEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthTransactor{contract: contract}, nil
}

// NewZetaNonEthFilterer creates a new log filterer instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaNonEthFilterer, error) {
	contract, err := bindZetaNonEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthFilterer{contract: contract}, nil
}

// bindZetaNonEth binds a generic wrapper to an already deployed contract.
func bindZetaNonEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZetaNonEthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaNonEth *ZetaNonEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaNonEth.Contract.ZetaNonEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaNonEth *ZetaNonEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.ZetaNonEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaNonEth *ZetaNonEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.ZetaNonEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaNonEth *ZetaNonEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaNonEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaNonEth *ZetaNonEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaNonEth *ZetaNonEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.Allowance(&_ZetaNonEth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.Allowance(&_ZetaNonEth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.BalanceOf(&_ZetaNonEth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.BalanceOf(&_ZetaNonEth.CallOpts, account)
}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCaller) ConnectorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "connectorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthSession) ConnectorAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.ConnectorAddress(&_ZetaNonEth.CallOpts)
}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCallerSession) ConnectorAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.ConnectorAddress(&_ZetaNonEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaNonEth *ZetaNonEthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaNonEth *ZetaNonEthSession) Decimals() (uint8, error) {
	return _ZetaNonEth.Contract.Decimals(&_ZetaNonEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaNonEth *ZetaNonEthCallerSession) Decimals() (uint8, error) {
	return _ZetaNonEth.Contract.Decimals(&_ZetaNonEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaNonEth *ZetaNonEthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaNonEth *ZetaNonEthSession) Name() (string, error) {
	return _ZetaNonEth.Contract.Name(&_ZetaNonEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaNonEth *ZetaNonEthCallerSession) Name() (string, error) {
	return _ZetaNonEth.Contract.Name(&_ZetaNonEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaNonEth *ZetaNonEthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaNonEth *ZetaNonEthSession) Symbol() (string, error) {
	return _ZetaNonEth.Contract.Symbol(&_ZetaNonEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaNonEth *ZetaNonEthCallerSession) Symbol() (string, error) {
	return _ZetaNonEth.Contract.Symbol(&_ZetaNonEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaNonEth *ZetaNonEthSession) TotalSupply() (*big.Int, error) {
	return _ZetaNonEth.Contract.TotalSupply(&_ZetaNonEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCallerSession) TotalSupply() (*big.Int, error) {
	return _ZetaNonEth.Contract.TotalSupply(&_ZetaNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthSession) TssAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddress(&_ZetaNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCallerSession) TssAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddress(&_ZetaNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaNonEth *ZetaNonEthCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaNonEth *ZetaNonEthSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddressUpdater(&_ZetaNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaNonEth *ZetaNonEthCallerSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddressUpdater(&_ZetaNonEth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Approve(&_ZetaNonEth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Approve(&_ZetaNonEth.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Burn(&_ZetaNonEth.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Burn(&_ZetaNonEth.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.BurnFrom(&_ZetaNonEth.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.BurnFrom(&_ZetaNonEth.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.DecreaseAllowance(&_ZetaNonEth.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.DecreaseAllowance(&_ZetaNonEth.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.IncreaseAllowance(&_ZetaNonEth.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.IncreaseAllowance(&_ZetaNonEth.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_ZetaNonEth *ZetaNonEthSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Mint(&_ZetaNonEth.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Mint(&_ZetaNonEth.TransactOpts, mintee, value, internalSendHash)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaNonEth *ZetaNonEthTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaNonEth *ZetaNonEthSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaNonEth.Contract.RenounceTssAddressUpdater(&_ZetaNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaNonEth.Contract.RenounceTssAddressUpdater(&_ZetaNonEth.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Transfer(&_ZetaNonEth.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Transfer(&_ZetaNonEth.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.TransferFrom(&_ZetaNonEth.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.TransferFrom(&_ZetaNonEth.TransactOpts, from, to, amount)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) UpdateTssAndConnectorAddresses(opts *bind.TransactOpts, tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "updateTssAndConnectorAddresses", tssAddress_, connectorAddress_)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_ZetaNonEth *ZetaNonEthSession) UpdateTssAndConnectorAddresses(tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.UpdateTssAndConnectorAddresses(&_ZetaNonEth.TransactOpts, tssAddress_, connectorAddress_)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) UpdateTssAndConnectorAddresses(tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.UpdateTssAndConnectorAddresses(&_ZetaNonEth.TransactOpts, tssAddress_, connectorAddress_)
}

// ZetaNonEthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZetaNonEth contract.
type ZetaNonEthApprovalIterator struct {
	Event *ZetaNonEthApproval // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthApproval)
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
		it.Event = new(ZetaNonEthApproval)
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
func (it *ZetaNonEthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthApproval represents a Approval event raised by the ZetaNonEth contract.
type ZetaNonEthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZetaNonEthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthApprovalIterator{contract: _ZetaNonEth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZetaNonEthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthApproval)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseApproval(log types.Log) (*ZetaNonEthApproval, error) {
	event := new(ZetaNonEthApproval)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthBurntIterator is returned from FilterBurnt and is used to iterate over the raw logs and unpacked data for Burnt events raised by the ZetaNonEth contract.
type ZetaNonEthBurntIterator struct {
	Event *ZetaNonEthBurnt // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthBurnt)
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
		it.Event = new(ZetaNonEthBurnt)
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
func (it *ZetaNonEthBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthBurnt represents a Burnt event raised by the ZetaNonEth contract.
type ZetaNonEthBurnt struct {
	Burnee common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnt is a free log retrieval operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterBurnt(opts *bind.FilterOpts, burnee []common.Address) (*ZetaNonEthBurntIterator, error) {

	var burneeRule []interface{}
	for _, burneeItem := range burnee {
		burneeRule = append(burneeRule, burneeItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Burnt", burneeRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthBurntIterator{contract: _ZetaNonEth.contract, event: "Burnt", logs: logs, sub: sub}, nil
}

// WatchBurnt is a free log subscription operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchBurnt(opts *bind.WatchOpts, sink chan<- *ZetaNonEthBurnt, burnee []common.Address) (event.Subscription, error) {

	var burneeRule []interface{}
	for _, burneeItem := range burnee {
		burneeRule = append(burneeRule, burneeItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Burnt", burneeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthBurnt)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Burnt", log); err != nil {
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

// ParseBurnt is a log parse operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseBurnt(log types.Log) (*ZetaNonEthBurnt, error) {
	event := new(ZetaNonEthBurnt)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Burnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the ZetaNonEth contract.
type ZetaNonEthMintedIterator struct {
	Event *ZetaNonEthMinted // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthMinted)
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
		it.Event = new(ZetaNonEthMinted)
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
func (it *ZetaNonEthMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthMinted represents a Minted event raised by the ZetaNonEth contract.
type ZetaNonEthMinted struct {
	Mintee           common.Address
	Amount           *big.Int
	InternalSendHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterMinted(opts *bind.FilterOpts, mintee []common.Address, internalSendHash [][32]byte) (*ZetaNonEthMintedIterator, error) {

	var minteeRule []interface{}
	for _, minteeItem := range mintee {
		minteeRule = append(minteeRule, minteeItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Minted", minteeRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthMintedIterator{contract: _ZetaNonEth.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ZetaNonEthMinted, mintee []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var minteeRule []interface{}
	for _, minteeItem := range mintee {
		minteeRule = append(minteeRule, minteeItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Minted", minteeRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthMinted)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseMinted(log types.Log) (*ZetaNonEthMinted, error) {
	event := new(ZetaNonEthMinted)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZetaNonEth contract.
type ZetaNonEthTransferIterator struct {
	Event *ZetaNonEthTransfer // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthTransfer)
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
		it.Event = new(ZetaNonEthTransfer)
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
func (it *ZetaNonEthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthTransfer represents a Transfer event raised by the ZetaNonEth contract.
type ZetaNonEthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZetaNonEthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthTransferIterator{contract: _ZetaNonEth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZetaNonEthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthTransfer)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseTransfer(log types.Log) (*ZetaNonEthTransfer, error) {
	event := new(ZetaNonEthTransfer)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
