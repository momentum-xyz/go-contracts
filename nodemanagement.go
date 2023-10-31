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

// NodeManagementNode is an auto generated low-level Go binding around an user-defined struct.
type NodeManagementNode struct {
	NodeId      *big.Int
	Name        string
	Hostname    string
	Owner       common.Address
	Pubkey      []byte
	NodeAccount common.Address
}

// NodemanagementMetaData contains all meta data concerning the Nodemanagement contract.
var NodemanagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_fee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdatedEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_fee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdatedMom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from_node_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to_node_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"}],\"name\":\"NodeMgmtEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"old_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"old_hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"new_hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"old_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"new_name\",\"type\":\"string\"}],\"name\":\"NodeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"addNodeWithEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"addNodeWithMom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeMom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"node_account\",\"type\":\"address\"}],\"internalType\":\"structNodeManagement.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"}],\"name\":\"getNodeForTheOdyssey\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"node_account\",\"type\":\"address\"}],\"internalType\":\"structNodeManagement.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_odyssey_nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeMom\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mom_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mom_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"node_from_odyssey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"node_account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"odyssey_nft\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"odysseys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"odysseys_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"}],\"name\":\"removeMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"}],\"name\":\"removeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"challenge\",\"type\":\"bytes\"}],\"name\":\"setNodeMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"odyssey_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"challenge\",\"type\":\"bytes\"}],\"name\":\"setOdysseyMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"updateNodeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"node_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"new_pubkey\",\"type\":\"bytes\"}],\"name\":\"updateNodePubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_feeEth\",\"type\":\"uint256\"}],\"name\":\"update_feeEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"new_feeMom\",\"type\":\"uint256\"}],\"name\":\"update_feeMom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NodemanagementABI is the input ABI used to generate the binding from.
// Deprecated: Use NodemanagementMetaData.ABI instead.
var NodemanagementABI = NodemanagementMetaData.ABI

// Nodemanagement is an auto generated Go binding around an Ethereum contract.
type Nodemanagement struct {
	NodemanagementCaller     // Read-only binding to the contract
	NodemanagementTransactor // Write-only binding to the contract
	NodemanagementFilterer   // Log filterer for contract events
}

// NodemanagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodemanagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodemanagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodemanagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodemanagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodemanagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodemanagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodemanagementSession struct {
	Contract     *Nodemanagement   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodemanagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodemanagementCallerSession struct {
	Contract *NodemanagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodemanagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodemanagementTransactorSession struct {
	Contract     *NodemanagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodemanagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodemanagementRaw struct {
	Contract *Nodemanagement // Generic contract binding to access the raw methods on
}

// NodemanagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodemanagementCallerRaw struct {
	Contract *NodemanagementCaller // Generic read-only contract binding to access the raw methods on
}

// NodemanagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodemanagementTransactorRaw struct {
	Contract *NodemanagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodemanagement creates a new instance of Nodemanagement, bound to a specific deployed contract.
func NewNodemanagement(address common.Address, backend bind.ContractBackend) (*Nodemanagement, error) {
	contract, err := bindNodemanagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nodemanagement{NodemanagementCaller: NodemanagementCaller{contract: contract}, NodemanagementTransactor: NodemanagementTransactor{contract: contract}, NodemanagementFilterer: NodemanagementFilterer{contract: contract}}, nil
}

// NewNodemanagementCaller creates a new read-only instance of Nodemanagement, bound to a specific deployed contract.
func NewNodemanagementCaller(address common.Address, caller bind.ContractCaller) (*NodemanagementCaller, error) {
	contract, err := bindNodemanagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodemanagementCaller{contract: contract}, nil
}

// NewNodemanagementTransactor creates a new write-only instance of Nodemanagement, bound to a specific deployed contract.
func NewNodemanagementTransactor(address common.Address, transactor bind.ContractTransactor) (*NodemanagementTransactor, error) {
	contract, err := bindNodemanagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodemanagementTransactor{contract: contract}, nil
}

// NewNodemanagementFilterer creates a new log filterer instance of Nodemanagement, bound to a specific deployed contract.
func NewNodemanagementFilterer(address common.Address, filterer bind.ContractFilterer) (*NodemanagementFilterer, error) {
	contract, err := bindNodemanagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodemanagementFilterer{contract: contract}, nil
}

// bindNodemanagement binds a generic wrapper to an already deployed contract.
func bindNodemanagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodemanagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodemanagement *NodemanagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodemanagement.Contract.NodemanagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodemanagement *NodemanagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodemanagement.Contract.NodemanagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodemanagement *NodemanagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodemanagement.Contract.NodemanagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodemanagement *NodemanagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodemanagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodemanagement *NodemanagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodemanagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodemanagement *NodemanagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodemanagement.Contract.contract.Transact(opts, method, params...)
}

// FeeETH is a free data retrieval call binding the contract method 0x4078fca8.
//
// Solidity: function feeETH() view returns(uint256)
func (_Nodemanagement *NodemanagementCaller) FeeETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "feeETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeETH is a free data retrieval call binding the contract method 0x4078fca8.
//
// Solidity: function feeETH() view returns(uint256)
func (_Nodemanagement *NodemanagementSession) FeeETH() (*big.Int, error) {
	return _Nodemanagement.Contract.FeeETH(&_Nodemanagement.CallOpts)
}

// FeeETH is a free data retrieval call binding the contract method 0x4078fca8.
//
// Solidity: function feeETH() view returns(uint256)
func (_Nodemanagement *NodemanagementCallerSession) FeeETH() (*big.Int, error) {
	return _Nodemanagement.Contract.FeeETH(&_Nodemanagement.CallOpts)
}

// FeeMom is a free data retrieval call binding the contract method 0x47ad1b70.
//
// Solidity: function feeMom() view returns(uint256)
func (_Nodemanagement *NodemanagementCaller) FeeMom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "feeMom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeMom is a free data retrieval call binding the contract method 0x47ad1b70.
//
// Solidity: function feeMom() view returns(uint256)
func (_Nodemanagement *NodemanagementSession) FeeMom() (*big.Int, error) {
	return _Nodemanagement.Contract.FeeMom(&_Nodemanagement.CallOpts)
}

