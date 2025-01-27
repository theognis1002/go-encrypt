package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rc4"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EncryptionAlgorithm interface {
	Encrypt(data []byte, key []byte) ([]byte, error)
	Decrypt(data []byte, key []byte) ([]byte, error)
	GetKeySize() int
}

type AESAlgorithm struct{}
type DESAlgorithm struct{}
type RC4Algorithm struct{}

// AES implementation
func (a *AESAlgorithm) Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func (a *AESAlgorithm) Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func (a *AESAlgorithm) GetKeySize() int {
	return 16 // AES-128
}

// DES implementation
func (d *DESAlgorithm) Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)

	ciphertext := make([]byte, len(data))
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, data)

	// Prepend IV to ciphertext
	return append(iv, ciphertext...), nil
}

func (d *DESAlgorithm) Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := data[:blockSize]
	data = data[blockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	return pkcs7Unpadding(data)
}

func (d *DESAlgorithm) GetKeySize() int {
	return 8 // DES uses 8-byte keys
}

// RC4 implementation
func (r *RC4Algorithm) Encrypt(data []byte, key []byte) ([]byte, error) {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(data))
	cipher.XORKeyStream(result, data)
	return result, nil
}

func (r *RC4Algorithm) Decrypt(data []byte, key []byte) ([]byte, error) {
	return r.Encrypt(data, key) // RC4 encryption and decryption are the same operation
}

func (r *RC4Algorithm) GetKeySize() int {
	return 16 // Using 16 bytes for RC4 key
}

// Helper functions for padding
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	padding := int(data[length-1])
	return data[:length-padding], nil
}

func getAlgorithm(name string) (EncryptionAlgorithm, error) {
	switch strings.ToUpper(name) {
	case "AES":
		return &AESAlgorithm{}, nil
	case "DES":
		return &DESAlgorithm{}, nil
	case "RC4":
		return &RC4Algorithm{}, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s", name)
	}
}

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

	algorithm, err := getAlgorithm(algorithmName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if len(key) != algorithm.GetKeySize() {
		log.Fatalf("Key length must be %d bytes for %s", algorithm.GetKeySize(), algorithmName)
	}

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Encrypt
	encrypted, err := algorithm.Encrypt(data, key)
	if err != nil {
		log.Fatalf("Error encrypting: %v", err)
	}
	err = os.WriteFile(encryptedFile, encrypted, 0644)
	if err != nil {
		log.Fatalf("Error writing encrypted file: %v", err)
	}
	fmt.Println("File encrypted successfully")

	// Decrypt
	decrypted, err := algorithm.Decrypt(encrypted, key)
	if err != nil {
		log.Fatalf("Error decrypting: %v", err)
	}
	err = os.WriteFile(decryptedFile, decrypted, 0644)
	if err != nil {
		log.Fatalf("Error writing decrypted file: %v", err)
	}
	fmt.Println("File decrypted successfully")
}
