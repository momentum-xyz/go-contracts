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

// StakingStakedBy is an auto generated low-level Go binding around an user-defined struct.
type StakingStakedBy struct {
	User               common.Address
	TotalAmount        *big.Int
	DadAmount          *big.Int
	MomAmount          *big.Int
	Timestamp          *big.Int
	EffectiveTimestamp *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_claimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_staked\",\"type\":\"uint256\"}],\"name\":\"ClaimedUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_rewards_claimed\",\"type\":\"uint256\"}],\"name\":\"OdysseyRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"odyssey_from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"odyssey_to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStaking.Token\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_staked_from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_staked_to\",\"type\":\"uint256\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_rewards_claimed\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"RewardsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_staked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStaking.Token\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_staked\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"StateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_unstaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStaking.Token\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total_staked\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"}],\"name\":\"claim_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim_unstaked_tokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dad_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"odyssey\",\"type\":\"uint256\"}],\"name\":\"get_staked_by\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dad_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mom_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effective_timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structStaking.StakedBy[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_staked_odysseys\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mom_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dad_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_odyssey_nfts\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_rewards_calculation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locking_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mom_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"odyssey_nfts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"odysseys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_staked_into\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_stakers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staked_odysseys_index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from_odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to_odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Token\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"restake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewards_timeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Token\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"staked_by\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dad_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mom_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effective_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"staked_by_indexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"staked_odysseys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total_rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total_staked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dad_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mom_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_staked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.Token\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unstakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dad_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mom_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"untaking_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_locking_period\",\"type\":\"uint256\"}],\"name\":\"update_locking_period\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakers_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"odysseys_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"odysseys_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"update_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewards_timeout\",\"type\":\"uint256\"}],\"name\":\"update_rewards_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Staking *StakingSession) MANAGERROLE() ([32]byte, error) {
	return _Staking.Contract.MANAGERROLE(&_Staking.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) MANAGERROLE() ([32]byte, error) {
	return _Staking.Contract.MANAGERROLE(&_Staking.CallOpts)
}

// DadToken is a free data retrieval call binding the contract method 0x16a2d236.
//
// Solidity: function dad_token() view returns(address)
func (_Staking *StakingCaller) DadToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "dad_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DadToken is a free data retrieval call binding the contract method 0x16a2d236.
//
// Solidity: function dad_token() view returns(address)
func (_Staking *StakingSession) DadToken() (common.Address, error) {
	return _Staking.Contract.DadToken(&_Staking.CallOpts)
}

// DadToken is a free data retrieval call binding the contract method 0x16a2d236.
//
// Solidity: function dad_token() view returns(address)
func (_Staking *StakingCallerSession) DadToken() (common.Address, error) {
	return _Staking.Contract.DadToken(&_Staking.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// GetStakedBy is a free data retrieval call binding the contract method 0xe22619ec.
//
// Solidity: function get_staked_by(uint256 odyssey) view returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Staking *StakingCaller) GetStakedBy(opts *bind.CallOpts, odyssey *big.Int) ([]StakingStakedBy, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "get_staked_by", odyssey)

	if err != nil {
		return *new([]StakingStakedBy), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingStakedBy)).(*[]StakingStakedBy)

	return out0, err

}

// GetStakedBy is a free data retrieval call binding the contract method 0xe22619ec.
//
// Solidity: function get_staked_by(uint256 odyssey) view returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Staking *StakingSession) GetStakedBy(odyssey *big.Int) ([]StakingStakedBy, error) {
	return _Staking.Contract.GetStakedBy(&_Staking.CallOpts, odyssey)
}

// GetStakedBy is a free data retrieval call binding the contract method 0xe22619ec.
//
// Solidity: function get_staked_by(uint256 odyssey) view returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Staking *StakingCallerSession) GetStakedBy(odyssey *big.Int) ([]StakingStakedBy, error) {
	return _Staking.Contract.GetStakedBy(&_Staking.CallOpts, odyssey)
}

// GetStakedOdysseys is a free data retrieval call binding the contract method 0x27b5f0ba.
//
// Solidity: function get_staked_odysseys() view returns(uint256[])
func (_Staking *StakingCaller) GetStakedOdysseys(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "get_staked_odysseys")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakedOdysseys is a free data retrieval call binding the contract method 0x27b5f0ba.
//
// Solidity: function get_staked_odysseys() view returns(uint256[])
func (_Staking *StakingSession) GetStakedOdysseys() ([]*big.Int, error) {
	return _Staking.Contract.GetStakedOdysseys(&_Staking.CallOpts)
}

// GetStakedOdysseys is a free data retrieval call binding the contract method 0x27b5f0ba.
//
// Solidity: function get_staked_odysseys() view returns(uint256[])
func (_Staking *StakingCallerSession) GetStakedOdysseys() ([]*big.Int, error) {
	return _Staking.Contract.GetStakedOdysseys(&_Staking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// LastRewardsCalculation is a free data retrieval call binding the contract method 0x89f1c6b0.
//
// Solidity: function last_rewards_calculation() view returns(uint256)
func (_Staking *StakingCaller) LastRewardsCalculation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "last_rewards_calculation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardsCalculation is a free data retrieval call binding the contract method 0x89f1c6b0.
//
// Solidity: function last_rewards_calculation() view returns(uint256)
func (_Staking *StakingSession) LastRewardsCalculation() (*big.Int, error) {
	return _Staking.Contract.LastRewardsCalculation(&_Staking.CallOpts)
}

// LastRewardsCalculation is a free data retrieval call binding the contract method 0x89f1c6b0.
//
// Solidity: function last_rewards_calculation() view returns(uint256)
func (_Staking *StakingCallerSession) LastRewardsCalculation() (*big.Int, error) {
	return _Staking.Contract.LastRewardsCalculation(&_Staking.CallOpts)
}

// LockingPeriod is a free data retrieval call binding the contract method 0x6dbf218e.
//
// Solidity: function locking_period() view returns(uint256)
func (_Staking *StakingCaller) LockingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "locking_period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockingPeriod is a free data retrieval call binding the contract method 0x6dbf218e.
//
// Solidity: function locking_period() view returns(uint256)
func (_Staking *StakingSession) LockingPeriod() (*big.Int, error) {
	return _Staking.Contract.LockingPeriod(&_Staking.CallOpts)
}

// LockingPeriod is a free data retrieval call binding the contract method 0x6dbf218e.
//
// Solidity: function locking_period() view returns(uint256)
func (_Staking *StakingCallerSession) LockingPeriod() (*big.Int, error) {
	return _Staking.Contract.LockingPeriod(&_Staking.CallOpts)
}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Staking *StakingCaller) MomToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "mom_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Staking *StakingSession) MomToken() (common.Address, error) {
	return _Staking.Contract.MomToken(&_Staking.CallOpts)
}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Staking *StakingCallerSession) MomToken() (common.Address, error) {
	return _Staking.Contract.MomToken(&_Staking.CallOpts)
}

