// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cksaleabi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// CKSaleABI is the input ABI used to generate the binding from.
const CKSaleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastGen0SalePrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSaleClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0SaleCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"averageGen0SalePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// CKSale is an auto generated Go binding around an Ethereum contract.
type CKSale struct {
	CKSaleCaller     // Read-only binding to the contract
	CKSaleTransactor // Write-only binding to the contract
	CKSaleFilterer   // Log filterer for contract events
}

// CKSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CKSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CKSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CKSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CKSaleSession struct {
	Contract     *CKSale           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CKSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CKSaleCallerSession struct {
	Contract *CKSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CKSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CKSaleTransactorSession struct {
	Contract     *CKSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CKSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CKSaleRaw struct {
	Contract *CKSale // Generic contract binding to access the raw methods on
}

// CKSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CKSaleCallerRaw struct {
	Contract *CKSaleCaller // Generic read-only contract binding to access the raw methods on
}

// CKSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CKSaleTransactorRaw struct {
	Contract *CKSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCKSale creates a new instance of CKSale, bound to a specific deployed contract.
func NewCKSale(address common.Address, backend bind.ContractBackend) (*CKSale, error) {
	contract, err := bindCKSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CKSale{CKSaleCaller: CKSaleCaller{contract: contract}, CKSaleTransactor: CKSaleTransactor{contract: contract}, CKSaleFilterer: CKSaleFilterer{contract: contract}}, nil
}

// NewCKSaleCaller creates a new read-only instance of CKSale, bound to a specific deployed contract.
func NewCKSaleCaller(address common.Address, caller bind.ContractCaller) (*CKSaleCaller, error) {
	contract, err := bindCKSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CKSaleCaller{contract: contract}, nil
}

// NewCKSaleTransactor creates a new write-only instance of CKSale, bound to a specific deployed contract.
func NewCKSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CKSaleTransactor, error) {
	contract, err := bindCKSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CKSaleTransactor{contract: contract}, nil
}

// NewCKSaleFilterer creates a new log filterer instance of CKSale, bound to a specific deployed contract.
func NewCKSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CKSaleFilterer, error) {
	contract, err := bindCKSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CKSaleFilterer{contract: contract}, nil
}

