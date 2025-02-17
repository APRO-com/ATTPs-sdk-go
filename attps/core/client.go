package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/auth"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/auth/credentials"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/consts"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var (
	regJSONTypeCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	regXMLTypeCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// Client base Client
type Client struct {
	httpClient *http.Client
	credential auth.Credential
	signer     auth.Signer
}

// APIResult  request result
type APIResult struct {
	// The HTTPRequest used by this request
	Request *http.Request
	// The HTTPResponse obtained for this request
	Response *http.Response
}

// ClientOption HTTPClient core.Client initialization parameters
type ClientOption interface {
	// Apply the initialization parameters to DialSettings
	Apply(settings *DialSettings) error
}

// NewClient initializes an HTTPClient
func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	settings, err := initSettings(opts)
	if err != nil {
		return nil, fmt.Errorf("init client setting err:%v", err)
	}

	client := initClientWithSettings(ctx, settings)
	return client, nil
}

// SelectHeaderContentType select a content type from the available list.
func SelectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return consts.ApplicationJSON
	}
	if contains(contentTypes, consts.ApplicationJSON) {
		return consts.ApplicationJSON
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// Request Send request
//
// Compared with Get / Post / Put / Patch / Delete methods, this method can set more content
// In particular, if you need to set the Header for the current request, you can use this method
func (client *Client) Request(
	ctx context.Context,
	method, requestPath string,
	headerParams http.Header,
	queryParams url.Values,
	postBody interface{},
	contentType string,
) (result *APIResult, err error) {

	// Setup path and query parameters
	varURL, err := url.Parse(requestPath)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := varURL.Query()
	for k, values := range queryParams {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	// Encode the parameters.
	varURL.RawQuery = query.Encode()

	if postBody == nil {
		return client.doRequest(ctx, method, varURL.String(), headerParams, contentType, nil, "")
	}

	// Detect postBody type and set body content
	if contentType == "" {
		contentType = consts.ApplicationJSON
	}
	var body *bytes.Buffer
	body, err = setBody(postBody, contentType)
	if err != nil {
		return nil, err
	}
	return client.doRequest(ctx, method, varURL.String(), headerParams, contentType, body, body.String())
}

// UnMarshalResponse organizes the response packet into structured data
func UnMarshalResponse(httpResp *http.Response, resp interface{}) error {
	body, err := io.ReadAll(httpResp.Body)
	_ = httpResp.Body.Close()

	if err != nil {
		return err
	}

	httpResp.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
}

func initSettings(opts []ClientOption) (*DialSettings, error) {
	var (
		o   DialSettings
		err error
	)
	for _, opt := range opts {
		if err = opt.Apply(&o); err != nil {
			return nil, err
		}
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}
	return &o, nil
}

func initClientWithSettings(_ context.Context, settings *DialSettings) *Client {
	client := &Client{
		signer:     settings.Signer,
		credential: &credentials.Credentials{Signer: settings.Signer},
		httpClient: settings.HTTPClient,
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: consts.DefaultTimeout,
		}
	}
	return client
}

func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	switch b := body.(type) {
	case string:
		_, err = bodyBuf.WriteString(b)
	case *string:
		_, err = bodyBuf.WriteString(*b)
	case []byte:
		_, err = bodyBuf.Write(b)
	case **os.File:
		_, err = bodyBuf.ReadFrom(*b)
	case io.Reader:
		_, err = bodyBuf.ReadFrom(b)
	default:
		if regJSONTypeCheck.MatchString(contentType) {
			err = json.NewEncoder(bodyBuf).Encode(body)
		} else if regXMLTypeCheck.MatchString(contentType) {
			err = xml.NewEncoder(bodyBuf).Encode(body)
		}
	}
	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// contains is a case-insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

func (client *Client) doRequest(
	ctx context.Context,
	method string,
	requestURL string,
	header http.Header,
	contentType string,
	reqBody io.Reader,
	signBody string,
) (*APIResult, error) {

	var (
		err            error
		authIdentifier string
		request        *http.Request
	)

	// Construct Request
	if request, err = http.NewRequestWithContext(ctx, method, requestURL, reqBody); err != nil {
		return nil, err
	}

	// Header Setting Priority:
	// Fixed Headers > Per-Request Header Parameters

	// Add Request Header Parameters
	for key, values := range header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// Set Fixed Headers
	request.Header.Set(consts.Accept, "*/*")
	request.Header.Set(consts.ContentType, contentType)

	ua := fmt.Sprintf(consts.UserAgentFormat, consts.Version, runtime.GOOS, runtime.Version())
	request.Header.Set(consts.UserAgent, ua)

	// Set Auth Header
	if authIdentifier, err = client.credential.GenerateAuthorization(
		ctx,
		signBody,
	); err != nil {
		return nil, fmt.Errorf("generate authorization err:%s", err.Error())
	}

	if authIdentifier != "" {
		request.Header.Set(consts.Authorization, authIdentifier)
	}

	// Send HTTP Request
	result, err := client.doHTTP(request)
	if err != nil {
		return result, err
	}
	// Check if Success
	if err = checkResponse(result.Response); err != nil {
		return result, err
	}
	//// TODO: If the returned data in the future also has corresponding signature verification data, need Validate Signature
	//if err = client.validator.Validate(ctx, result.Response); err != nil {
	//	return result, err
	//}
	return result, nil
}

func (client *Client) doHTTP(req *http.Request) (result *APIResult, err error) {
	result = &APIResult{
		Request: req,
	}

	result.Response, err = client.httpClient.Do(req)
	return result, err
}

// checkResponse checks whether the request is successful
//
// When the status code of the http response packet is not in the range of 200-299, the corresponding error information will be returned, mainly including the http status code, response packet error code, and response packet error information prompt
func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = resp.Body.Close()

	resp.Body = io.NopCloser(bytes.NewBuffer(slurp))
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(slurp),
	}
	// Ignore JSON parsing errors and return apiError
	_ = json.Unmarshal(slurp, apiError)
	return apiError
}
