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
	IsActive          bool
	Settings          CommonAgentSettings
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

// AgentManagerMetaData contains all meta data concerning the AgentManager contract.
var AgentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentIsAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentIsRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderAgentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderMessageId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderMessageType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderPriority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentHeaderVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentSettingProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowedAgent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCallData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFactoryAgent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegisteredAgent\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"AgentProxySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentSettingsProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"AgentSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"acceptAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"acceptAgentSettingProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"allowedAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingDigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"allowedSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"changeAgentSettingProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingDigest\",\"type\":\"bytes32\"}],\"name\":\"getAgentConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"configBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getAgentConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"configBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getAgentConfigsCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"agentConfigIdxStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"agentConfigIdxEnd\",\"type\":\"uint64\"}],\"name\":\"getAgentConfigsInRange\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"configBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentConfig[]\",\"name\":\"agentConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAllowedAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRegisteringAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowedAgentsCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"allowedAgentIdxStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"allowedAgentIdxEnd\",\"type\":\"uint64\"}],\"name\":\"getAllowedAgentsInRange\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"allowedAgents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteringAgentsCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registeringAgentIdxStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"registeringAgentIdxEnd\",\"type\":\"uint64\"}],\"name\":\"getRegisteringAgentsInRange\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"registeringAgents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"}],\"name\":\"isValidMessageId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"}],\"name\":\"isValidSourceAgentId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"converterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"messageId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAgentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetAgentId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCommon.MessageType\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"enumCommon.Priority\",\"name\":\"priority\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AgentHeader\",\"name\":\"agentHeader\",\"type\":\"tuple\"}],\"internalType\":\"structCommon.AgentSettings\",\"name\":\"agentSettings\",\"type\":\"tuple\"}],\"name\":\"registerAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"setAgentProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"settingDigest\",\"type\":\"bytes32\"}],\"name\":\"signerThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validateDataConversion\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AgentManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentManagerMetaData.ABI instead.
var AgentManagerABI = AgentManagerMetaData.ABI

// AgentManager is an auto generated Go binding around an Ethereum contract.
type AgentManager struct {
	AgentManagerCaller     // Read-only binding to the contract
	AgentManagerTransactor // Write-only binding to the contract
	AgentManagerFilterer   // Log filterer for contract events
}

// AgentManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentManagerSession struct {
	Contract     *AgentManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentManagerCallerSession struct {
	Contract *AgentManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AgentManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentManagerTransactorSession struct {
	Contract     *AgentManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AgentManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentManagerRaw struct {
	Contract *AgentManager // Generic contract binding to access the raw methods on
}

// AgentManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentManagerCallerRaw struct {
	Contract *AgentManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AgentManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentManagerTransactorRaw struct {
	Contract *AgentManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentManager creates a new instance of AgentManager, bound to a specific deployed contract.
func NewAgentManager(address common.Address, backend bind.ContractBackend) (*AgentManager, error) {
	contract, err := bindAgentManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentManager{AgentManagerCaller: AgentManagerCaller{contract: contract}, AgentManagerTransactor: AgentManagerTransactor{contract: contract}, AgentManagerFilterer: AgentManagerFilterer{contract: contract}}, nil
}

// NewAgentManagerCaller creates a new read-only instance of AgentManager, bound to a specific deployed contract.
func NewAgentManagerCaller(address common.Address, caller bind.ContractCaller) (*AgentManagerCaller, error) {
	contract, err := bindAgentManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerCaller{contract: contract}, nil
}

// NewAgentManagerTransactor creates a new write-only instance of AgentManager, bound to a specific deployed contract.
func NewAgentManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentManagerTransactor, error) {
	contract, err := bindAgentManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerTransactor{contract: contract}, nil
}

// NewAgentManagerFilterer creates a new log filterer instance of AgentManager, bound to a specific deployed contract.
func NewAgentManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentManagerFilterer, error) {
	contract, err := bindAgentManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentManagerFilterer{contract: contract}, nil
}

// bindAgentManager binds a generic wrapper to an already deployed contract.
func bindAgentManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManager *AgentManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManager.Contract.AgentManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManager *AgentManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManager.Contract.AgentManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManager *AgentManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManager.Contract.AgentManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManager *AgentManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManager *AgentManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManager *AgentManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManager.Contract.contract.Transact(opts, method, params...)
}

// AgentProxy is a free data retrieval call binding the contract method 0x81070369.
//
// Solidity: function agentProxy() view returns(address)
func (_AgentManager *AgentManagerCaller) AgentProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "agentProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentProxy is a free data retrieval call binding the contract method 0x81070369.
//
// Solidity: function agentProxy() view returns(address)
func (_AgentManager *AgentManagerSession) AgentProxy() (common.Address, error) {
	return _AgentManager.Contract.AgentProxy(&_AgentManager.CallOpts)
}

// AgentProxy is a free data retrieval call binding the contract method 0x81070369.
//
// Solidity: function agentProxy() view returns(address)
func (_AgentManager *AgentManagerCallerSession) AgentProxy() (common.Address, error) {
	return _AgentManager.Contract.AgentProxy(&_AgentManager.CallOpts)
}

// AgentVersion is a free data retrieval call binding the contract method 0x10c59c72.
//
// Solidity: function agentVersion() pure returns(string)
func (_AgentManager *AgentManagerCaller) AgentVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "agentVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AgentVersion is a free data retrieval call binding the contract method 0x10c59c72.
//
// Solidity: function agentVersion() pure returns(string)
func (_AgentManager *AgentManagerSession) AgentVersion() (string, error) {
	return _AgentManager.Contract.AgentVersion(&_AgentManager.CallOpts)
}

