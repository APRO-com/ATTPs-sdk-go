// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agent_proxy

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

// CommonAgentHeader is an auto generated low-level Go binding around an user-defined struct.
type CommonAgentHeader struct {
	Version         string
	MessageId       string
	SourceAgentId   string
	SourceAgentName string
	TargetAgentId   string
	Timestamp       *big.Int
	MessageType     uint8
	Priority        uint8
	Ttl             *big.Int
}

// CommonAgentSettings is an auto generated low-level Go binding around an user-defined struct.
type CommonAgentSettings struct {
	Signers          []common.Address
	Threshold        uint8
	ConverterAddress common.Address
	AgentHeader      CommonAgentHeader
}

// CommonMessagePayload is an auto generated low-level Go binding around an user-defined struct.
type CommonMessagePayload struct {
	Data     []byte
	DataHash [32]byte
	Proofs   CommonProofs
	Metadata CommonMetadata
}

// CommonMetadata is an auto generated low-level Go binding around an user-defined struct.
type CommonMetadata struct {
	ContentType string
	Encoding    string
	Compression string
}

// CommonProofs is an auto generated low-level Go binding around an user-defined struct.
type CommonProofs struct {
	ZkProof        []byte
	MerkleProof    []byte
	SignatureProof []byte
}

// AgentProxyMetaData contains all meta data concerning the AgentProxy contract.
var AgentProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAgentFactoryOrManager\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFactory\",\"type\":\"address\"}],\"name\":\"AgentFactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"AgentManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"createAndRegisterAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"setAgentFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setAgentManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingsDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"zkProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatureProof\",\"type\":\"bytes\"}],\"internalType\":\"structCommon.Proofs\",\"name\":\"proofs\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"contentType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encoding\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"compression\",\"type\":\"string\"}],\"internalType\":\"structCommon.Metadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AgentProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentProxyMetaData.ABI instead.
var AgentProxyABI = AgentProxyMetaData.ABI

// AgentProxy is an auto generated Go binding around an Ethereum contract.
type AgentProxy struct {
	AgentProxyCaller     // Read-only binding to the contract
	AgentProxyTransactor // Write-only binding to the contract
	AgentProxyFilterer   // Log filterer for contract events
}

// AgentProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentProxySession struct {
	Contract     *AgentProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentProxyCallerSession struct {
	Contract *AgentProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AgentProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentProxyTransactorSession struct {
	Contract     *AgentProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AgentProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentProxyRaw struct {
	Contract *AgentProxy // Generic contract binding to access the raw methods on
}

// AgentProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentProxyCallerRaw struct {
	Contract *AgentProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AgentProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentProxyTransactorRaw struct {
	Contract *AgentProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentProxy creates a new instance of AgentProxy, bound to a specific deployed contract.
func NewAgentProxy(address common.Address, backend bind.ContractBackend) (*AgentProxy, error) {
	contract, err := bindAgentProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentProxy{AgentProxyCaller: AgentProxyCaller{contract: contract}, AgentProxyTransactor: AgentProxyTransactor{contract: contract}, AgentProxyFilterer: AgentProxyFilterer{contract: contract}}, nil
}

// NewAgentProxyCaller creates a new read-only instance of AgentProxy, bound to a specific deployed contract.
func NewAgentProxyCaller(address common.Address, caller bind.ContractCaller) (*AgentProxyCaller, error) {
	contract, err := bindAgentProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentProxyCaller{contract: contract}, nil
}

// NewAgentProxyTransactor creates a new write-only instance of AgentProxy, bound to a specific deployed contract.
func NewAgentProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentProxyTransactor, error) {
	contract, err := bindAgentProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentProxyTransactor{contract: contract}, nil
}

// NewAgentProxyFilterer creates a new log filterer instance of AgentProxy, bound to a specific deployed contract.
func NewAgentProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentProxyFilterer, error) {
	contract, err := bindAgentProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentProxyFilterer{contract: contract}, nil
}

// bindAgentProxy binds a generic wrapper to an already deployed contract.
func bindAgentProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (util) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentProxy *AgentProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentProxy.Contract.AgentProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentProxy *AgentProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentProxy.Contract.AgentProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentProxy *AgentProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentProxy.Contract.AgentProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (util) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentProxy *AgentProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentProxy *AgentProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentProxy *AgentProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentProxy.Contract.contract.Transact(opts, method, params...)
}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_AgentProxy *AgentProxyCaller) AgentFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "agentFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_AgentProxy *AgentProxySession) AgentFactory() (common.Address, error) {
	return _AgentProxy.Contract.AgentFactory(&_AgentProxy.CallOpts)
}

