package util

import (
	"github.com/google/uuid"
	"math/big"
)

func IsUUIDv4(input string) bool {
	parsedUUID, err := uuid.Parse(input)
	if err != nil {
		return false
	}
	return parsedUUID.Version() == 4
}

func NewUUIDv4() string {
	return uuid.New().String()
}

func IsValid13DigitTimestamp(timestamp *big.Int) bool {
	minTimestamps := big.NewInt(1000000000000)
	maxTimestamps := big.NewInt(9999999999999)

	if timestamp.Cmp(minTimestamps) < 0 || timestamp.Cmp(maxTimestamps) > 0 {
		return false
	}
	return true
}
