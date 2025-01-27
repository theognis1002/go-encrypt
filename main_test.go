package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rc4"
	"io"
	"testing"
)

func TestAESEncryption(t *testing.T) {
	key := []byte("1234567890123456") // 16-byte key
	data := []byte("Hello, World!")

	// Encrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatalf("Failed to create AES cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatalf("Failed to create GCM: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		t.Fatalf("Failed to generate nonce: %v", err)
	}

	encrypted := gcm.Seal(nonce, nonce, data, nil)

	// Decrypt
	block2, _ := aes.NewCipher(key)
	gcm2, _ := cipher.NewGCM(block2)
	nonceSize := gcm2.NonceSize()
	nonce2, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	decrypted, err := gcm2.Open(nil, nonce2, ciphertext, nil)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("Decrypted data doesn't match original")
	}
}

func TestDESEncryption(t *testing.T) {
	key := []byte("12345678") // 8-byte key
	data := []byte("Hello, World!")

	// Encrypt
	block, err := des.NewCipher(key)
	if err != nil {
		t.Fatalf("Failed to create DES cipher: %v", err)
	}

	// Pad the data to match block size
	blockSize := block.BlockSize()
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	data = append(data, padtext...)

	// Create IV and encrypt
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		t.Fatalf("Failed to generate IV: %v", err)
	}

	encrypted := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, data)

	// Decrypt
	mode = cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	// Remove padding
	unpadding := int(decrypted[len(decrypted)-1])
	decrypted = decrypted[:len(decrypted)-unpadding]

	if !bytes.Equal([]byte("Hello, World!"), decrypted) {
		t.Error("Decrypted data doesn't match original")
	}
}

func TestRC4Encryption(t *testing.T) {
	key := []byte("1234567890123456") // 16-byte key
	data := []byte("Hello, World!")

	// Encrypt
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		t.Fatalf("Failed to create RC4 cipher: %v", err)
	}

	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)

	// Decrypt (RC4 is symmetric)
	cipher2, _ := rc4.NewCipher(key)
	decrypted := make([]byte, len(encrypted))
	cipher2.XORKeyStream(decrypted, encrypted)

	if !bytes.Equal(data, decrypted) {
		t.Error("Decrypted data doesn't match original")
	}
}
