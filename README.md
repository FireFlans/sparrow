# SPARROW - Security Policy AdministRation and Retrieval Over the Web 
![logo](documentation/assets/logo.jpg)
SPARROW is a web service built with Gin-Gonic that processes XML Security Policy Information Files (as described in NATO's STANAG 4774) and returns JSON informations about value domains. 
This project is designed to handle SPIF XML Files and provide a structured JSON output for other services dealing with security labels.
In addition it provides a security labels playground and an administration interface

## Features

- **SPIF Parsing**: Load any XML SPIF file, parse it and access data through REST API
- **Security Label Handling**: Build STANAG 4774 labels (coming soon), convert them in JSON in a full (coming soon) or in a simplified format
- **Security Label Playground (scheduled)**: Build compliant security labels from the provided SPIFs or from your own security policies
- **SPIF Administration (scheduled)**: Administrate your security policies from a unique web interface

## Getting Started

### Prerequisites

- Go 1.23+

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/FireFlans/sparrow.git
   cd sparrow
   ```
2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
### Running the Server
1. **Start the Server**:
   ```bash
   go run main.go
   ```
2. **Access the API**:
   - The server runs on `http://localhost:8080`.
   - To access API documentation, go to `http://localhost:8080/documentation/index.html`
## Testing
To run the tests, start the server and use the following command:
```bash
cd test && go test -v
```
## Adding your own SPIF files

SPIF files are located in `config/spifs`\
You add your own in this folder

## Contributing

If you feel something is missing, don't hesitate to open an issue