// OdysseyNfts is a free data retrieval call binding the contract method 0x995c4a3c.
//
// Solidity: function odyssey_nfts() view returns(address)
func (_Staking *StakingCaller) OdysseyNfts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "odyssey_nfts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OdysseyNfts is a free data retrieval call binding the contract method 0x995c4a3c.
//
// Solidity: function odyssey_nfts() view returns(address)
func (_Staking *StakingSession) OdysseyNfts() (common.Address, error) {
	return _Staking.Contract.OdysseyNfts(&_Staking.CallOpts)
}

// OdysseyNfts is a free data retrieval call binding the contract method 0x995c4a3c.
//
// Solidity: function odyssey_nfts() view returns(address)
func (_Staking *StakingCallerSession) OdysseyNfts() (common.Address, error) {
	return _Staking.Contract.OdysseyNfts(&_Staking.CallOpts)
}

// Odysseys is a free data retrieval call binding the contract method 0x6efdf56c.
//
// Solidity: function odysseys(uint256 ) view returns(uint256 odyssey_id, uint256 total_staked_into, uint256 total_stakers, uint256 total_rewards, uint256 staked_odysseys_index)
func (_Staking *StakingCaller) Odysseys(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OdysseyId           *big.Int
	TotalStakedInto     *big.Int
	TotalStakers        *big.Int
	TotalRewards        *big.Int
	StakedOdysseysIndex *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "odysseys", arg0)

	outstruct := new(struct {
		OdysseyId           *big.Int
		TotalStakedInto     *big.Int
		TotalStakers        *big.Int
		TotalRewards        *big.Int
		StakedOdysseysIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OdysseyId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStakedInto = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalStakers = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalRewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StakedOdysseysIndex = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Odysseys is a free data retrieval call binding the contract method 0x6efdf56c.
//
// Solidity: function odysseys(uint256 ) view returns(uint256 odyssey_id, uint256 total_staked_into, uint256 total_stakers, uint256 total_rewards, uint256 staked_odysseys_index)
func (_Staking *StakingSession) Odysseys(arg0 *big.Int) (struct {
	OdysseyId           *big.Int
	TotalStakedInto     *big.Int
	TotalStakers        *big.Int
	TotalRewards        *big.Int
	StakedOdysseysIndex *big.Int
}, error) {
	return _Staking.Contract.Odysseys(&_Staking.CallOpts, arg0)
}

// Odysseys is a free data retrieval call binding the contract method 0x6efdf56c.
//
// Solidity: function odysseys(uint256 ) view returns(uint256 odyssey_id, uint256 total_staked_into, uint256 total_stakers, uint256 total_rewards, uint256 staked_odysseys_index)
func (_Staking *StakingCallerSession) Odysseys(arg0 *big.Int) (struct {
	OdysseyId           *big.Int
	TotalStakedInto     *big.Int
	TotalStakers        *big.Int
	TotalRewards        *big.Int
	StakedOdysseysIndex *big.Int
}, error) {
	return _Staking.Contract.Odysseys(&_Staking.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Staking *StakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Staking *StakingSession) ProxiableUUID() ([32]byte, error) {
	return _Staking.Contract.ProxiableUUID(&_Staking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Staking *StakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Staking.Contract.ProxiableUUID(&_Staking.CallOpts)
}

// RewardsTimeout is a free data retrieval call binding the contract method 0xb143c584.
//
// Solidity: function rewards_timeout() view returns(uint256)
func (_Staking *StakingCaller) RewardsTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewards_timeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsTimeout is a free data retrieval call binding the contract method 0xb143c584.
//
// Solidity: function rewards_timeout() view returns(uint256)
func (_Staking *StakingSession) RewardsTimeout() (*big.Int, error) {
	return _Staking.Contract.RewardsTimeout(&_Staking.CallOpts)
}

// RewardsTimeout is a free data retrieval call binding the contract method 0xb143c584.
//
// Solidity: function rewards_timeout() view returns(uint256)
func (_Staking *StakingCallerSession) RewardsTimeout() (*big.Int, error) {
	return _Staking.Contract.RewardsTimeout(&_Staking.CallOpts)
}

// StakedBy is a free data retrieval call binding the contract method 0xa2e9c3be.
//
// Solidity: function staked_by(uint256 , uint256 ) view returns(address user, uint256 total_amount, uint256 dad_amount, uint256 mom_amount, uint256 timestamp, uint256 effective_timestamp)
func (_Staking *StakingCaller) StakedBy(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	User               common.Address
	TotalAmount        *big.Int
	DadAmount          *big.Int
	MomAmount          *big.Int
	Timestamp          *big.Int
	EffectiveTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "staked_by", arg0, arg1)

	outstruct := new(struct {
		User               common.Address
		TotalAmount        *big.Int
		DadAmount          *big.Int
		MomAmount          *big.Int
		Timestamp          *big.Int
		EffectiveTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DadAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MomAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EffectiveTimestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakedBy is a free data retrieval call binding the contract method 0xa2e9c3be.
//
// Solidity: function staked_by(uint256 , uint256 ) view returns(address user, uint256 total_amount, uint256 dad_amount, uint256 mom_amount, uint256 timestamp, uint256 effective_timestamp)
func (_Staking *StakingSession) StakedBy(arg0 *big.Int, arg1 *big.Int) (struct {
	User               common.Address
	TotalAmount        *big.Int
	DadAmount          *big.Int
	MomAmount          *big.Int
	Timestamp          *big.Int
	EffectiveTimestamp *big.Int
}, error) {
	return _Staking.Contract.StakedBy(&_Staking.CallOpts, arg0, arg1)
}

// StakedBy is a free data retrieval call binding the contract method 0xa2e9c3be.
//
// Solidity: function staked_by(uint256 , uint256 ) view returns(address user, uint256 total_amount, uint256 dad_amount, uint256 mom_amount, uint256 timestamp, uint256 effective_timestamp)
func (_Staking *StakingCallerSession) StakedBy(arg0 *big.Int, arg1 *big.Int) (struct {
	User               common.Address
	TotalAmount        *big.Int
	DadAmount          *big.Int
	MomAmount          *big.Int
	Timestamp          *big.Int
	EffectiveTimestamp *big.Int
}, error) {
	return _Staking.Contract.StakedBy(&_Staking.CallOpts, arg0, arg1)
}

// StakedByIndexes is a free data retrieval call binding the contract method 0x67925344.
//
// Solidity: function staked_by_indexes(uint256 , address ) view returns(uint256)
func (_Staking *StakingCaller) StakedByIndexes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "staked_by_indexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedByIndexes is a free data retrieval call binding the contract method 0x67925344.
//
// Solidity: function staked_by_indexes(uint256 , address ) view returns(uint256)
func (_Staking *StakingSession) StakedByIndexes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Staking.Contract.StakedByIndexes(&_Staking.CallOpts, arg0, arg1)
}

// StakedByIndexes is a free data retrieval call binding the contract method 0x67925344.
//
// Solidity: function staked_by_indexes(uint256 , address ) view returns(uint256)
func (_Staking *StakingCallerSession) StakedByIndexes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Staking.Contract.StakedByIndexes(&_Staking.CallOpts, arg0, arg1)
}

// StakedOdysseys is a free data retrieval call binding the contract method 0xc02bdf70.
//
// Solidity: function staked_odysseys(uint256 ) view returns(uint256)
func (_Staking *StakingCaller) StakedOdysseys(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "staked_odysseys", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedOdysseys is a free data retrieval call binding the contract method 0xc02bdf70.
//
// Solidity: function staked_odysseys(uint256 ) view returns(uint256)
func (_Staking *StakingSession) StakedOdysseys(arg0 *big.Int) (*big.Int, error) {
	return _Staking.Contract.StakedOdysseys(&_Staking.CallOpts, arg0)
}

// StakedOdysseys is a free data retrieval call binding the contract method 0xc02bdf70.
//
// Solidity: function staked_odysseys(uint256 ) view returns(uint256)
func (_Staking *StakingCallerSession) StakedOdysseys(arg0 *big.Int) (*big.Int, error) {
	return _Staking.Contract.StakedOdysseys(&_Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address user, uint256 total_rewards, uint256 total_staked, uint256 dad_amount, uint256 mom_amount)
func (_Staking *StakingCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	User         common.Address
	TotalRewards *big.Int
	TotalStaked  *big.Int
	DadAmount    *big.Int
	MomAmount    *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		User         common.Address
		TotalRewards *big.Int
		TotalStaked  *big.Int
		DadAmount    *big.Int
		MomAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DadAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MomAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address user, uint256 total_rewards, uint256 total_staked, uint256 dad_amount, uint256 mom_amount)
func (_Staking *StakingSession) Stakers(arg0 common.Address) (struct {
	User         common.Address
	TotalRewards *big.Int
	TotalStaked  *big.Int
	DadAmount    *big.Int
	MomAmount    *big.Int
}, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address user, uint256 total_rewards, uint256 total_staked, uint256 dad_amount, uint256 mom_amount)
func (_Staking *StakingCallerSession) Stakers(arg0 common.Address) (struct {
	User         common.Address
	TotalRewards *big.Int
	TotalStaked  *big.Int
	DadAmount    *big.Int
	MomAmount    *big.Int
}, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// TotalStaked is a free data retrieval call binding the contract method 0xaf7568dd.
//
// Solidity: function total_staked() view returns(uint256)
func (_Staking *StakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "total_staked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0xaf7568dd.
//
// Solidity: function total_staked() view returns(uint256)
func (_Staking *StakingSession) TotalStaked() (*big.Int, error) {
	return _Staking.Contract.TotalStaked(&_Staking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0xaf7568dd.
//
// Solidity: function total_staked() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStaked() (*big.Int, error) {
	return _Staking.Contract.TotalStaked(&_Staking.CallOpts)
}

// Unstakes is a free data retrieval call binding the contract method 0xdfef94b4.
//
// Solidity: function unstakes(address , uint256 ) view returns(uint256 dad_amount, uint256 mom_amount, uint256 untaking_timestamp)
func (_Staking *StakingCaller) Unstakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	DadAmount         *big.Int
	MomAmount         *big.Int
	UntakingTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "unstakes", arg0, arg1)

	outstruct := new(struct {
		DadAmount         *big.Int
		MomAmount         *big.Int
		UntakingTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DadAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MomAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UntakingTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Unstakes is a free data retrieval call binding the contract method 0xdfef94b4.
//
// Solidity: function unstakes(address , uint256 ) view returns(uint256 dad_amount, uint256 mom_amount, uint256 untaking_timestamp)
func (_Staking *StakingSession) Unstakes(arg0 common.Address, arg1 *big.Int) (struct {
	DadAmount         *big.Int
	MomAmount         *big.Int
	UntakingTimestamp *big.Int
}, error) {
	return _Staking.Contract.Unstakes(&_Staking.CallOpts, arg0, arg1)
}

// Unstakes is a free data retrieval call binding the contract method 0xdfef94b4.
//
// Solidity: function unstakes(address , uint256 ) view returns(uint256 dad_amount, uint256 mom_amount, uint256 untaking_timestamp)
func (_Staking *StakingCallerSession) Unstakes(arg0 common.Address, arg1 *big.Int) (struct {
	DadAmount         *big.Int
	MomAmount         *big.Int
	UntakingTimestamp *big.Int
}, error) {
	return _Staking.Contract.Unstakes(&_Staking.CallOpts, arg0, arg1)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xb5d3ef17.
//
// Solidity: function claim_rewards(uint256 odyssey_id) returns()
func (_Staking *StakingTransactor) ClaimRewards(opts *bind.TransactOpts, odyssey_id *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claim_rewards", odyssey_id)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xb5d3ef17.
//
// Solidity: function claim_rewards(uint256 odyssey_id) returns()
func (_Staking *StakingSession) ClaimRewards(odyssey_id *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ClaimRewards(&_Staking.TransactOpts, odyssey_id)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xb5d3ef17.
//
// Solidity: function claim_rewards(uint256 odyssey_id) returns()
func (_Staking *StakingTransactorSession) ClaimRewards(odyssey_id *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ClaimRewards(&_Staking.TransactOpts, odyssey_id)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_Staking *StakingTransactor) ClaimRewards0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claim_rewards0")
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_Staking *StakingSession) ClaimRewards0() (*types.Transaction, error) {
	return _Staking.Contract.ClaimRewards0(&_Staking.TransactOpts)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_Staking *StakingTransactorSession) ClaimRewards0() (*types.Transaction, error) {
	return _Staking.Contract.ClaimRewards0(&_Staking.TransactOpts)
}

// ClaimUnstakedTokens is a paid mutator transaction binding the contract method 0x2fb0e9f6.
//
// Solidity: function claim_unstaked_tokens() returns()
func (_Staking *StakingTransactor) ClaimUnstakedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claim_unstaked_tokens")
}

// ClaimUnstakedTokens is a paid mutator transaction binding the contract method 0x2fb0e9f6.
//
// Solidity: function claim_unstaked_tokens() returns()
func (_Staking *StakingSession) ClaimUnstakedTokens() (*types.Transaction, error) {
	return _Staking.Contract.ClaimUnstakedTokens(&_Staking.TransactOpts)
}

// ClaimUnstakedTokens is a paid mutator transaction binding the contract method 0x2fb0e9f6.
//
// Solidity: function claim_unstaked_tokens() returns()
func (_Staking *StakingTransactorSession) ClaimUnstakedTokens() (*types.Transaction, error) {
	return _Staking.Contract.ClaimUnstakedTokens(&_Staking.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _mom_token, address _dad_token, address _odyssey_nfts) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _mom_token common.Address, _dad_token common.Address, _odyssey_nfts common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _mom_token, _dad_token, _odyssey_nfts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _mom_token, address _dad_token, address _odyssey_nfts) returns()
func (_Staking *StakingSession) Initialize(_mom_token common.Address, _dad_token common.Address, _odyssey_nfts common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _mom_token, _dad_token, _odyssey_nfts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _mom_token, address _dad_token, address _odyssey_nfts) returns()
func (_Staking *StakingTransactorSession) Initialize(_mom_token common.Address, _dad_token common.Address, _odyssey_nfts common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _mom_token, _dad_token, _odyssey_nfts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, account)
}

// Restake is a paid mutator transaction binding the contract method 0x09f082a7.
//
// Solidity: function restake(uint256 from_odyssey_id, uint256 to_odyssey_id, uint256 amount, uint8 token) returns()
func (_Staking *StakingTransactor) Restake(opts *bind.TransactOpts, from_odyssey_id *big.Int, to_odyssey_id *big.Int, amount *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "restake", from_odyssey_id, to_odyssey_id, amount, token)
}

// Restake is a paid mutator transaction binding the contract method 0x09f082a7.
//
// Solidity: function restake(uint256 from_odyssey_id, uint256 to_odyssey_id, uint256 amount, uint8 token) returns()
func (_Staking *StakingSession) Restake(from_odyssey_id *big.Int, to_odyssey_id *big.Int, amount *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.Contract.Restake(&_Staking.TransactOpts, from_odyssey_id, to_odyssey_id, amount, token)
}

// Restake is a paid mutator transaction binding the contract method 0x09f082a7.
//
// Solidity: function restake(uint256 from_odyssey_id, uint256 to_odyssey_id, uint256 amount, uint8 token) returns()
func (_Staking *StakingTransactorSession) Restake(from_odyssey_id *big.Int, to_odyssey_id *big.Int, amount *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.Contract.Restake(&_Staking.TransactOpts, from_odyssey_id, to_odyssey_id, amount, token)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// Stake is a paid mutator transaction binding the contract method 0x748e6856.
//
// Solidity: function stake(uint256 odyssey_id, uint256 amount, uint8 token) payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, odyssey_id *big.Int, amount *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", odyssey_id, amount, token)
}

// Stake is a paid mutator transaction binding the contract method 0x748e6856.
//
// Solidity: function stake(uint256 odyssey_id, uint256 amount, uint8 token) payable returns()
func (_Staking *StakingSession) Stake(odyssey_id *big.Int, amount *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, odyssey_id, amount, token)
}

// Stake is a paid mutator transaction binding the contract method 0x748e6856.
//
// Solidity: function stake(uint256 odyssey_id, uint256 amount, uint8 token) payable returns()
func (_Staking *StakingTransactorSession) Stake(odyssey_id *big.Int, amount *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, odyssey_id, amount, token)
}

// Unstake is a paid mutator transaction binding the contract method 0xbc9699f1.
//
// Solidity: function unstake(uint256 odyssey_id, uint8 token) returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts, odyssey_id *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake", odyssey_id, token)
}

// Unstake is a paid mutator transaction binding the contract method 0xbc9699f1.
//
// Solidity: function unstake(uint256 odyssey_id, uint8 token) returns()
func (_Staking *StakingSession) Unstake(odyssey_id *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, odyssey_id, token)
}

// Unstake is a paid mutator transaction binding the contract method 0xbc9699f1.
//
// Solidity: function unstake(uint256 odyssey_id, uint8 token) returns()
func (_Staking *StakingTransactorSession) Unstake(odyssey_id *big.Int, token uint8) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, odyssey_id, token)
}

// UpdateLockingPeriod is a paid mutator transaction binding the contract method 0xcef47851.
//
// Solidity: function update_locking_period(uint256 _locking_period) returns()
func (_Staking *StakingTransactor) UpdateLockingPeriod(opts *bind.TransactOpts, _locking_period *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "update_locking_period", _locking_period)
}

// UpdateLockingPeriod is a paid mutator transaction binding the contract method 0xcef47851.
//
// Solidity: function update_locking_period(uint256 _locking_period) returns()
func (_Staking *StakingSession) UpdateLockingPeriod(_locking_period *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateLockingPeriod(&_Staking.TransactOpts, _locking_period)
}

// UpdateLockingPeriod is a paid mutator transaction binding the contract method 0xcef47851.
//
// Solidity: function update_locking_period(uint256 _locking_period) returns()
func (_Staking *StakingTransactorSession) UpdateLockingPeriod(_locking_period *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateLockingPeriod(&_Staking.TransactOpts, _locking_period)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x96137177.
//
// Solidity: function update_rewards(address[] addresses, uint256[] stakers_amounts, uint256[] odysseys_ids, uint256[] odysseys_amounts, uint256 timestamp) returns()
func (_Staking *StakingTransactor) UpdateRewards(opts *bind.TransactOpts, addresses []common.Address, stakers_amounts []*big.Int, odysseys_ids []*big.Int, odysseys_amounts []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "update_rewards", addresses, stakers_amounts, odysseys_ids, odysseys_amounts, timestamp)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x96137177.
//
// Solidity: function update_rewards(address[] addresses, uint256[] stakers_amounts, uint256[] odysseys_ids, uint256[] odysseys_amounts, uint256 timestamp) returns()
func (_Staking *StakingSession) UpdateRewards(addresses []common.Address, stakers_amounts []*big.Int, odysseys_ids []*big.Int, odysseys_amounts []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateRewards(&_Staking.TransactOpts, addresses, stakers_amounts, odysseys_ids, odysseys_amounts, timestamp)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x96137177.
//
// Solidity: function update_rewards(address[] addresses, uint256[] stakers_amounts, uint256[] odysseys_ids, uint256[] odysseys_amounts, uint256 timestamp) returns()
func (_Staking *StakingTransactorSession) UpdateRewards(addresses []common.Address, stakers_amounts []*big.Int, odysseys_ids []*big.Int, odysseys_amounts []*big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateRewards(&_Staking.TransactOpts, addresses, stakers_amounts, odysseys_ids, odysseys_amounts, timestamp)
}

// UpdateRewardsTimeout is a paid mutator transaction binding the contract method 0x969f7d46.
//
// Solidity: function update_rewards_timeout(uint256 _rewards_timeout) returns()
func (_Staking *StakingTransactor) UpdateRewardsTimeout(opts *bind.TransactOpts, _rewards_timeout *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "update_rewards_timeout", _rewards_timeout)
}

// UpdateRewardsTimeout is a paid mutator transaction binding the contract method 0x969f7d46.
//
// Solidity: function update_rewards_timeout(uint256 _rewards_timeout) returns()
func (_Staking *StakingSession) UpdateRewardsTimeout(_rewards_timeout *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateRewardsTimeout(&_Staking.TransactOpts, _rewards_timeout)
}

// UpdateRewardsTimeout is a paid mutator transaction binding the contract method 0x969f7d46.
//
// Solidity: function update_rewards_timeout(uint256 _rewards_timeout) returns()
func (_Staking *StakingTransactorSession) UpdateRewardsTimeout(_rewards_timeout *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateRewardsTimeout(&_Staking.TransactOpts, _rewards_timeout)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Staking *StakingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Staking *StakingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpgradeTo(&_Staking.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Staking *StakingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpgradeTo(&_Staking.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Staking *StakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Staking *StakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.Contract.UpgradeToAndCall(&_Staking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Staking *StakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.Contract.UpgradeToAndCall(&_Staking.TransactOpts, newImplementation, data)
}

// StakingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Staking contract.
type StakingAdminChangedIterator struct {
	Event *StakingAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAdminChanged)
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
		it.Event = new(StakingAdminChanged)
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
func (it *StakingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAdminChanged represents a AdminChanged event raised by the Staking contract.
type StakingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Staking *StakingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StakingAdminChangedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StakingAdminChangedIterator{contract: _Staking.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Staking *StakingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAdminChanged)
				if err := _Staking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Staking *StakingFilterer) ParseAdminChanged(log types.Log) (*StakingAdminChanged, error) {
	event := new(StakingAdminChanged)
	if err := _Staking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Staking contract.
type StakingBeaconUpgradedIterator struct {
	Event *StakingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBeaconUpgraded)
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
		it.Event = new(StakingBeaconUpgraded)
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
func (it *StakingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBeaconUpgraded represents a BeaconUpgraded event raised by the Staking contract.
type StakingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Staking *StakingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StakingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StakingBeaconUpgradedIterator{contract: _Staking.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Staking *StakingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StakingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBeaconUpgraded)
				if err := _Staking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Staking *StakingFilterer) ParseBeaconUpgraded(log types.Log) (*StakingBeaconUpgraded, error) {
	event := new(StakingBeaconUpgraded)
	if err := _Staking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingClaimedUnstakedIterator is returned from FilterClaimedUnstaked and is used to iterate over the raw logs and unpacked data for ClaimedUnstaked events raised by the Staking contract.
type StakingClaimedUnstakedIterator struct {
	Event *StakingClaimedUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingClaimedUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimedUnstaked)
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
		it.Event = new(StakingClaimedUnstaked)
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
func (it *StakingClaimedUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimedUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimedUnstaked represents a ClaimedUnstaked event raised by the Staking contract.
type StakingClaimedUnstaked struct {
	User         common.Address
	TotalClaimed *big.Int
	TotalStaked  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimedUnstaked is a free log retrieval operation binding the contract event 0x6eeaedc56832e751596e30e289453c8568daad613b8050876671fa4cced0fb8c.
//
// Solidity: event ClaimedUnstaked(address indexed user, uint256 total_claimed, uint256 total_staked)
func (_Staking *StakingFilterer) FilterClaimedUnstaked(opts *bind.FilterOpts, user []common.Address) (*StakingClaimedUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ClaimedUnstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingClaimedUnstakedIterator{contract: _Staking.contract, event: "ClaimedUnstaked", logs: logs, sub: sub}, nil
}

// WatchClaimedUnstaked is a free log subscription operation binding the contract event 0x6eeaedc56832e751596e30e289453c8568daad613b8050876671fa4cced0fb8c.
//
// Solidity: event ClaimedUnstaked(address indexed user, uint256 total_claimed, uint256 total_staked)
func (_Staking *StakingFilterer) WatchClaimedUnstaked(opts *bind.WatchOpts, sink chan<- *StakingClaimedUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ClaimedUnstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimedUnstaked)
				if err := _Staking.contract.UnpackLog(event, "ClaimedUnstaked", log); err != nil {
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

// ParseClaimedUnstaked is a log parse operation binding the contract event 0x6eeaedc56832e751596e30e289453c8568daad613b8050876671fa4cced0fb8c.
//
// Solidity: event ClaimedUnstaked(address indexed user, uint256 total_claimed, uint256 total_staked)
func (_Staking *StakingFilterer) ParseClaimedUnstaked(log types.Log) (*StakingClaimedUnstaked, error) {
	event := new(StakingClaimedUnstaked)
	if err := _Staking.contract.UnpackLog(event, "ClaimedUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOdysseyRewardsClaimedIterator is returned from FilterOdysseyRewardsClaimed and is used to iterate over the raw logs and unpacked data for OdysseyRewardsClaimed events raised by the Staking contract.
type StakingOdysseyRewardsClaimedIterator struct {
	Event *StakingOdysseyRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakingOdysseyRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOdysseyRewardsClaimed)
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
		it.Event = new(StakingOdysseyRewardsClaimed)
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
func (it *StakingOdysseyRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOdysseyRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOdysseyRewardsClaimed represents a OdysseyRewardsClaimed event raised by the Staking contract.
type StakingOdysseyRewardsClaimed struct {
	OdysseyId           *big.Int
	TotalRewardsClaimed *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOdysseyRewardsClaimed is a free log retrieval operation binding the contract event 0x2bb8f99e42894785a1520f311e41259fa055076d23bbf66400646dcdc86d080f.
//
// Solidity: event OdysseyRewardsClaimed(uint256 indexed odyssey_id, uint256 total_rewards_claimed)
func (_Staking *StakingFilterer) FilterOdysseyRewardsClaimed(opts *bind.FilterOpts, odyssey_id []*big.Int) (*StakingOdysseyRewardsClaimedIterator, error) {

	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OdysseyRewardsClaimed", odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return &StakingOdysseyRewardsClaimedIterator{contract: _Staking.contract, event: "OdysseyRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchOdysseyRewardsClaimed is a free log subscription operation binding the contract event 0x2bb8f99e42894785a1520f311e41259fa055076d23bbf66400646dcdc86d080f.
//
// Solidity: event OdysseyRewardsClaimed(uint256 indexed odyssey_id, uint256 total_rewards_claimed)
func (_Staking *StakingFilterer) WatchOdysseyRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakingOdysseyRewardsClaimed, odyssey_id []*big.Int) (event.Subscription, error) {

	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OdysseyRewardsClaimed", odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOdysseyRewardsClaimed)
				if err := _Staking.contract.UnpackLog(event, "OdysseyRewardsClaimed", log); err != nil {
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

// ParseOdysseyRewardsClaimed is a log parse operation binding the contract event 0x2bb8f99e42894785a1520f311e41259fa055076d23bbf66400646dcdc86d080f.
//
// Solidity: event OdysseyRewardsClaimed(uint256 indexed odyssey_id, uint256 total_rewards_claimed)
func (_Staking *StakingFilterer) ParseOdysseyRewardsClaimed(log types.Log) (*StakingOdysseyRewardsClaimed, error) {
	event := new(StakingOdysseyRewardsClaimed)
	if err := _Staking.contract.UnpackLog(event, "OdysseyRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the Staking contract.
type StakingRestakeIterator struct {
	Event *StakingRestake // Event containing the contract specifics and raw log

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
func (it *StakingRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRestake)
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
		it.Event = new(StakingRestake)
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
func (it *StakingRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRestake represents a Restake event raised by the Staking contract.
type StakingRestake struct {
	User            common.Address
	OdysseyFrom     *big.Int
	OdysseyTo       *big.Int
	Amount          *big.Int
	Token           uint8
	TotalStakedFrom *big.Int
	TotalStakedTo   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0xf5b31220c1acf678c8a904f3a060a4b4a0bc4b9e42c9a49d318142bf4c282672.
//
// Solidity: event Restake(address indexed user, uint256 indexed odyssey_from, uint256 indexed odyssey_to, uint256 amount, uint8 token, uint256 total_staked_from, uint256 total_staked_to)
func (_Staking *StakingFilterer) FilterRestake(opts *bind.FilterOpts, user []common.Address, odyssey_from []*big.Int, odyssey_to []*big.Int) (*StakingRestakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var odyssey_fromRule []interface{}
	for _, odyssey_fromItem := range odyssey_from {
		odyssey_fromRule = append(odyssey_fromRule, odyssey_fromItem)
	}
	var odyssey_toRule []interface{}
	for _, odyssey_toItem := range odyssey_to {
		odyssey_toRule = append(odyssey_toRule, odyssey_toItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Restake", userRule, odyssey_fromRule, odyssey_toRule)
	if err != nil {
		return nil, err
	}
	return &StakingRestakeIterator{contract: _Staking.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0xf5b31220c1acf678c8a904f3a060a4b4a0bc4b9e42c9a49d318142bf4c282672.
//
// Solidity: event Restake(address indexed user, uint256 indexed odyssey_from, uint256 indexed odyssey_to, uint256 amount, uint8 token, uint256 total_staked_from, uint256 total_staked_to)
func (_Staking *StakingFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *StakingRestake, user []common.Address, odyssey_from []*big.Int, odyssey_to []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var odyssey_fromRule []interface{}
	for _, odyssey_fromItem := range odyssey_from {
		odyssey_fromRule = append(odyssey_fromRule, odyssey_fromItem)
	}
	var odyssey_toRule []interface{}
	for _, odyssey_toItem := range odyssey_to {
		odyssey_toRule = append(odyssey_toRule, odyssey_toItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Restake", userRule, odyssey_fromRule, odyssey_toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRestake)
				if err := _Staking.contract.UnpackLog(event, "Restake", log); err != nil {
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

// ParseRestake is a log parse operation binding the contract event 0xf5b31220c1acf678c8a904f3a060a4b4a0bc4b9e42c9a49d318142bf4c282672.
//
// Solidity: event Restake(address indexed user, uint256 indexed odyssey_from, uint256 indexed odyssey_to, uint256 amount, uint8 token, uint256 total_staked_from, uint256 total_staked_to)
func (_Staking *StakingFilterer) ParseRestake(log types.Log) (*StakingRestake, error) {
	event := new(StakingRestake)
	if err := _Staking.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Staking contract.
type StakingRewardsClaimedIterator struct {
	Event *StakingRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsClaimed)
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
		it.Event = new(StakingRewardsClaimed)
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
func (it *StakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsClaimed represents a RewardsClaimed event raised by the Staking contract.
type StakingRewardsClaimed struct {
	User                common.Address
	TotalRewardsClaimed *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 total_rewards_claimed)
func (_Staking *StakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address) (*StakingRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardsClaimedIterator{contract: _Staking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 total_rewards_claimed)
func (_Staking *StakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakingRewardsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsClaimed)
				if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 total_rewards_claimed)
func (_Staking *StakingFilterer) ParseRewardsClaimed(log types.Log) (*StakingRewardsClaimed, error) {
	event := new(StakingRewardsClaimed)
	if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsUpdatedIterator is returned from FilterRewardsUpdated and is used to iterate over the raw logs and unpacked data for RewardsUpdated events raised by the Staking contract.
type StakingRewardsUpdatedIterator struct {
	Event *StakingRewardsUpdated // Event containing the contract specifics and raw log

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
func (it *StakingRewardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsUpdated)
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
		it.Event = new(StakingRewardsUpdated)
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
func (it *StakingRewardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsUpdated represents a RewardsUpdated event raised by the Staking contract.
type StakingRewardsUpdated struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsUpdated is a free log retrieval operation binding the contract event 0xcd7c861d63893659123f817b192b8621ff1923a7c4589c260173cedcf9e986a6.
//
// Solidity: event RewardsUpdated(uint256 timestamp, uint256 blocknumber)
func (_Staking *StakingFilterer) FilterRewardsUpdated(opts *bind.FilterOpts) (*StakingRewardsUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingRewardsUpdatedIterator{contract: _Staking.contract, event: "RewardsUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsUpdated is a free log subscription operation binding the contract event 0xcd7c861d63893659123f817b192b8621ff1923a7c4589c260173cedcf9e986a6.
//
// Solidity: event RewardsUpdated(uint256 timestamp, uint256 blocknumber)
func (_Staking *StakingFilterer) WatchRewardsUpdated(opts *bind.WatchOpts, sink chan<- *StakingRewardsUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsUpdated)
				if err := _Staking.contract.UnpackLog(event, "RewardsUpdated", log); err != nil {
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

// ParseRewardsUpdated is a log parse operation binding the contract event 0xcd7c861d63893659123f817b192b8621ff1923a7c4589c260173cedcf9e986a6.
//
// Solidity: event RewardsUpdated(uint256 timestamp, uint256 blocknumber)
func (_Staking *StakingFilterer) ParseRewardsUpdated(log types.Log) (*StakingRewardsUpdated, error) {
	event := new(StakingRewardsUpdated)
	if err := _Staking.contract.UnpackLog(event, "RewardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Staking contract.
type StakingRoleAdminChangedIterator struct {
	Event *StakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleAdminChanged)
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
		it.Event = new(StakingRoleAdminChanged)
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
func (it *StakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleAdminChanged represents a RoleAdminChanged event raised by the Staking contract.
type StakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleAdminChangedIterator{contract: _Staking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleAdminChanged)
				if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleAdminChanged(log types.Log) (*StakingRoleAdminChanged, error) {
	event := new(StakingRoleAdminChanged)
	if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Staking contract.
type StakingRoleGrantedIterator struct {
	Event *StakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleGranted)
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
		it.Event = new(StakingRoleGranted)
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
func (it *StakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleGranted represents a RoleGranted event raised by the Staking contract.
type StakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleGrantedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleGrantedIterator{contract: _Staking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleGranted)
				if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleGranted(log types.Log) (*StakingRoleGranted, error) {
	event := new(StakingRoleGranted)
	if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Staking contract.
type StakingRoleRevokedIterator struct {
	Event *StakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleRevoked)
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
		it.Event = new(StakingRoleRevoked)
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
func (it *StakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleRevoked represents a RoleRevoked event raised by the Staking contract.
type StakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleRevokedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleRevokedIterator{contract: _Staking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleRevoked)
				if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleRevoked(log types.Log) (*StakingRoleRevoked, error) {
	event := new(StakingRoleRevoked)
	if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Staking contract.
type StakingStakeIterator struct {
	Event *StakingStake // Event containing the contract specifics and raw log

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
func (it *StakingStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStake)
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
		it.Event = new(StakingStake)
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
func (it *StakingStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStake represents a Stake event raised by the Staking contract.
type StakingStake struct {
	User         common.Address
	OdysseyId    *big.Int
	AmountStaked *big.Int
	Token        uint8
	TotalStaked  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0xc1a5c084e329922f8ec7c1f9d388080b5ce185980b99830170a9bd911e56ce7d.
//
// Solidity: event Stake(address indexed user, uint256 indexed odyssey_id, uint256 amount_staked, uint8 token, uint256 total_staked)
func (_Staking *StakingFilterer) FilterStake(opts *bind.FilterOpts, user []common.Address, odyssey_id []*big.Int) (*StakingStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Stake", userRule, odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeIterator{contract: _Staking.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0xc1a5c084e329922f8ec7c1f9d388080b5ce185980b99830170a9bd911e56ce7d.
//
// Solidity: event Stake(address indexed user, uint256 indexed odyssey_id, uint256 amount_staked, uint8 token, uint256 total_staked)
func (_Staking *StakingFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakingStake, user []common.Address, odyssey_id []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Stake", userRule, odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStake)
				if err := _Staking.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0xc1a5c084e329922f8ec7c1f9d388080b5ce185980b99830170a9bd911e56ce7d.
//
// Solidity: event Stake(address indexed user, uint256 indexed odyssey_id, uint256 amount_staked, uint8 token, uint256 total_staked)
func (_Staking *StakingFilterer) ParseStake(log types.Log) (*StakingStake, error) {
	event := new(StakingStake)
	if err := _Staking.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStateUpdatedIterator is returned from FilterStateUpdated and is used to iterate over the raw logs and unpacked data for StateUpdated events raised by the Staking contract.
type StakingStateUpdatedIterator struct {
	Event *StakingStateUpdated // Event containing the contract specifics and raw log

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
func (it *StakingStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStateUpdated)
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
		it.Event = new(StakingStateUpdated)
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
func (it *StakingStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStateUpdated represents a StateUpdated event raised by the Staking contract.
type StakingStateUpdated struct {
	State common.Hash
	From  *big.Int
	To    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateUpdated is a free log retrieval operation binding the contract event 0x348b835bd624f424a1ce88f82b74506124327d9b949e2b159c2744cb21310648.
//
// Solidity: event StateUpdated(string indexed state, uint256 from, uint256 to)
func (_Staking *StakingFilterer) FilterStateUpdated(opts *bind.FilterOpts, state []string) (*StakingStateUpdatedIterator, error) {

	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StateUpdated", stateRule)
	if err != nil {
		return nil, err
	}
	return &StakingStateUpdatedIterator{contract: _Staking.contract, event: "StateUpdated", logs: logs, sub: sub}, nil
}

// WatchStateUpdated is a free log subscription operation binding the contract event 0x348b835bd624f424a1ce88f82b74506124327d9b949e2b159c2744cb21310648.
//
// Solidity: event StateUpdated(string indexed state, uint256 from, uint256 to)
func (_Staking *StakingFilterer) WatchStateUpdated(opts *bind.WatchOpts, sink chan<- *StakingStateUpdated, state []string) (event.Subscription, error) {

	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StateUpdated", stateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStateUpdated)
				if err := _Staking.contract.UnpackLog(event, "StateUpdated", log); err != nil {
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

// ParseStateUpdated is a log parse operation binding the contract event 0x348b835bd624f424a1ce88f82b74506124327d9b949e2b159c2744cb21310648.
//
// Solidity: event StateUpdated(string indexed state, uint256 from, uint256 to)
func (_Staking *StakingFilterer) ParseStateUpdated(log types.Log) (*StakingStateUpdated, error) {
	event := new(StakingStateUpdated)
	if err := _Staking.contract.UnpackLog(event, "StateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the Staking contract.
type StakingUnstakeIterator struct {
	Event *StakingUnstake // Event containing the contract specifics and raw log

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
func (it *StakingUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstake)
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
		it.Event = new(StakingUnstake)
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
func (it *StakingUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstake represents a Unstake event raised by the Staking contract.
type StakingUnstake struct {
	User           common.Address
	OdysseyId      *big.Int
	AmountUnstaked *big.Int
	Token          uint8
	TotalStaked    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0xcd1b86c650bf0663e252966651bfd5312669a3bb6e38353a141fa368541e3735.
//
// Solidity: event Unstake(address indexed user, uint256 indexed odyssey_id, uint256 amount_unstaked, uint8 token, uint256 total_staked)
func (_Staking *StakingFilterer) FilterUnstake(opts *bind.FilterOpts, user []common.Address, odyssey_id []*big.Int) (*StakingUnstakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstake", userRule, odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakeIterator{contract: _Staking.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0xcd1b86c650bf0663e252966651bfd5312669a3bb6e38353a141fa368541e3735.
//
// Solidity: event Unstake(address indexed user, uint256 indexed odyssey_id, uint256 amount_unstaked, uint8 token, uint256 total_staked)
func (_Staking *StakingFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *StakingUnstake, user []common.Address, odyssey_id []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstake", userRule, odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstake)
				if err := _Staking.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0xcd1b86c650bf0663e252966651bfd5312669a3bb6e38353a141fa368541e3735.
//
// Solidity: event Unstake(address indexed user, uint256 indexed odyssey_id, uint256 amount_unstaked, uint8 token, uint256 total_staked)
func (_Staking *StakingFilterer) ParseUnstake(log types.Log) (*StakingUnstake, error) {
	event := new(StakingUnstake)
	if err := _Staking.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Staking contract.
type StakingUpgradedIterator struct {
	Event *StakingUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUpgraded)
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
		it.Event = new(StakingUpgraded)
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
func (it *StakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUpgraded represents a Upgraded event raised by the Staking contract.
type StakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Staking *StakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingUpgradedIterator{contract: _Staking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Staking *StakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUpgraded)
				if err := _Staking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUpgraded(log types.Log) (*StakingUpgraded, error) {
	event := new(StakingUpgraded)
	if err := _Staking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
