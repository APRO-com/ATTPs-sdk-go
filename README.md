# APRO AI Agent SDK for Go

The **APRO AI Agent SDK for Go** is a Go-based SDK designed to interact with the APRO AI Agent system. It provides tools and methods for communicating with the APRO AI Agent platform, sending transactions, managing agent settings, and more.

This SDK is intended for developers building Go applications that need to interact with the **APRO AI Agent**.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use the AI Agent SDK for Go, you need to install the package via `go get`. Run the following command to install the SDK:

```bash
go get github.com/APRO-com/ai-agent-sdk-go
```

## Usage
### Basic Example
Here's an example of how to use the SDK to interact with the AI Agent system:
```
func main() {
    // Initialize the AI Agent client with the test network
    client, err := ai_agent_sdk.NewClient(config.BSC_TEST_NET)
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    // Get the version of the AI Agent
    version, err := client.GetVersion("your-proxy-address")
    if err != nil {
        log.Fatalf("Error getting version: %v", err)
    }

    fmt.Println("AI Agent Version:", version)
}
```

For more examples, see [client_test.go](sdk%2Fclient_test.go)

## Contributing

We welcome contributions to the AI Agent SDK for Go. To contribute:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-name).
3. Make your changes.
4. Commit your changes (git commit -am 'Add new feature').
5. Push to the branch (git push origin feature-name).
6. Open a pull request.

Please ensure your code adheres to the existing code style and includes tests where appropriate.

## License

This project is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html).