// FeeMom is a free data retrieval call binding the contract method 0x47ad1b70.
//
// Solidity: function feeMom() view returns(uint256)
func (_Nodemanagement *NodemanagementCallerSession) FeeMom() (*big.Int, error) {
	return _Nodemanagement.Contract.FeeMom(&_Nodemanagement.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 node_id) view returns((uint256,string,string,address,bytes,address))
func (_Nodemanagement *NodemanagementCaller) GetNode(opts *bind.CallOpts, node_id *big.Int) (NodeManagementNode, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "getNode", node_id)

	if err != nil {
		return *new(NodeManagementNode), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeManagementNode)).(*NodeManagementNode)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 node_id) view returns((uint256,string,string,address,bytes,address))
func (_Nodemanagement *NodemanagementSession) GetNode(node_id *big.Int) (NodeManagementNode, error) {
	return _Nodemanagement.Contract.GetNode(&_Nodemanagement.CallOpts, node_id)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 node_id) view returns((uint256,string,string,address,bytes,address))
func (_Nodemanagement *NodemanagementCallerSession) GetNode(node_id *big.Int) (NodeManagementNode, error) {
	return _Nodemanagement.Contract.GetNode(&_Nodemanagement.CallOpts, node_id)
}

// GetNodeForTheOdyssey is a free data retrieval call binding the contract method 0x2752c044.
//
// Solidity: function getNodeForTheOdyssey(uint256 odyssey_id) view returns((uint256,string,string,address,bytes,address))
func (_Nodemanagement *NodemanagementCaller) GetNodeForTheOdyssey(opts *bind.CallOpts, odyssey_id *big.Int) (NodeManagementNode, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "getNodeForTheOdyssey", odyssey_id)

	if err != nil {
		return *new(NodeManagementNode), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeManagementNode)).(*NodeManagementNode)

	return out0, err

}

// GetNodeForTheOdyssey is a free data retrieval call binding the contract method 0x2752c044.
//
// Solidity: function getNodeForTheOdyssey(uint256 odyssey_id) view returns((uint256,string,string,address,bytes,address))
func (_Nodemanagement *NodemanagementSession) GetNodeForTheOdyssey(odyssey_id *big.Int) (NodeManagementNode, error) {
	return _Nodemanagement.Contract.GetNodeForTheOdyssey(&_Nodemanagement.CallOpts, odyssey_id)
}

// GetNodeForTheOdyssey is a free data retrieval call binding the contract method 0x2752c044.
//
// Solidity: function getNodeForTheOdyssey(uint256 odyssey_id) view returns((uint256,string,string,address,bytes,address))
func (_Nodemanagement *NodemanagementCallerSession) GetNodeForTheOdyssey(odyssey_id *big.Int) (NodeManagementNode, error) {
	return _Nodemanagement.Contract.GetNodeForTheOdyssey(&_Nodemanagement.CallOpts, odyssey_id)
}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Nodemanagement *NodemanagementCaller) MomToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "mom_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Nodemanagement *NodemanagementSession) MomToken() (common.Address, error) {
	return _Nodemanagement.Contract.MomToken(&_Nodemanagement.CallOpts)
}

// MomToken is a free data retrieval call binding the contract method 0x3a2bb837.
//
// Solidity: function mom_token() view returns(address)
func (_Nodemanagement *NodemanagementCallerSession) MomToken() (common.Address, error) {
	return _Nodemanagement.Contract.MomToken(&_Nodemanagement.CallOpts)
}

