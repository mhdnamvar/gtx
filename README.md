# GTX - Transaction Management System

GTX is a high-performance transaction management system written in Go, designed for processing ISO 8583 financial messages. It provides a comprehensive solution for handling credit card transactions, ATM communications, and other financial message processing requirements.

## Features

- **ISO 8583 Message Processing**: Complete implementation of ISO 8583 message encoding/decoding
- **Multiple Encoding Support**: ASCII, EBCDIC, and Binary encoding formats
- **Network Communication**: TCP-based client/server architecture for message exchange
- **Protocol Import**: Import jPOS ISO packager XML definitions to generate Go code
- **Echo Server**: Built-in echo server for testing and development
- **REST API**: HTTP REST API server for integration
- **Configurable**: Flexible configuration for different ISO 8583 variants

## Installation

### Prerequisites

- Go 1.19 or higher
- Git

### Build from Source

```bash
# Clone the repository
git clone https://github.com/mhdnamvar/gtx.git
cd gtx

# Initialize Go module and install dependencies
go mod init github.com/mhdnamvar/gtx
go mod tidy

# Build the application
go build -o gtx .
```

## Usage

### Available Commands

GTX provides several commands for different operations:

```bash
# Show version information
./gtx version

# Run servers
./gtx run [options]

# Import jPOS protocol definitions
./gtx import <xml-file> <protocol-name>
```

### Running Servers

#### Echo Server
Start an echo server for testing message communication:

```bash
# Run echo server on default port 8583
./gtx run --echo-server

# Run echo server on custom port
./gtx run --echo-server --echo-server-port 9000
```

#### REST API Server
Start a REST API server:

```bash
# Run REST server on default port 8584
./gtx run --rest-server

# Run REST server on custom port
./gtx run --rest-server --rest-server-port 8080
```

#### Run All Servers
Start both echo and REST servers simultaneously:

```bash
./gtx run --all
```

### Protocol Import

Import jPOS ISO packager XML files to generate Go code:

```bash
# Import a protocol definition
./gtx import protocol-iso87binary.xml MyBinary87

# Force overwrite existing protocol
./gtx import --force protocol-iso87ascii.xml MyAscii87
```

This will generate a new Go file in `codec/iso8583/` with the protocol specification.

## Configuration

GTX uses YAML configuration files. Create a `gtx.yaml` file in your project directory:

```yaml
server:
  port: 8081
```

You can also configure the system programmatically or through JSON configuration files when using the network components.

## Project Structure

```
gtx/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command and configuration
│   ├── run.go             # Server commands
│   ├── import.go          # Protocol import functionality
│   └── version.go         # Version command
├── codec/                 # Message encoding/decoding
│   ├── iso8583/          # ISO 8583 implementation
│   │   ├── IsoMsg.go     # ISO message structure
│   │   ├── IsoData.go    # Data type definitions
│   │   └── Default*.go   # Default protocol specifications
│   └── tlv/              # TLV (Tag-Length-Value) support
├── crypto/               # Cryptographic functions
├── net/                  # Network communication
│   ├── conf.go          # Configuration handling
│   ├── net.go           # Network utilities
│   ├── isoserver.go     # ISO server implementation
│   └── isoclient.go     # ISO client implementation
├── utils/               # Utility functions
├── main.go             # Application entry point
└── gtx.yaml           # Configuration file
```

## ISO 8583 Support

GTX provides comprehensive support for ISO 8583 message processing:

### Supported Field Types

- Fixed-length fields (ASCII, EBCDIC, Binary)
- Variable-length fields with LL, LLL, LLLL prefixes
- Numeric, alphabetic, and alphanumeric content
- Bitmap handling for field presence indication
- Amount fields with currency support

### Encoding Formats

- **ASCII**: Standard ASCII text encoding
- **EBCDIC**: IBM EBCDIC encoding for mainframe compatibility
- **Binary**: Raw binary data encoding

### Example Usage

```go
import (
    "github.com/mhdnamvar/gtx/codec/iso8583"
)

// Create a new ISO message
msg := iso8583.NewIsoMsg()

// Set message type indicator
msg.Set(0, "0200")

// Set primary account number
msg.Set(2, "4111111111111111")

// Set processing code
msg.Set(3, "000000")

// Set transaction amount
msg.Set(4, "000000001000")

// Encode the message
encoded, err := msg.Encode(iso8583.DefaultBinary87)
if err != nil {
    log.Fatal(err)
}
```

## Development

### Building

```bash
# Build the application
go build -o gtx .

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

### Testing

The project includes comprehensive tests for the ISO 8583 codec:

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./codec/iso8583/

# Run with verbose output
go test -v ./codec/iso8583/
```

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style

- Follow standard Go formatting (`go fmt`)
- Write comprehensive tests for new features
- Document public APIs with comments
- Use meaningful variable and function names

## Examples

### Simple Echo Client

```go
package main

import (
    "fmt"
    "net"
    "github.com/mhdnamvar/gtx/net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8583")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // Send a message
    message := []byte("Hello, GTX!")
    net.WriteMessage(conn, message)

    // Read the response
    response, err := net.ReadMessage(conn)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Response: %s\n", response)
}
```

### ISO 8583 Message Processing

```go
package main

import (
    "fmt"
    "github.com/mhdnamvar/gtx/codec/iso8583"
)

func main() {
    // Create and populate an ISO message
    msg := iso8583.NewIsoMsg()
    msg.Set(0, "0200")  // Authorization request
    msg.Set(2, "4111111111111111")  // PAN
    msg.Set(3, "000000")  // Processing code
    msg.Set(4, "000000001000")  // Amount
    msg.Set(11, "123456")  // System trace audit number

    // Display the message
    fmt.Println(msg.String())

    // Encode using default binary specification
    encoded, err := msg.Encode(iso8583.DefaultBinary87)
    if err != nil {
        fmt.Printf("Error encoding: %v\n", err)
        return
    }

    fmt.Printf("Encoded message length: %d bytes\n", len(encoded))
}
```

## License

This project is open source. Please check the repository for license information.

## Support

For questions, issues, or contributions:

- Create an issue on GitHub
- Submit a pull request
- Contact the maintainers

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Uses [Gin](https://github.com/gin-gonic/gin) for REST API server
- Inspired by jPOS for ISO 8583 protocol handling

---

GTX provides a robust foundation for financial transaction processing in Go, offering the flexibility and performance needed for modern payment systems.