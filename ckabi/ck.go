// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ckabi

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

// CKABI is the input ABI used to generate the binding from.
const CKABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPromoKitty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v2Address\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"createGen0Auction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getKitty\",\"outputs\":[{\"name\":\"isGestating\",\"type\":\"bool\"},{\"name\":\"isReady\",\"type\":\"bool\"},{\"name\":\"cooldownIndex\",\"type\":\"uint256\"},{\"name\":\"nextActionAt\",\"type\":\"uint256\"},{\"name\":\"siringWithId\",\"type\":\"uint256\"},{\"name\":\"birthTime\",\"type\":\"uint256\"},{\"name\":\"matronId\",\"type\":\"uint256\"},{\"name\":\"sireId\",\"type\":\"uint256\"},{\"name\":\"generation\",\"type\":\"uint256\"},{\"name\":\"genes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0CreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// CK is an auto generated Go binding around an Ethereum contract.
type CK struct {
	CKCaller     // Read-only binding to the contract
	CKTransactor // Write-only binding to the contract
	CKFilterer   // Log filterer for contract events
}

// CKCaller is an auto generated read-only Go binding around an Ethereum contract.
type CKCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CKTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CKFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CKSession struct {
	Contract     *CK               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CKCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CKCallerSession struct {
	Contract *CKCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CKTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CKTransactorSession struct {
	Contract     *CKTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CKRaw is an auto generated low-level Go binding around an Ethereum contract.
type CKRaw struct {
	Contract *CK // Generic contract binding to access the raw methods on
}

// CKCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CKCallerRaw struct {
	Contract *CKCaller // Generic read-only contract binding to access the raw methods on
}

// CKTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CKTransactorRaw struct {
	Contract *CKTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCK creates a new instance of CK, bound to a specific deployed contract.
func NewCK(address common.Address, backend bind.ContractBackend) (*CK, error) {
	contract, err := bindCK(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CK{CKCaller: CKCaller{contract: contract}, CKTransactor: CKTransactor{contract: contract}, CKFilterer: CKFilterer{contract: contract}}, nil
}

// NewCKCaller creates a new read-only instance of CK, bound to a specific deployed contract.
func NewCKCaller(address common.Address, caller bind.ContractCaller) (*CKCaller, error) {
	contract, err := bindCK(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CKCaller{contract: contract}, nil
}

// NewCKTransactor creates a new write-only instance of CK, bound to a specific deployed contract.
func NewCKTransactor(address common.Address, transactor bind.ContractTransactor) (*CKTransactor, error) {
	contract, err := bindCK(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CKTransactor{contract: contract}, nil
}

// NewCKFilterer creates a new log filterer instance of CK, bound to a specific deployed contract.
func NewCKFilterer(address common.Address, filterer bind.ContractFilterer) (*CKFilterer, error) {
	contract, err := bindCK(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CKFilterer{contract: contract}, nil
}

// bindCK binds a generic wrapper to an already deployed contract.
func bindCK(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CKABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CK *CKRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CK.Contract.CKCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CK *CKRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CK.Contract.CKTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CK *CKRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CK.Contract.CKTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CK *CKCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CK.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CK *CKTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CK.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CK *CKTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CK.Contract.contract.Transact(opts, method, params...)
}

// GEN0_AUCTION_DURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_CK *CKCaller) GEN0_AUCTION_DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "GEN0_AUCTION_DURATION")
	return *ret0, err
}

// GEN0_AUCTION_DURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_CK *CKSession) GEN0_AUCTION_DURATION() (*big.Int, error) {
	return _CK.Contract.GEN0_AUCTION_DURATION(&_CK.CallOpts)
}

// GEN0_AUCTION_DURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() constant returns(uint256)
func (_CK *CKCallerSession) GEN0_AUCTION_DURATION() (*big.Int, error) {
	return _CK.Contract.GEN0_AUCTION_DURATION(&_CK.CallOpts)
}

// GEN0_CREATION_LIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_CK *CKCaller) GEN0_CREATION_LIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "GEN0_CREATION_LIMIT")
	return *ret0, err
}

// GEN0_CREATION_LIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_CK *CKSession) GEN0_CREATION_LIMIT() (*big.Int, error) {
	return _CK.Contract.GEN0_CREATION_LIMIT(&_CK.CallOpts)
}

// GEN0_CREATION_LIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() constant returns(uint256)
func (_CK *CKCallerSession) GEN0_CREATION_LIMIT() (*big.Int, error) {
	return _CK.Contract.GEN0_CREATION_LIMIT(&_CK.CallOpts)
}

// GEN0_STARTING_PRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_CK *CKCaller) GEN0_STARTING_PRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "GEN0_STARTING_PRICE")
	return *ret0, err
}

// GEN0_STARTING_PRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_CK *CKSession) GEN0_STARTING_PRICE() (*big.Int, error) {
	return _CK.Contract.GEN0_STARTING_PRICE(&_CK.CallOpts)
}