// NodeFromOdyssey is a free data retrieval call binding the contract method 0xc5e8000e.
//
// Solidity: function node_from_odyssey(uint256 ) view returns(uint256 node_id, uint256 index)
func (_Nodemanagement *NodemanagementCaller) NodeFromOdyssey(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NodeId *big.Int
	Index  *big.Int
}, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "node_from_odyssey", arg0)

	outstruct := new(struct {
		NodeId *big.Int
		Index  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NodeFromOdyssey is a free data retrieval call binding the contract method 0xc5e8000e.
//
// Solidity: function node_from_odyssey(uint256 ) view returns(uint256 node_id, uint256 index)
func (_Nodemanagement *NodemanagementSession) NodeFromOdyssey(arg0 *big.Int) (struct {
	NodeId *big.Int
	Index  *big.Int
}, error) {
	return _Nodemanagement.Contract.NodeFromOdyssey(&_Nodemanagement.CallOpts, arg0)
}

// NodeFromOdyssey is a free data retrieval call binding the contract method 0xc5e8000e.
//
// Solidity: function node_from_odyssey(uint256 ) view returns(uint256 node_id, uint256 index)
func (_Nodemanagement *NodemanagementCallerSession) NodeFromOdyssey(arg0 *big.Int) (struct {
	NodeId *big.Int
	Index  *big.Int
}, error) {
	return _Nodemanagement.Contract.NodeFromOdyssey(&_Nodemanagement.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(uint256 node_id, string name, string hostname, address owner, bytes pubkey, address node_account)
func (_Nodemanagement *NodemanagementCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NodeId      *big.Int
	Name        string
	Hostname    string
	Owner       common.Address
	Pubkey      []byte
	NodeAccount common.Address
}, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		NodeId      *big.Int
		Name        string
		Hostname    string
		Owner       common.Address
		Pubkey      []byte
		NodeAccount common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Hostname = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Pubkey = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.NodeAccount = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(uint256 node_id, string name, string hostname, address owner, bytes pubkey, address node_account)
func (_Nodemanagement *NodemanagementSession) Nodes(arg0 *big.Int) (struct {
	NodeId      *big.Int
	Name        string
	Hostname    string
	Owner       common.Address
	Pubkey      []byte
	NodeAccount common.Address
}, error) {
	return _Nodemanagement.Contract.Nodes(&_Nodemanagement.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(uint256 node_id, string name, string hostname, address owner, bytes pubkey, address node_account)
func (_Nodemanagement *NodemanagementCallerSession) Nodes(arg0 *big.Int) (struct {
	NodeId      *big.Int
	Name        string
	Hostname    string
	Owner       common.Address
	Pubkey      []byte
	NodeAccount common.Address
}, error) {
	return _Nodemanagement.Contract.Nodes(&_Nodemanagement.CallOpts, arg0)
}

// NodesCount is a free data retrieval call binding the contract method 0xf1a3c5b3.
//
// Solidity: function nodesCount() view returns(uint256)
func (_Nodemanagement *NodemanagementCaller) NodesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "nodesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodesCount is a free data retrieval call binding the contract method 0xf1a3c5b3.
//
// Solidity: function nodesCount() view returns(uint256)
func (_Nodemanagement *NodemanagementSession) NodesCount() (*big.Int, error) {
	return _Nodemanagement.Contract.NodesCount(&_Nodemanagement.CallOpts)
}

// NodesCount is a free data retrieval call binding the contract method 0xf1a3c5b3.
//
// Solidity: function nodesCount() view returns(uint256)
func (_Nodemanagement *NodemanagementCallerSession) NodesCount() (*big.Int, error) {
	return _Nodemanagement.Contract.NodesCount(&_Nodemanagement.CallOpts)
}

// NodesIndex is a free data retrieval call binding the contract method 0xf452ba40.
//
// Solidity: function nodes_index(uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementCaller) NodesIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "nodes_index", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodesIndex is a free data retrieval call binding the contract method 0xf452ba40.
//
// Solidity: function nodes_index(uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementSession) NodesIndex(arg0 *big.Int) (*big.Int, error) {
	return _Nodemanagement.Contract.NodesIndex(&_Nodemanagement.CallOpts, arg0)
}

// NodesIndex is a free data retrieval call binding the contract method 0xf452ba40.
//
// Solidity: function nodes_index(uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementCallerSession) NodesIndex(arg0 *big.Int) (*big.Int, error) {
	return _Nodemanagement.Contract.NodesIndex(&_Nodemanagement.CallOpts, arg0)
}

// OdysseyNft is a free data retrieval call binding the contract method 0x98edac36.
//
// Solidity: function odyssey_nft() view returns(address)
func (_Nodemanagement *NodemanagementCaller) OdysseyNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "odyssey_nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OdysseyNft is a free data retrieval call binding the contract method 0x98edac36.
//
// Solidity: function odyssey_nft() view returns(address)
func (_Nodemanagement *NodemanagementSession) OdysseyNft() (common.Address, error) {
	return _Nodemanagement.Contract.OdysseyNft(&_Nodemanagement.CallOpts)
}

// OdysseyNft is a free data retrieval call binding the contract method 0x98edac36.
//
// Solidity: function odyssey_nft() view returns(address)
func (_Nodemanagement *NodemanagementCallerSession) OdysseyNft() (common.Address, error) {
	return _Nodemanagement.Contract.OdysseyNft(&_Nodemanagement.CallOpts)
}

// Odysseys is a free data retrieval call binding the contract method 0x1c257a13.
//
// Solidity: function odysseys(uint256 , uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementCaller) Odysseys(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "odysseys", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Odysseys is a free data retrieval call binding the contract method 0x1c257a13.
//
// Solidity: function odysseys(uint256 , uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementSession) Odysseys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Nodemanagement.Contract.Odysseys(&_Nodemanagement.CallOpts, arg0, arg1)
}

// Odysseys is a free data retrieval call binding the contract method 0x1c257a13.
//
// Solidity: function odysseys(uint256 , uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementCallerSession) Odysseys(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Nodemanagement.Contract.Odysseys(&_Nodemanagement.CallOpts, arg0, arg1)
}

// OdysseysIndex is a free data retrieval call binding the contract method 0x16c859bb.
//
// Solidity: function odysseys_index(uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementCaller) OdysseysIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "odysseys_index", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OdysseysIndex is a free data retrieval call binding the contract method 0x16c859bb.
//
// Solidity: function odysseys_index(uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementSession) OdysseysIndex(arg0 *big.Int) (*big.Int, error) {
	return _Nodemanagement.Contract.OdysseysIndex(&_Nodemanagement.CallOpts, arg0)
}

// OdysseysIndex is a free data retrieval call binding the contract method 0x16c859bb.
//
// Solidity: function odysseys_index(uint256 ) view returns(uint256)
func (_Nodemanagement *NodemanagementCallerSession) OdysseysIndex(arg0 *big.Int) (*big.Int, error) {
	return _Nodemanagement.Contract.OdysseysIndex(&_Nodemanagement.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodemanagement *NodemanagementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodemanagement *NodemanagementSession) Owner() (common.Address, error) {
	return _Nodemanagement.Contract.Owner(&_Nodemanagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodemanagement *NodemanagementCallerSession) Owner() (common.Address, error) {
	return _Nodemanagement.Contract.Owner(&_Nodemanagement.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodemanagement *NodemanagementCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodemanagement *NodemanagementSession) ProxiableUUID() ([32]byte, error) {
	return _Nodemanagement.Contract.ProxiableUUID(&_Nodemanagement.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodemanagement *NodemanagementCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Nodemanagement.Contract.ProxiableUUID(&_Nodemanagement.CallOpts)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(uint8, bytes32, bytes32)
func (_Nodemanagement *NodemanagementCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (uint8, [32]byte, [32]byte, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "splitSignature", sig)

	if err != nil {
		return *new(uint8), *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return out0, out1, out2, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(uint8, bytes32, bytes32)
func (_Nodemanagement *NodemanagementSession) SplitSignature(sig []byte) (uint8, [32]byte, [32]byte, error) {
	return _Nodemanagement.Contract.SplitSignature(&_Nodemanagement.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(uint8, bytes32, bytes32)
func (_Nodemanagement *NodemanagementCallerSession) SplitSignature(sig []byte) (uint8, [32]byte, [32]byte, error) {
	return _Nodemanagement.Contract.SplitSignature(&_Nodemanagement.CallOpts, sig)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Nodemanagement *NodemanagementCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodemanagement.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Nodemanagement *NodemanagementSession) Treasury() (common.Address, error) {
	return _Nodemanagement.Contract.Treasury(&_Nodemanagement.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Nodemanagement *NodemanagementCallerSession) Treasury() (common.Address, error) {
	return _Nodemanagement.Contract.Treasury(&_Nodemanagement.CallOpts)
}

// AddNodeWithEth is a paid mutator transaction binding the contract method 0x69541c1c.
//
// Solidity: function addNodeWithEth(uint256 node_id, string hostname, string name, bytes pubkey) returns()
func (_Nodemanagement *NodemanagementTransactor) AddNodeWithEth(opts *bind.TransactOpts, node_id *big.Int, hostname string, name string, pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "addNodeWithEth", node_id, hostname, name, pubkey)
}

// AddNodeWithEth is a paid mutator transaction binding the contract method 0x69541c1c.
//
// Solidity: function addNodeWithEth(uint256 node_id, string hostname, string name, bytes pubkey) returns()
func (_Nodemanagement *NodemanagementSession) AddNodeWithEth(node_id *big.Int, hostname string, name string, pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.AddNodeWithEth(&_Nodemanagement.TransactOpts, node_id, hostname, name, pubkey)
}

// AddNodeWithEth is a paid mutator transaction binding the contract method 0x69541c1c.
//
// Solidity: function addNodeWithEth(uint256 node_id, string hostname, string name, bytes pubkey) returns()
func (_Nodemanagement *NodemanagementTransactorSession) AddNodeWithEth(node_id *big.Int, hostname string, name string, pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.AddNodeWithEth(&_Nodemanagement.TransactOpts, node_id, hostname, name, pubkey)
}

// AddNodeWithMom is a paid mutator transaction binding the contract method 0x36fe2fd4.
//
// Solidity: function addNodeWithMom(uint256 node_id, string hostname, string name, bytes pubkey) returns()
func (_Nodemanagement *NodemanagementTransactor) AddNodeWithMom(opts *bind.TransactOpts, node_id *big.Int, hostname string, name string, pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "addNodeWithMom", node_id, hostname, name, pubkey)
}

// AddNodeWithMom is a paid mutator transaction binding the contract method 0x36fe2fd4.
//
// Solidity: function addNodeWithMom(uint256 node_id, string hostname, string name, bytes pubkey) returns()
func (_Nodemanagement *NodemanagementSession) AddNodeWithMom(node_id *big.Int, hostname string, name string, pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.AddNodeWithMom(&_Nodemanagement.TransactOpts, node_id, hostname, name, pubkey)
}

// AddNodeWithMom is a paid mutator transaction binding the contract method 0x36fe2fd4.
//
// Solidity: function addNodeWithMom(uint256 node_id, string hostname, string name, bytes pubkey) returns()
func (_Nodemanagement *NodemanagementTransactorSession) AddNodeWithMom(node_id *big.Int, hostname string, name string, pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.AddNodeWithMom(&_Nodemanagement.TransactOpts, node_id, hostname, name, pubkey)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _odyssey_nft, address _treasury, uint256 _feeEth, uint256 _feeMom, address _mom_token) returns()
func (_Nodemanagement *NodemanagementTransactor) Initialize(opts *bind.TransactOpts, _odyssey_nft common.Address, _treasury common.Address, _feeEth *big.Int, _feeMom *big.Int, _mom_token common.Address) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "initialize", _odyssey_nft, _treasury, _feeEth, _feeMom, _mom_token)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _odyssey_nft, address _treasury, uint256 _feeEth, uint256 _feeMom, address _mom_token) returns()
func (_Nodemanagement *NodemanagementSession) Initialize(_odyssey_nft common.Address, _treasury common.Address, _feeEth *big.Int, _feeMom *big.Int, _mom_token common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.Initialize(&_Nodemanagement.TransactOpts, _odyssey_nft, _treasury, _feeEth, _feeMom, _mom_token)
}

// Initialize is a paid mutator transaction binding the contract method 0x57fb25cc.
//
// Solidity: function initialize(address _odyssey_nft, address _treasury, uint256 _feeEth, uint256 _feeMom, address _mom_token) returns()
func (_Nodemanagement *NodemanagementTransactorSession) Initialize(_odyssey_nft common.Address, _treasury common.Address, _feeEth *big.Int, _feeMom *big.Int, _mom_token common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.Initialize(&_Nodemanagement.TransactOpts, _odyssey_nft, _treasury, _feeEth, _feeMom, _mom_token)
}

// RemoveMapping is a paid mutator transaction binding the contract method 0x57bed438.
//
// Solidity: function removeMapping(uint256 node_id, uint256 odyssey_id) returns()
func (_Nodemanagement *NodemanagementTransactor) RemoveMapping(opts *bind.TransactOpts, node_id *big.Int, odyssey_id *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "removeMapping", node_id, odyssey_id)
}

// RemoveMapping is a paid mutator transaction binding the contract method 0x57bed438.
//
// Solidity: function removeMapping(uint256 node_id, uint256 odyssey_id) returns()
func (_Nodemanagement *NodemanagementSession) RemoveMapping(node_id *big.Int, odyssey_id *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.RemoveMapping(&_Nodemanagement.TransactOpts, node_id, odyssey_id)
}

// RemoveMapping is a paid mutator transaction binding the contract method 0x57bed438.
//
// Solidity: function removeMapping(uint256 node_id, uint256 odyssey_id) returns()
func (_Nodemanagement *NodemanagementTransactorSession) RemoveMapping(node_id *big.Int, odyssey_id *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.RemoveMapping(&_Nodemanagement.TransactOpts, node_id, odyssey_id)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 node_id) returns()
func (_Nodemanagement *NodemanagementTransactor) RemoveNode(opts *bind.TransactOpts, node_id *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "removeNode", node_id)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 node_id) returns()
func (_Nodemanagement *NodemanagementSession) RemoveNode(node_id *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.RemoveNode(&_Nodemanagement.TransactOpts, node_id)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 node_id) returns()
func (_Nodemanagement *NodemanagementTransactorSession) RemoveNode(node_id *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.RemoveNode(&_Nodemanagement.TransactOpts, node_id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodemanagement *NodemanagementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodemanagement *NodemanagementSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nodemanagement.Contract.RenounceOwnership(&_Nodemanagement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodemanagement *NodemanagementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nodemanagement.Contract.RenounceOwnership(&_Nodemanagement.TransactOpts)
}

// SetNodeMapping is a paid mutator transaction binding the contract method 0x76b32e37.
//
// Solidity: function setNodeMapping(uint256 node_id, uint256 odyssey_id, bytes challenge) returns()
func (_Nodemanagement *NodemanagementTransactor) SetNodeMapping(opts *bind.TransactOpts, node_id *big.Int, odyssey_id *big.Int, challenge []byte) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "setNodeMapping", node_id, odyssey_id, challenge)
}

// SetNodeMapping is a paid mutator transaction binding the contract method 0x76b32e37.
//
// Solidity: function setNodeMapping(uint256 node_id, uint256 odyssey_id, bytes challenge) returns()
func (_Nodemanagement *NodemanagementSession) SetNodeMapping(node_id *big.Int, odyssey_id *big.Int, challenge []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.SetNodeMapping(&_Nodemanagement.TransactOpts, node_id, odyssey_id, challenge)
}

// SetNodeMapping is a paid mutator transaction binding the contract method 0x76b32e37.
//
// Solidity: function setNodeMapping(uint256 node_id, uint256 odyssey_id, bytes challenge) returns()
func (_Nodemanagement *NodemanagementTransactorSession) SetNodeMapping(node_id *big.Int, odyssey_id *big.Int, challenge []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.SetNodeMapping(&_Nodemanagement.TransactOpts, node_id, odyssey_id, challenge)
}

// SetOdysseyMapping is a paid mutator transaction binding the contract method 0x805e408c.
//
// Solidity: function setOdysseyMapping(uint256 node_id, uint256 odyssey_id, bytes challenge) returns()
func (_Nodemanagement *NodemanagementTransactor) SetOdysseyMapping(opts *bind.TransactOpts, node_id *big.Int, odyssey_id *big.Int, challenge []byte) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "setOdysseyMapping", node_id, odyssey_id, challenge)
}

// SetOdysseyMapping is a paid mutator transaction binding the contract method 0x805e408c.
//
// Solidity: function setOdysseyMapping(uint256 node_id, uint256 odyssey_id, bytes challenge) returns()
func (_Nodemanagement *NodemanagementSession) SetOdysseyMapping(node_id *big.Int, odyssey_id *big.Int, challenge []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.SetOdysseyMapping(&_Nodemanagement.TransactOpts, node_id, odyssey_id, challenge)
}

// SetOdysseyMapping is a paid mutator transaction binding the contract method 0x805e408c.
//
// Solidity: function setOdysseyMapping(uint256 node_id, uint256 odyssey_id, bytes challenge) returns()
func (_Nodemanagement *NodemanagementTransactorSession) SetOdysseyMapping(node_id *big.Int, odyssey_id *big.Int, challenge []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.SetOdysseyMapping(&_Nodemanagement.TransactOpts, node_id, odyssey_id, challenge)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodemanagement *NodemanagementTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodemanagement *NodemanagementSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.TransferOwnership(&_Nodemanagement.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodemanagement *NodemanagementTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.TransferOwnership(&_Nodemanagement.TransactOpts, newOwner)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x6cc0243e.
//
// Solidity: function updateNode(uint256 node_id, string hostname, string name) returns()
func (_Nodemanagement *NodemanagementTransactor) UpdateNode(opts *bind.TransactOpts, node_id *big.Int, hostname string, name string) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "updateNode", node_id, hostname, name)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x6cc0243e.
//
// Solidity: function updateNode(uint256 node_id, string hostname, string name) returns()
func (_Nodemanagement *NodemanagementSession) UpdateNode(node_id *big.Int, hostname string, name string) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateNode(&_Nodemanagement.TransactOpts, node_id, hostname, name)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x6cc0243e.
//
// Solidity: function updateNode(uint256 node_id, string hostname, string name) returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpdateNode(node_id *big.Int, hostname string, name string) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateNode(&_Nodemanagement.TransactOpts, node_id, hostname, name)
}

// UpdateNodeOwner is a paid mutator transaction binding the contract method 0x97749506.
//
// Solidity: function updateNodeOwner(uint256 node_id, address new_owner) returns()
func (_Nodemanagement *NodemanagementTransactor) UpdateNodeOwner(opts *bind.TransactOpts, node_id *big.Int, new_owner common.Address) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "updateNodeOwner", node_id, new_owner)
}

// UpdateNodeOwner is a paid mutator transaction binding the contract method 0x97749506.
//
// Solidity: function updateNodeOwner(uint256 node_id, address new_owner) returns()
func (_Nodemanagement *NodemanagementSession) UpdateNodeOwner(node_id *big.Int, new_owner common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateNodeOwner(&_Nodemanagement.TransactOpts, node_id, new_owner)
}

// UpdateNodeOwner is a paid mutator transaction binding the contract method 0x97749506.
//
// Solidity: function updateNodeOwner(uint256 node_id, address new_owner) returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpdateNodeOwner(node_id *big.Int, new_owner common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateNodeOwner(&_Nodemanagement.TransactOpts, node_id, new_owner)
}

