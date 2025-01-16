// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agent_manager

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

// CommonAgentConfig is an auto generated low-level Go binding around an user-defined struct.
type CommonAgentConfig struct {
	ConfigDigest      [32]byte
	ConfigBlockNumber uint32
	IsActive bool
	Settings CommonAgentSettings
}

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

// AgentProxyMetaData contains all meta data concerning the AgentProxy contract.
var AgentProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentIsAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentIsRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderAgentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderMessageId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderMessageType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderPriority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentSettingProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowedAgent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCallData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFactoryAgent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegisteredAgent\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"AgentProxySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentSettingsProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"acceptAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"acceptAgentSettingProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"allowedAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingDigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"}],\"name\":\"allowedSinger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"changeAgentSettingProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingDigest\",\"type\":\"bytes32\"}],\"name\":\"getAgentConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"configBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getAgentConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"configBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getAgentConfigsCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"agentConfigIdxStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"agentConfigIdxEnd\",\"type\":\"uint64\"}],\"name\":\"getAgentConfigsInRange\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"configBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentConfig[]\",\"name\":\"agentConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAllowedAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRegisteringAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowedAgentsCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"allowedAgentIdxStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"allowedAgentIdxEnd\",\"type\":\"uint64\"}],\"name\":\"getAllowedAgentsInRange\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedAgents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteringAgentsCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registeringAgentIdxStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"registeringAgentIdxEnd\",\"type\":\"uint64\"}],\"name\":\"getRegisteringAgentsInRange\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"registeringAgents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"registerAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"setAgentProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingDigest\",\"type\":\"bytes32\"}],\"name\":\"signerThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validateDataConversion\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AgentProxy is a free data retrieval call binding the contract method 0x81070369.
//
// Solidity: function agentProxy() view returns(address)
func (_AgentProxy *AgentProxyCaller) AgentProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "agentProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentProxy is a free data retrieval call binding the contract method 0x81070369.
//
// Solidity: function agentProxy() view returns(address)
func (_AgentProxy *AgentProxySession) AgentProxy() (common.Address, error) {
	return _AgentProxy.Contract.AgentProxy(&_AgentProxy.CallOpts)
}

// AgentProxy is a free data retrieval call binding the contract method 0x81070369.
//
// Solidity: function agentProxy() view returns(address)
func (_AgentProxy *AgentProxyCallerSession) AgentProxy() (common.Address, error) {
	return _AgentProxy.Contract.AgentProxy(&_AgentProxy.CallOpts)
}

// AgentVersion is a free data retrieval call binding the contract method 0x10c59c72.
//
// Solidity: function agentVersion() pure returns(string)
func (_AgentProxy *AgentProxyCaller) AgentVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "agentVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AgentVersion is a free data retrieval call binding the contract method 0x10c59c72.
//
// Solidity: function agentVersion() pure returns(string)
func (_AgentProxy *AgentProxySession) AgentVersion() (string, error) {
	return _AgentProxy.Contract.AgentVersion(&_AgentProxy.CallOpts)
}

// AgentVersion is a free data retrieval call binding the contract method 0x10c59c72.
//
// Solidity: function agentVersion() pure returns(string)
func (_AgentProxy *AgentProxyCallerSession) AgentVersion() (string, error) {
	return _AgentProxy.Contract.AgentVersion(&_AgentProxy.CallOpts)
}