// AgentFactory is a free data retrieval call binding the contract method 0x7df107ea.
//
// Solidity: function agentFactory() view returns(address)
func (_AgentProxy *AgentProxyCallerSession) AgentFactory() (common.Address, error) {
	return _AgentProxy.Contract.AgentFactory(&_AgentProxy.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_AgentProxy *AgentProxyCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_AgentProxy *AgentProxySession) AgentManager() (common.Address, error) {
	return _AgentProxy.Contract.AgentManager(&_AgentProxy.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_AgentProxy *AgentProxyCallerSession) AgentManager() (common.Address, error) {
	return _AgentProxy.Contract.AgentManager(&_AgentProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentProxy *AgentProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentProxy *AgentProxySession) Owner() (common.Address, error) {
	return _AgentProxy.Contract.Owner(&_AgentProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentProxy *AgentProxyCallerSession) Owner() (common.Address, error) {
	return _AgentProxy.Contract.Owner(&_AgentProxy.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_AgentProxy *AgentProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_AgentProxy *AgentProxySession) TypeAndVersion() (string, error) {
	return _AgentProxy.Contract.TypeAndVersion(&_AgentProxy.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_AgentProxy *AgentProxyCallerSession) TypeAndVersion() (string, error) {
	return _AgentProxy.Contract.TypeAndVersion(&_AgentProxy.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentProxy *AgentProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentProxy *AgentProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentProxy.Contract.AcceptOwnership(&_AgentProxy.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentProxy *AgentProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentProxy.Contract.AcceptOwnership(&_AgentProxy.TransactOpts)
}

// CreateAndRegisterAgent is a paid mutator transaction binding the contract method 0x4a5a04e9.
//
// Solidity: function createAndRegisterAgent((address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxyTransactor) CreateAndRegisterAgent(opts *bind.TransactOpts, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "createAndRegisterAgent", agentSettings)
}

// CreateAndRegisterAgent is a paid mutator transaction binding the contract method 0x4a5a04e9.
//
// Solidity: function createAndRegisterAgent((address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxySession) CreateAndRegisterAgent(agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.Contract.CreateAndRegisterAgent(&_AgentProxy.TransactOpts, agentSettings)
}

// CreateAndRegisterAgent is a paid mutator transaction binding the contract method 0x4a5a04e9.
//
// Solidity: function createAndRegisterAgent((address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxyTransactorSession) CreateAndRegisterAgent(agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.Contract.CreateAndRegisterAgent(&_AgentProxy.TransactOpts, agentSettings)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address factory) returns()
func (_AgentProxy *AgentProxyTransactor) SetAgentFactory(opts *bind.TransactOpts, factory common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "setAgentFactory", factory)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address factory) returns()
func (_AgentProxy *AgentProxySession) SetAgentFactory(factory common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.SetAgentFactory(&_AgentProxy.TransactOpts, factory)
}

// SetAgentFactory is a paid mutator transaction binding the contract method 0x7caf9073.
//
// Solidity: function setAgentFactory(address factory) returns()
func (_AgentProxy *AgentProxyTransactorSession) SetAgentFactory(factory common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.SetAgentFactory(&_AgentProxy.TransactOpts, factory)
}

// SetAgentManager is a paid mutator transaction binding the contract method 0xd35e9420.
//
// Solidity: function setAgentManager(address manager) returns()
func (_AgentProxy *AgentProxyTransactor) SetAgentManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "setAgentManager", manager)
}

// SetAgentManager is a paid mutator transaction binding the contract method 0xd35e9420.
//
// Solidity: function setAgentManager(address manager) returns()
func (_AgentProxy *AgentProxySession) SetAgentManager(manager common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.SetAgentManager(&_AgentProxy.TransactOpts, manager)
}

// SetAgentManager is a paid mutator transaction binding the contract method 0xd35e9420.
//
// Solidity: function setAgentManager(address manager) returns()
func (_AgentProxy *AgentProxyTransactorSession) SetAgentManager(manager common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.SetAgentManager(&_AgentProxy.TransactOpts, manager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AgentProxy *AgentProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AgentProxy *AgentProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.TransferOwnership(&_AgentProxy.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AgentProxy *AgentProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.TransferOwnership(&_AgentProxy.TransactOpts, to)
}

// Verify is a paid mutator transaction binding the contract method 0x59dabe6a.
//
// Solidity: function verify(address agent, bytes32 settingsDigest, (bytes,bytes32,(bytes,bytes,bytes),(string,string,string)) payload) returns()
func (_AgentProxy *AgentProxyTransactor) Verify(opts *bind.TransactOpts, agent common.Address, settingsDigest [32]byte, payload CommonMessagePayload) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "verify", agent, settingsDigest, payload)
}

// Verify is a paid mutator transaction binding the contract method 0x59dabe6a.
//
// Solidity: function verify(address agent, bytes32 settingsDigest, (bytes,bytes32,(bytes,bytes,bytes),(string,string,string)) payload) returns()
func (_AgentProxy *AgentProxySession) Verify(agent common.Address, settingsDigest [32]byte, payload CommonMessagePayload) (*types.Transaction, error) {
	return _AgentProxy.Contract.Verify(&_AgentProxy.TransactOpts, agent, settingsDigest, payload)
}

// Verify is a paid mutator transaction binding the contract method 0x59dabe6a.
//
// Solidity: function verify(address agent, bytes32 settingsDigest, (bytes,bytes32,(bytes,bytes,bytes),(string,string,string)) payload) returns()
func (_AgentProxy *AgentProxyTransactorSession) Verify(agent common.Address, settingsDigest [32]byte, payload CommonMessagePayload) (*types.Transaction, error) {
	return _AgentProxy.Contract.Verify(&_AgentProxy.TransactOpts, agent, settingsDigest, payload)
}

// AgentProxyAgentFactorySetIterator is returned from FilterAgentFactorySet and is used to iterate over the raw logs and unpacked data for AgentFactorySet events raised by the AgentProxy contract.
type AgentProxyAgentFactorySetIterator struct {
	Event *AgentProxyAgentFactorySet // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentFactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentFactorySet)
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
		it.Event = new(AgentProxyAgentFactorySet)
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
func (it *AgentProxyAgentFactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentFactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentFactorySet represents a AgentFactorySet event raised by the AgentProxy contract.
type AgentProxyAgentFactorySet struct {
	OldFactory common.Address
	NewFactory common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentFactorySet is a free log retrieval operation binding the contract event 0xc6de8d526ba74f96d5911c55f3292f7197efb360bf2d40f3d3886d1ff7d2d8ed.
//
// Solidity: event AgentFactorySet(address oldFactory, address newFactory)
func (_AgentProxy *AgentProxyFilterer) FilterAgentFactorySet(opts *bind.FilterOpts) (*AgentProxyAgentFactorySetIterator, error) {

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentFactorySet")
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentFactorySetIterator{contract: _AgentProxy.contract, event: "AgentFactorySet", logs: logs, sub: sub}, nil
}

// WatchAgentFactorySet is a free log subscription operation binding the contract event 0xc6de8d526ba74f96d5911c55f3292f7197efb360bf2d40f3d3886d1ff7d2d8ed.
//
// Solidity: event AgentFactorySet(address oldFactory, address newFactory)
func (_AgentProxy *AgentProxyFilterer) WatchAgentFactorySet(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentFactorySet) (event.Subscription, error) {

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentFactorySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentFactorySet)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentFactorySet", log); err != nil {
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

// ParseAgentFactorySet is a log parse operation binding the contract event 0xc6de8d526ba74f96d5911c55f3292f7197efb360bf2d40f3d3886d1ff7d2d8ed.
//
// Solidity: event AgentFactorySet(address oldFactory, address newFactory)
func (_AgentProxy *AgentProxyFilterer) ParseAgentFactorySet(log types.Log) (*AgentProxyAgentFactorySet, error) {
	event := new(AgentProxyAgentFactorySet)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentFactorySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyAgentManagerSetIterator is returned from FilterAgentManagerSet and is used to iterate over the raw logs and unpacked data for AgentManagerSet events raised by the AgentProxy contract.
type AgentProxyAgentManagerSetIterator struct {
	Event *AgentProxyAgentManagerSet // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentManagerSet)
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
		it.Event = new(AgentProxyAgentManagerSet)
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
func (it *AgentProxyAgentManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentManagerSet represents a AgentManagerSet event raised by the AgentProxy contract.
type AgentProxyAgentManagerSet struct {
	OldManager common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentManagerSet is a free log retrieval operation binding the contract event 0x69e3e65eac755283674d54b2543ca7562ae3eeb2a051d4f22f7de3d486053db9.
//
// Solidity: event AgentManagerSet(address oldManager, address newManager)
func (_AgentProxy *AgentProxyFilterer) FilterAgentManagerSet(opts *bind.FilterOpts) (*AgentProxyAgentManagerSetIterator, error) {

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentManagerSet")
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentManagerSetIterator{contract: _AgentProxy.contract, event: "AgentManagerSet", logs: logs, sub: sub}, nil
}

// WatchAgentManagerSet is a free log subscription operation binding the contract event 0x69e3e65eac755283674d54b2543ca7562ae3eeb2a051d4f22f7de3d486053db9.
//
// Solidity: event AgentManagerSet(address oldManager, address newManager)
func (_AgentProxy *AgentProxyFilterer) WatchAgentManagerSet(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentManagerSet) (event.Subscription, error) {

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentManagerSet)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentManagerSet", log); err != nil {
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

// ParseAgentManagerSet is a log parse operation binding the contract event 0x69e3e65eac755283674d54b2543ca7562ae3eeb2a051d4f22f7de3d486053db9.
//
// Solidity: event AgentManagerSet(address oldManager, address newManager)
func (_AgentProxy *AgentProxyFilterer) ParseAgentManagerSet(log types.Log) (*AgentProxyAgentManagerSet, error) {
	event := new(AgentProxyAgentManagerSet)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the AgentProxy contract.
type AgentProxyOwnershipTransferRequestedIterator struct {
	Event *AgentProxyOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *AgentProxyOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyOwnershipTransferRequested)
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
		it.Event = new(AgentProxyOwnershipTransferRequested)
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
func (it *AgentProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the AgentProxy contract.
type AgentProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AgentProxy *AgentProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AgentProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyOwnershipTransferRequestedIterator{contract: _AgentProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AgentProxy *AgentProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AgentProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyOwnershipTransferRequested)
				if err := _AgentProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AgentProxy *AgentProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*AgentProxyOwnershipTransferRequested, error) {
	event := new(AgentProxyOwnershipTransferRequested)
	if err := _AgentProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentProxy contract.
type AgentProxyOwnershipTransferredIterator struct {
	Event *AgentProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AgentProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyOwnershipTransferred)
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
		it.Event = new(AgentProxyOwnershipTransferred)
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
func (it *AgentProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyOwnershipTransferred represents a OwnershipTransferred event raised by the AgentProxy contract.
type AgentProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AgentProxy *AgentProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AgentProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyOwnershipTransferredIterator{contract: _AgentProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AgentProxy *AgentProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyOwnershipTransferred)
				if err := _AgentProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AgentProxy *AgentProxyFilterer) ParseOwnershipTransferred(log types.Log) (*AgentProxyOwnershipTransferred, error) {
	event := new(AgentProxyOwnershipTransferred)
	if err := _AgentProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
