// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package odysseyabi

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

// VestingMetaData contains all meta data concerning the Vesting contract.
var VestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dad_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_starting_date\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"last_claim_date\",\"type\":\"uint256\"}],\"name\":\"HolderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mom\",\"type\":\"address\"}],\"name\":\"MOMAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dad_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end_date\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"holders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"last_claim_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_tokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mom_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeem_tokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mom_token\",\"type\":\"address\"}],\"name\":\"set_mom_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"starting_date\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"update_holder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VestingABI is the input ABI used to generate the binding from.
// Deprecated: Use VestingMetaData.ABI instead.
var VestingABI = VestingMetaData.ABI

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VestingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vesting *VestingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vesting *VestingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vesting.Contract.DEFAULTADMINROLE(&_Vesting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vesting *VestingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vesting.Contract.DEFAULTADMINROLE(&_Vesting.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_Vesting *VestingCaller) UPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_Vesting *VestingSession) UPDATERROLE() ([32]byte, error) {
	return _Vesting.Contract.UPDATERROLE(&_Vesting.CallOpts)
}

// UPDATERROLE is a free data retrieval call binding the contract method 0x47e63380.
//
// Solidity: function UPDATER_ROLE() view returns(bytes32)
func (_Vesting *VestingCallerSession) UPDATERROLE() ([32]byte, error) {
	return _Vesting.Contract.UPDATERROLE(&_Vesting.CallOpts)
}

// DadToken is a free data retrieval call binding the contract method 0x16a2d236.
//
// Solidity: function dad_token() view returns(address)
func (_Vesting *VestingCaller) DadToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "dad_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DadToken is a free data retrieval call binding the contract method 0x16a2d236.
//
// Solidity: function dad_token() view returns(address)
func (_Vesting *VestingSession) DadToken() (common.Address, error) {
	return _Vesting.Contract.DadToken(&_Vesting.CallOpts)
}

// DadToken is a free data retrieval call binding the contract method 0x16a2d236.
//
// Solidity: function dad_token() view returns(address)
func (_Vesting *VestingCallerSession) DadToken() (common.Address, error) {
	return _Vesting.Contract.DadToken(&_Vesting.CallOpts)
}

// EndDate is a free data retrieval call binding the contract method 0x33bbae93.
//
// Solidity: function end_date() view returns(uint256)
func (_Vesting *VestingCaller) EndDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "end_date")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndDate is a free data retrieval call binding the contract method 0x33bbae93.
//
// Solidity: function end_date() view returns(uint256)
func (_Vesting *VestingSession) EndDate() (*big.Int, error) {
	return _Vesting.Contract.EndDate(&_Vesting.CallOpts)
}

