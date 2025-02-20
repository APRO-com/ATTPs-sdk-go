package vrf

import (
	"context"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/option"
	"github.com/APRO-com/ATTPs-sdk-go/util"
	"log"
	"testing"
	"time"
)

func TestVRF_Request(t *testing.T) {
	baseAPIServer := "http://10.0.54.95:8888" // base API address

	ctx := context.Background()

	opts := []core.ClientOption{
		option.WithNullAuthCipher(),
	}
	client, err := core.NewClient(ctx, baseAPIServer, opts...)
	if err != nil {
		log.Printf("new client err:%s", err)
		return
	}

	svc := VRF{Client: client}
	providers, err := svc.Provider(ctx)
	if err != nil {
		t.Errorf("call Provider err:%s", err)
		return
	}

	version := int64(1)
	targetAgentID := util.NewUUIDv4()
	customerSeed := util.SecureRandomString(4)
	keyHash := providers[0].KeyHash
	callbackUri := "http://127.0.0.1:8888/api/vrf/proof"

	requestTimestamp := time.Now().Unix()
	requestID, err := new(VRF).CalRequestID(version, targetAgentID, customerSeed, requestTimestamp, callbackUri)
	if err != nil {
		t.Errorf("call CalRequestID err:%s", err)
		return
	}
	random, err := svc.Request(ctx,
		VRFRequest{
			Version:          core.Int64(version),
			TargetAgentID:    core.String(targetAgentID),
			ClientSeed:       core.String(customerSeed),
			KeyHash:          core.String(*keyHash),
			RequestTimestamp: core.Int64(requestTimestamp),
			RequestID:        core.String(requestID),
			CallbackURI:      core.String(callbackUri),
		},
	)

	if err != nil {
		t.Errorf("call Request err:%s", err)
		return
	} else {
		t.Log(random)
	}

	// query proof
	proof, err := svc.QueryProof(ctx, requestID)
	if err != nil {
		t.Errorf("call QueryProof err:%s", err)
		return
	} else {
		t.Log(proof.Marshal())
	}
}

func TestVRF_VerifyProof(t *testing.T) {
	status, err := new(VRF).VerifyVRFProof(&VRFProof{
		RequestId: core.String("297ad27913e541f293700c63960f9e6bdc840b2982ec6be363c0852a8c7c403b"),
		Proof: &Proof{
			PublicX: core.String("0xed3bace23c5e17652e174c835fb72bf53ee306b3406a26890221b4cef7500f88"),
			PublicY: core.String("0xe57a6f571288ccffdcda5e8a7a1f87bf97bd17be084895d0fce17ad5e335286e"),
			GammaX:  core.String("0x7ce22e7667f955f5dcc805a5bae7f78d21d0cb04eb5190f3b8e20b68a45d0b87"),
			GammaY:  core.String("0xc8f9d9e8d5e4eb22adf379df733a8b1ce4edf26a2ca9a4a3d8a07cb3e3dffd9"),
			C:       core.String("0x45945e1b7362a7026df893d39496eb838b6d85264f56899182269be4d53d6fe"),
			S:       core.String("0x7ebf871ad068ce4bbe04cd726e359334581881b4e78da352b7ac413ebcf90a2"),
			Seed:    core.String("0xd3ea21873da2909f9f732966278cc022d523006ea574a58b324b83c0c08a5346"),
			Output:  core.String("0x11449014f7e3fb46f190149f5c147242300ccdb0e77a58fa53e01972939a3f14"),
		},
	})
	if err != nil || !status {
		t.Errorf("call VerifyProof err:%s", err)
		return
	}
}

func TestVRF_CalRequestID(t *testing.T) {
	benchmarkRequestID := "6f71619f1e6ea42616c9bbdc8fe001511e0c37b72373dc259857b29c1e61597c"
	version := int64(1)
	targetAgentID := "f2464336-fbcf-4603-bda5-ce65c0318fb6"
	customerSeed := "1234"
	callbackUri := "http://127.0.0.1:8888/api/vrf/proof"
	requestTimestamp := int64(1739265192)

	requestID, err := new(VRF).CalRequestID(version, targetAgentID, customerSeed, requestTimestamp, callbackUri)
	if err != nil {
		t.Errorf("call CalRequestID err:%s", err)
		return
	}
	if requestID != benchmarkRequestID {
		t.Errorf("requestID:%s != %s", requestID, benchmarkRequestID)
		return
	}
}