// AllowedAgent is a free data retrieval call binding the contract method 0x1c774d11.
//
// Solidity: function allowedAgent(address agent) view returns(bool)
func (_AgentProxy *AgentProxyCaller) AllowedAgent(opts *bind.CallOpts, agent common.Address) (bool, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "allowedAgent", agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAgent is a free data retrieval call binding the contract method 0x1c774d11.
//
// Solidity: function allowedAgent(address agent) view returns(bool)
func (_AgentProxy *AgentProxySession) AllowedAgent(agent common.Address) (bool, error) {
	return _AgentProxy.Contract.AllowedAgent(&_AgentProxy.CallOpts, agent)
}

// AllowedAgent is a free data retrieval call binding the contract method 0x1c774d11.
//
// Solidity: function allowedAgent(address agent) view returns(bool)
func (_AgentProxy *AgentProxyCallerSession) AllowedAgent(agent common.Address) (bool, error) {
	return _AgentProxy.Contract.AllowedAgent(&_AgentProxy.CallOpts, agent)
}

// AllowedSinger is a free data retrieval call binding the contract method 0xfac2ac49.
//
// Solidity: function allowedSinger(address agent, bytes32 settingDigest, address singer) view returns(bool)
func (_AgentProxy *AgentProxyCaller) AllowedSinger(opts *bind.CallOpts, agent common.Address, settingDigest [32]byte, singer common.Address) (bool, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "allowedSinger", agent, settingDigest, singer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedSinger is a free data retrieval call binding the contract method 0xfac2ac49.
//
// Solidity: function allowedSinger(address agent, bytes32 settingDigest, address singer) view returns(bool)
func (_AgentProxy *AgentProxySession) AllowedSinger(agent common.Address, settingDigest [32]byte, singer common.Address) (bool, error) {
	return _AgentProxy.Contract.AllowedSinger(&_AgentProxy.CallOpts, agent, settingDigest, singer)
}

// AllowedSinger is a free data retrieval call binding the contract method 0xfac2ac49.
//
// Solidity: function allowedSinger(address agent, bytes32 settingDigest, address singer) view returns(bool)
func (_AgentProxy *AgentProxyCallerSession) AllowedSinger(agent common.Address, settingDigest [32]byte, singer common.Address) (bool, error) {
	return _AgentProxy.Contract.AllowedSinger(&_AgentProxy.CallOpts, agent, settingDigest, singer)
}

// GetAgentConfig is a free data retrieval call binding the contract method 0xc703f3f9.
//
// Solidity: function getAgentConfig(address agent, bytes32 settingDigest) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256))))
func (_AgentProxy *AgentProxyCaller) GetAgentConfig(opts *bind.CallOpts, agent common.Address, settingDigest [32]byte) (CommonAgentConfig, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAgentConfig", agent, settingDigest)

	if err != nil {
		return *new(CommonAgentConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommonAgentConfig)).(*CommonAgentConfig)

	return out0, err

}

// GetAgentConfig is a free data retrieval call binding the contract method 0xc703f3f9.
//
// Solidity: function getAgentConfig(address agent, bytes32 settingDigest) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256))))
func (_AgentProxy *AgentProxySession) GetAgentConfig(agent common.Address, settingDigest [32]byte) (CommonAgentConfig, error) {
	return _AgentProxy.Contract.GetAgentConfig(&_AgentProxy.CallOpts, agent, settingDigest)
}

// GetAgentConfig is a free data retrieval call binding the contract method 0xc703f3f9.
//
// Solidity: function getAgentConfig(address agent, bytes32 settingDigest) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256))))
func (_AgentProxy *AgentProxyCallerSession) GetAgentConfig(agent common.Address, settingDigest [32]byte) (CommonAgentConfig, error) {
	return _AgentProxy.Contract.GetAgentConfig(&_AgentProxy.CallOpts, agent, settingDigest)
}

// GetAgentConfigs is a free data retrieval call binding the contract method 0x6ccfcf63.
//
// Solidity: function getAgentConfigs(address agent) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[])
func (_AgentProxy *AgentProxyCaller) GetAgentConfigs(opts *bind.CallOpts, agent common.Address) ([]CommonAgentConfig, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAgentConfigs", agent)

	if err != nil {
		return *new([]CommonAgentConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]CommonAgentConfig)).(*[]CommonAgentConfig)

	return out0, err

}

// GetAgentConfigs is a free data retrieval call binding the contract method 0x6ccfcf63.
//
// Solidity: function getAgentConfigs(address agent) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[])
func (_AgentProxy *AgentProxySession) GetAgentConfigs(agent common.Address) ([]CommonAgentConfig, error) {
	return _AgentProxy.Contract.GetAgentConfigs(&_AgentProxy.CallOpts, agent)
}

// GetAgentConfigs is a free data retrieval call binding the contract method 0x6ccfcf63.
//
// Solidity: function getAgentConfigs(address agent) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[])
func (_AgentProxy *AgentProxyCallerSession) GetAgentConfigs(agent common.Address) ([]CommonAgentConfig, error) {
	return _AgentProxy.Contract.GetAgentConfigs(&_AgentProxy.CallOpts, agent)
}

// GetAgentConfigsCount is a free data retrieval call binding the contract method 0xd2c0823a.
//
// Solidity: function getAgentConfigsCount(address agent) view returns(uint64)
func (_AgentProxy *AgentProxyCaller) GetAgentConfigsCount(opts *bind.CallOpts, agent common.Address) (uint64, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAgentConfigsCount", agent)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAgentConfigsCount is a free data retrieval call binding the contract method 0xd2c0823a.
//
// Solidity: function getAgentConfigsCount(address agent) view returns(uint64)
func (_AgentProxy *AgentProxySession) GetAgentConfigsCount(agent common.Address) (uint64, error) {
	return _AgentProxy.Contract.GetAgentConfigsCount(&_AgentProxy.CallOpts, agent)
}

// GetAgentConfigsCount is a free data retrieval call binding the contract method 0xd2c0823a.
//
// Solidity: function getAgentConfigsCount(address agent) view returns(uint64)
func (_AgentProxy *AgentProxyCallerSession) GetAgentConfigsCount(agent common.Address) (uint64, error) {
	return _AgentProxy.Contract.GetAgentConfigsCount(&_AgentProxy.CallOpts, agent)
}

// GetAgentConfigsInRange is a free data retrieval call binding the contract method 0x1b429f1d.
//
// Solidity: function getAgentConfigsInRange(address agent, uint64 agentConfigIdxStart, uint64 agentConfigIdxEnd) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[] agentConfigs)
func (_AgentProxy *AgentProxyCaller) GetAgentConfigsInRange(opts *bind.CallOpts, agent common.Address, agentConfigIdxStart uint64, agentConfigIdxEnd uint64) ([]CommonAgentConfig, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAgentConfigsInRange", agent, agentConfigIdxStart, agentConfigIdxEnd)

	if err != nil {
		return *new([]CommonAgentConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]CommonAgentConfig)).(*[]CommonAgentConfig)

	return out0, err

}

// GetAgentConfigsInRange is a free data retrieval call binding the contract method 0x1b429f1d.
//
// Solidity: function getAgentConfigsInRange(address agent, uint64 agentConfigIdxStart, uint64 agentConfigIdxEnd) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[] agentConfigs)
func (_AgentProxy *AgentProxySession) GetAgentConfigsInRange(agent common.Address, agentConfigIdxStart uint64, agentConfigIdxEnd uint64) ([]CommonAgentConfig, error) {
	return _AgentProxy.Contract.GetAgentConfigsInRange(&_AgentProxy.CallOpts, agent, agentConfigIdxStart, agentConfigIdxEnd)
}

// GetAgentConfigsInRange is a free data retrieval call binding the contract method 0x1b429f1d.
//
// Solidity: function getAgentConfigsInRange(address agent, uint64 agentConfigIdxStart, uint64 agentConfigIdxEnd) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[] agentConfigs)
func (_AgentProxy *AgentProxyCallerSession) GetAgentConfigsInRange(agent common.Address, agentConfigIdxStart uint64, agentConfigIdxEnd uint64) ([]CommonAgentConfig, error) {
	return _AgentProxy.Contract.GetAgentConfigsInRange(&_AgentProxy.CallOpts, agent, agentConfigIdxStart, agentConfigIdxEnd)
}

// GetAllAllowedAgents is a free data retrieval call binding the contract method 0xac3e7680.
//
// Solidity: function getAllAllowedAgents() view returns(address[])
func (_AgentProxy *AgentProxyCaller) GetAllAllowedAgents(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAllAllowedAgents")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAllowedAgents is a free data retrieval call binding the contract method 0xac3e7680.
//
// Solidity: function getAllAllowedAgents() view returns(address[])
func (_AgentProxy *AgentProxySession) GetAllAllowedAgents() ([]common.Address, error) {
	return _AgentProxy.Contract.GetAllAllowedAgents(&_AgentProxy.CallOpts)
}

// GetAllAllowedAgents is a free data retrieval call binding the contract method 0xac3e7680.
//
// Solidity: function getAllAllowedAgents() view returns(address[])
func (_AgentProxy *AgentProxyCallerSession) GetAllAllowedAgents() ([]common.Address, error) {
	return _AgentProxy.Contract.GetAllAllowedAgents(&_AgentProxy.CallOpts)
}

// GetAllRegisteringAgents is a free data retrieval call binding the contract method 0xa73e07b4.
//
// Solidity: function getAllRegisteringAgents() view returns(address[])
func (_AgentProxy *AgentProxyCaller) GetAllRegisteringAgents(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAllRegisteringAgents")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllRegisteringAgents is a free data retrieval call binding the contract method 0xa73e07b4.
//
// Solidity: function getAllRegisteringAgents() view returns(address[])
func (_AgentProxy *AgentProxySession) GetAllRegisteringAgents() ([]common.Address, error) {
	return _AgentProxy.Contract.GetAllRegisteringAgents(&_AgentProxy.CallOpts)
}

// GetAllRegisteringAgents is a free data retrieval call binding the contract method 0xa73e07b4.
//
// Solidity: function getAllRegisteringAgents() view returns(address[])
func (_AgentProxy *AgentProxyCallerSession) GetAllRegisteringAgents() ([]common.Address, error) {
	return _AgentProxy.Contract.GetAllRegisteringAgents(&_AgentProxy.CallOpts)
}

// GetAllowedAgentsCount is a free data retrieval call binding the contract method 0xa251f4fb.
//
// Solidity: function getAllowedAgentsCount() view returns(uint64)
func (_AgentProxy *AgentProxyCaller) GetAllowedAgentsCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAllowedAgentsCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllowedAgentsCount is a free data retrieval call binding the contract method 0xa251f4fb.
//
// Solidity: function getAllowedAgentsCount() view returns(uint64)
func (_AgentProxy *AgentProxySession) GetAllowedAgentsCount() (uint64, error) {
	return _AgentProxy.Contract.GetAllowedAgentsCount(&_AgentProxy.CallOpts)
}

// GetAllowedAgentsCount is a free data retrieval call binding the contract method 0xa251f4fb.
//
// Solidity: function getAllowedAgentsCount() view returns(uint64)
func (_AgentProxy *AgentProxyCallerSession) GetAllowedAgentsCount() (uint64, error) {
	return _AgentProxy.Contract.GetAllowedAgentsCount(&_AgentProxy.CallOpts)
}

// GetAllowedAgentsInRange is a free data retrieval call binding the contract method 0x8eefb433.
//
// Solidity: function getAllowedAgentsInRange(uint64 allowedAgentIdxStart, uint64 allowedAgentIdxEnd) view returns(address[] allowedAgents)
func (_AgentProxy *AgentProxyCaller) GetAllowedAgentsInRange(opts *bind.CallOpts, allowedAgentIdxStart uint64, allowedAgentIdxEnd uint64) ([]common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getAllowedAgentsInRange", allowedAgentIdxStart, allowedAgentIdxEnd)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedAgentsInRange is a free data retrieval call binding the contract method 0x8eefb433.
//
// Solidity: function getAllowedAgentsInRange(uint64 allowedAgentIdxStart, uint64 allowedAgentIdxEnd) view returns(address[] allowedAgents)
func (_AgentProxy *AgentProxySession) GetAllowedAgentsInRange(allowedAgentIdxStart uint64, allowedAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentProxy.Contract.GetAllowedAgentsInRange(&_AgentProxy.CallOpts, allowedAgentIdxStart, allowedAgentIdxEnd)
}

// GetAllowedAgentsInRange is a free data retrieval call binding the contract method 0x8eefb433.
//
// Solidity: function getAllowedAgentsInRange(uint64 allowedAgentIdxStart, uint64 allowedAgentIdxEnd) view returns(address[] allowedAgents)
func (_AgentProxy *AgentProxyCallerSession) GetAllowedAgentsInRange(allowedAgentIdxStart uint64, allowedAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentProxy.Contract.GetAllowedAgentsInRange(&_AgentProxy.CallOpts, allowedAgentIdxStart, allowedAgentIdxEnd)
}

// GetRegisteringAgentsCount is a free data retrieval call binding the contract method 0xb2833451.
//
// Solidity: function getRegisteringAgentsCount() view returns(uint64)
func (_AgentProxy *AgentProxyCaller) GetRegisteringAgentsCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getRegisteringAgentsCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRegisteringAgentsCount is a free data retrieval call binding the contract method 0xb2833451.
//
// Solidity: function getRegisteringAgentsCount() view returns(uint64)
func (_AgentProxy *AgentProxySession) GetRegisteringAgentsCount() (uint64, error) {
	return _AgentProxy.Contract.GetRegisteringAgentsCount(&_AgentProxy.CallOpts)
}

// GetRegisteringAgentsCount is a free data retrieval call binding the contract method 0xb2833451.
//
// Solidity: function getRegisteringAgentsCount() view returns(uint64)
func (_AgentProxy *AgentProxyCallerSession) GetRegisteringAgentsCount() (uint64, error) {
	return _AgentProxy.Contract.GetRegisteringAgentsCount(&_AgentProxy.CallOpts)
}

// GetRegisteringAgentsInRange is a free data retrieval call binding the contract method 0xdb971453.
//
// Solidity: function getRegisteringAgentsInRange(uint64 registeringAgentIdxStart, uint64 registeringAgentIdxEnd) view returns(address[] registeringAgents)
func (_AgentProxy *AgentProxyCaller) GetRegisteringAgentsInRange(opts *bind.CallOpts, registeringAgentIdxStart uint64, registeringAgentIdxEnd uint64) ([]common.Address, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "getRegisteringAgentsInRange", registeringAgentIdxStart, registeringAgentIdxEnd)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteringAgentsInRange is a free data retrieval call binding the contract method 0xdb971453.
//
// Solidity: function getRegisteringAgentsInRange(uint64 registeringAgentIdxStart, uint64 registeringAgentIdxEnd) view returns(address[] registeringAgents)
func (_AgentProxy *AgentProxySession) GetRegisteringAgentsInRange(registeringAgentIdxStart uint64, registeringAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentProxy.Contract.GetRegisteringAgentsInRange(&_AgentProxy.CallOpts, registeringAgentIdxStart, registeringAgentIdxEnd)
}

// GetRegisteringAgentsInRange is a free data retrieval call binding the contract method 0xdb971453.
//
// Solidity: function getRegisteringAgentsInRange(uint64 registeringAgentIdxStart, uint64 registeringAgentIdxEnd) view returns(address[] registeringAgents)
func (_AgentProxy *AgentProxyCallerSession) GetRegisteringAgentsInRange(registeringAgentIdxStart uint64, registeringAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentProxy.Contract.GetRegisteringAgentsInRange(&_AgentProxy.CallOpts, registeringAgentIdxStart, registeringAgentIdxEnd)
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

// SignerThreshold is a free data retrieval call binding the contract method 0x55200a8d.
//
// Solidity: function signerThreshold(address agent, bytes32 settingDigest) view returns(uint8)
func (_AgentProxy *AgentProxyCaller) SignerThreshold(opts *bind.CallOpts, agent common.Address, settingDigest [32]byte) (uint8, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "signerThreshold", agent, settingDigest)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SignerThreshold is a free data retrieval call binding the contract method 0x55200a8d.
//
// Solidity: function signerThreshold(address agent, bytes32 settingDigest) view returns(uint8)
func (_AgentProxy *AgentProxySession) SignerThreshold(agent common.Address, settingDigest [32]byte) (uint8, error) {
	return _AgentProxy.Contract.SignerThreshold(&_AgentProxy.CallOpts, agent, settingDigest)
}

// SignerThreshold is a free data retrieval call binding the contract method 0x55200a8d.
//
// Solidity: function signerThreshold(address agent, bytes32 settingDigest) view returns(uint8)
func (_AgentProxy *AgentProxyCallerSession) SignerThreshold(agent common.Address, settingDigest [32]byte) (uint8, error) {
	return _AgentProxy.Contract.SignerThreshold(&_AgentProxy.CallOpts, agent, settingDigest)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
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
// Solidity: function typeAndVersion() view returns(string)
func (_AgentProxy *AgentProxySession) TypeAndVersion() (string, error) {
	return _AgentProxy.Contract.TypeAndVersion(&_AgentProxy.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AgentProxy *AgentProxyCallerSession) TypeAndVersion() (string, error) {
	return _AgentProxy.Contract.TypeAndVersion(&_AgentProxy.CallOpts)
}

// ValidateDataConversion is a free data retrieval call binding the contract method 0x6a4c5139.
//
// Solidity: function validateDataConversion(address agent, bytes data) view returns(bytes)
func (_AgentProxy *AgentProxyCaller) ValidateDataConversion(opts *bind.CallOpts, agent common.Address, data []byte) ([]byte, error) {
	var out []interface{}
	err := _AgentProxy.contract.Call(opts, &out, "validateDataConversion", agent, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateDataConversion is a free data retrieval call binding the contract method 0x6a4c5139.
//
// Solidity: function validateDataConversion(address agent, bytes data) view returns(bytes)
func (_AgentProxy *AgentProxySession) ValidateDataConversion(agent common.Address, data []byte) ([]byte, error) {
	return _AgentProxy.Contract.ValidateDataConversion(&_AgentProxy.CallOpts, agent, data)
}

// ValidateDataConversion is a free data retrieval call binding the contract method 0x6a4c5139.
//
// Solidity: function validateDataConversion(address agent, bytes data) view returns(bytes)
func (_AgentProxy *AgentProxyCallerSession) ValidateDataConversion(agent common.Address, data []byte) ([]byte, error) {
	return _AgentProxy.Contract.ValidateDataConversion(&_AgentProxy.CallOpts, agent, data)
}

// AcceptAgent is a paid mutator transaction binding the contract method 0x906f61df.
//
// Solidity: function acceptAgent(address agent) returns()
func (_AgentProxy *AgentProxyTransactor) AcceptAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "acceptAgent", agent)
}

// AcceptAgent is a paid mutator transaction binding the contract method 0x906f61df.
//
// Solidity: function acceptAgent(address agent) returns()
func (_AgentProxy *AgentProxySession) AcceptAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.AcceptAgent(&_AgentProxy.TransactOpts, agent)
}

// AcceptAgent is a paid mutator transaction binding the contract method 0x906f61df.
//
// Solidity: function acceptAgent(address agent) returns()
func (_AgentProxy *AgentProxyTransactorSession) AcceptAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.AcceptAgent(&_AgentProxy.TransactOpts, agent)
}

// AcceptAgentSettingProposal is a paid mutator transaction binding the contract method 0x93771c84.
//
// Solidity: function acceptAgentSettingProposal(address agent) returns()
func (_AgentProxy *AgentProxyTransactor) AcceptAgentSettingProposal(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "acceptAgentSettingProposal", agent)
}

// AcceptAgentSettingProposal is a paid mutator transaction binding the contract method 0x93771c84.
//
// Solidity: function acceptAgentSettingProposal(address agent) returns()
func (_AgentProxy *AgentProxySession) AcceptAgentSettingProposal(agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.AcceptAgentSettingProposal(&_AgentProxy.TransactOpts, agent)
}

// AcceptAgentSettingProposal is a paid mutator transaction binding the contract method 0x93771c84.
//
// Solidity: function acceptAgentSettingProposal(address agent) returns()
func (_AgentProxy *AgentProxyTransactorSession) AcceptAgentSettingProposal(agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.AcceptAgentSettingProposal(&_AgentProxy.TransactOpts, agent)
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

// ChangeAgentSettingProposal is a paid mutator transaction binding the contract method 0x69fd4d8e.
//
// Solidity: function changeAgentSettingProposal(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxyTransactor) ChangeAgentSettingProposal(opts *bind.TransactOpts, agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "changeAgentSettingProposal", agent, agentSettings)
}

// ChangeAgentSettingProposal is a paid mutator transaction binding the contract method 0x69fd4d8e.
//
// Solidity: function changeAgentSettingProposal(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxySession) ChangeAgentSettingProposal(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.Contract.ChangeAgentSettingProposal(&_AgentProxy.TransactOpts, agent, agentSettings)
}

// ChangeAgentSettingProposal is a paid mutator transaction binding the contract method 0x69fd4d8e.
//
// Solidity: function changeAgentSettingProposal(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxyTransactorSession) ChangeAgentSettingProposal(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.Contract.ChangeAgentSettingProposal(&_AgentProxy.TransactOpts, agent, agentSettings)
}

// RegisterAgent is a paid mutator transaction binding the contract method 0x352f76e8.
//
// Solidity: function registerAgent(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxyTransactor) RegisterAgent(opts *bind.TransactOpts, agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "registerAgent", agent, agentSettings)
}

// RegisterAgent is a paid mutator transaction binding the contract method 0x352f76e8.
//
// Solidity: function registerAgent(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxySession) RegisterAgent(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.Contract.RegisterAgent(&_AgentProxy.TransactOpts, agent, agentSettings)
}

