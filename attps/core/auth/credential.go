// Package auth Go SDK Security verification related interfaces
package auth

import "context"

// Credential Request header Authorization information generator
type Credential interface {
	GenerateAuthorization(ctx context.Context, signBody string) (string, error)
}