// EndDate is a free data retrieval call binding the contract method 0x33bbae93.
//
// Solidity: function end_date() view returns(uint256)
func (_Vesting *VestingCallerSession) EndDate() (*big.Int, error) {
	return _Vesting.Contract.EndDate(&_Vesting.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vesting *VestingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vesting *VestingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vesting.Contract.GetRoleAdmin(&_Vesting.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vesting *VestingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vesting.Contract.GetRoleAdmin(&_Vesting.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vesting *VestingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vesting *VestingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vesting.Contract.HasRole(&_Vesting.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vesting *VestingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vesting.Contract.HasRole(&_Vesting.CallOpts, role, account)
}

// Holders is a free data retrieval call binding the contract method 0x18a5bbdc.
//
// Solidity: function holders(address ) view returns(uint256 last_claim_date, uint256 total_tokens)
func (_Vesting *VestingCaller) Holders(opts *bind.CallOpts, arg0 common.Address) (struct {
	LastClaimDate *big.Int
	TotalTokens   *big.Int
}, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "holders", arg0)

	outstruct := new(struct {
		LastClaimDate *big.Int
		TotalTokens   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastClaimDate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Holders is a free data retrieval call binding the contract method 0x18a5bbdc.
//
// Solidity: function holders(address ) view returns(uint256 last_claim_date, uint256 total_tokens)
func (_Vesting *VestingSession) Holders(arg0 common.Address) (struct {
	LastClaimDate *big.Int
	TotalTokens   *big.Int
}, error) {
	return _Vesting.Contract.Holders(&_Vesting.CallOpts, arg0)
}

// Holders is a free data retrieval call binding the contract method 0x18a5bbdc.
//
// Solidity: function holders(address ) view returns(uint256 last_claim_date, uint256 total_tokens)
func (_Vesting *VestingCallerSession) Holders(arg0 common.Address) (struct {
	LastClaimDate *big.Int
	TotalTokens   *big.Int
}, error) {
	return _Vesting.Contract.Holders(&_Vesting.CallOpts, arg0)
}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Vesting *VestingCaller) MomToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "mom_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Vesting *VestingSession) MomToken() (common.Address, error) {
	return _Vesting.Contract.MomToken(&_Vesting.CallOpts)
}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Vesting *VestingCallerSession) MomToken() (common.Address, error) {
	return _Vesting.Contract.MomToken(&_Vesting.CallOpts)
}

// StartingDate is a free data retrieval call binding the contract method 0x9e6286a6.
//
// Solidity: function starting_date() view returns(uint256)
func (_Vesting *VestingCaller) StartingDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "starting_date")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingDate is a free data retrieval call binding the contract method 0x9e6286a6.
//
// Solidity: function starting_date() view returns(uint256)
func (_Vesting *VestingSession) StartingDate() (*big.Int, error) {
	return _Vesting.Contract.StartingDate(&_Vesting.CallOpts)
}

// StartingDate is a free data retrieval call binding the contract method 0x9e6286a6.
//
// Solidity: function starting_date() view returns(uint256)
func (_Vesting *VestingCallerSession) StartingDate() (*big.Int, error) {
	return _Vesting.Contract.StartingDate(&_Vesting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vesting *VestingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vesting *VestingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vesting.Contract.SupportsInterface(&_Vesting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vesting *VestingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vesting.Contract.SupportsInterface(&_Vesting.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vesting *VestingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vesting *VestingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.GrantRole(&_Vesting.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vesting *VestingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.GrantRole(&_Vesting.TransactOpts, role, account)
}

// RedeemTokens is a paid mutator transaction binding the contract method 0xf790ea98.
//
// Solidity: function redeem_tokens() returns()
func (_Vesting *VestingTransactor) RedeemTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "redeem_tokens")
}

// RedeemTokens is a paid mutator transaction binding the contract method 0xf790ea98.
//
// Solidity: function redeem_tokens() returns()
func (_Vesting *VestingSession) RedeemTokens() (*types.Transaction, error) {
	return _Vesting.Contract.RedeemTokens(&_Vesting.TransactOpts)
}

// RedeemTokens is a paid mutator transaction binding the contract method 0xf790ea98.
//
// Solidity: function redeem_tokens() returns()
func (_Vesting *VestingTransactorSession) RedeemTokens() (*types.Transaction, error) {
	return _Vesting.Contract.RedeemTokens(&_Vesting.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vesting *VestingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vesting *VestingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.RenounceRole(&_Vesting.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vesting *VestingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.RenounceRole(&_Vesting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vesting *VestingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vesting *VestingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.RevokeRole(&_Vesting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vesting *VestingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.RevokeRole(&_Vesting.TransactOpts, role, account)
}

// SetMomAddress is a paid mutator transaction binding the contract method 0x4700e993.
//
// Solidity: function set_mom_address(address _mom_token) returns()
func (_Vesting *VestingTransactor) SetMomAddress(opts *bind.TransactOpts, _mom_token common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "set_mom_address", _mom_token)
}

// SetMomAddress is a paid mutator transaction binding the contract method 0x4700e993.
//
// Solidity: function set_mom_address(address _mom_token) returns()
func (_Vesting *VestingSession) SetMomAddress(_mom_token common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.SetMomAddress(&_Vesting.TransactOpts, _mom_token)
}

// SetMomAddress is a paid mutator transaction binding the contract method 0x4700e993.
//
// Solidity: function set_mom_address(address _mom_token) returns()
func (_Vesting *VestingTransactorSession) SetMomAddress(_mom_token common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.SetMomAddress(&_Vesting.TransactOpts, _mom_token)
}

// UpdateHolder is a paid mutator transaction binding the contract method 0xb7fe5839.
//
// Solidity: function update_holder(address holder, uint256 amount) returns()
func (_Vesting *VestingTransactor) UpdateHolder(opts *bind.TransactOpts, holder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "update_holder", holder, amount)
}

// UpdateHolder is a paid mutator transaction binding the contract method 0xb7fe5839.
//
// Solidity: function update_holder(address holder, uint256 amount) returns()
func (_Vesting *VestingSession) UpdateHolder(holder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.UpdateHolder(&_Vesting.TransactOpts, holder, amount)
}

// UpdateHolder is a paid mutator transaction binding the contract method 0xb7fe5839.
//
// Solidity: function update_holder(address holder, uint256 amount) returns()
func (_Vesting *VestingTransactorSession) UpdateHolder(holder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.UpdateHolder(&_Vesting.TransactOpts, holder, amount)
}

// VestingHolderUpdatedIterator is returned from FilterHolderUpdated and is used to iterate over the raw logs and unpacked data for HolderUpdated events raised by the Vesting contract.
type VestingHolderUpdatedIterator struct {
	Event *VestingHolderUpdated // Event containing the contract specifics and raw log

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
func (it *VestingHolderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingHolderUpdated)
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
		it.Event = new(VestingHolderUpdated)
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
func (it *VestingHolderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingHolderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingHolderUpdated represents a HolderUpdated event raised by the Vesting contract.
type VestingHolderUpdated struct {
	Holder        common.Address
	Amount        *big.Int
	LastClaimDate *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHolderUpdated is a free log retrieval operation binding the contract event 0xbe90c177b758726eaac3c47d0956232b12f1f5252248c1d0fa129b6fb35c8d00.
//
// Solidity: event HolderUpdated(address holder, uint256 amount, uint256 last_claim_date)
func (_Vesting *VestingFilterer) FilterHolderUpdated(opts *bind.FilterOpts) (*VestingHolderUpdatedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "HolderUpdated")
	if err != nil {
		return nil, err
	}
	return &VestingHolderUpdatedIterator{contract: _Vesting.contract, event: "HolderUpdated", logs: logs, sub: sub}, nil
}

// WatchHolderUpdated is a free log subscription operation binding the contract event 0xbe90c177b758726eaac3c47d0956232b12f1f5252248c1d0fa129b6fb35c8d00.
//
// Solidity: event HolderUpdated(address holder, uint256 amount, uint256 last_claim_date)
func (_Vesting *VestingFilterer) WatchHolderUpdated(opts *bind.WatchOpts, sink chan<- *VestingHolderUpdated) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "HolderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingHolderUpdated)
				if err := _Vesting.contract.UnpackLog(event, "HolderUpdated", log); err != nil {
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

// ParseHolderUpdated is a log parse operation binding the contract event 0xbe90c177b758726eaac3c47d0956232b12f1f5252248c1d0fa129b6fb35c8d00.
//
// Solidity: event HolderUpdated(address holder, uint256 amount, uint256 last_claim_date)
func (_Vesting *VestingFilterer) ParseHolderUpdated(log types.Log) (*VestingHolderUpdated, error) {
	event := new(VestingHolderUpdated)
	if err := _Vesting.contract.UnpackLog(event, "HolderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingMOMAddressUpdatedIterator is returned from FilterMOMAddressUpdated and is used to iterate over the raw logs and unpacked data for MOMAddressUpdated events raised by the Vesting contract.
type VestingMOMAddressUpdatedIterator struct {
	Event *VestingMOMAddressUpdated // Event containing the contract specifics and raw log

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
func (it *VestingMOMAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingMOMAddressUpdated)
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
		it.Event = new(VestingMOMAddressUpdated)
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
func (it *VestingMOMAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingMOMAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingMOMAddressUpdated represents a MOMAddressUpdated event raised by the Vesting contract.
type VestingMOMAddressUpdated struct {
	Mom common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMOMAddressUpdated is a free log retrieval operation binding the contract event 0x09e1d4fb8192dbd7602a29a119b236a9666c85bd825b70ef31d5a5396cdadd67.
//
// Solidity: event MOMAddressUpdated(address mom)
func (_Vesting *VestingFilterer) FilterMOMAddressUpdated(opts *bind.FilterOpts) (*VestingMOMAddressUpdatedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "MOMAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &VestingMOMAddressUpdatedIterator{contract: _Vesting.contract, event: "MOMAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchMOMAddressUpdated is a free log subscription operation binding the contract event 0x09e1d4fb8192dbd7602a29a119b236a9666c85bd825b70ef31d5a5396cdadd67.
//
// Solidity: event MOMAddressUpdated(address mom)
func (_Vesting *VestingFilterer) WatchMOMAddressUpdated(opts *bind.WatchOpts, sink chan<- *VestingMOMAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "MOMAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingMOMAddressUpdated)
				if err := _Vesting.contract.UnpackLog(event, "MOMAddressUpdated", log); err != nil {
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

// ParseMOMAddressUpdated is a log parse operation binding the contract event 0x09e1d4fb8192dbd7602a29a119b236a9666c85bd825b70ef31d5a5396cdadd67.
//
// Solidity: event MOMAddressUpdated(address mom)
func (_Vesting *VestingFilterer) ParseMOMAddressUpdated(log types.Log) (*VestingMOMAddressUpdated, error) {
	event := new(VestingMOMAddressUpdated)
	if err := _Vesting.contract.UnpackLog(event, "MOMAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the Vesting contract.
type VestingRedeemedIterator struct {
	Event *VestingRedeemed // Event containing the contract specifics and raw log

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
func (it *VestingRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingRedeemed)
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
		it.Event = new(VestingRedeemed)
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
func (it *VestingRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingRedeemed represents a Redeemed event raised by the Vesting contract.
type VestingRedeemed struct {
	Holder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x4896181ff8f4543cc00db9fe9b6fb7e6f032b7eb772c72ab1ec1b4d2e03b9369.
//
// Solidity: event Redeemed(address holder, uint256 amount)
func (_Vesting *VestingFilterer) FilterRedeemed(opts *bind.FilterOpts) (*VestingRedeemedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "Redeemed")
	if err != nil {
		return nil, err
	}
	return &VestingRedeemedIterator{contract: _Vesting.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x4896181ff8f4543cc00db9fe9b6fb7e6f032b7eb772c72ab1ec1b4d2e03b9369.
//
// Solidity: event Redeemed(address holder, uint256 amount)
func (_Vesting *VestingFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *VestingRedeemed) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "Redeemed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingRedeemed)
				if err := _Vesting.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x4896181ff8f4543cc00db9fe9b6fb7e6f032b7eb772c72ab1ec1b4d2e03b9369.
//
// Solidity: event Redeemed(address holder, uint256 amount)
func (_Vesting *VestingFilterer) ParseRedeemed(log types.Log) (*VestingRedeemed, error) {
	event := new(VestingRedeemed)
	if err := _Vesting.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vesting contract.
type VestingRoleAdminChangedIterator struct {
	Event *VestingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VestingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingRoleAdminChanged)
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
		it.Event = new(VestingRoleAdminChanged)
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
func (it *VestingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingRoleAdminChanged represents a RoleAdminChanged event raised by the Vesting contract.
type VestingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vesting *VestingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VestingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VestingRoleAdminChangedIterator{contract: _Vesting.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vesting *VestingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VestingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingRoleAdminChanged)
				if err := _Vesting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vesting *VestingFilterer) ParseRoleAdminChanged(log types.Log) (*VestingRoleAdminChanged, error) {
	event := new(VestingRoleAdminChanged)
	if err := _Vesting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vesting contract.
type VestingRoleGrantedIterator struct {
	Event *VestingRoleGranted // Event containing the contract specifics and raw log

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
func (it *VestingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingRoleGranted)
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
		it.Event = new(VestingRoleGranted)
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
func (it *VestingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingRoleGranted represents a RoleGranted event raised by the Vesting contract.
type VestingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vesting *VestingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VestingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VestingRoleGrantedIterator{contract: _Vesting.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vesting *VestingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VestingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingRoleGranted)
				if err := _Vesting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vesting *VestingFilterer) ParseRoleGranted(log types.Log) (*VestingRoleGranted, error) {
	event := new(VestingRoleGranted)
	if err := _Vesting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vesting contract.
type VestingRoleRevokedIterator struct {
	Event *VestingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VestingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingRoleRevoked)
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
		it.Event = new(VestingRoleRevoked)
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
func (it *VestingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingRoleRevoked represents a RoleRevoked event raised by the Vesting contract.
type VestingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vesting *VestingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VestingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VestingRoleRevokedIterator{contract: _Vesting.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vesting *VestingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VestingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingRoleRevoked)
				if err := _Vesting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vesting *VestingFilterer) ParseRoleRevoked(log types.Log) (*VestingRoleRevoked, error) {
	event := new(VestingRoleRevoked)
	if err := _Vesting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