// GEN0_STARTING_PRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() constant returns(uint256)
func (_CK *CKCallerSession) GEN0_STARTING_PRICE() (*big.Int, error) {
	return _CK.Contract.GEN0_STARTING_PRICE(&_CK.CallOpts)
}

// PROMO_CREATION_LIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_CK *CKCaller) PROMO_CREATION_LIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "PROMO_CREATION_LIMIT")
	return *ret0, err
}

// PROMO_CREATION_LIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_CK *CKSession) PROMO_CREATION_LIMIT() (*big.Int, error) {
	return _CK.Contract.PROMO_CREATION_LIMIT(&_CK.CallOpts)
}

// PROMO_CREATION_LIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() constant returns(uint256)
func (_CK *CKCallerSession) PROMO_CREATION_LIMIT() (*big.Int, error) {
	return _CK.Contract.PROMO_CREATION_LIMIT(&_CK.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_CK *CKCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "autoBirthFee")
	return *ret0, err
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_CK *CKSession) AutoBirthFee() (*big.Int, error) {
	return _CK.Contract.AutoBirthFee(&_CK.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() constant returns(uint256)
func (_CK *CKCallerSession) AutoBirthFee() (*big.Int, error) {
	return _CK.Contract.AutoBirthFee(&_CK.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CK *CKCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CK *CKSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CK.Contract.BalanceOf(&_CK.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CK *CKCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CK.Contract.BalanceOf(&_CK.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_CK *CKCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "canBreedWith", _matronId, _sireId)
	return *ret0, err
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_CK *CKSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _CK.Contract.CanBreedWith(&_CK.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(_matronId uint256, _sireId uint256) constant returns(bool)
func (_CK *CKCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _CK.Contract.CanBreedWith(&_CK.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_CK *CKCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_CK *CKSession) CeoAddress() (common.Address, error) {
	return _CK.Contract.CeoAddress(&_CK.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_CK *CKCallerSession) CeoAddress() (common.Address, error) {
	return _CK.Contract.CeoAddress(&_CK.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_CK *CKCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_CK *CKSession) CfoAddress() (common.Address, error) {
	return _CK.Contract.CfoAddress(&_CK.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_CK *CKCallerSession) CfoAddress() (common.Address, error) {
	return _CK.Contract.CfoAddress(&_CK.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_CK *CKCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_CK *CKSession) CooAddress() (common.Address, error) {
	return _CK.Contract.CooAddress(&_CK.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_CK *CKCallerSession) CooAddress() (common.Address, error) {
	return _CK.Contract.CooAddress(&_CK.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_CK *CKCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "cooldowns", arg0)
	return *ret0, err
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_CK *CKSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _CK.Contract.Cooldowns(&_CK.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns( uint256) constant returns(uint32)
func (_CK *CKCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _CK.Contract.Cooldowns(&_CK.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_CK *CKCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "erc721Metadata")
	return *ret0, err
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_CK *CKSession) Erc721Metadata() (common.Address, error) {
	return _CK.Contract.Erc721Metadata(&_CK.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() constant returns(address)
func (_CK *CKCallerSession) Erc721Metadata() (common.Address, error) {
	return _CK.Contract.Erc721Metadata(&_CK.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_CK *CKCaller) Gen0CreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "gen0CreatedCount")
	return *ret0, err
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_CK *CKSession) Gen0CreatedCount() (*big.Int, error) {
	return _CK.Contract.Gen0CreatedCount(&_CK.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() constant returns(uint256)
func (_CK *CKCallerSession) Gen0CreatedCount() (*big.Int, error) {
	return _CK.Contract.Gen0CreatedCount(&_CK.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_CK *CKCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "geneScience")
	return *ret0, err
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_CK *CKSession) GeneScience() (common.Address, error) {
	return _CK.Contract.GeneScience(&_CK.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() constant returns(address)
func (_CK *CKCallerSession) GeneScience() (common.Address, error) {
	return _CK.Contract.GeneScience(&_CK.CallOpts)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(_id uint256) constant returns(isGestating bool, isReady bool, cooldownIndex uint256, nextActionAt uint256, siringWithId uint256, birthTime uint256, matronId uint256, sireId uint256, generation uint256, genes uint256)
func (_CK *CKCaller) GetKitty(opts *bind.CallOpts, _id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	ret := new(struct {
		IsGestating   bool
		IsReady       bool
		CooldownIndex *big.Int
		NextActionAt  *big.Int
		SiringWithId  *big.Int
		BirthTime     *big.Int
		MatronId      *big.Int
		SireId        *big.Int
		Generation    *big.Int
		Genes         *big.Int
	})
	out := ret
	err := _CK.contract.Call(opts, out, "getKitty", _id)
	return *ret, err
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(_id uint256) constant returns(isGestating bool, isReady bool, cooldownIndex uint256, nextActionAt uint256, siringWithId uint256, birthTime uint256, matronId uint256, sireId uint256, generation uint256, genes uint256)
func (_CK *CKSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _CK.Contract.GetKitty(&_CK.CallOpts, _id)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(_id uint256) constant returns(isGestating bool, isReady bool, cooldownIndex uint256, nextActionAt uint256, siringWithId uint256, birthTime uint256, matronId uint256, sireId uint256, generation uint256, genes uint256)
func (_CK *CKCallerSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _CK.Contract.GetKitty(&_CK.CallOpts, _id)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_CK *CKCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "isPregnant", _kittyId)
	return *ret0, err
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_CK *CKSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _CK.Contract.IsPregnant(&_CK.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(_kittyId uint256) constant returns(bool)
func (_CK *CKCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _CK.Contract.IsPregnant(&_CK.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_CK *CKCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "isReadyToBreed", _kittyId)
	return *ret0, err
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_CK *CKSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _CK.Contract.IsReadyToBreed(&_CK.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(_kittyId uint256) constant returns(bool)
func (_CK *CKCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _CK.Contract.IsReadyToBreed(&_CK.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_CK *CKCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "kittyIndexToApproved", arg0)
	return *ret0, err
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_CK *CKSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _CK.Contract.KittyIndexToApproved(&_CK.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved( uint256) constant returns(address)
func (_CK *CKCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _CK.Contract.KittyIndexToApproved(&_CK.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_CK *CKCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "kittyIndexToOwner", arg0)
	return *ret0, err
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_CK *CKSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CK.Contract.KittyIndexToOwner(&_CK.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner( uint256) constant returns(address)
func (_CK *CKCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CK.Contract.KittyIndexToOwner(&_CK.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CK *CKCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CK *CKSession) Name() (string, error) {
	return _CK.Contract.Name(&_CK.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CK *CKCallerSession) Name() (string, error) {
	return _CK.Contract.Name(&_CK.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CK *CKCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CK *CKSession) NewContractAddress() (common.Address, error) {
	return _CK.Contract.NewContractAddress(&_CK.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CK *CKCallerSession) NewContractAddress() (common.Address, error) {
	return _CK.Contract.NewContractAddress(&_CK.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CK *CKCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CK *CKSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CK.Contract.OwnerOf(&_CK.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CK *CKCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CK.Contract.OwnerOf(&_CK.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CK *CKCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CK *CKSession) Paused() (bool, error) {
	return _CK.Contract.Paused(&_CK.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CK *CKCallerSession) Paused() (bool, error) {
	return _CK.Contract.Paused(&_CK.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_CK *CKCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "pregnantKitties")
	return *ret0, err
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_CK *CKSession) PregnantKitties() (*big.Int, error) {
	return _CK.Contract.PregnantKitties(&_CK.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() constant returns(uint256)
func (_CK *CKCallerSession) PregnantKitties() (*big.Int, error) {
	return _CK.Contract.PregnantKitties(&_CK.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_CK *CKCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "promoCreatedCount")
	return *ret0, err
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_CK *CKSession) PromoCreatedCount() (*big.Int, error) {
	return _CK.Contract.PromoCreatedCount(&_CK.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() constant returns(uint256)
func (_CK *CKCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _CK.Contract.PromoCreatedCount(&_CK.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_CK *CKCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "saleAuction")
	return *ret0, err
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_CK *CKSession) SaleAuction() (common.Address, error) {
	return _CK.Contract.SaleAuction(&_CK.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() constant returns(address)
func (_CK *CKCallerSession) SaleAuction() (common.Address, error) {
	return _CK.Contract.SaleAuction(&_CK.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_CK *CKCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "secondsPerBlock")
	return *ret0, err
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_CK *CKSession) SecondsPerBlock() (*big.Int, error) {
	return _CK.Contract.SecondsPerBlock(&_CK.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() constant returns(uint256)
func (_CK *CKCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _CK.Contract.SecondsPerBlock(&_CK.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_CK *CKCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "sireAllowedToAddress", arg0)
	return *ret0, err
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_CK *CKSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _CK.Contract.SireAllowedToAddress(&_CK.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress( uint256) constant returns(address)
func (_CK *CKCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _CK.Contract.SireAllowedToAddress(&_CK.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_CK *CKCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "siringAuction")
	return *ret0, err
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_CK *CKSession) SiringAuction() (common.Address, error) {
	return _CK.Contract.SiringAuction(&_CK.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() constant returns(address)
func (_CK *CKCallerSession) SiringAuction() (common.Address, error) {
	return _CK.Contract.SiringAuction(&_CK.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_CK *CKCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_CK *CKSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _CK.Contract.SupportsInterface(&_CK.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceID bytes4) constant returns(bool)
func (_CK *CKCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _CK.Contract.SupportsInterface(&_CK.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CK *CKCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CK *CKSession) Symbol() (string, error) {
	return _CK.Contract.Symbol(&_CK.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CK *CKCallerSession) Symbol() (string, error) {
	return _CK.Contract.Symbol(&_CK.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_CK *CKCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "tokenMetadata", _tokenId, _preferredTransport)
	return *ret0, err
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_CK *CKSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _CK.Contract.TokenMetadata(&_CK.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(_tokenId uint256, _preferredTransport string) constant returns(infoUrl string)
func (_CK *CKCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _CK.Contract.TokenMetadata(&_CK.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CK *CKCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CK *CKSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CK.Contract.TokensOfOwner(&_CK.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CK *CKCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CK.Contract.TokensOfOwner(&_CK.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CK *CKCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CK.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CK *CKSession) TotalSupply() (*big.Int, error) {
	return _CK.Contract.TotalSupply(&_CK.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CK *CKCallerSession) TotalSupply() (*big.Int, error) {
	return _CK.Contract.TotalSupply(&_CK.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_CK *CKTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_CK *CKSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.Approve(&_CK.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_CK *CKTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.Approve(&_CK.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_CK *CKTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_CK *CKSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.ApproveSiring(&_CK.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(_addr address, _sireId uint256) returns()
func (_CK *CKTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.ApproveSiring(&_CK.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_CK *CKTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_CK *CKSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.BidOnSiringAuction(&_CK.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(_sireId uint256, _matronId uint256) returns()
func (_CK *CKTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.BidOnSiringAuction(&_CK.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_CK *CKTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_CK *CKSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.BreedWithAuto(&_CK.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(_matronId uint256, _sireId uint256) returns()
func (_CK *CKTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.BreedWithAuto(&_CK.TransactOpts, _matronId, _sireId)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_CK *CKTransactor) CreateGen0Auction(opts *bind.TransactOpts, _genes *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "createGen0Auction", _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_CK *CKSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _CK.Contract.CreateGen0Auction(&_CK.TransactOpts, _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(_genes uint256) returns()
func (_CK *CKTransactorSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _CK.Contract.CreateGen0Auction(&_CK.TransactOpts, _genes)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_CK *CKTransactor) CreatePromoKitty(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "createPromoKitty", _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_CK *CKSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _CK.Contract.CreatePromoKitty(&_CK.TransactOpts, _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(_genes uint256, _owner address) returns()
func (_CK *CKTransactorSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _CK.Contract.CreatePromoKitty(&_CK.TransactOpts, _genes, _owner)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_CK *CKTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_CK *CKSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _CK.Contract.CreateSaleAuction(&_CK.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_CK *CKTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _CK.Contract.CreateSaleAuction(&_CK.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_CK *CKTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_CK *CKSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _CK.Contract.CreateSiringAuction(&_CK.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(_kittyId uint256, _startingPrice uint256, _endingPrice uint256, _duration uint256) returns()
func (_CK *CKTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _CK.Contract.CreateSiringAuction(&_CK.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_CK *CKTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_CK *CKSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.GiveBirth(&_CK.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(_matronId uint256) returns(uint256)
func (_CK *CKTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.GiveBirth(&_CK.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CK *CKTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CK *CKSession) Pause() (*types.Transaction, error) {
	return _CK.Contract.Pause(&_CK.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CK *CKTransactorSession) Pause() (*types.Transaction, error) {
	return _CK.Contract.Pause(&_CK.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_CK *CKTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_CK *CKSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _CK.Contract.SetAutoBirthFee(&_CK.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(val uint256) returns()
func (_CK *CKTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _CK.Contract.SetAutoBirthFee(&_CK.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_CK *CKTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_CK *CKSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetCEO(&_CK.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(_newCEO address) returns()
func (_CK *CKTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetCEO(&_CK.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_CK *CKTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_CK *CKSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetCFO(&_CK.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(_newCFO address) returns()
func (_CK *CKTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetCFO(&_CK.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_CK *CKTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_CK *CKSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetCOO(&_CK.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(_newCOO address) returns()
func (_CK *CKTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetCOO(&_CK.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_CK *CKTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_CK *CKSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetGeneScienceAddress(&_CK.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(_address address) returns()
func (_CK *CKTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetGeneScienceAddress(&_CK.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_CK *CKTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_CK *CKSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetMetadataAddress(&_CK.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(_contractAddress address) returns()
func (_CK *CKTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetMetadataAddress(&_CK.TransactOpts, _contractAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_CK *CKTransactor) SetNewAddress(opts *bind.TransactOpts, _v2Address common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setNewAddress", _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_CK *CKSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetNewAddress(&_CK.TransactOpts, _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_v2Address address) returns()
func (_CK *CKTransactorSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetNewAddress(&_CK.TransactOpts, _v2Address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_CK *CKTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_CK *CKSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetSaleAuctionAddress(&_CK.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(_address address) returns()
func (_CK *CKTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetSaleAuctionAddress(&_CK.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_CK *CKTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_CK *CKSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _CK.Contract.SetSecondsPerBlock(&_CK.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(secs uint256) returns()
func (_CK *CKTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _CK.Contract.SetSecondsPerBlock(&_CK.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_CK *CKTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_CK *CKSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetSiringAuctionAddress(&_CK.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(_address address) returns()
func (_CK *CKTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _CK.Contract.SetSiringAuctionAddress(&_CK.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_CK *CKTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_CK *CKSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.Transfer(&_CK.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_CK *CKTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.Transfer(&_CK.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_CK *CKTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_CK *CKSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.TransferFrom(&_CK.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_CK *CKTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _CK.Contract.TransferFrom(&_CK.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CK *CKTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CK *CKSession) Unpause() (*types.Transaction, error) {
	return _CK.Contract.Unpause(&_CK.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CK *CKTransactorSession) Unpause() (*types.Transaction, error) {
	return _CK.Contract.Unpause(&_CK.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_CK *CKTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_CK *CKSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _CK.Contract.WithdrawAuctionBalances(&_CK.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_CK *CKTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _CK.Contract.WithdrawAuctionBalances(&_CK.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CK *CKTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CK.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CK *CKSession) WithdrawBalance() (*types.Transaction, error) {
	return _CK.Contract.WithdrawBalance(&_CK.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CK *CKTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _CK.Contract.WithdrawBalance(&_CK.TransactOpts)
}

// CKApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CK contract.
type CKApprovalIterator struct {
	Event *CKApproval // Event containing the contract specifics and raw log

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
func (it *CKApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKApproval)
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
		it.Event = new(CKApproval)
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
func (it *CKApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKApproval represents a Approval event raised by the CK contract.
type CKApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner address, approved address, tokenId uint256)
func (_CK *CKFilterer) FilterApproval(opts *bind.FilterOpts) (*CKApprovalIterator, error) {

	logs, sub, err := _CK.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &CKApprovalIterator{contract: _CK.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner address, approved address, tokenId uint256)
func (_CK *CKFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CKApproval) (event.Subscription, error) {

	logs, sub, err := _CK.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKApproval)
				if err := _CK.contract.UnpackLog(event, "Approval", log); err != nil {
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

// CKBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the CK contract.
type CKBirthIterator struct {
	Event *CKBirth // Event containing the contract specifics and raw log

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
func (it *CKBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKBirth)
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
		it.Event = new(CKBirth)
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
func (it *CKBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKBirth represents a Birth event raised by the CK contract.
type CKBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_CK *CKFilterer) FilterBirth(opts *bind.FilterOpts) (*CKBirthIterator, error) {

	logs, sub, err := _CK.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &CKBirthIterator{contract: _CK.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(owner address, kittyId uint256, matronId uint256, sireId uint256, genes uint256)
func (_CK *CKFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *CKBirth) (event.Subscription, error) {

	logs, sub, err := _CK.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKBirth)
				if err := _CK.contract.UnpackLog(event, "Birth", log); err != nil {
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

// CKContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the CK contract.
type CKContractUpgradeIterator struct {
	Event *CKContractUpgrade // Event containing the contract specifics and raw log

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
func (it *CKContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKContractUpgrade)
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
		it.Event = new(CKContractUpgrade)
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
func (it *CKContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKContractUpgrade represents a ContractUpgrade event raised by the CK contract.
type CKContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CK *CKFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*CKContractUpgradeIterator, error) {

	logs, sub, err := _CK.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &CKContractUpgradeIterator{contract: _CK.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CK *CKFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *CKContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _CK.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKContractUpgrade)
				if err := _CK.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// CKPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the CK contract.
type CKPregnantIterator struct {
	Event *CKPregnant // Event containing the contract specifics and raw log

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
func (it *CKPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPregnant)
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
		it.Event = new(CKPregnant)
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
func (it *CKPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPregnant represents a Pregnant event raised by the CK contract.
type CKPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_CK *CKFilterer) FilterPregnant(opts *bind.FilterOpts) (*CKPregnantIterator, error) {

	logs, sub, err := _CK.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &CKPregnantIterator{contract: _CK.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(owner address, matronId uint256, sireId uint256, cooldownEndBlock uint256)
func (_CK *CKFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *CKPregnant) (event.Subscription, error) {

	logs, sub, err := _CK.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPregnant)
				if err := _CK.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// CKTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CK contract.
type CKTransferIterator struct {
	Event *CKTransfer // Event containing the contract specifics and raw log

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
func (it *CKTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKTransfer)
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
		it.Event = new(CKTransfer)
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
func (it *CKTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKTransfer represents a Transfer event raised by the CK contract.
type CKTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenId uint256)
func (_CK *CKFilterer) FilterTransfer(opts *bind.FilterOpts) (*CKTransferIterator, error) {

	logs, sub, err := _CK.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &CKTransferIterator{contract: _CK.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenId uint256)
func (_CK *CKFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CKTransfer) (event.Subscription, error) {

	logs, sub, err := _CK.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKTransfer)
				if err := _CK.contract.UnpackLog(event, "Transfer", log); err != nil {
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
