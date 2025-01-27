package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/theognis1002/go-encrypt/test"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	algorithmName := os.Getenv("ALGORITHM")
	key := []byte(os.Getenv("KEY"))
	inputFile := os.Getenv("INPUT_FILE")
	encryptedFile := os.Getenv("ENCRYPTED_FILE")
	decryptedFile := os.Getenv("DECRYPTED_FILE")

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	var encrypted, decrypted []byte
	var errEncrypt error

	switch algorithmName {
	case "AES":
		encrypted, errEncrypt = test.EncryptAES(data, key)
	case "DES":
		encrypted, errEncrypt = test.EncryptDES(data, key)
	case "RC4":
		encrypted, errEncrypt = test.EncryptRC4(data, key)
	default:
		log.Fatalf("Unsupported algorithm: %s", algorithmName)
	}

	if errEncrypt != nil {
		log.Fatalf("Error encrypting: %v", errEncrypt)
	}
	err = os.WriteFile(encryptedFile, encrypted, 0644)
	if err != nil {
		log.Fatalf("Error writing encrypted file: %v", err)
	}
	fmt.Println("File encrypted successfully")

	// Decrypt
	decrypted, err = test.Decrypt(encrypted, key)
	if err != nil {
		log.Fatalf("Error decrypting: %v", err)
	}
	err = os.WriteFile(decryptedFile, decrypted, 0644)
	if err != nil {
		log.Fatalf("Error writing decrypted file: %v", err)
	}
	fmt.Println("File decrypted successfully")
}
