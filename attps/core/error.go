package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// APIError standard error structure
type APIError struct {
	StatusCode int         // HTTP status code of the response message
	Header     http.Header // Header information of the response message
	Body       string      // Original body of the response message
	Code       string      `json:"code"`             // The error code information after the body of the response message is parsed. It only exists when it does not meet expectations or a system error occurs.
	Message    string      `json:"message"`          // The text description of the body of the response message after parsing. It only exists when it does not meet expectations or a system error occurs.
	Detail     interface{} `json:"detail,omitempty"` // Detailed information after the body of the response message is parsed. It only exists when it does not meet expectations or a system error occurs.
}

// Error Output APIError
func (e *APIError) Error() string {
	var buf bytes.Buffer
	_, _ = fmt.Fprintf(&buf, "error http response:[StatusCode: %d Code: \"%s\"", e.StatusCode, e.Code)
	if e.Message != "" {
		_, _ = fmt.Fprintf(&buf, "\nMessage: %s", e.Message)
	}
	if e.Detail != nil {
		var detailBuf bytes.Buffer
		enc := json.NewEncoder(&detailBuf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(e.Detail); err == nil {
			_, _ = fmt.Fprint(&buf, "\nDetail:")
			_, _ = fmt.Fprintf(&buf, "\n%s", strings.TrimSpace(detailBuf.String()))
		}
	}
	if len(e.Header) > 0 {
		_, _ = fmt.Fprint(&buf, "\nHeader:")
		for key, value := range e.Header {
			_, _ = fmt.Fprintf(&buf, "\n - %v=%v", key, value)
		}
	}
	_, _ = fmt.Fprintf(&buf, "]")
	return buf.String()
}

// IsAPIError determines whether the current error is a *APIError with a specific code
//
// Returns false if the error type is other than *APIError or the code does not match
func IsAPIError(err error, code string) bool {
	if ne, ok := err.(*APIError); ok {
		return ne.Code == code
	}
	return false
}
