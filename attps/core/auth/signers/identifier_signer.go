// Package signers Digital Signature Generator
package signers

import (
	"context"
	"fmt"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/auth"
)

// IdentifierSigner Identifier digital signature generator
type IdentifierSigner struct {
	Identifier string
}

func (s *IdentifierSigner) Sign(_ context.Context, message string, timestamp int64) (*auth.SignatureResult, error) {
	if s.Identifier == "" {
		return nil, fmt.Errorf("you must set secretKey to use IdentifierSigner")
	}
	return &auth.SignatureResult{AgentId: "", Signature: s.Identifier}, nil
}

func (s *IdentifierSigner) Algorithm() string {
	return "Identifier"
}
