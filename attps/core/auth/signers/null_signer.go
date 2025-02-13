// Package signers
package signers

import (
	"context"
	"github.com/APRO-com/ai-agent-sdk-go/attps/core/auth"
)

// NULLSigner NULL digital signature generator
type NULLSigner struct{}

func (s *NULLSigner) Sign(_ context.Context, message string, timestamp int64) (*auth.SignatureResult, error) {
	return &auth.SignatureResult{AgentId: "", Signature: ""}, nil
}

func (s *NULLSigner) Algorithm() string {
	return "Null"
}
