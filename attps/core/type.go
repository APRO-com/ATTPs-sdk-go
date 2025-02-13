package core

import (
	"encoding/hex"
	"time"
)

// Time copies the time.Time object and returns a pointer to the copy
func Time(t time.Time) *time.Time {
	return &t
}

// String copies the string object and returns a pointer to the copy
func String(s string) *string {
	return &s
}

// String copies the string object and returns a pointer to the copy
func HexBytes(s string) *[]byte {
	b, _ := hex.DecodeString(s)
	return &b
}

// Bool copies the bool object and returns a pointer to the copy
func Bool(b bool) *bool {
	return &b
}

// Float64 copies the float64 object and returns a pointer to the copy
func Float64(f float64) *float64 {
	return &f
}

// Float32 copies the float32 object and returns a pointer to the copy
func Float32(f float32) *float32 {
	return &f
}

// Int64 copies the int64 object and returns a pointer to the copy
func Int64(i int64) *int64 {
	return &i
}

// Int32 copies the int64 object and returns a pointer to the copy
func Int32(i int32) *int32 {
	return &i
}