// UpdateNodePubkey is a paid mutator transaction binding the contract method 0x64d7d365.
//
// Solidity: function updateNodePubkey(uint256 node_id, bytes new_pubkey) returns()
func (_Nodemanagement *NodemanagementTransactor) UpdateNodePubkey(opts *bind.TransactOpts, node_id *big.Int, new_pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "updateNodePubkey", node_id, new_pubkey)
}

// UpdateNodePubkey is a paid mutator transaction binding the contract method 0x64d7d365.
//
// Solidity: function updateNodePubkey(uint256 node_id, bytes new_pubkey) returns()
func (_Nodemanagement *NodemanagementSession) UpdateNodePubkey(node_id *big.Int, new_pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateNodePubkey(&_Nodemanagement.TransactOpts, node_id, new_pubkey)
}

// UpdateNodePubkey is a paid mutator transaction binding the contract method 0x64d7d365.
//
// Solidity: function updateNodePubkey(uint256 node_id, bytes new_pubkey) returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpdateNodePubkey(node_id *big.Int, new_pubkey []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateNodePubkey(&_Nodemanagement.TransactOpts, node_id, new_pubkey)
}

// UpdateFeeEth is a paid mutator transaction binding the contract method 0x84182ceb.
//
// Solidity: function update_feeEth(uint256 new_feeEth) returns()
func (_Nodemanagement *NodemanagementTransactor) UpdateFeeEth(opts *bind.TransactOpts, new_feeEth *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "update_feeEth", new_feeEth)
}

