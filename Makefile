### Makefile

```Makefile
# Makefile for the Go File Encryption and Decryption project

# The binary name
BINARY_NAME = file-crypto

# Build the Go binary
build:
	@echo "Building the Go binary..."
	go build -o $(BINARY_NAME) main.go

# Run the Go application with encryption
encrypt:
	@echo "Encrypting the file..."
	./$(BINARY_NAME) encrypt example.csv encrypted.csv

# Run the Go application with decryption
decrypt:
	@echo "Decrypting the file..."
	./$(BINARY_NAME) decrypt encrypted.csv decrypted.csv

# Clean up build files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Default target: build the binary and run tests
all: build test
