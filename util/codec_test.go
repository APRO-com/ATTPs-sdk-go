package util

import (
	"encoding/hex"
	"testing"
)

func TestStringToKeccak256(t *testing.T) {
	tests := []struct {
		input    string
		expected string // Keccak256 hash as hex string
	}{
		{
			input:    "hello world",
			expected: "47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := StringToKeccak256(test.input)
			resultHex := hex.EncodeToString(result)

			if resultHex != test.expected {
				t.Errorf("For input '%s', expected '%s', got '%s'", test.input, test.expected, resultHex)
			}
		})
	}
}

func TestHexStringToKeccak256(t *testing.T) {
	tests := []struct {
		name      string
		hexString string
		expected  string
		expectErr bool
	}{
		{
			name:      "hello world",
			hexString: "0x68656c6c6f20776f726c64",
			expected:  "47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad",
			expectErr: false,
		},
		{
			name:      "hello world",
			hexString: "68656c6c6f20776f726c64",
			expected:  "47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad",
			expectErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := HexStringToKeccak256(test.hexString)

			if (err != nil) != test.expectErr {
				t.Fatalf("expected error: %v, got: %v", test.expectErr, err)
			}

			if err == nil && result != test.expected {
				t.Errorf("for input '%s', expected hash '%s', but got '%s'", test.hexString, test.expected, result)
			}
		})
	}
}
