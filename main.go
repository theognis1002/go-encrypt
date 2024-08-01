package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	key := []byte("example key 1234") // 16 bytes key for AES-128
	inputFile := "example.csv"
	encryptedFile := "encrypted.csv"
	decryptedFile := "decrypted.csv"

	// Encrypt the file
	err := encryptFile(inputFile, encryptedFile, key)
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}
	fmt.Println("File encrypted successfully.")

	// Decrypt the file
	err = decryptFile(encryptedFile, decryptedFile, key)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		return
	}
	fmt.Println("File decrypted successfully.")
}

// encryptFile encrypts the file at inputPath and writes the encrypted data to outputPath
func encryptFile(inputPath, outputPath string, key []byte) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Generate a new AES cipher using our 16, 24 or 32 bytes long key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Encrypt the data using Seal (which also appends the nonce and the encrypted data together)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return os.WriteFile(outputPath, ciphertext, 0644)
}

// decryptFile decrypts the file at inputPath and writes the decrypted data to outputPath
func decryptFile(inputPath, outputPath string, key []byte) error {
	ciphertext, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, plaintext, 0644)
}