// AgentVersion is a free data retrieval call binding the contract method 0x10c59c72.
//
// Solidity: function agentVersion() pure returns(string)
func (_AgentManager *AgentManagerCallerSession) AgentVersion() (string, error) {
	return _AgentManager.Contract.AgentVersion(&_AgentManager.CallOpts)
}

// AllowedAgent is a free data retrieval call binding the contract method 0x1c774d11.
//
// Solidity: function allowedAgent(address agent) view returns(bool)
func (_AgentManager *AgentManagerCaller) AllowedAgent(opts *bind.CallOpts, agent common.Address) (bool, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "allowedAgent", agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAgent is a free data retrieval call binding the contract method 0x1c774d11.
//
// Solidity: function allowedAgent(address agent) view returns(bool)
func (_AgentManager *AgentManagerSession) AllowedAgent(agent common.Address) (bool, error) {
	return _AgentManager.Contract.AllowedAgent(&_AgentManager.CallOpts, agent)
}

// AllowedAgent is a free data retrieval call binding the contract method 0x1c774d11.
//
// Solidity: function allowedAgent(address agent) view returns(bool)
func (_AgentManager *AgentManagerCallerSession) AllowedAgent(agent common.Address) (bool, error) {
	return _AgentManager.Contract.AllowedAgent(&_AgentManager.CallOpts, agent)
}

// AllowedSigner is a free data retrieval call binding the contract method 0x825a10b1.
//
// Solidity: function allowedSigner(address agent, bytes32 settingDigest, address signer) view returns(bool)
func (_AgentManager *AgentManagerCaller) AllowedSigner(opts *bind.CallOpts, agent common.Address, settingDigest [32]byte, signer common.Address) (bool, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "allowedSigner", agent, settingDigest, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedSigner is a free data retrieval call binding the contract method 0x825a10b1.
//
// Solidity: function allowedSigner(address agent, bytes32 settingDigest, address signer) view returns(bool)
func (_AgentManager *AgentManagerSession) AllowedSigner(agent common.Address, settingDigest [32]byte, signer common.Address) (bool, error) {
	return _AgentManager.Contract.AllowedSigner(&_AgentManager.CallOpts, agent, settingDigest, signer)
}

// AllowedSigner is a free data retrieval call binding the contract method 0x825a10b1.
//
// Solidity: function allowedSigner(address agent, bytes32 settingDigest, address signer) view returns(bool)
func (_AgentManager *AgentManagerCallerSession) AllowedSigner(agent common.Address, settingDigest [32]byte, signer common.Address) (bool, error) {
	return _AgentManager.Contract.AllowedSigner(&_AgentManager.CallOpts, agent, settingDigest, signer)
}

// GetAgentConfig is a free data retrieval call binding the contract method 0xc703f3f9.
//
// Solidity: function getAgentConfig(address agent, bytes32 settingDigest) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256))))
func (_AgentManager *AgentManagerCaller) GetAgentConfig(opts *bind.CallOpts, agent common.Address, settingDigest [32]byte) (CommonAgentConfig, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAgentConfig", agent, settingDigest)

	if err != nil {
		return *new(CommonAgentConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CommonAgentConfig)).(*CommonAgentConfig)

	return out0, err

}

// GetAgentConfig is a free data retrieval call binding the contract method 0xc703f3f9.
//
// Solidity: function getAgentConfig(address agent, bytes32 settingDigest) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256))))
func (_AgentManager *AgentManagerSession) GetAgentConfig(agent common.Address, settingDigest [32]byte) (CommonAgentConfig, error) {
	return _AgentManager.Contract.GetAgentConfig(&_AgentManager.CallOpts, agent, settingDigest)
}

// GetAgentConfig is a free data retrieval call binding the contract method 0xc703f3f9.
//
// Solidity: function getAgentConfig(address agent, bytes32 settingDigest) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256))))
func (_AgentManager *AgentManagerCallerSession) GetAgentConfig(agent common.Address, settingDigest [32]byte) (CommonAgentConfig, error) {
	return _AgentManager.Contract.GetAgentConfig(&_AgentManager.CallOpts, agent, settingDigest)
}

// GetAgentConfigs is a free data retrieval call binding the contract method 0x6ccfcf63.
//
// Solidity: function getAgentConfigs(address agent) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[])
func (_AgentManager *AgentManagerCaller) GetAgentConfigs(opts *bind.CallOpts, agent common.Address) ([]CommonAgentConfig, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAgentConfigs", agent)

	if err != nil {
		return *new([]CommonAgentConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]CommonAgentConfig)).(*[]CommonAgentConfig)

	return out0, err

}

// GetAgentConfigs is a free data retrieval call binding the contract method 0x6ccfcf63.
//
// Solidity: function getAgentConfigs(address agent) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[])
func (_AgentManager *AgentManagerSession) GetAgentConfigs(agent common.Address) ([]CommonAgentConfig, error) {
	return _AgentManager.Contract.GetAgentConfigs(&_AgentManager.CallOpts, agent)
}

// GetAgentConfigs is a free data retrieval call binding the contract method 0x6ccfcf63.
//
// Solidity: function getAgentConfigs(address agent) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[])
func (_AgentManager *AgentManagerCallerSession) GetAgentConfigs(agent common.Address) ([]CommonAgentConfig, error) {
	return _AgentManager.Contract.GetAgentConfigs(&_AgentManager.CallOpts, agent)
}

// GetAgentConfigsCount is a free data retrieval call binding the contract method 0xd2c0823a.
//
// Solidity: function getAgentConfigsCount(address agent) view returns(uint64)
func (_AgentManager *AgentManagerCaller) GetAgentConfigsCount(opts *bind.CallOpts, agent common.Address) (uint64, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAgentConfigsCount", agent)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAgentConfigsCount is a free data retrieval call binding the contract method 0xd2c0823a.
//
// Solidity: function getAgentConfigsCount(address agent) view returns(uint64)
func (_AgentManager *AgentManagerSession) GetAgentConfigsCount(agent common.Address) (uint64, error) {
	return _AgentManager.Contract.GetAgentConfigsCount(&_AgentManager.CallOpts, agent)
}

// GetAgentConfigsCount is a free data retrieval call binding the contract method 0xd2c0823a.
//
// Solidity: function getAgentConfigsCount(address agent) view returns(uint64)
func (_AgentManager *AgentManagerCallerSession) GetAgentConfigsCount(agent common.Address) (uint64, error) {
	return _AgentManager.Contract.GetAgentConfigsCount(&_AgentManager.CallOpts, agent)
}

// GetAgentConfigsInRange is a free data retrieval call binding the contract method 0x1b429f1d.
//
// Solidity: function getAgentConfigsInRange(address agent, uint64 agentConfigIdxStart, uint64 agentConfigIdxEnd) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[] agentConfigs)
func (_AgentManager *AgentManagerCaller) GetAgentConfigsInRange(opts *bind.CallOpts, agent common.Address, agentConfigIdxStart uint64, agentConfigIdxEnd uint64) ([]CommonAgentConfig, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAgentConfigsInRange", agent, agentConfigIdxStart, agentConfigIdxEnd)

	if err != nil {
		return *new([]CommonAgentConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]CommonAgentConfig)).(*[]CommonAgentConfig)

	return out0, err

}

// GetAgentConfigsInRange is a free data retrieval call binding the contract method 0x1b429f1d.
//
// Solidity: function getAgentConfigsInRange(address agent, uint64 agentConfigIdxStart, uint64 agentConfigIdxEnd) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[] agentConfigs)
func (_AgentManager *AgentManagerSession) GetAgentConfigsInRange(agent common.Address, agentConfigIdxStart uint64, agentConfigIdxEnd uint64) ([]CommonAgentConfig, error) {
	return _AgentManager.Contract.GetAgentConfigsInRange(&_AgentManager.CallOpts, agent, agentConfigIdxStart, agentConfigIdxEnd)
}

// GetAgentConfigsInRange is a free data retrieval call binding the contract method 0x1b429f1d.
//
// Solidity: function getAgentConfigsInRange(address agent, uint64 agentConfigIdxStart, uint64 agentConfigIdxEnd) view returns((bytes32,uint32,bool,(address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)))[] agentConfigs)
func (_AgentManager *AgentManagerCallerSession) GetAgentConfigsInRange(agent common.Address, agentConfigIdxStart uint64, agentConfigIdxEnd uint64) ([]CommonAgentConfig, error) {
	return _AgentManager.Contract.GetAgentConfigsInRange(&_AgentManager.CallOpts, agent, agentConfigIdxStart, agentConfigIdxEnd)
}

// GetAllAllowedAgents is a free data retrieval call binding the contract method 0xac3e7680.
//
// Solidity: function getAllAllowedAgents() view returns(address[])
func (_AgentManager *AgentManagerCaller) GetAllAllowedAgents(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAllAllowedAgents")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAllowedAgents is a free data retrieval call binding the contract method 0xac3e7680.
//
// Solidity: function getAllAllowedAgents() view returns(address[])
func (_AgentManager *AgentManagerSession) GetAllAllowedAgents() ([]common.Address, error) {
	return _AgentManager.Contract.GetAllAllowedAgents(&_AgentManager.CallOpts)
}

// GetAllAllowedAgents is a free data retrieval call binding the contract method 0xac3e7680.
//
// Solidity: function getAllAllowedAgents() view returns(address[])
func (_AgentManager *AgentManagerCallerSession) GetAllAllowedAgents() ([]common.Address, error) {
	return _AgentManager.Contract.GetAllAllowedAgents(&_AgentManager.CallOpts)
}

// GetAllRegisteringAgents is a free data retrieval call binding the contract method 0xa73e07b4.
//
// Solidity: function getAllRegisteringAgents() view returns(address[])
func (_AgentManager *AgentManagerCaller) GetAllRegisteringAgents(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAllRegisteringAgents")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllRegisteringAgents is a free data retrieval call binding the contract method 0xa73e07b4.
//
// Solidity: function getAllRegisteringAgents() view returns(address[])
func (_AgentManager *AgentManagerSession) GetAllRegisteringAgents() ([]common.Address, error) {
	return _AgentManager.Contract.GetAllRegisteringAgents(&_AgentManager.CallOpts)
}

// GetAllRegisteringAgents is a free data retrieval call binding the contract method 0xa73e07b4.
//
// Solidity: function getAllRegisteringAgents() view returns(address[])
func (_AgentManager *AgentManagerCallerSession) GetAllRegisteringAgents() ([]common.Address, error) {
	return _AgentManager.Contract.GetAllRegisteringAgents(&_AgentManager.CallOpts)
}

// GetAllowedAgentsCount is a free data retrieval call binding the contract method 0xa251f4fb.
//
// Solidity: function getAllowedAgentsCount() view returns(uint64)
func (_AgentManager *AgentManagerCaller) GetAllowedAgentsCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAllowedAgentsCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllowedAgentsCount is a free data retrieval call binding the contract method 0xa251f4fb.
//
// Solidity: function getAllowedAgentsCount() view returns(uint64)
func (_AgentManager *AgentManagerSession) GetAllowedAgentsCount() (uint64, error) {
	return _AgentManager.Contract.GetAllowedAgentsCount(&_AgentManager.CallOpts)
}

// GetAllowedAgentsCount is a free data retrieval call binding the contract method 0xa251f4fb.
//
// Solidity: function getAllowedAgentsCount() view returns(uint64)
func (_AgentManager *AgentManagerCallerSession) GetAllowedAgentsCount() (uint64, error) {
	return _AgentManager.Contract.GetAllowedAgentsCount(&_AgentManager.CallOpts)
}

// GetAllowedAgentsInRange is a free data retrieval call binding the contract method 0x8eefb433.
//
// Solidity: function getAllowedAgentsInRange(uint64 allowedAgentIdxStart, uint64 allowedAgentIdxEnd) view returns(address[] allowedAgents)
func (_AgentManager *AgentManagerCaller) GetAllowedAgentsInRange(opts *bind.CallOpts, allowedAgentIdxStart uint64, allowedAgentIdxEnd uint64) ([]common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAllowedAgentsInRange", allowedAgentIdxStart, allowedAgentIdxEnd)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedAgentsInRange is a free data retrieval call binding the contract method 0x8eefb433.
//
// Solidity: function getAllowedAgentsInRange(uint64 allowedAgentIdxStart, uint64 allowedAgentIdxEnd) view returns(address[] allowedAgents)
func (_AgentManager *AgentManagerSession) GetAllowedAgentsInRange(allowedAgentIdxStart uint64, allowedAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentManager.Contract.GetAllowedAgentsInRange(&_AgentManager.CallOpts, allowedAgentIdxStart, allowedAgentIdxEnd)
}

// GetAllowedAgentsInRange is a free data retrieval call binding the contract method 0x8eefb433.
//
// Solidity: function getAllowedAgentsInRange(uint64 allowedAgentIdxStart, uint64 allowedAgentIdxEnd) view returns(address[] allowedAgents)
func (_AgentManager *AgentManagerCallerSession) GetAllowedAgentsInRange(allowedAgentIdxStart uint64, allowedAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentManager.Contract.GetAllowedAgentsInRange(&_AgentManager.CallOpts, allowedAgentIdxStart, allowedAgentIdxEnd)
}

// GetRegisteringAgentsCount is a free data retrieval call binding the contract method 0xb2833451.
//
// Solidity: function getRegisteringAgentsCount() view returns(uint64)
func (_AgentManager *AgentManagerCaller) GetRegisteringAgentsCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getRegisteringAgentsCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRegisteringAgentsCount is a free data retrieval call binding the contract method 0xb2833451.
//
// Solidity: function getRegisteringAgentsCount() view returns(uint64)
func (_AgentManager *AgentManagerSession) GetRegisteringAgentsCount() (uint64, error) {
	return _AgentManager.Contract.GetRegisteringAgentsCount(&_AgentManager.CallOpts)
}

// GetRegisteringAgentsCount is a free data retrieval call binding the contract method 0xb2833451.
//
// Solidity: function getRegisteringAgentsCount() view returns(uint64)
func (_AgentManager *AgentManagerCallerSession) GetRegisteringAgentsCount() (uint64, error) {
	return _AgentManager.Contract.GetRegisteringAgentsCount(&_AgentManager.CallOpts)
}

// GetRegisteringAgentsInRange is a free data retrieval call binding the contract method 0xdb971453.
//
// Solidity: function getRegisteringAgentsInRange(uint64 registeringAgentIdxStart, uint64 registeringAgentIdxEnd) view returns(address[] registeringAgents)
func (_AgentManager *AgentManagerCaller) GetRegisteringAgentsInRange(opts *bind.CallOpts, registeringAgentIdxStart uint64, registeringAgentIdxEnd uint64) ([]common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getRegisteringAgentsInRange", registeringAgentIdxStart, registeringAgentIdxEnd)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteringAgentsInRange is a free data retrieval call binding the contract method 0xdb971453.
//
// Solidity: function getRegisteringAgentsInRange(uint64 registeringAgentIdxStart, uint64 registeringAgentIdxEnd) view returns(address[] registeringAgents)
func (_AgentManager *AgentManagerSession) GetRegisteringAgentsInRange(registeringAgentIdxStart uint64, registeringAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentManager.Contract.GetRegisteringAgentsInRange(&_AgentManager.CallOpts, registeringAgentIdxStart, registeringAgentIdxEnd)
}

// GetRegisteringAgentsInRange is a free data retrieval call binding the contract method 0xdb971453.
//
// Solidity: function getRegisteringAgentsInRange(uint64 registeringAgentIdxStart, uint64 registeringAgentIdxEnd) view returns(address[] registeringAgents)
func (_AgentManager *AgentManagerCallerSession) GetRegisteringAgentsInRange(registeringAgentIdxStart uint64, registeringAgentIdxEnd uint64) ([]common.Address, error) {
	return _AgentManager.Contract.GetRegisteringAgentsInRange(&_AgentManager.CallOpts, registeringAgentIdxStart, registeringAgentIdxEnd)
}

// IsValidMessageId is a free data retrieval call binding the contract method 0x4dfbe7c6.
//
// Solidity: function isValidMessageId(string messageId) pure returns(bool)
func (_AgentManager *AgentManagerCaller) IsValidMessageId(opts *bind.CallOpts, messageId string) (bool, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "isValidMessageId", messageId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidMessageId is a free data retrieval call binding the contract method 0x4dfbe7c6.
//
// Solidity: function isValidMessageId(string messageId) pure returns(bool)
func (_AgentManager *AgentManagerSession) IsValidMessageId(messageId string) (bool, error) {
	return _AgentManager.Contract.IsValidMessageId(&_AgentManager.CallOpts, messageId)
}

// IsValidMessageId is a free data retrieval call binding the contract method 0x4dfbe7c6.
//
// Solidity: function isValidMessageId(string messageId) pure returns(bool)
func (_AgentManager *AgentManagerCallerSession) IsValidMessageId(messageId string) (bool, error) {
	return _AgentManager.Contract.IsValidMessageId(&_AgentManager.CallOpts, messageId)
}

// IsValidSourceAgentId is a free data retrieval call binding the contract method 0x0e553cae.
//
// Solidity: function isValidSourceAgentId(string sourceAgentId) view returns(bool)
func (_AgentManager *AgentManagerCaller) IsValidSourceAgentId(opts *bind.CallOpts, sourceAgentId string) (bool, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "isValidSourceAgentId", sourceAgentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSourceAgentId is a free data retrieval call binding the contract method 0x0e553cae.
//
// Solidity: function isValidSourceAgentId(string sourceAgentId) view returns(bool)
func (_AgentManager *AgentManagerSession) IsValidSourceAgentId(sourceAgentId string) (bool, error) {
	return _AgentManager.Contract.IsValidSourceAgentId(&_AgentManager.CallOpts, sourceAgentId)
}

// IsValidSourceAgentId is a free data retrieval call binding the contract method 0x0e553cae.
//
// Solidity: function isValidSourceAgentId(string sourceAgentId) view returns(bool)
func (_AgentManager *AgentManagerCallerSession) IsValidSourceAgentId(sourceAgentId string) (bool, error) {
	return _AgentManager.Contract.IsValidSourceAgentId(&_AgentManager.CallOpts, sourceAgentId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManager *AgentManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManager *AgentManagerSession) Owner() (common.Address, error) {
	return _AgentManager.Contract.Owner(&_AgentManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManager *AgentManagerCallerSession) Owner() (common.Address, error) {
	return _AgentManager.Contract.Owner(&_AgentManager.CallOpts)
}

// SignerThreshold is a free data retrieval call binding the contract method 0x55200a8d.
//
// Solidity: function signerThreshold(address agent, bytes32 settingDigest) view returns(uint8)
func (_AgentManager *AgentManagerCaller) SignerThreshold(opts *bind.CallOpts, agent common.Address, settingDigest [32]byte) (uint8, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "signerThreshold", agent, settingDigest)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SignerThreshold is a free data retrieval call binding the contract method 0x55200a8d.
//
// Solidity: function signerThreshold(address agent, bytes32 settingDigest) view returns(uint8)
func (_AgentManager *AgentManagerSession) SignerThreshold(agent common.Address, settingDigest [32]byte) (uint8, error) {
	return _AgentManager.Contract.SignerThreshold(&_AgentManager.CallOpts, agent, settingDigest)
}

// SignerThreshold is a free data retrieval call binding the contract method 0x55200a8d.
//
// Solidity: function signerThreshold(address agent, bytes32 settingDigest) view returns(uint8)
func (_AgentManager *AgentManagerCallerSession) SignerThreshold(agent common.Address, settingDigest [32]byte) (uint8, error) {
	return _AgentManager.Contract.SignerThreshold(&_AgentManager.CallOpts, agent, settingDigest)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AgentManager *AgentManagerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AgentManager *AgentManagerSession) TypeAndVersion() (string, error) {
	return _AgentManager.Contract.TypeAndVersion(&_AgentManager.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AgentManager *AgentManagerCallerSession) TypeAndVersion() (string, error) {
	return _AgentManager.Contract.TypeAndVersion(&_AgentManager.CallOpts)
}

// ValidateDataConversion is a free data retrieval call binding the contract method 0x6a4c5139.
//
// Solidity: function validateDataConversion(address agent, bytes data) view returns(bytes)
func (_AgentManager *AgentManagerCaller) ValidateDataConversion(opts *bind.CallOpts, agent common.Address, data []byte) ([]byte, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "validateDataConversion", agent, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateDataConversion is a free data retrieval call binding the contract method 0x6a4c5139.
//
// Solidity: function validateDataConversion(address agent, bytes data) view returns(bytes)
func (_AgentManager *AgentManagerSession) ValidateDataConversion(agent common.Address, data []byte) ([]byte, error) {
	return _AgentManager.Contract.ValidateDataConversion(&_AgentManager.CallOpts, agent, data)
}

// ValidateDataConversion is a free data retrieval call binding the contract method 0x6a4c5139.
//
// Solidity: function validateDataConversion(address agent, bytes data) view returns(bytes)
func (_AgentManager *AgentManagerCallerSession) ValidateDataConversion(agent common.Address, data []byte) ([]byte, error) {
	return _AgentManager.Contract.ValidateDataConversion(&_AgentManager.CallOpts, agent, data)
}

// AcceptAgent is a paid mutator transaction binding the contract method 0x906f61df.
//
// Solidity: function acceptAgent(address agent) returns()
func (_AgentManager *AgentManagerTransactor) AcceptAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "acceptAgent", agent)
}

// AcceptAgent is a paid mutator transaction binding the contract method 0x906f61df.
//
// Solidity: function acceptAgent(address agent) returns()
func (_AgentManager *AgentManagerSession) AcceptAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.AcceptAgent(&_AgentManager.TransactOpts, agent)
}

// AcceptAgent is a paid mutator transaction binding the contract method 0x906f61df.
//
// Solidity: function acceptAgent(address agent) returns()
func (_AgentManager *AgentManagerTransactorSession) AcceptAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.AcceptAgent(&_AgentManager.TransactOpts, agent)
}

// AcceptAgentSettingProposal is a paid mutator transaction binding the contract method 0x93771c84.
//
// Solidity: function acceptAgentSettingProposal(address agent) returns()
func (_AgentManager *AgentManagerTransactor) AcceptAgentSettingProposal(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "acceptAgentSettingProposal", agent)
}

// AcceptAgentSettingProposal is a paid mutator transaction binding the contract method 0x93771c84.
//
// Solidity: function acceptAgentSettingProposal(address agent) returns()
func (_AgentManager *AgentManagerSession) AcceptAgentSettingProposal(agent common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.AcceptAgentSettingProposal(&_AgentManager.TransactOpts, agent)
}

// AcceptAgentSettingProposal is a paid mutator transaction binding the contract method 0x93771c84.
//
// Solidity: function acceptAgentSettingProposal(address agent) returns()
func (_AgentManager *AgentManagerTransactorSession) AcceptAgentSettingProposal(agent common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.AcceptAgentSettingProposal(&_AgentManager.TransactOpts, agent)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentManager *AgentManagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentManager *AgentManagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentManager.Contract.AcceptOwnership(&_AgentManager.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentManager *AgentManagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentManager.Contract.AcceptOwnership(&_AgentManager.TransactOpts)
}

// ChangeAgentSettingProposal is a paid mutator transaction binding the contract method 0x69fd4d8e.
//
// Solidity: function changeAgentSettingProposal(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentManager *AgentManagerTransactor) ChangeAgentSettingProposal(opts *bind.TransactOpts, agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "changeAgentSettingProposal", agent, agentSettings)
}

// ChangeAgentSettingProposal is a paid mutator transaction binding the contract method 0x69fd4d8e.
//
// Solidity: function changeAgentSettingProposal(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentManager *AgentManagerSession) ChangeAgentSettingProposal(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentManager.Contract.ChangeAgentSettingProposal(&_AgentManager.TransactOpts, agent, agentSettings)
}

// ChangeAgentSettingProposal is a paid mutator transaction binding the contract method 0x69fd4d8e.
//
// Solidity: function changeAgentSettingProposal(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentManager *AgentManagerTransactorSession) ChangeAgentSettingProposal(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentManager.Contract.ChangeAgentSettingProposal(&_AgentManager.TransactOpts, agent, agentSettings)
}

// RegisterAgent is a paid mutator transaction binding the contract method 0x352f76e8.
//
// Solidity: function registerAgent(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentManager *AgentManagerTransactor) RegisterAgent(opts *bind.TransactOpts, agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "registerAgent", agent, agentSettings)
}

// RegisterAgent is a paid mutator transaction binding the contract method 0x352f76e8.
//
// Solidity: function registerAgent(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentManager *AgentManagerSession) RegisterAgent(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentManager.Contract.RegisterAgent(&_AgentManager.TransactOpts, agent, agentSettings)
}

// RegisterAgent is a paid mutator transaction binding the contract method 0x352f76e8.
//
// Solidity: function registerAgent(address agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings) returns()
func (_AgentManager *AgentManagerTransactorSession) RegisterAgent(agent common.Address, agentSettings CommonAgentSettings) (*types.Transaction, error) {
	return _AgentManager.Contract.RegisterAgent(&_AgentManager.TransactOpts, agent, agentSettings)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns()
func (_AgentManager *AgentManagerTransactor) RemoveAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "removeAgent", agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns()
func (_AgentManager *AgentManagerSession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.RemoveAgent(&_AgentManager.TransactOpts, agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns()
func (_AgentManager *AgentManagerTransactorSession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.RemoveAgent(&_AgentManager.TransactOpts, agent)
}

// SetAgentProxy is a paid mutator transaction binding the contract method 0xb78e6848.
//
// Solidity: function setAgentProxy(address proxy) returns()
func (_AgentManager *AgentManagerTransactor) SetAgentProxy(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "setAgentProxy", proxy)
}

// SetAgentProxy is a paid mutator transaction binding the contract method 0xb78e6848.
//
// Solidity: function setAgentProxy(address proxy) returns()
func (_AgentManager *AgentManagerSession) SetAgentProxy(proxy common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.SetAgentProxy(&_AgentManager.TransactOpts, proxy)
}

// SetAgentProxy is a paid mutator transaction binding the contract method 0xb78e6848.
//
// Solidity: function setAgentProxy(address proxy) returns()
func (_AgentManager *AgentManagerTransactorSession) SetAgentProxy(proxy common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.SetAgentProxy(&_AgentManager.TransactOpts, proxy)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AgentManager *AgentManagerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AgentManager *AgentManagerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.TransferOwnership(&_AgentManager.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AgentManager *AgentManagerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.TransferOwnership(&_AgentManager.TransactOpts, to)
}

// AgentManagerAgentAcceptedIterator is returned from FilterAgentAccepted and is used to iterate over the raw logs and unpacked data for AgentAccepted events raised by the AgentManager contract.
type AgentManagerAgentAcceptedIterator struct {
	Event *AgentManagerAgentAccepted // Event containing the contract specifics and raw log

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
func (it *AgentManagerAgentAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAgentAccepted)
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
		it.Event = new(AgentManagerAgentAccepted)
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
func (it *AgentManagerAgentAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAgentAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAgentAccepted represents a AgentAccepted event raised by the AgentManager contract.
type AgentManagerAgentAccepted struct {
	Agent         common.Address
	Digest        [32]byte
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentAccepted is a free log retrieval operation binding the contract event 0xa36e3cae6a00fea59cd678be00a81b5254290a4a713fd275c29d2e4225d5ee2d.
//
// Solidity: event AgentAccepted(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) FilterAgentAccepted(opts *bind.FilterOpts, agent []common.Address, digest [][32]byte) (*AgentManagerAgentAcceptedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AgentAccepted", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerAgentAcceptedIterator{contract: _AgentManager.contract, event: "AgentAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentAccepted is a free log subscription operation binding the contract event 0xa36e3cae6a00fea59cd678be00a81b5254290a4a713fd275c29d2e4225d5ee2d.
//
// Solidity: event AgentAccepted(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) WatchAgentAccepted(opts *bind.WatchOpts, sink chan<- *AgentManagerAgentAccepted, agent []common.Address, digest [][32]byte) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AgentAccepted", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAgentAccepted)
				if err := _AgentManager.contract.UnpackLog(event, "AgentAccepted", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseAgentAccepted(log types.Log) (*AgentManagerAgentAccepted, error) {
	event := new(AgentManagerAgentAccepted)
	if err := _AgentManager.contract.UnpackLog(event, "AgentAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerAgentProxySetIterator is returned from FilterAgentProxySet and is used to iterate over the raw logs and unpacked data for AgentProxySet events raised by the AgentManager contract.
type AgentManagerAgentProxySetIterator struct {
	Event *AgentManagerAgentProxySet // Event containing the contract specifics and raw log

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
func (it *AgentManagerAgentProxySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAgentProxySet)
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
		it.Event = new(AgentManagerAgentProxySet)
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
func (it *AgentManagerAgentProxySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAgentProxySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAgentProxySet represents a AgentProxySet event raised by the AgentManager contract.
type AgentManagerAgentProxySet struct {
	OldProxy common.Address
	NewProxy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentProxySet is a free log retrieval operation binding the contract event 0x027832fe8b1cfd252a8b83993255c855b482dd5caa679217465ff9106fe13cba.
//
// Solidity: event AgentProxySet(address oldProxy, address newProxy)
func (_AgentManager *AgentManagerFilterer) FilterAgentProxySet(opts *bind.FilterOpts) (*AgentManagerAgentProxySetIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AgentProxySet")
	if err != nil {
		return nil, err
	}
	return &AgentManagerAgentProxySetIterator{contract: _AgentManager.contract, event: "AgentProxySet", logs: logs, sub: sub}, nil
}

// WatchAgentProxySet is a free log subscription operation binding the contract event 0x027832fe8b1cfd252a8b83993255c855b482dd5caa679217465ff9106fe13cba.
//
// Solidity: event AgentProxySet(address oldProxy, address newProxy)
func (_AgentManager *AgentManagerFilterer) WatchAgentProxySet(opts *bind.WatchOpts, sink chan<- *AgentManagerAgentProxySet) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AgentProxySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAgentProxySet)
				if err := _AgentManager.contract.UnpackLog(event, "AgentProxySet", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseAgentProxySet(log types.Log) (*AgentManagerAgentProxySet, error) {
	event := new(AgentManagerAgentProxySet)
	if err := _AgentManager.contract.UnpackLog(event, "AgentProxySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerAgentRegisteredIterator is returned from FilterAgentRegistered and is used to iterate over the raw logs and unpacked data for AgentRegistered events raised by the AgentManager contract.
type AgentManagerAgentRegisteredIterator struct {
	Event *AgentManagerAgentRegistered // Event containing the contract specifics and raw log

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
func (it *AgentManagerAgentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAgentRegistered)
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
		it.Event = new(AgentManagerAgentRegistered)
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
func (it *AgentManagerAgentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAgentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAgentRegistered represents a AgentRegistered event raised by the AgentManager contract.
type AgentManagerAgentRegistered struct {
	Agent         common.Address
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentRegistered is a free log retrieval operation binding the contract event 0x8f9f5acc73d739a9bfdd8cd64d17b9e9bb10217585d472d17444462339bccaed.
//
// Solidity: event AgentRegistered(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) FilterAgentRegistered(opts *bind.FilterOpts, agent []common.Address) (*AgentManagerAgentRegisteredIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AgentRegistered", agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerAgentRegisteredIterator{contract: _AgentManager.contract, event: "AgentRegistered", logs: logs, sub: sub}, nil
}

// WatchAgentRegistered is a free log subscription operation binding the contract event 0x8f9f5acc73d739a9bfdd8cd64d17b9e9bb10217585d472d17444462339bccaed.
//
// Solidity: event AgentRegistered(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) WatchAgentRegistered(opts *bind.WatchOpts, sink chan<- *AgentManagerAgentRegistered, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AgentRegistered", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAgentRegistered)
				if err := _AgentManager.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseAgentRegistered(log types.Log) (*AgentManagerAgentRegistered, error) {
	event := new(AgentManagerAgentRegistered)
	if err := _AgentManager.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the AgentManager contract.
type AgentManagerAgentRemovedIterator struct {
	Event *AgentManagerAgentRemoved // Event containing the contract specifics and raw log

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
func (it *AgentManagerAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAgentRemoved)
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
		it.Event = new(AgentManagerAgentRemoved)
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
func (it *AgentManagerAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAgentRemoved represents a AgentRemoved event raised by the AgentManager contract.
type AgentManagerAgentRemoved struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed agent)
func (_AgentManager *AgentManagerFilterer) FilterAgentRemoved(opts *bind.FilterOpts, agent []common.Address) (*AgentManagerAgentRemovedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AgentRemoved", agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerAgentRemovedIterator{contract: _AgentManager.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed agent)
func (_AgentManager *AgentManagerFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *AgentManagerAgentRemoved, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AgentRemoved", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAgentRemoved)
				if err := _AgentManager.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseAgentRemoved(log types.Log) (*AgentManagerAgentRemoved, error) {
	event := new(AgentManagerAgentRemoved)
	if err := _AgentManager.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerAgentSettingsProposedIterator is returned from FilterAgentSettingsProposed and is used to iterate over the raw logs and unpacked data for AgentSettingsProposed events raised by the AgentManager contract.
type AgentManagerAgentSettingsProposedIterator struct {
	Event *AgentManagerAgentSettingsProposed // Event containing the contract specifics and raw log

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
func (it *AgentManagerAgentSettingsProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAgentSettingsProposed)
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
		it.Event = new(AgentManagerAgentSettingsProposed)
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
func (it *AgentManagerAgentSettingsProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAgentSettingsProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAgentSettingsProposed represents a AgentSettingsProposed event raised by the AgentManager contract.
type AgentManagerAgentSettingsProposed struct {
	Agent         common.Address
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentSettingsProposed is a free log retrieval operation binding the contract event 0x35f9d2cf50394e31e48ee780e6471bb9e0e3cceadc62d25d1349a94bdd7a9378.
//
// Solidity: event AgentSettingsProposed(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) FilterAgentSettingsProposed(opts *bind.FilterOpts, agent []common.Address) (*AgentManagerAgentSettingsProposedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AgentSettingsProposed", agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerAgentSettingsProposedIterator{contract: _AgentManager.contract, event: "AgentSettingsProposed", logs: logs, sub: sub}, nil
}

// WatchAgentSettingsProposed is a free log subscription operation binding the contract event 0x35f9d2cf50394e31e48ee780e6471bb9e0e3cceadc62d25d1349a94bdd7a9378.
//
// Solidity: event AgentSettingsProposed(address indexed agent, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) WatchAgentSettingsProposed(opts *bind.WatchOpts, sink chan<- *AgentManagerAgentSettingsProposed, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AgentSettingsProposed", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAgentSettingsProposed)
				if err := _AgentManager.contract.UnpackLog(event, "AgentSettingsProposed", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseAgentSettingsProposed(log types.Log) (*AgentManagerAgentSettingsProposed, error) {
	event := new(AgentManagerAgentSettingsProposed)
	if err := _AgentManager.contract.UnpackLog(event, "AgentSettingsProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerAgentSettingsUpdatedIterator is returned from FilterAgentSettingsUpdated and is used to iterate over the raw logs and unpacked data for AgentSettingsUpdated events raised by the AgentManager contract.
type AgentManagerAgentSettingsUpdatedIterator struct {
	Event *AgentManagerAgentSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *AgentManagerAgentSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAgentSettingsUpdated)
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
		it.Event = new(AgentManagerAgentSettingsUpdated)
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
func (it *AgentManagerAgentSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAgentSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAgentSettingsUpdated represents a AgentSettingsUpdated event raised by the AgentManager contract.
type AgentManagerAgentSettingsUpdated struct {
	Agent         common.Address
	Digest        [32]byte
	AgentSettings CommonAgentSettings
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentSettingsUpdated is a free log retrieval operation binding the contract event 0x9cbec0da5ddc48f6a581101583fb4487d38599559a940e8ebf031857881a79a1.
//
// Solidity: event AgentSettingsUpdated(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) FilterAgentSettingsUpdated(opts *bind.FilterOpts, agent []common.Address, digest [][32]byte) (*AgentManagerAgentSettingsUpdatedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AgentSettingsUpdated", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerAgentSettingsUpdatedIterator{contract: _AgentManager.contract, event: "AgentSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentSettingsUpdated is a free log subscription operation binding the contract event 0x9cbec0da5ddc48f6a581101583fb4487d38599559a940e8ebf031857881a79a1.
//
// Solidity: event AgentSettingsUpdated(address indexed agent, bytes32 indexed digest, (address[],uint8,address,(string,string,string,string,string,uint256,uint8,uint8,uint256)) agentSettings)
func (_AgentManager *AgentManagerFilterer) WatchAgentSettingsUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerAgentSettingsUpdated, agent []common.Address, digest [][32]byte) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AgentSettingsUpdated", agentRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAgentSettingsUpdated)
				if err := _AgentManager.contract.UnpackLog(event, "AgentSettingsUpdated", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseAgentSettingsUpdated(log types.Log) (*AgentManagerAgentSettingsUpdated, error) {
	event := new(AgentManagerAgentSettingsUpdated)
	if err := _AgentManager.contract.UnpackLog(event, "AgentSettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the AgentManager contract.
type AgentManagerOwnershipTransferRequestedIterator struct {
	Event *AgentManagerOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *AgentManagerOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerOwnershipTransferRequested)
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
		it.Event = new(AgentManagerOwnershipTransferRequested)
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
func (it *AgentManagerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the AgentManager contract.
type AgentManagerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AgentManager *AgentManagerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AgentManagerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerOwnershipTransferRequestedIterator{contract: _AgentManager.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AgentManager *AgentManagerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AgentManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerOwnershipTransferRequested)
				if err := _AgentManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseOwnershipTransferRequested(log types.Log) (*AgentManagerOwnershipTransferRequested, error) {
	event := new(AgentManagerOwnershipTransferRequested)
	if err := _AgentManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentManager contract.
type AgentManagerOwnershipTransferredIterator struct {
	Event *AgentManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AgentManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerOwnershipTransferred)
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
		it.Event = new(AgentManagerOwnershipTransferred)
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
func (it *AgentManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AgentManager contract.
type AgentManagerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AgentManager *AgentManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AgentManagerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerOwnershipTransferredIterator{contract: _AgentManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AgentManager *AgentManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerOwnershipTransferred)
				if err := _AgentManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AgentManager *AgentManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AgentManagerOwnershipTransferred, error) {
	event := new(AgentManagerOwnershipTransferred)
	if err := _AgentManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
