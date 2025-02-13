package vrf

import (
	"github.com/APRO-com/ai-agent-sdk-go/attps/service"
)

// VRFRequest Random Request
type VRFRequest struct {
	Version          *int64  `json:"version"`
	TargetAgentID    *string `json:"target_agent_id"`
	ClientSeed       *string `json:"client_seed"`
	KeyHash          *string `json:"key_hash"`
	RequestTimestamp *int64  `json:"request_timestamp"`
	RequestID        *string `json:"request_id"`
	CallbackURI      *string `json:"callback_uri"`
}

// VRFResponse Random Response
type VRFResponse struct {
	service.BaseResponse
	Result *string `json:"result"`
}

// Provider Provider
type Provider struct {
	Address *string `json:"address"`
	KeyHash *string `json:"keyHash"`
}

type ProviderResponse struct {
	service.BaseResponse
	Result *[]Provider `json:"result"`
}

type Proof struct {
	PublicX *string `json:"publicX"`
	PublicY *string `json:"publicY"`
	GammaX  *string `json:"gammaX"`
	GammaY  *string `json:"gammaY"`
	C       *string `json:"c"`
	S       *string `json:"s"`
	Seed    *string `json:"seed"`
	Output  *string `json:"output"`
}

// VRFProof VRF Proof
type VRFProof struct {
	RequestId *string `json:"requestId"`
	Proof     *Proof  `json:"proof"`
}

type VRFProofResponse struct {
	service.BaseResponse
	Result *VRFProof `json:"result"`
}
