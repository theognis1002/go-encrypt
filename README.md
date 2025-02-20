# Multi-Algorithm File Encryption & Decryption in Go

This project provides a robust implementation of file encryption and decryption using AES-GCM in Go.

## Features

- AES (Advanced Encryption Standard) with GCM mode
- Secure key handling
- File-based encryption/decryption
- Command-line interface
- Makefile support

## Requirements

- Go 1.21 or later

## Supported Algorithm

### AES (Advanced Encryption Standard)

- Key size: 16 bytes (AES-128)
- Modern, secure block cipher
- Uses GCM mode for authenticated encryption
- Recommended for secure data encryption

## Usage

### Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/go-encrypt
   cd go-encrypt/
   ```

### Running the Program

1. **Build the program:**

   ```sh
   make build
   ```

2. **Encrypt a file:**

   ```sh
   make encrypt
   ```

3. **Decrypt a file:**
   ```sh
   make decrypt
   ```

### Command Line Options

The program supports the following flags:

```sh
-key string    16-byte encryption/decryption key
-input string  input file path
-encrypt string output path for encrypted file
-decrypt string output path for decrypted file
```

### Example Usage

1. **Direct command line usage:**

   ```sh
   ./go-encrypt -key "1234567890123456" -input "example.txt" -encrypt "encrypted.bin"
   ./go-encrypt -key "1234567890123456" -input "encrypted.bin" -decrypt "decrypted.txt"
   ```

2. **Using Makefile (with default values):**

   ```sh
   make encrypt  # Uses KEY=1234567890123456 and INPUT_FILE=example.csv
   make decrypt
   ```

## Security Considerations

- Always use strong, random keys (16 bytes for AES-128)
- Keep your keys secure and never commit them to version control
- The implementation uses AES-GCM which provides both confidentiality and authenticity

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Testing

### Running Tests

1. **Run all tests:**

   ```sh
   make test
   ```

2. **Run tests with coverage:**

   ```sh
   make test-coverage
   ```

   This will generate a coverage report in HTML format (coverage.html)

3. **Run specific tests:**

   ```sh
   go test -v ./... -run TestName
   ```

4. **Run tests for a specific package:**

   ```sh
   go test ./pkg/encryption -v
   ```

5. **Run tests with race detection:**

   ```sh
   go test -race ./...
   ```

### Test Organization

Tests are organized in the following structure:

```
├── pkg/
│   ├── encryption/
│   │   ├── aes_test.go
│   │   ├── des_test.go
│   │   └── rc4_test.go
│   └── utils/
│       └── utils_test.go
```

Each test file corresponds to its implementation file and follows Go's standard testing conventions.

### Test Coverage

The test suite includes:

- Unit tests for each encryption algorithm (AES, DES, RC4)
- Key size validation tests
- Encryption/decryption round-trip tests
- Error handling tests

### Writing Tests

When contributing new features, please ensure:

- All new code is thoroughly tested
- Tests are placed in the appropriate `_test.go` file
- Test coverage is maintained or improved
- Tests are clear and well-documented
