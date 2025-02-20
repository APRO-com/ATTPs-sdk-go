# APRO ATTPs SDK for Go

The **APRO ATTPs SDK for Go** is a Go-based SDK designed to interact with the APRO ATTPs system. It provides tools and methods for communicating with the APRO ATTPs platform, sending transactions, managing agent settings, and more.

This SDK is intended for developers building Go applications that need to interact with the **APRO ATTPs **.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use the ATTPs SDK for Go, you need to install the package via `go get`. Run the following command to install the SDK:

```bash
go get github.com/APRO-com/ATTPs-sdk-go
```

## Usage
### Basic Verify Example
Here's an example of how to use the SDK to interact with the ATTPs system:
```
import (
	"fmt"
	"github.com/APRO-com/ATTPs-sdk-go/attps/verify"
	"github.com/APRO-com/ATTPs-sdk-go/attps/verify/config"
	"log"
)

func main() {
	// Initialize the ATTPs verify client with the test network
	client, err := verify.NewClient(config.BSC_TEST_NET)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Get the version of the ATTPs verify
	version, err := client.GetVersion("your-proxy-address")
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}

	fmt.Println("ATTPs verify Version:", version)
}
```

For more examples, see [client_test.go](attps/verify/client_test.go)

### VRF Example
Here's an example of how to use the SDK to interact with the ATTPs VRF system:
```
import (
	"github.com/APRO-com/ATTPs-sdk-go/attps/core"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/option"
	"github.com/APRO-com/ATTPs-sdk-go/attps/vrf"
	"github.com/APRO-com/ATTPs-sdk-go/util"

	"context"
	"log"
	"time"
)

func main() {
	baseAPIServer := "http://10.0.54.95:8888" // base API address
    
	ctx := context.Background()

	opts := []core.ClientOption{
		option.WithNullAuthCipher(),
	}
	client, err := core.NewClient(ctx, baseAPIServer, opts...)
	if err != nil {
		log.Fatalf("new client err:%s", err)
		return
	}

	svc := vrf.VRF{Client: client}
	providers, err := svc.Provider(ctx)
	if err != nil {
		log.Fatalf("call Provider err:%s", err)
		return
	}

	version := int64(1)
	targetAgentID := util.NewUUIDv4()
	customerSeed := util.SecureRandomString(4)
	keyHash := providers[0].KeyHash
	callbackUri := "http://127.0.0.1:8888/api/vrf/proof"

	requestTimestamp := time.Now().Unix()
	requestID, err := new(vrf.VRF).CalRequestID(version, targetAgentID, customerSeed, requestTimestamp, callbackUri)
	if err != nil {
		log.Fatalf("call CalRequestID err:%s", err)
		return
	}
	random, err := svc.Request(ctx,
		vrf.VRFRequest{
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
		log.Fatalf("call Request err:%s", err)
		return
	} else {
		log.Print(random)
	}

	// query proof
	proof, err := svc.QueryProof(ctx, requestID)
	if err != nil {
		log.Fatalf("call QueryProof err:%s", err)
		return
	} else {
		log.Print(proof.Marshal())
	}
}

```

For more examples, see [vrf_test.go](attps/vrf/vrf_test.go)

In customer service, it is necessary to implement a callback to obtain the VRF proof and perform verification. Specifically, implement the /api/vrf/proof request notification interface. Refer to VRFProof for the request format, and the response should follow the BaseResponse format.

## Contributing

We welcome contributions to the ATTPs SDK for Go. To contribute:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-name).
3. Make your changes.
4. Commit your changes (git commit -am 'Add new feature').
5. Push to the branch (git push origin feature-name).
6. Open a pull request.

Please ensure your code adheres to the existing code style and includes tests where appropriate.

## License

This project is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html).