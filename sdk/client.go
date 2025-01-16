package sdk

import (
	"ai-agent-go-sdk/config"
	"ai-agent-go-sdk/contract/agent_manager"
	"ai-agent-go-sdk/contract/agent_proxy"
	"ai-agent-go-sdk/util"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
	agentSettings agent_proxy.CommonAgentSettings) (*types.Transaction, error) {
	version, err := aiAgentClient.GetVersion(proxyAddr)
	if err != nil {
		return nil, err
	}
	if agentSettings.AgentHeader.Version == "" {
		agentSettings.AgentHeader.Version = version
	} else if agentSettings.AgentHeader.Version != version {
		return nil, fmt.Errorf("agentSettings version is not matched with agentProxy")
	}
	agentProxy, err := agent_proxy.NewAgentProxy(common.HexToAddress(proxyAddr), aiAgentClient.Client)
	tx, err := agentProxy.CreateAndRegisterAgent(opts, agentSettings)
	if err != nil {
		return nil, err
	}
	return tx, nil
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
	agentManager, err := agent_manager.NewAgentProxy(common.HexToAddress(managerAddr), aiAgentClient.Client)
	if err != nil {
		return "", err
	}
	version, err := agentManager.AgentVersion(nil)
	if err != nil {
		return "", err
	}
	return version, nil
}
