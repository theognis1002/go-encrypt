# File Encryption & Decryption in Go

This project provides a simple implementation of file encryption and decryption using the AES-GCM (Galois/Counter Mode) in Go. It demonstrates how to securely encrypt and decrypt files using a secret key.

## Requirements

- Go 1.16 or later

## Usage

### Setup

1. **Clone the repository:**

   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Set the secret key:**

   Update the key variable in the main.go file with your 16, 24, or 32 bytes key for AES-128, AES-192, or AES-256 respectively.

## Makefile Commands

The Makefile includes commands to build, encrypt, and decrypt files, as well as to clean up the build files.

1. **Build the project:**

   ```sh
   make build
   ```

   This command compiles the Go program into a binary named file-crypto.

2. **Encrypt a file:**

   ```sh
   make encrypt
   ```

   This command encrypts the file example.csv into `encrypted.csv`. You can change the input and output filenames by editing the Makefile or running the binary directly with the desired arguments.

3. **Decrypt a file:**

   ```sh
   make decrypt
   ```

   This command decrypts the file `encrypted.csv` into decrypted.csv.

4. **Clean up build files:**

   ```sh
   make clean
   ```

   This command removes the binary and any other build-related files.

## Example Usage

### Encrypting a file

To encrypt `example.csv` into `encrypted.csv`, run:

    ```sh
    make encrypt
    ```

### Decrypting a file

To decrypt `encrypted.csv` into `decrypted.csv`, run:

    ```sh
    make decrypt
    ```
