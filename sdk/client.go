package sdk

import (
	"ai-agent-go-sdk/config"
	"ai-agent-go-sdk/contract/agent_manager"
	"ai-agent-go-sdk/contract/agent_proxy"
	"ai-agent-go-sdk/contract/converter"
	"ai-agent-go-sdk/util"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type AiAgentClient struct {
	Client  *ethclient.Client
	Rpc     string
	ChainID *big.Int
}

func NewClient(chainInfo config.ChainInfo) (*AiAgentClient, error) {
	var client = AiAgentClient{
		Rpc:     chainInfo.Url,
		ChainID: chainInfo.ChainID,
	}
	ethClient, err := ethclient.Dial(client.Rpc)
	if err != nil {
		return nil, err
	}
	client.Client = ethClient

	return &client, nil
}

func (aiAgentClient *AiAgentClient) RegisterAgent(
	proxyAddr string,
	opts *bind.TransactOpts,
	agentSettings agent_proxy.CommonAgentSettings) (tx *types.Transaction, agentAddr string, err error) {
	version, err := aiAgentClient.GetVersion(proxyAddr)
	if err != nil {
		return tx, "", err
	}
	if agentSettings.AgentHeader.Version == "" {
		agentSettings.AgentHeader.Version = version
	} else if agentSettings.AgentHeader.Version != version {
		return nil, "", fmt.Errorf("agentSettings version is not matched with agentProxy")
	}
	isValid, err := aiAgentClient.isValidSourceAgentId(proxyAddr, agentSettings.AgentHeader.SourceAgentId)
	if err != nil {
		return tx, "", err
	}
	if !isValid {
		return nil, "", fmt.Errorf("agentSettings.SourceAgentId is not valid")
	}
	agentProxy, err := agent_proxy.NewAgentProxy(common.HexToAddress(proxyAddr), aiAgentClient.Client)
	tx, err = agentProxy.CreateAndRegisterAgent(opts, agentSettings)
	if err != nil {
		return tx, "", err
	}

	txReceipt, err := util.WaitForTransactionReceipt(aiAgentClient.Client, tx.Hash(), 10, time.Second*6)
	if err != nil {
		return tx, "", fmt.Errorf("failed to fetch transaction receipt: %v", err)
	}

	for _, log := range txReceipt.Logs {
		if log.Topics[0].Hex() == crypto.Keccak256Hash([]byte(config.AgentCreatedEvent)).Hex() {
			agentAddr = common.HexToAddress(log.Topics[2].Hex()).Hex()
		}
	}
	return tx, agentAddr, nil
}

func (aiAgentClient *AiAgentClient) Verify(
	proxyAddr string,
	opts *bind.TransactOpts,
	agent string,
	agentDigest string,
	messagPayload agent_proxy.CommonMessagePayload) (*types.Transaction, error) {
	agentProxy, err := agent_proxy.NewAgentProxy(common.HexToAddress(proxyAddr), aiAgentClient.Client)
	if err != nil {
		return nil, err
	}
	digest, err := util.HexString2Byte32(agentDigest)
	if err != nil {
		return nil, err
	}
	tx, err := agentProxy.Verify(opts, common.HexToAddress(agent), digest, messagPayload)

	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (aiAgentClient *AiAgentClient) Converter(converterAddr, data string) (string, error) {
	newConverter, err := converter.NewConverter(common.HexToAddress(converterAddr), aiAgentClient.Client)
	if err != nil {
		return "", err
	}
	newData, err := newConverter.Converter(nil, common.FromHex(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(newData), nil
}

func (aiAgentClient *AiAgentClient) isValidSourceAgentId(proxy, agentId string) (bool, error) {
	managerAddr, err := aiAgentClient.GetManager(proxy)
	if err != nil {
		return false, err
	}
	agentManager, err := agent_manager.NewAgentManager(common.HexToAddress(managerAddr), aiAgentClient.Client)
	if err != nil {
		return false, err
	}
	return agentManager.IsValidSourceAgentId(nil, agentId)
}

func (aiAgentClient *AiAgentClient) GetManager(proxy string) (string, error) {
	agentProxy, err := agent_proxy.NewAgentProxy(common.HexToAddress(proxy), aiAgentClient.Client)
	if err != nil {
		return "", err
	}
	manager, err := agentProxy.AgentManager(nil)
	if err != nil {
		return "", err
	}
	return manager.String(), nil
}

func (aiAgentClient *AiAgentClient) GetVersion(proxy string) (string, error) {
	managerAddr, err := aiAgentClient.GetManager(proxy)
	if err != nil {
		return "", err
	}
	agentManager, err := agent_manager.NewAgentManager(common.HexToAddress(managerAddr), aiAgentClient.Client)
	if err != nil {
		return "", err
	}
	version, err := agentManager.AgentVersion(nil)
	if err != nil {
		return "", err
	}
	return version, nil
}
