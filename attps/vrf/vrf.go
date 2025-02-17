package vrf

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/consts"
	"github.com/APRO-com/ATTPs-sdk-go/attps/service"
	"github.com/APRO-com/ATTPs-sdk-go/attps/vrf/proof"
	"github.com/APRO-com/ATTPs-sdk-go/attps/vrf/proof/secp256k1"
	"github.com/APRO-com/ATTPs-sdk-go/util"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

type VRF service.Service

// Provider get a vrf random provider list
func (v *VRF) Provider(ctx context.Context) (providers []Provider, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.APIBaseServer + "/api/vrf/provider"

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err := v.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, err
	}

	resp := new(ProviderResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, err
	}

	if resp.BaseResponse.Code != 0 {
		return nil, errors.New(*resp.Message)
	}

	return *resp.Result, nil
}

// Request get a vrf random number and verify the returned proof data
func (v *VRF) Request(ctx context.Context, req VRFRequest) (random string, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.APIBaseServer + "/api/vrf/request"

	// Setup Body Params and check
	localVarPostBody = &req

	if err = v.checkRequestParams(req); err != nil {
		return "", err
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err := v.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return "", err
	}

	resp := new(VRFResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return "", err
	}

	if resp.BaseResponse.Code != 0 {
		return "", errors.New(*resp.Message)
	}

	return *resp.Result, nil
}

// QueryProof get a vrf random proof by request id
func (v *VRF) QueryProof(ctx context.Context, requestID string) (proof *VRFProof, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.APIBaseServer + "/api/vrf/query"

	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("request_id", requestID)

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err := v.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, err
	}

	resp := new(VRFProofResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, err
	}

	if resp.BaseResponse.Code != 0 {
		return nil, errors.New(*resp.Message)
	}

	// check proof
	status, err := v.VerifyVRFProof(resp.Result)
	if err != nil || status != true {
		return resp.Result, err
	}

	return resp.Result, nil
}

func (v *VRF) CalRequestID(version int64, targetAgentId string, customerSeed string, requestTimestamp int64, callbackUri string) (string, error) {
	// Concatenate byte arrays
	t := append(util.LongToBytes(version), []byte(targetAgentId)...)
	t = append(t, *core.HexBytes(customerSeed)...)
	t = append(t, util.LongToBytes(requestTimestamp)...)
	t = append(t, []byte(callbackUri)...)

	// Compute Keccak256 hash
	hash := crypto.Keccak256(t)
	requestId := hex.EncodeToString(hash[:])

	return requestId, nil
}

func (v *VRF) VerifyVRFProof(vpr *VRFProof) (status bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	pubXBig := new(big.Int)
	pubYBig := new(big.Int)
	gammaXBig := new(big.Int)
	gammaYBig := new(big.Int)

	cBig := new(big.Int)
	sBig := new(big.Int)
	seedBig := new(big.Int)
	outputBig := new(big.Int)

	pubXBig.SetString(strings.TrimPrefix(*vpr.Proof.PublicX, "0x"), 16)
	pubYBig.SetString(strings.TrimPrefix(*vpr.Proof.PublicY, "0x"), 16)

	gammaXBig.SetString(strings.TrimPrefix(*vpr.Proof.GammaX, "0x"), 16)
	gammaYBig.SetString(strings.TrimPrefix(*vpr.Proof.GammaY, "0x"), 16)

	cBig.SetString(strings.TrimPrefix(*vpr.Proof.C, "0x"), 16)
	sBig.SetString(strings.TrimPrefix(*vpr.Proof.S, "0x"), 16)
	seedBig.SetString(strings.TrimPrefix(*vpr.Proof.Seed, "0x"), 16)
	outputBig.SetString(strings.TrimPrefix(*vpr.Proof.Output, "0x"), 16)

	public := secp256k1.SetCoordinates(pubXBig, pubYBig)
	gamma := secp256k1.SetCoordinates(gammaXBig, gammaYBig)

	pf := proof.Proof{
		PublicKey: public,
		Gamma:     gamma,
		C:         cBig,
		S:         sBig,
		Seed:      seedBig,
		Output:    outputBig,
	}
	status, err = pf.VerifyVRFProof()
	if err != nil {
		return false, err
	}

	if !status {
		return false, errors.New("invalid proof")
	}

	return true, nil
}

func (v *VRF) checkRequestParams(req VRFRequest) error {
	if *req.Version != consts.VRFVersion {
		return errors.New(fmt.Sprintf("vrf version mismatch, must %d", consts.VRFVersion))
	}
	if !util.IsUUIDv4(*req.TargetAgentID) {
		return errors.New(fmt.Sprintf("invalid target agent id: %s", *req.TargetAgentID))
	}
	if !util.IsHexString(*req.ClientSeed) {
		return errors.New(fmt.Sprintf("invalid client seed: %s", *req.ClientSeed))
	}
	if !util.IsHexString(*req.KeyHash) {
		return errors.New(fmt.Sprintf("invalid key hash: %s", *req.KeyHash))
	}
	//if !util.IsValid13DigitTimestamp(big.NewInt(*req.RequestTimestamp)) {
	//	return errors.New(fmt.Sprintf("invalid request timestamp: %d", *req.RequestTimestamp))
	//}

	requestID, err := v.CalRequestID(*req.Version, *req.TargetAgentID, *req.ClientSeed, *req.RequestTimestamp, *req.CallbackURI)
	if err != nil {
		return err
	}
	if *req.RequestID != requestID {
		return errors.New(fmt.Sprintf("invalid request ID: %s", *req.RequestID))
	}

	return nil
}

func (v *VRFProof) Marshal() (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
