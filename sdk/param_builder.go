package sdk

import (
	"ai-agent-go-sdk/contract/agent_proxy"
	"ai-agent-go-sdk/util"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"math/big"
)

func BuildRegisterAgentData(
	signers []string,
	threshold uint8,
	converterAddress string,
	version string,
	messageId, sourceAgentId, sourceAgentName, targetAgentId string,
	timestamp *big.Int,
	messageType uint8,
	priority uint8,
	ttl *big.Int,
) (agent_proxy.CommonAgentSettings, error) {
	if len(signers) < 1 {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("signers number should greater than 1")
	}
	if int(threshold) > len(signers) {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("threshold should not greater than signer length")
	}
	if !common.IsHexAddress(converterAddress) {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("converterAddress is invalid")
	}
	if messageId == "" {
		messageId = uuid.New().String()
	}
	if sourceAgentId == "" {
		sourceAgentId = uuid.New().String()
	}
	if targetAgentId == "" {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("targetAgentId should not be nil")
	}
	if !util.IsUUIDv4(messageId) || !util.IsUUIDv4(sourceAgentId) || !util.IsUUIDv4(targetAgentId) {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("messageId|sourceAgentId|targetAgentId should be valid UUID")
	}
	if util.IsValid13DigitTimestamp(timestamp) {
		timestamp = new(big.Int).Div(timestamp, big.NewInt(1000))
	}
	if messageType > 2 || messageType < 0 {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("invalild messageType")
	}
	if priority > 2 || priority < 0 {
		return agent_proxy.CommonAgentSettings{}, fmt.Errorf("invalild priority")
	}

	signersAddr := make([]common.Address, len(signers))
	for _, signer := range signers {
		signersAddr = append(signersAddr, common.HexToAddress(signer))
	}

	return agent_proxy.CommonAgentSettings{
		Signers:          signersAddr,
		Threshold:        threshold,
		ConverterAddress: common.HexToAddress(converterAddress),
		AgentHeader: agent_proxy.CommonAgentHeader{
			Version:         version,
			MessageId:       messageId,
			SourceAgentId:   sourceAgentId,
			SourceAgentName: sourceAgentName,
			TargetAgentId:   targetAgentId,
			Timestamp:       timestamp,
			MessageType:     messageType,
			Priority:        priority,
			Ttl:             ttl,
		},
	}, nil
}

func BuildVerifyData(
	data,
	dataHash string,
	signatures []byte,
	metadata agent_proxy.CommonMetadata) (agent_proxy.CommonMessagePayload, error) {
	dataHashByte, err := util.HexString2Byte32(dataHash)
	if err != nil {
		return agent_proxy.CommonMessagePayload{}, err
	}
	return agent_proxy.CommonMessagePayload{
		Data:     common.Hex2Bytes(data),
		DataHash: dataHashByte,
		Proofs: agent_proxy.CommonProofs{
			ZkProof:        []byte{},
			MerkleProof:    []byte{},
			SignatureProof: signatures,
		},
		Metadata: metadata,
	}, nil
}
