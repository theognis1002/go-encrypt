package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	key := []byte(os.Getenv("KEY"))
	inputFile := os.Getenv("INPUT_FILE")
	encryptedFile := os.Getenv("ENCRYPTED_FILE")
	decryptedFile := os.Getenv("DECRYPTED_FILE")

	// Check if the key length is correct
	if len(key) != 16 {
		log.Fatalf("Key length must be 16 bytes")
	}

	fmt.Println("Key:", string(key))
	fmt.Println("Input File:", inputFile)
	fmt.Println("Encrypted File:", encryptedFile)
	fmt.Println("Decrypted File:", decryptedFile)

	// Encrypt the file
	err = encryptFile(inputFile, encryptedFile, key)
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
