package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func HexToECDSA(prikey string) (*ecdsa.PrivateKey, error) {
	if len(prikey) > 1 && prikey[:2] == "0x" {
		prikey = prikey[2:]
	}
	return crypto.HexToECDSA(prikey)
}
