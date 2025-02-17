package option

import (
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/auth"
	"net/http"

	"github.com/APRO-com/ATTPs-sdk-go/attps/core"
)

// region SignerOption

// withSignerOption sets the Signer for the Client
type withSignerOption struct {
	Signer auth.Signer
}

// Apply adds the configuration to core.DialSettings
func (w withSignerOption) Apply(o *core.DialSettings) error {
	o.Signer = w.Signer
	return nil
}

// WithSigner returns a ClientOption with a specified signer
func WithSigner(signer auth.Signer) core.ClientOption {
	return withSignerOption{Signer: signer}
}

// endregion

// region HTTPClientOption

// withHTTPClientOption sets HTTPClient for Client
type withHTTPClientOption struct {
	Client *http.Client
}

// Apply adds the configuration to core.DialSettings
func (w withHTTPClientOption) Apply(o *core.DialSettings) error {
	o.HTTPClient = w.Client
	return nil
}

// WithHTTPClient returns a ClientOption that specifies the network communication as HttpClient. After specifying, use the http.client automatically created by the user. If the user does not create one, help the user
// Create a default http.client
func WithHTTPClient(client *http.Client) core.ClientOption {
	return withHTTPClientOption{Client: client}
}

// endregion
