package sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/APRO-com/ai-agent-sdk-go/config"
	"github.com/APRO-com/ai-agent-sdk-go/contract/agent_proxy"
	"github.com/APRO-com/ai-agent-sdk-go/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err, "create client failed")

	opts, err := util.DefaultTransactOpts(aiAgentClient.Client, privateKeyHex, aiAgentClient.ChainID, big.NewInt(5000000000), 0)
	assert.NoError(t, err, "init transactOpt failed")

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
	assert.NoError(t, err, "BuildRegisterAgentData failed")

	tx, agent, err := aiAgentClient.RegisterAgent(proxyAddress, opts, agentSettings)
	assert.NoError(t, err, "RegisterAgent failed")

	assert.NotNil(t, tx, "tx is nil")
	assert.NotEmpty(t, tx.Hash().String(), "txHash is nil")

	fmt.Println("txHash:", tx.Hash().String())
	fmt.Println("agent address:", agent)
}

func Test_verify(t *testing.T) {
	agent := "0x96532A01295b81268F201dD3EAB4A281543a8547"
	agentDigest := "0x0100238D5B38AEE38000F1C928B8F291E79F8CB243410E3B617800DFB96007A8"

	// Create client
	aiAgentClient, err := NewClient(config.BSC_TEST_NET)
	assert.NoError(t, err, "Failed to create client")

	// Generate data hash
	data := hex.EncodeToString([]byte("hello world"))
	dataHash, err := util.HexStringToKeccak256(data)
	assert.NoError(t, err, "Failed to generate data hash")
	dataHashBytes, err := hex.DecodeString(dataHash)
	assert.NoError(t, err, "Failed to parse data hash")

	var signatures [][]byte
	for _, prikey := range signerPrikeys {
		prikeyObj, err := util.HexToECDSA(prikey)
		if err != nil {
			assert.NoError(t, err, "Failed to init private key")
		}
		signature, err := crypto.Sign(dataHashBytes, prikeyObj)
		if err != nil {
			assert.NoError(t, err, "Failed to generate signatures")
		}
		signatures = append(signatures, signature)
	}

	// Encode signatures
	signBytes, err := util.EncodeSignatures(signatures)
	assert.NoError(t, err, "Failed to encode signatures")

	// Build verification data
	verifyData, err := BuildVerifyData(data, dataHash, signBytes, agent_proxy.CommonMetadata{
		ContentType: "",
		Encoding:    "",
		Compression: "",
	})
	assert.NoError(t, err, "Failed to build verification data")

	// Set transaction options
	opts, err := util.DefaultTransactOpts(aiAgentClient.Client, privateKeyHex, aiAgentClient.ChainID, big.NewInt(5000000000), 0)
	assert.NoError(t, err, "Failed to get transaction options")

	// Call blockchain verify method
	tx, err := aiAgentClient.Verify(proxyAddress, opts, agent, agentDigest, verifyData)
	assert.NoError(t, err, "Blockchain verification failed")

	// Output transaction hash
	fmt.Println("Transaction Hash:", tx.Hash().String())
}

func TestClient_Converter(t *testing.T) {
	converterAddr := "0x24c36e9996eb84138Ed7cAa483B4c59FF7640E5C"
	data := "0x0006e706cf7ab41fa599311eb3de68be869198ce62aef1cd079475ca50e5b3f60000000000000000000000000000000000000000000000000000000002b1bf0e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002a0000101000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001200003665949c883f9e0f6f002eac32e00bd59dfe6c34e92a91c37d6a8322d6489000000000000000000000000000000000000000000000000000000006762677d000000000000000000000000000000000000000000000000000000006762677d000000000000000000000000000000000000000000000000000003128629ec0800000000000000000000000000000000000000000000000004db732547630000000000000000000000000000000000000000000000000000000000006763b8fd0000000000000000000000000000000000000000000015f0f60671beb95cc0000000000000000000000000000000000000000000000015f083baa654a7b900000000000000000000000000000000000000000000000015f103ec7cb057ea80000000000000000000000000000000000000000000000000000000000000000003b64f7e72208147bb898e8b215d0997967bef0219263726c76995d8a19107d6ba5306a176474f9ccdb1bc5841f97e0592013e404e15b0de0839b81d0efb26179f222e0191269a8560ebd9096707d225bc606d61466b85d8568d7620a3b59a73e800000000000000000000000000000000000000000000000000000000000000037cae0f05c1bf8353eb5db27635f02b40a534d4192099de445764891198231c597a303cd15f302dafbb1263eb6e8e19cbacea985c66c6fed3231fd84a84ebe0276f69f481fe7808c339a04ceb905bb49980846c8ceb89a27b1c09713cb356f773"
	expectConverterData := "257529fe335f30941728adc69975a20af0eb5419fba583cdfbede8bba5a6f8ac0006e706cf7ab41fa599311eb3de68be869198ce62aef1cd079475ca50e5b3f60000000000000000000000000000000000000000000000000000000002b1bf0e0000000000000000000000000000000000000000000000000000000000000000"
	// Create a new client
	aiAgentClient, err := NewClient(config.BSC_TEST_NET)
	assert.NoError(t, err, "Failed to create aiAgentClient client")
	converterData, err := aiAgentClient.Converter(converterAddr, data)
	assert.NoError(t, err, "Failed to call converter")
	assert.Equal(t, expectConverterData, converterData)
}

func TestClient_IsValidateAgentID(t *testing.T) {
	// Create a new client
	aiAgentClient, err := NewClient(config.BSC_TEST_NET)
	assert.NoError(t, err, "Failed to create aiAgentClient client")

	// Get the manager for the proxy address
	//isValid, err := aiAgentClient.isValidSourceAgentId(proxyAddress, "087d5233-3256-43c5-8ce2-0ad705357c4c")
	isValid, err := aiAgentClient.isValidSourceAgentId(proxyAddress, "0ccdc715-5386-4483-8a6b-8b7e319f7b0f")
	assert.NoError(t, err, "Failed to call isValidSourceAgentId for proxy address")
	assert.NotNil(t, isValid, "isValid should not be nil")
	fmt.Println("isValid:", isValid)
}

func TestClient_GetVersion(t *testing.T) {
	// Create a new client
	aiAgentClient, err := NewClient(config.BSC_TEST_NET)
	assert.NoError(t, err, "Failed to create aiAgentClient client")

	// Get the manager for the proxy address
	manager, err := aiAgentClient.GetManager(proxyAddress)
	assert.NoError(t, err, "Failed to get manager for proxy address")
	assert.NotNil(t, manager, "Manager should not be nil")
	fmt.Println("Manager:", manager)

	// Get the version for the proxy address
	version, err := aiAgentClient.GetVersion(proxyAddress)
	assert.NoError(t, err, "Failed to get version for proxy address")
	assert.NotEmpty(t, version, "Version should not be empty")
	fmt.Println("Version:", version)
}
