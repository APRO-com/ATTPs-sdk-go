// Package auth Go SDK Security verification related interfaces
package auth

import "context"

// SignatureResult Digital signature result
type SignatureResult struct {
	AgentId   string // Agent Id
	Signature string // Signature content
}

// Signer Digital Signature Generator
type Signer interface {
	// Sign the message
	Sign(ctx context.Context, message string, timestamp int64) (*SignatureResult, error)
	// Algorithm Returns the signature algorithm used
	Algorithm() string
}
