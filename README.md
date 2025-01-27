# Multi-Algorithm File Encryption & Decryption in Go

This project provides a robust implementation of file encryption and decryption using multiple encryption algorithms in Go. It supports AES-GCM, DES-CBC, and RC4 encryption methods.

## Features

- Multiple encryption algorithms:
  - AES (Advanced Encryption Standard) with GCM mode
  - DES (Data Encryption Standard) with CBC mode
  - RC4 (Rivest Cipher 4) stream cipher
- Secure key handling
- File-based encryption/decryption
- Environment variable configuration

## Requirements

- Go 1.21 or later
- Environment variables configuration

## Supported Algorithms

### AES (Advanced Encryption Standard)

- Key size: 16 bytes (AES-128)
- Modern, secure block cipher
- Recommended for most use cases
- Uses GCM mode for authenticated encryption

### DES (Data Encryption Standard)

- Key size: 8 bytes
- Legacy block cipher
- Not recommended for new applications
- Uses CBC mode with PKCS7 padding

### RC4 (Rivest Cipher 4)

- Variable key size (using 16 bytes in this implementation)
- Stream cipher
- Fast but not cryptographically secure
- Not recommended for sensitive data

## Usage

### Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/go-encrypt
   cd go-encrypt/
   ```

2. **Configure the environment:**
   Create a `.env` file with the following variables:

   ```
   ALGORITHM=AES  # Can be AES, DES, or RC4
   KEY=your16bytesecret  # Key size depends on algorithm
   INPUT_FILE=example.txt
   ENCRYPTED_FILE=encrypted.bin
   DECRYPTED_FILE=decrypted.txt
   ```

   Key size requirements:

   - AES: 16 bytes
   - DES: 8 bytes
   - RC4: 16 bytes

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

### Example Usage

1. **Using AES encryption:**

   ```sh
   ALGORITHM=AES KEY=1234567890123456 make encrypt
   ```

2. **Using DES encryption:**

   ```sh
   ALGORITHM=DES KEY=12345678 make encrypt
   ```

3. **Using RC4 encryption:**
   ```sh
   ALGORITHM=RC4 KEY=1234567890123456 make encrypt
   ```

## Security Considerations

- AES is the recommended algorithm for new applications
- DES is included for legacy compatibility but should not be used for new applications
- RC4 is included for educational purposes but should not be used for sensitive data
- Always use strong, random keys
- Keep your keys secure and never commit them to version control

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