// RegisterAgent is a paid mutator transaction binding the contract method 0x352f76e8.
//
// Solidity: function registerAgent(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentProxy *AgentProxyTransactorSession) RegisterAgent(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentProxy.Contract.RegisterAgent(&_AgentProxy.TransactOpts, agent, agentSettings)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns()
func (_AgentProxy *AgentProxyTransactor) RemoveAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "removeAgent", agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns()
func (_AgentProxy *AgentProxySession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.RemoveAgent(&_AgentProxy.TransactOpts, agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns()
func (_AgentProxy *AgentProxyTransactorSession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.RemoveAgent(&_AgentProxy.TransactOpts, agent)
}

// SetAgentProxy is a paid mutator transaction binding the contract method 0xb78e6848.
//
// Solidity: function setAgentProxy(address proxy) returns()
func (_AgentProxy *AgentProxyTransactor) SetAgentProxy(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _AgentProxy.contract.Transact(opts, "setAgentProxy", proxy)
}

// SetAgentProxy is a paid mutator transaction binding the contract method 0xb78e6848.
//
// Solidity: function setAgentProxy(address proxy) returns()
func (_AgentProxy *AgentProxySession) SetAgentProxy(proxy common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.SetAgentProxy(&_AgentProxy.TransactOpts, proxy)
}

// SetAgentProxy is a paid mutator transaction binding the contract method 0xb78e6848.
//
// Solidity: function setAgentProxy(address proxy) returns()
func (_AgentProxy *AgentProxyTransactorSession) SetAgentProxy(proxy common.Address) (*types.Transaction, error) {
	return _AgentProxy.Contract.SetAgentProxy(&_AgentProxy.TransactOpts, proxy)
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

// AgentProxyAgentAcceptedIterator is returned from FilterAgentAccepted and is used to iterate over the raw logs and unpacked data for AgentAccepted events raised by the AgentProxy contract.
type AgentProxyAgentAcceptedIterator struct {
	Event *AgentProxyAgentAccepted // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentAccepted)
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
		it.Event = new(AgentProxyAgentAccepted)
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
func (it *AgentProxyAgentAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentAccepted represents a AgentAccepted event raised by the AgentProxy contract.
type AgentProxyAgentAccepted struct {
	Agent         common.Address
	Digest        [32]byte
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentAccepted is a free log retrieval operation binding the contract event 0xa36e3cae6a00fea59cd678be00a81b5254290a4a713fd275c29d2e4225d5ee2d.
//
// Solidity: event AgentAccepted(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) FilterAgentAccepted(opts *bind.FilterOpts, agent []common.Address, digest [][32]byte) (*AgentProxyAgentAcceptedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentAccepted", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentAcceptedIterator{contract: _AgentProxy.contract, event: "AgentAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentAccepted is a free log subscription operation binding the contract event 0xa36e3cae6a00fea59cd678be00a81b5254290a4a713fd275c29d2e4225d5ee2d.
//
// Solidity: event AgentAccepted(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) WatchAgentAccepted(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentAccepted, agent []common.Address, digest [][32]byte) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentAccepted", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentAccepted)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentAccepted", log); err != nil {
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

// ParseAgentAccepted is a log parse operation binding the contract event 0xa36e3cae6a00fea59cd678be00a81b5254290a4a713fd275c29d2e4225d5ee2d.
//
// Solidity: event AgentAccepted(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) ParseAgentAccepted(log types.Log) (*AgentProxyAgentAccepted, error) {
	event := new(AgentProxyAgentAccepted)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyAgentProxySetIterator is returned from FilterAgentProxySet and is used to iterate over the raw logs and unpacked data for AgentProxySet events raised by the AgentProxy contract.
type AgentProxyAgentProxySetIterator struct {
	Event *AgentProxyAgentProxySet // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentProxySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentProxySet)
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
		it.Event = new(AgentProxyAgentProxySet)
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
func (it *AgentProxyAgentProxySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentProxySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentProxySet represents a AgentProxySet event raised by the AgentProxy contract.
type AgentProxyAgentProxySet struct {
	OldProxy common.Address
	NewProxy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentProxySet is a free log retrieval operation binding the contract event 0x027832fe8b1cfd252a8b83993255c855b482dd5caa679217465ff9106fe13cba.
//
// Solidity: event AgentProxySet(address oldProxy, address newProxy)
func (_AgentProxy *AgentProxyFilterer) FilterAgentProxySet(opts *bind.FilterOpts) (*AgentProxyAgentProxySetIterator, error) {

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentProxySet")
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentProxySetIterator{contract: _AgentProxy.contract, event: "AgentProxySet", logs: logs, sub: sub}, nil
}

// WatchAgentProxySet is a free log subscription operation binding the contract event 0x027832fe8b1cfd252a8b83993255c855b482dd5caa679217465ff9106fe13cba.
//
// Solidity: event AgentProxySet(address oldProxy, address newProxy)
func (_AgentProxy *AgentProxyFilterer) WatchAgentProxySet(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentProxySet) (event.Subscription, error) {

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentProxySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentProxySet)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentProxySet", log); err != nil {
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

// ParseAgentProxySet is a log parse operation binding the contract event 0x027832fe8b1cfd252a8b83993255c855b482dd5caa679217465ff9106fe13cba.
//
// Solidity: event AgentProxySet(address oldProxy, address newProxy)
func (_AgentProxy *AgentProxyFilterer) ParseAgentProxySet(log types.Log) (*AgentProxyAgentProxySet, error) {
	event := new(AgentProxyAgentProxySet)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentProxySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyAgentRegisteredIterator is returned from FilterAgentRegistered and is used to iterate over the raw logs and unpacked data for AgentRegistered events raised by the AgentProxy contract.
type AgentProxyAgentRegisteredIterator struct {
	Event *AgentProxyAgentRegistered // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentRegistered)
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
		it.Event = new(AgentProxyAgentRegistered)
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
func (it *AgentProxyAgentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentRegistered represents a AgentRegistered event raised by the AgentProxy contract.
type AgentProxyAgentRegistered struct {
	Agent         common.Address
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentRegistered is a free log retrieval operation binding the contract event 0x8f9f5acc73d739a9bfdd8cd64d17b9e9bb10217585d472d17444462339bccaed.
//
// Solidity: event AgentRegistered(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) FilterAgentRegistered(opts *bind.FilterOpts, agent []common.Address) (*AgentProxyAgentRegisteredIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentRegistered", agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentRegisteredIterator{contract: _AgentProxy.contract, event: "AgentRegistered", logs: logs, sub: sub}, nil
}

// WatchAgentRegistered is a free log subscription operation binding the contract event 0x8f9f5acc73d739a9bfdd8cd64d17b9e9bb10217585d472d17444462339bccaed.
//
// Solidity: event AgentRegistered(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) WatchAgentRegistered(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentRegistered, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentRegistered", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentRegistered)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
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

// ParseAgentRegistered is a log parse operation binding the contract event 0x8f9f5acc73d739a9bfdd8cd64d17b9e9bb10217585d472d17444462339bccaed.
//
// Solidity: event AgentRegistered(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) ParseAgentRegistered(log types.Log) (*AgentProxyAgentRegistered, error) {
	event := new(AgentProxyAgentRegistered)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the AgentProxy contract.
type AgentProxyAgentRemovedIterator struct {
	Event *AgentProxyAgentRemoved // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentRemoved)
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
		it.Event = new(AgentProxyAgentRemoved)
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
func (it *AgentProxyAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentRemoved represents a AgentRemoved event raised by the AgentProxy contract.
type AgentProxyAgentRemoved struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed agent)
func (_AgentProxy *AgentProxyFilterer) FilterAgentRemoved(opts *bind.FilterOpts, agent []common.Address) (*AgentProxyAgentRemovedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentRemoved", agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentRemovedIterator{contract: _AgentProxy.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed agent)
func (_AgentProxy *AgentProxyFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentRemoved, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentRemoved", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentRemoved)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed agent)
func (_AgentProxy *AgentProxyFilterer) ParseAgentRemoved(log types.Log) (*AgentProxyAgentRemoved, error) {
	event := new(AgentProxyAgentRemoved)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyAgentSettingsProposedIterator is returned from FilterAgentSettingsProposed and is used to iterate over the raw logs and unpacked data for AgentSettingsProposed events raised by the AgentProxy contract.
type AgentProxyAgentSettingsProposedIterator struct {
	Event *AgentProxyAgentSettingsProposed // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentSettingsProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentSettingsProposed)
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
		it.Event = new(AgentProxyAgentSettingsProposed)
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
func (it *AgentProxyAgentSettingsProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentSettingsProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentSettingsProposed represents a AgentSettingsProposed event raised by the AgentProxy contract.
type AgentProxyAgentSettingsProposed struct {
	Agent         common.Address
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentSettingsProposed is a free log retrieval operation binding the contract event 0x35f9d2cf50394e31e48ee780e6471bb9e0e3cceadc62d25d1349a94bdd7a9378.
//
// Solidity: event AgentSettingsProposed(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) FilterAgentSettingsProposed(opts *bind.FilterOpts, agent []common.Address) (*AgentProxyAgentSettingsProposedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentSettingsProposed", agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentSettingsProposedIterator{contract: _AgentProxy.contract, event: "AgentSettingsProposed", logs: logs, sub: sub}, nil
}

// WatchAgentSettingsProposed is a free log subscription operation binding the contract event 0x35f9d2cf50394e31e48ee780e6471bb9e0e3cceadc62d25d1349a94bdd7a9378.
//
// Solidity: event AgentSettingsProposed(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) WatchAgentSettingsProposed(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentSettingsProposed, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentSettingsProposed", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentSettingsProposed)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentSettingsProposed", log); err != nil {
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

// ParseAgentSettingsProposed is a log parse operation binding the contract event 0x35f9d2cf50394e31e48ee780e6471bb9e0e3cceadc62d25d1349a94bdd7a9378.
//
// Solidity: event AgentSettingsProposed(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) ParseAgentSettingsProposed(log types.Log) (*AgentProxyAgentSettingsProposed, error) {
	event := new(AgentProxyAgentSettingsProposed)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentSettingsProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentProxyAgentSettingsUpdatedIterator is returned from FilterAgentSettingsUpdated and is used to iterate over the raw logs and unpacked data for AgentSettingsUpdated events raised by the AgentProxy contract.
type AgentProxyAgentSettingsUpdatedIterator struct {
	Event *AgentProxyAgentSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *AgentProxyAgentSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentProxyAgentSettingsUpdated)
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
		it.Event = new(AgentProxyAgentSettingsUpdated)
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
func (it *AgentProxyAgentSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentProxyAgentSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentProxyAgentSettingsUpdated represents a AgentSettingsUpdated event raised by the AgentProxy contract.
type AgentProxyAgentSettingsUpdated struct {
	Agent         common.Address
	Digest        [32]byte
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentSettingsUpdated is a free log retrieval operation binding the contract event 0x9cbec0da5ddc48f6a581101583fb4487d38599559a940e8ebf031857881a79a1.
//
// Solidity: event AgentSettingsUpdated(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) FilterAgentSettingsUpdated(opts *bind.FilterOpts, agent []common.Address, digest [][32]byte) (*AgentProxyAgentSettingsUpdatedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentProxy.contract.FilterLogs(opts, "AgentSettingsUpdated", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &AgentProxyAgentSettingsUpdatedIterator{contract: _AgentProxy.contract, event: "AgentSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentSettingsUpdated is a free log subscription operation binding the contract event 0x9cbec0da5ddc48f6a581101583fb4487d38599559a940e8ebf031857881a79a1.
//
// Solidity: event AgentSettingsUpdated(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) WatchAgentSettingsUpdated(opts *bind.WatchOpts, sink chan<- *AgentProxyAgentSettingsUpdated, agent []common.Address, digest [][32]byte) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentProxy.contract.WatchLogs(opts, "AgentSettingsUpdated", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentProxyAgentSettingsUpdated)
				if err := _AgentProxy.contract.UnpackLog(event, "AgentSettingsUpdated", log); err != nil {
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

// ParseAgentSettingsUpdated is a log parse operation binding the contract event 0x9cbec0da5ddc48f6a581101583fb4487d38599559a940e8ebf031857881a79a1.
//
// Solidity: event AgentSettingsUpdated(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentProxy *AgentProxyFilterer) ParseAgentSettingsUpdated(log types.Log) (*AgentProxyAgentSettingsUpdated, error) {
	event := new(AgentProxyAgentSettingsUpdated)
	if err := _AgentProxy.contract.UnpackLog(event, "AgentSettingsUpdated", log); err != nil {
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