// bindCKSale binds a generic wrapper to an already deployed contract.
func bindCKSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CKSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CKSale *CKSaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CKSale.Contract.CKSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CKSale *CKSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKSale.Contract.CKSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CKSale *CKSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CKSale.Contract.CKSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CKSale *CKSaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CKSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CKSale *CKSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CKSale *CKSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CKSale.Contract.contract.Transact(opts, method, params...)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_CKSale *CKSaleCaller) AverageGen0SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "averageGen0SalePrice")
	return *ret0, err
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_CKSale *CKSaleSession) AverageGen0SalePrice() (*big.Int, error) {
	return _CKSale.Contract.AverageGen0SalePrice(&_CKSale.CallOpts)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() constant returns(uint256)
func (_CKSale *CKSaleCallerSession) AverageGen0SalePrice() (*big.Int, error) {
	return _CKSale.Contract.AverageGen0SalePrice(&_CKSale.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_CKSale *CKSaleCaller) Gen0SaleCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "gen0SaleCount")
	return *ret0, err
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_CKSale *CKSaleSession) Gen0SaleCount() (*big.Int, error) {
	return _CKSale.Contract.Gen0SaleCount(&_CKSale.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() constant returns(uint256)
func (_CKSale *CKSaleCallerSession) Gen0SaleCount() (*big.Int, error) {
	return _CKSale.Contract.Gen0SaleCount(&_CKSale.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_CKSale *CKSaleCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	ret := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	out := ret
	err := _CKSale.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_CKSale *CKSaleSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _CKSale.Contract.GetAuction(&_CKSale.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(_tokenId uint256) constant returns(seller address, startingPrice uint256, endingPrice uint256, duration uint256, startedAt uint256)
func (_CKSale *CKSaleCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _CKSale.Contract.GetAuction(&_CKSale.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_CKSale *CKSaleCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_CKSale *CKSaleSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _CKSale.Contract.GetCurrentPrice(&_CKSale.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(_tokenId uint256) constant returns(uint256)
func (_CKSale *CKSaleCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _CKSale.Contract.GetCurrentPrice(&_CKSale.CallOpts, _tokenId)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_CKSale *CKSaleCaller) IsSaleClockAuction(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "isSaleClockAuction")
	return *ret0, err
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_CKSale *CKSaleSession) IsSaleClockAuction() (bool, error) {
	return _CKSale.Contract.IsSaleClockAuction(&_CKSale.CallOpts)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() constant returns(bool)
func (_CKSale *CKSaleCallerSession) IsSaleClockAuction() (bool, error) {
	return _CKSale.Contract.IsSaleClockAuction(&_CKSale.CallOpts)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_CKSale *CKSaleCaller) LastGen0SalePrices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "lastGen0SalePrices", arg0)
	return *ret0, err
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_CKSale *CKSaleSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _CKSale.Contract.LastGen0SalePrices(&_CKSale.CallOpts, arg0)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices( uint256) constant returns(uint256)
func (_CKSale *CKSaleCallerSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _CKSale.Contract.LastGen0SalePrices(&_CKSale.CallOpts, arg0)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_CKSale *CKSaleCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_CKSale *CKSaleSession) NonFungibleContract() (common.Address, error) {
	return _CKSale.Contract.NonFungibleContract(&_CKSale.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_CKSale *CKSaleCallerSession) NonFungibleContract() (common.Address, error) {
	return _CKSale.Contract.NonFungibleContract(&_CKSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CKSale *CKSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CKSale *CKSaleSession) Owner() (common.Address, error) {
	return _CKSale.Contract.Owner(&_CKSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CKSale *CKSaleCallerSession) Owner() (common.Address, error) {
	return _CKSale.Contract.Owner(&_CKSale.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_CKSale *CKSaleCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_CKSale *CKSaleSession) OwnerCut() (*big.Int, error) {
	return _CKSale.Contract.OwnerCut(&_CKSale.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_CKSale *CKSaleCallerSession) OwnerCut() (*big.Int, error) {
	return _CKSale.Contract.OwnerCut(&_CKSale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CKSale *CKSaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CKSale.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CKSale *CKSaleSession) Paused() (bool, error) {
	return _CKSale.Contract.Paused(&_CKSale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CKSale *CKSaleCallerSession) Paused() (bool, error) {
	return _CKSale.Contract.Paused(&_CKSale.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_CKSale *CKSaleTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_CKSale *CKSaleSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.Contract.Bid(&_CKSale.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(_tokenId uint256) returns()
func (_CKSale *CKSaleTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.Contract.Bid(&_CKSale.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_CKSale *CKSaleTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_CKSale *CKSaleSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.Contract.CancelAuction(&_CKSale.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(_tokenId uint256) returns()
func (_CKSale *CKSaleTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.Contract.CancelAuction(&_CKSale.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_CKSale *CKSaleTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_CKSale *CKSaleSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.Contract.CancelAuctionWhenPaused(&_CKSale.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(_tokenId uint256) returns()
func (_CKSale *CKSaleTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _CKSale.Contract.CancelAuctionWhenPaused(&_CKSale.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_CKSale *CKSaleTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_CKSale *CKSaleSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _CKSale.Contract.CreateAuction(&_CKSale.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(_tokenId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256, _seller address) returns()
func (_CKSale *CKSaleTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _CKSale.Contract.CreateAuction(&_CKSale.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_CKSale *CKSaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_CKSale *CKSaleSession) Pause() (*types.Transaction, error) {
	return _CKSale.Contract.Pause(&_CKSale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_CKSale *CKSaleTransactorSession) Pause() (*types.Transaction, error) {
	return _CKSale.Contract.Pause(&_CKSale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CKSale *CKSaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CKSale *CKSaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CKSale.Contract.TransferOwnership(&_CKSale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CKSale *CKSaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CKSale.Contract.TransferOwnership(&_CKSale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_CKSale *CKSaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_CKSale *CKSaleSession) Unpause() (*types.Transaction, error) {
	return _CKSale.Contract.Unpause(&_CKSale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_CKSale *CKSaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _CKSale.Contract.Unpause(&_CKSale.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CKSale *CKSaleTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKSale.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CKSale *CKSaleSession) WithdrawBalance() (*types.Transaction, error) {
	return _CKSale.Contract.WithdrawBalance(&_CKSale.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CKSale *CKSaleTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _CKSale.Contract.WithdrawBalance(&_CKSale.TransactOpts)
}

// CKSaleAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the CKSale contract.
type CKSaleAuctionCancelledIterator struct {
	Event *CKSaleAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *CKSaleAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKSaleAuctionCancelled)
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
		it.Event = new(CKSaleAuctionCancelled)
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
func (it *CKSaleAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKSaleAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKSaleAuctionCancelled represents a AuctionCancelled event raised by the CKSale contract.
type CKSaleAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(tokenId uint256)
func (_CKSale *CKSaleFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*CKSaleAuctionCancelledIterator, error) {

	logs, sub, err := _CKSale.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &CKSaleAuctionCancelledIterator{contract: _CKSale.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(tokenId uint256)
func (_CKSale *CKSaleFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *CKSaleAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _CKSale.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKSaleAuctionCancelled)
				if err := _CKSale.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// CKSaleAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the CKSale contract.
type CKSaleAuctionCreatedIterator struct {
	Event *CKSaleAuctionCreated // Event containing the contract specifics and raw log

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
func (it *CKSaleAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKSaleAuctionCreated)
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
		it.Event = new(CKSaleAuctionCreated)
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
func (it *CKSaleAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKSaleAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKSaleAuctionCreated represents a AuctionCreated event raised by the CKSale contract.
type CKSaleAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_CKSale *CKSaleFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*CKSaleAuctionCreatedIterator, error) {

	logs, sub, err := _CKSale.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &CKSaleAuctionCreatedIterator{contract: _CKSale.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(tokenId uint256, startingPrice uint256, endingPrice uint256, duration uint256)
func (_CKSale *CKSaleFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *CKSaleAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _CKSale.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKSaleAuctionCreated)
				if err := _CKSale.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// CKSaleAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the CKSale contract.
type CKSaleAuctionSuccessfulIterator struct {
	Event *CKSaleAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *CKSaleAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKSaleAuctionSuccessful)
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
		it.Event = new(CKSaleAuctionSuccessful)
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
func (it *CKSaleAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKSaleAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKSaleAuctionSuccessful represents a AuctionSuccessful event raised by the CKSale contract.
type CKSaleAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_CKSale *CKSaleFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*CKSaleAuctionSuccessfulIterator, error) {

	logs, sub, err := _CKSale.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &CKSaleAuctionSuccessfulIterator{contract: _CKSale.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(tokenId uint256, totalPrice uint256, winner address)
func (_CKSale *CKSaleFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *CKSaleAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _CKSale.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKSaleAuctionSuccessful)
				if err := _CKSale.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// CKSalePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the CKSale contract.
type CKSalePauseIterator struct {
	Event *CKSalePause // Event containing the contract specifics and raw log

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
func (it *CKSalePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKSalePause)
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
		it.Event = new(CKSalePause)
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
func (it *CKSalePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKSalePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKSalePause represents a Pause event raised by the CKSale contract.
type CKSalePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CKSale *CKSaleFilterer) FilterPause(opts *bind.FilterOpts) (*CKSalePauseIterator, error) {

	logs, sub, err := _CKSale.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CKSalePauseIterator{contract: _CKSale.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CKSale *CKSaleFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CKSalePause) (event.Subscription, error) {

	logs, sub, err := _CKSale.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKSalePause)
				if err := _CKSale.contract.UnpackLog(event, "Pause", log); err != nil {
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

// CKSaleUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the CKSale contract.
type CKSaleUnpauseIterator struct {
	Event *CKSaleUnpause // Event containing the contract specifics and raw log

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
func (it *CKSaleUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKSaleUnpause)
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
		it.Event = new(CKSaleUnpause)
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
func (it *CKSaleUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKSaleUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKSaleUnpause represents a Unpause event raised by the CKSale contract.
type CKSaleUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CKSale *CKSaleFilterer) FilterUnpause(opts *bind.FilterOpts) (*CKSaleUnpauseIterator, error) {

	logs, sub, err := _CKSale.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CKSaleUnpauseIterator{contract: _CKSale.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CKSale *CKSaleFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CKSaleUnpause) (event.Subscription, error) {

	logs, sub, err := _CKSale.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKSaleUnpause)
				if err := _CKSale.contract.UnpackLog(event, "Unpause", log); err != nil {
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
