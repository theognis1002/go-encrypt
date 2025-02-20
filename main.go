package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	keyPtr := flag.String("key", "", "16-byte encryption/decryption key")
	inputPtr := flag.String("input", "", "input file path")
	encryptPtr := flag.String("encrypt", "", "output path for encrypted file")
	decryptPtr := flag.String("decrypt", "", "output path for decrypted file")

	flag.Parse()

	// Validate required flags
	if *keyPtr == "" || *inputPtr == "" || (*encryptPtr == "" && *decryptPtr == "") {
		flag.Usage()
		log.Fatalf("Missing required flags")
	}

	key := []byte(*keyPtr)
	inputFile := *inputPtr

	if len(key) != 16 {
		log.Fatalf("Key length must be 16 bytes")
	}

	err := os.MkdirAll("output", 0755)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	fmt.Println("Key:", string(key))
	fmt.Println("Input File:", inputFile)

	if *encryptPtr != "" {
		outputPath := filepath.Join("output", *encryptPtr)
		err = encryptFile(inputFile, outputPath, key)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully to:", outputPath)
	}

	if *decryptPtr != "" {
		inputForDecrypt := inputFile
		if *encryptPtr != "" {
			inputForDecrypt = filepath.Join("output", *encryptPtr)
		}
		outputPath := filepath.Join("output", *decryptPtr)
		err = decryptFile(inputForDecrypt, outputPath, key)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully to:", outputPath)
	}
}

// encryptFile encrypts the file at inputPath and writes the encrypted data to outputPath
func encryptFile(inputPath, outputPath string, key []byte) error {
	data, err := os.ReadFile(inputPath)
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

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

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	return os.WriteFile(outputPath, plaintext, 0644)
}