// UpdateFeeEth is a paid mutator transaction binding the contract method 0x84182ceb.
//
// Solidity: function update_feeEth(uint256 new_feeEth) returns()
func (_Nodemanagement *NodemanagementSession) UpdateFeeEth(new_feeEth *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateFeeEth(&_Nodemanagement.TransactOpts, new_feeEth)
}

// UpdateFeeEth is a paid mutator transaction binding the contract method 0x84182ceb.
//
// Solidity: function update_feeEth(uint256 new_feeEth) returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpdateFeeEth(new_feeEth *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateFeeEth(&_Nodemanagement.TransactOpts, new_feeEth)
}

// UpdateFeeMom is a paid mutator transaction binding the contract method 0x9a900650.
//
// Solidity: function update_feeMom(uint256 new_feeMom) returns()
func (_Nodemanagement *NodemanagementTransactor) UpdateFeeMom(opts *bind.TransactOpts, new_feeMom *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "update_feeMom", new_feeMom)
}

// UpdateFeeMom is a paid mutator transaction binding the contract method 0x9a900650.
//
// Solidity: function update_feeMom(uint256 new_feeMom) returns()
func (_Nodemanagement *NodemanagementSession) UpdateFeeMom(new_feeMom *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateFeeMom(&_Nodemanagement.TransactOpts, new_feeMom)
}

