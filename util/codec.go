package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"net/url"
	"regexp"
)

func HexString2Byte32(hexString string) ([32]byte, error) {
	if len(hexString) > 2 && hexString[:2] == "0x" {
		hexString = hexString[2:]
	}
	var fixedArray [32]byte
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return fixedArray, err
	}
	if len(bytes) != 32 {
		return fixedArray, err
	}
	copy(fixedArray[:], bytes)
	return fixedArray, nil
}

func StringToKeccak256(input string) []byte {
	return crypto.Keccak256([]byte(input))
}

func HexStringToKeccak256(hexString string) (string, error) {
	// Check if the string starts with "0x" and remove it
	if len(hexString) > 1 && hexString[:2] == "0x" {
		hexString = hexString[2:]
	}
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return "", fmt.Errorf("invalid hex string: %v", err)
	}

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash), nil
}

type SignatureData struct {
	R [32]byte
	S [32]byte
	V uint8
}

func EncodeSignatureString(signatures []string) ([]byte, error) {
	var signatureList [][]byte
	for _, signature := range signatures {
		signByte := common.FromHex(signature)
		signatureList = append(signatureList, signByte)
	}
	return EncodeSignatures(signatureList)
}

func EncodeSignatures(signatures [][]byte) ([]byte, error) {
	var signatureList []SignatureData
	for _, signature := range signatures {
		if len(signature) != 65 {
			return nil, fmt.Errorf("invalid signature")
		}
		var sigData SignatureData
		copy(sigData.R[:], signature[:32])
		copy(sigData.S[:], signature[32:64])
		sigData.V = signature[64]
		signatureList = append(signatureList, sigData)
	}
	return EncodeSignatureDatas(signatureList)
}

func EncodeSignatureDatas(signatures []SignatureData) ([]byte, error) {
	// Convert signatures to separate slices
	var rsArray []common.Hash
	var ssArray []common.Hash
	var vArray []uint8

	for _, sig := range signatures {
		rsArray = append(rsArray, common.BytesToHash(sig.R[:]))
		ssArray = append(ssArray, common.BytesToHash(sig.S[:]))
		vArray = append(vArray, sig.V) // V needs to subtract 27
	}

	// Define the ABI types
	rsType, err := abi.NewType("bytes32[]", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsType: %w", err)
	}
	ssType, err := abi.NewType("bytes32[]", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ssType: %w", err)
	}
	vType, err := abi.NewType("uint8[]", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create vType: %w", err)
	}

	// Create the ABI arguments definition
	arguments := abi.Arguments{
		{Type: rsType},
		{Type: ssType},
		{Type: vType},
	}

	// Pack the parameters into ABI-encoded bytes
	encoded, err := arguments.Pack(rsArray, ssArray, vArray)
	if err != nil {
		return nil, fmt.Errorf("failed to pack arguments: %w", err)
	}

	return encoded, nil
}

// LongToBytes Convert long (int64) to byte array
func LongToBytes(n int64) []byte {
	b := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		b[7-i] = byte(n >> (i * 8))
	}
	return b
}

// IsHexString Use regular expressions to determine if a string is a hexadecimal string
func IsHexString(s string) bool {
	match, _ := regexp.MatchString("^[0-9a-fA-F]+$", s)
	return match
}

func SecureRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length/2)
	randomBytes := make([]byte, length/2)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	for i := range result {
		result[i] = charset[randomBytes[i]%byte(len(charset))]
	}
	return hex.EncodeToString(result)
}

// IsValidHTTPBaseURL checks if the input is a valid HTTP(S) base URL
func IsValidHTTPBaseURL(input string) bool {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return false
	}

	// Must have a valid scheme (http or https)
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Must have a valid Host
	if parsedURL.Host == "" {
		return false
	}

	return true
}
