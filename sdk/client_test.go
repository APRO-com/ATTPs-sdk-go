package sdk

import (
	"ai-agent-go-sdk/config"
	"ai-agent-go-sdk/contract/agent_proxy"
	"ai-agent-go-sdk/util"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
	"time"
)

var (
	proxyAddress = "0x5E787A4131Cf9fC902C99235df5C8314C816DA11"
	signers      = []string{
		"0x17344AF2aB888dBCa1bd2c3E3d2d5E4F771FEf58",
		"0x4173964DAAaB016b5F94833f983B9090a4B2396C",
		"0x6824086c0d6daeD91acd2e9EE960c6463210d0f0",
	}

	privateKeyHex = ""
	signerPrikeys = []string{
		"",
		"",
		"",
	}
)

func TestClient_RegisterAgent(t *testing.T) {
	aiAgentClient, err := NewClient(config.BSC_TEST_NET)
	if err != nil {
		panic(err)
	}
	opts, err := util.DefaultTransactOpts(aiAgentClient.Client, privateKeyHex, aiAgentClient.ChainID, big.NewInt(5000000000), 0)
	if err != nil {
		panic(err)
	}
	agentSettings, err := BuildRegisterAgentData(
		signers,
		2,
		config.ZERO_ADDRESS,
		"1.0",
		util.NewUUIDv4(),
		util.NewUUIDv4(),
		"go sdk test",
		util.NewUUIDv4(),
		big.NewInt(time.Now().Unix()),
		1,
		1,
		big.NewInt(3600),
	)

	if err != nil {
		panic(err)
	}

	tx, err := aiAgentClient.RegisterAgent(proxyAddress, opts, agentSettings)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Hash().String())
}

func Test_verify(t *testing.T) {
	agent := "0x96532A01295b81268F201dD3EAB4A281543a8547"
	agentDigest := "0x0100238D5B38AEE38000F1C928B8F291E79F8CB243410E3B617800DFB96007A8"

	aiAgentClient, err := NewClient(config.BSC_TEST_NET)
	if err != nil {
		panic(err)
	}
	data := hex.EncodeToString([]byte("hello world"))
	dataHash, err := util.HexStringToKeccak256(data)
	if err != nil {
		panic(err)
	}
	dataHashBytes, err := util.HexString2Byte32(dataHash)
	if err != nil {
		panic(err)
	}

	signatures := [][]byte{}
	for _, prikey := range signerPrikeys {
		prikeyObj, err := crypto.HexToECDSA(prikey)
		if err != nil {
			panic(err)
		}
		signature, err := crypto.Sign(dataHashBytes[:], prikeyObj)
		if err != nil {
			panic(err)
		}
		signatures = append(signatures, signature)
	}

	signBytes, err := util.EncodeSignatures(signatures)

	if err != nil {
		panic(err)
	}

	verifyData, err := BuildVerifyData(data, dataHash, signBytes, agent_proxy.CommonMetadata{
		ContentType: "",
		Encoding:    "",
		Compression: "",
	})
	if err != nil {
		panic(err)
	}
	opts, err := util.DefaultTransactOpts(aiAgentClient.Client, privateKeyHex, aiAgentClient.ChainID, big.NewInt(5000000000), 10000000)
	if err != nil {
		panic(err)
	}

	tx, err := aiAgentClient.Verify(proxyAddress, opts, agent, agentDigest, verifyData)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Hash().String())

}

func TestClient_GetVersion(t *testing.T) {
	sdk, err := NewClient(config.BSC_TEST_NET)
	if err != nil {
		panic(err)
	}
	manager, err := sdk.GetManager(proxyAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(manager)

	version, err := sdk.GetVersion(proxyAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(version)
}