// UpdateFeeMom is a paid mutator transaction binding the contract method 0x9a900650.
//
// Solidity: function update_feeMom(uint256 new_feeMom) returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpdateFeeMom(new_feeMom *big.Int) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpdateFeeMom(&_Nodemanagement.TransactOpts, new_feeMom)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodemanagement *NodemanagementTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodemanagement *NodemanagementSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpgradeTo(&_Nodemanagement.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpgradeTo(&_Nodemanagement.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodemanagement *NodemanagementTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodemanagement.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodemanagement *NodemanagementSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpgradeToAndCall(&_Nodemanagement.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodemanagement *NodemanagementTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodemanagement.Contract.UpgradeToAndCall(&_Nodemanagement.TransactOpts, newImplementation, data)
}

// NodemanagementAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Nodemanagement contract.
type NodemanagementAdminChangedIterator struct {
	Event *NodemanagementAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodemanagementAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementAdminChanged)
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
		it.Event = new(NodemanagementAdminChanged)
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
func (it *NodemanagementAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementAdminChanged represents a AdminChanged event raised by the Nodemanagement contract.
type NodemanagementAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Nodemanagement *NodemanagementFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NodemanagementAdminChangedIterator, error) {

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NodemanagementAdminChangedIterator{contract: _Nodemanagement.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Nodemanagement *NodemanagementFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NodemanagementAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementAdminChanged)
				if err := _Nodemanagement.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Nodemanagement *NodemanagementFilterer) ParseAdminChanged(log types.Log) (*NodemanagementAdminChanged, error) {
	event := new(NodemanagementAdminChanged)
	if err := _Nodemanagement.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Nodemanagement contract.
type NodemanagementBeaconUpgradedIterator struct {
	Event *NodemanagementBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NodemanagementBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementBeaconUpgraded)
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
		it.Event = new(NodemanagementBeaconUpgraded)
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
func (it *NodemanagementBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementBeaconUpgraded represents a BeaconUpgraded event raised by the Nodemanagement contract.
type NodemanagementBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Nodemanagement *NodemanagementFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NodemanagementBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NodemanagementBeaconUpgradedIterator{contract: _Nodemanagement.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Nodemanagement *NodemanagementFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NodemanagementBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementBeaconUpgraded)
				if err := _Nodemanagement.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Nodemanagement *NodemanagementFilterer) ParseBeaconUpgraded(log types.Log) (*NodemanagementBeaconUpgraded, error) {
	event := new(NodemanagementBeaconUpgraded)
	if err := _Nodemanagement.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementFeeUpdatedEthIterator is returned from FilterFeeUpdatedEth and is used to iterate over the raw logs and unpacked data for FeeUpdatedEth events raised by the Nodemanagement contract.
type NodemanagementFeeUpdatedEthIterator struct {
	Event *NodemanagementFeeUpdatedEth // Event containing the contract specifics and raw log

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
func (it *NodemanagementFeeUpdatedEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementFeeUpdatedEth)
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
		it.Event = new(NodemanagementFeeUpdatedEth)
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
func (it *NodemanagementFeeUpdatedEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementFeeUpdatedEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementFeeUpdatedEth represents a FeeUpdatedEth event raised by the Nodemanagement contract.
type NodemanagementFeeUpdatedEth struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdatedEth is a free log retrieval operation binding the contract event 0xabad68c916849b2d3dc65063fb53673e2afbca565ccafa6d879693ee2dc6eae1.
//
// Solidity: event FeeUpdatedEth(uint256 old_fee, uint256 new_fee)
func (_Nodemanagement *NodemanagementFilterer) FilterFeeUpdatedEth(opts *bind.FilterOpts) (*NodemanagementFeeUpdatedEthIterator, error) {

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "FeeUpdatedEth")
	if err != nil {
		return nil, err
	}
	return &NodemanagementFeeUpdatedEthIterator{contract: _Nodemanagement.contract, event: "FeeUpdatedEth", logs: logs, sub: sub}, nil
}

// WatchFeeUpdatedEth is a free log subscription operation binding the contract event 0xabad68c916849b2d3dc65063fb53673e2afbca565ccafa6d879693ee2dc6eae1.
//
// Solidity: event FeeUpdatedEth(uint256 old_fee, uint256 new_fee)
func (_Nodemanagement *NodemanagementFilterer) WatchFeeUpdatedEth(opts *bind.WatchOpts, sink chan<- *NodemanagementFeeUpdatedEth) (event.Subscription, error) {

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "FeeUpdatedEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementFeeUpdatedEth)
				if err := _Nodemanagement.contract.UnpackLog(event, "FeeUpdatedEth", log); err != nil {
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

// ParseFeeUpdatedEth is a log parse operation binding the contract event 0xabad68c916849b2d3dc65063fb53673e2afbca565ccafa6d879693ee2dc6eae1.
//
// Solidity: event FeeUpdatedEth(uint256 old_fee, uint256 new_fee)
func (_Nodemanagement *NodemanagementFilterer) ParseFeeUpdatedEth(log types.Log) (*NodemanagementFeeUpdatedEth, error) {
	event := new(NodemanagementFeeUpdatedEth)
	if err := _Nodemanagement.contract.UnpackLog(event, "FeeUpdatedEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementFeeUpdatedMomIterator is returned from FilterFeeUpdatedMom and is used to iterate over the raw logs and unpacked data for FeeUpdatedMom events raised by the Nodemanagement contract.
type NodemanagementFeeUpdatedMomIterator struct {
	Event *NodemanagementFeeUpdatedMom // Event containing the contract specifics and raw log

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
func (it *NodemanagementFeeUpdatedMomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementFeeUpdatedMom)
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
		it.Event = new(NodemanagementFeeUpdatedMom)
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
func (it *NodemanagementFeeUpdatedMomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementFeeUpdatedMomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementFeeUpdatedMom represents a FeeUpdatedMom event raised by the Nodemanagement contract.
type NodemanagementFeeUpdatedMom struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdatedMom is a free log retrieval operation binding the contract event 0x8173aa58a259e63faf706f8a21a9f36cc8177d015a0a272c554bce99ef47188b.
//
// Solidity: event FeeUpdatedMom(uint256 old_fee, uint256 new_fee)
func (_Nodemanagement *NodemanagementFilterer) FilterFeeUpdatedMom(opts *bind.FilterOpts) (*NodemanagementFeeUpdatedMomIterator, error) {

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "FeeUpdatedMom")
	if err != nil {
		return nil, err
	}
	return &NodemanagementFeeUpdatedMomIterator{contract: _Nodemanagement.contract, event: "FeeUpdatedMom", logs: logs, sub: sub}, nil
}

// WatchFeeUpdatedMom is a free log subscription operation binding the contract event 0x8173aa58a259e63faf706f8a21a9f36cc8177d015a0a272c554bce99ef47188b.
//
// Solidity: event FeeUpdatedMom(uint256 old_fee, uint256 new_fee)
func (_Nodemanagement *NodemanagementFilterer) WatchFeeUpdatedMom(opts *bind.WatchOpts, sink chan<- *NodemanagementFeeUpdatedMom) (event.Subscription, error) {

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "FeeUpdatedMom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementFeeUpdatedMom)
				if err := _Nodemanagement.contract.UnpackLog(event, "FeeUpdatedMom", log); err != nil {
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

// ParseFeeUpdatedMom is a log parse operation binding the contract event 0x8173aa58a259e63faf706f8a21a9f36cc8177d015a0a272c554bce99ef47188b.
//
// Solidity: event FeeUpdatedMom(uint256 old_fee, uint256 new_fee)
func (_Nodemanagement *NodemanagementFilterer) ParseFeeUpdatedMom(log types.Log) (*NodemanagementFeeUpdatedMom, error) {
	event := new(NodemanagementFeeUpdatedMom)
	if err := _Nodemanagement.contract.UnpackLog(event, "FeeUpdatedMom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Nodemanagement contract.
type NodemanagementInitializedIterator struct {
	Event *NodemanagementInitialized // Event containing the contract specifics and raw log

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
func (it *NodemanagementInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementInitialized)
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
		it.Event = new(NodemanagementInitialized)
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
func (it *NodemanagementInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementInitialized represents a Initialized event raised by the Nodemanagement contract.
type NodemanagementInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodemanagement *NodemanagementFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodemanagementInitializedIterator, error) {

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodemanagementInitializedIterator{contract: _Nodemanagement.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodemanagement *NodemanagementFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodemanagementInitialized) (event.Subscription, error) {

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementInitialized)
				if err := _Nodemanagement.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Nodemanagement *NodemanagementFilterer) ParseInitialized(log types.Log) (*NodemanagementInitialized, error) {
	event := new(NodemanagementInitialized)
	if err := _Nodemanagement.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementNodeMgmtEventIterator is returned from FilterNodeMgmtEvent and is used to iterate over the raw logs and unpacked data for NodeMgmtEvent events raised by the Nodemanagement contract.
type NodemanagementNodeMgmtEventIterator struct {
	Event *NodemanagementNodeMgmtEvent // Event containing the contract specifics and raw log

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
func (it *NodemanagementNodeMgmtEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementNodeMgmtEvent)
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
		it.Event = new(NodemanagementNodeMgmtEvent)
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
func (it *NodemanagementNodeMgmtEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementNodeMgmtEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementNodeMgmtEvent represents a NodeMgmtEvent event raised by the Nodemanagement contract.
type NodemanagementNodeMgmtEvent struct {
	FromNodeId *big.Int
	ToNodeId   *big.Int
	OdysseyId  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeMgmtEvent is a free log retrieval operation binding the contract event 0xe56a80af793b5747dedcba9783e630d0affdfbda849f9f8e55f3e9888c1d67f6.
//
// Solidity: event NodeMgmtEvent(uint256 indexed from_node_id, uint256 indexed to_node_id, uint256 indexed odyssey_id)
func (_Nodemanagement *NodemanagementFilterer) FilterNodeMgmtEvent(opts *bind.FilterOpts, from_node_id []*big.Int, to_node_id []*big.Int, odyssey_id []*big.Int) (*NodemanagementNodeMgmtEventIterator, error) {

	var from_node_idRule []interface{}
	for _, from_node_idItem := range from_node_id {
		from_node_idRule = append(from_node_idRule, from_node_idItem)
	}
	var to_node_idRule []interface{}
	for _, to_node_idItem := range to_node_id {
		to_node_idRule = append(to_node_idRule, to_node_idItem)
	}
	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "NodeMgmtEvent", from_node_idRule, to_node_idRule, odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return &NodemanagementNodeMgmtEventIterator{contract: _Nodemanagement.contract, event: "NodeMgmtEvent", logs: logs, sub: sub}, nil
}

// WatchNodeMgmtEvent is a free log subscription operation binding the contract event 0xe56a80af793b5747dedcba9783e630d0affdfbda849f9f8e55f3e9888c1d67f6.
//
// Solidity: event NodeMgmtEvent(uint256 indexed from_node_id, uint256 indexed to_node_id, uint256 indexed odyssey_id)
func (_Nodemanagement *NodemanagementFilterer) WatchNodeMgmtEvent(opts *bind.WatchOpts, sink chan<- *NodemanagementNodeMgmtEvent, from_node_id []*big.Int, to_node_id []*big.Int, odyssey_id []*big.Int) (event.Subscription, error) {

	var from_node_idRule []interface{}
	for _, from_node_idItem := range from_node_id {
		from_node_idRule = append(from_node_idRule, from_node_idItem)
	}
	var to_node_idRule []interface{}
	for _, to_node_idItem := range to_node_id {
		to_node_idRule = append(to_node_idRule, to_node_idItem)
	}
	var odyssey_idRule []interface{}
	for _, odyssey_idItem := range odyssey_id {
		odyssey_idRule = append(odyssey_idRule, odyssey_idItem)
	}

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "NodeMgmtEvent", from_node_idRule, to_node_idRule, odyssey_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementNodeMgmtEvent)
				if err := _Nodemanagement.contract.UnpackLog(event, "NodeMgmtEvent", log); err != nil {
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

// ParseNodeMgmtEvent is a log parse operation binding the contract event 0xe56a80af793b5747dedcba9783e630d0affdfbda849f9f8e55f3e9888c1d67f6.
//
// Solidity: event NodeMgmtEvent(uint256 indexed from_node_id, uint256 indexed to_node_id, uint256 indexed odyssey_id)
func (_Nodemanagement *NodemanagementFilterer) ParseNodeMgmtEvent(log types.Log) (*NodemanagementNodeMgmtEvent, error) {
	event := new(NodemanagementNodeMgmtEvent)
	if err := _Nodemanagement.contract.UnpackLog(event, "NodeMgmtEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the Nodemanagement contract.
type NodemanagementNodeRemovedIterator struct {
	Event *NodemanagementNodeRemoved // Event containing the contract specifics and raw log

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
func (it *NodemanagementNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementNodeRemoved)
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
		it.Event = new(NodemanagementNodeRemoved)
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
func (it *NodemanagementNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementNodeRemoved represents a NodeRemoved event raised by the Nodemanagement contract.
type NodemanagementNodeRemoved struct {
	NodeId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0x2b61f53058df44855c08a7304979066cb409f86a3dca910e2275b9e85635d1f9.
//
// Solidity: event NodeRemoved(uint256 indexed node_id)
func (_Nodemanagement *NodemanagementFilterer) FilterNodeRemoved(opts *bind.FilterOpts, node_id []*big.Int) (*NodemanagementNodeRemovedIterator, error) {

	var node_idRule []interface{}
	for _, node_idItem := range node_id {
		node_idRule = append(node_idRule, node_idItem)
	}

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "NodeRemoved", node_idRule)
	if err != nil {
		return nil, err
	}
	return &NodemanagementNodeRemovedIterator{contract: _Nodemanagement.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0x2b61f53058df44855c08a7304979066cb409f86a3dca910e2275b9e85635d1f9.
//
// Solidity: event NodeRemoved(uint256 indexed node_id)
func (_Nodemanagement *NodemanagementFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *NodemanagementNodeRemoved, node_id []*big.Int) (event.Subscription, error) {

	var node_idRule []interface{}
	for _, node_idItem := range node_id {
		node_idRule = append(node_idRule, node_idItem)
	}

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "NodeRemoved", node_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementNodeRemoved)
				if err := _Nodemanagement.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// ParseNodeRemoved is a log parse operation binding the contract event 0x2b61f53058df44855c08a7304979066cb409f86a3dca910e2275b9e85635d1f9.
//
// Solidity: event NodeRemoved(uint256 indexed node_id)
func (_Nodemanagement *NodemanagementFilterer) ParseNodeRemoved(log types.Log) (*NodemanagementNodeRemoved, error) {
	event := new(NodemanagementNodeRemoved)
	if err := _Nodemanagement.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementNodeUpdatedIterator is returned from FilterNodeUpdated and is used to iterate over the raw logs and unpacked data for NodeUpdated events raised by the Nodemanagement contract.
type NodemanagementNodeUpdatedIterator struct {
	Event *NodemanagementNodeUpdated // Event containing the contract specifics and raw log

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
func (it *NodemanagementNodeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementNodeUpdated)
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
		it.Event = new(NodemanagementNodeUpdated)
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
func (it *NodemanagementNodeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementNodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementNodeUpdated represents a NodeUpdated event raised by the Nodemanagement contract.
type NodemanagementNodeUpdated struct {
	NodeId      *big.Int
	OldOwner    common.Address
	NewOwner    common.Address
	OldHostname string
	NewHostname string
	OldName     string
	NewName     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeUpdated is a free log retrieval operation binding the contract event 0x7dc29048401c15b1b76c7fb99c378fb0bc6cd5f5248191e6b4dc5027cf65692a.
//
// Solidity: event NodeUpdated(uint256 indexed node_id, address old_owner, address new_owner, string old_hostname, string new_hostname, string old_name, string new_name)
func (_Nodemanagement *NodemanagementFilterer) FilterNodeUpdated(opts *bind.FilterOpts, node_id []*big.Int) (*NodemanagementNodeUpdatedIterator, error) {

	var node_idRule []interface{}
	for _, node_idItem := range node_id {
		node_idRule = append(node_idRule, node_idItem)
	}

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "NodeUpdated", node_idRule)
	if err != nil {
		return nil, err
	}
	return &NodemanagementNodeUpdatedIterator{contract: _Nodemanagement.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeUpdated is a free log subscription operation binding the contract event 0x7dc29048401c15b1b76c7fb99c378fb0bc6cd5f5248191e6b4dc5027cf65692a.
//
// Solidity: event NodeUpdated(uint256 indexed node_id, address old_owner, address new_owner, string old_hostname, string new_hostname, string old_name, string new_name)
func (_Nodemanagement *NodemanagementFilterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *NodemanagementNodeUpdated, node_id []*big.Int) (event.Subscription, error) {

	var node_idRule []interface{}
	for _, node_idItem := range node_id {
		node_idRule = append(node_idRule, node_idItem)
	}

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "NodeUpdated", node_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementNodeUpdated)
				if err := _Nodemanagement.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
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

// ParseNodeUpdated is a log parse operation binding the contract event 0x7dc29048401c15b1b76c7fb99c378fb0bc6cd5f5248191e6b4dc5027cf65692a.
//
// Solidity: event NodeUpdated(uint256 indexed node_id, address old_owner, address new_owner, string old_hostname, string new_hostname, string old_name, string new_name)
func (_Nodemanagement *NodemanagementFilterer) ParseNodeUpdated(log types.Log) (*NodemanagementNodeUpdated, error) {
	event := new(NodemanagementNodeUpdated)
	if err := _Nodemanagement.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nodemanagement contract.
type NodemanagementOwnershipTransferredIterator struct {
	Event *NodemanagementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodemanagementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementOwnershipTransferred)
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
		it.Event = new(NodemanagementOwnershipTransferred)
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
func (it *NodemanagementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementOwnershipTransferred represents a OwnershipTransferred event raised by the Nodemanagement contract.
type NodemanagementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodemanagement *NodemanagementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodemanagementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodemanagementOwnershipTransferredIterator{contract: _Nodemanagement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodemanagement *NodemanagementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodemanagementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementOwnershipTransferred)
				if err := _Nodemanagement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nodemanagement *NodemanagementFilterer) ParseOwnershipTransferred(log types.Log) (*NodemanagementOwnershipTransferred, error) {
	event := new(NodemanagementOwnershipTransferred)
	if err := _Nodemanagement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodemanagementUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Nodemanagement contract.
type NodemanagementUpgradedIterator struct {
	Event *NodemanagementUpgraded // Event containing the contract specifics and raw log

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
func (it *NodemanagementUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodemanagementUpgraded)
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
		it.Event = new(NodemanagementUpgraded)
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
func (it *NodemanagementUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodemanagementUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodemanagementUpgraded represents a Upgraded event raised by the Nodemanagement contract.
type NodemanagementUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Nodemanagement *NodemanagementFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodemanagementUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Nodemanagement.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodemanagementUpgradedIterator{contract: _Nodemanagement.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Nodemanagement *NodemanagementFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodemanagementUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Nodemanagement.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodemanagementUpgraded)
				if err := _Nodemanagement.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Nodemanagement *NodemanagementFilterer) ParseUpgraded(log types.Log) (*NodemanagementUpgraded, error) {
	event := new(NodemanagementUpgraded)
	if err := _Nodemanagement.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
