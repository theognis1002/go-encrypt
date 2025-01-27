package test

import (
	"bytes"
	"testing"
)

// Test data
var (
	testKey  = []byte("1234567890123456") // 16-byte key for AES
	testData = []byte("Hello, World!")
)

func TestAESEncryption(t *testing.T) {
	encrypted, err := EncryptAES(testData, testKey)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := DecryptAES(encrypted, testKey)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(testData, decrypted) {
		t.Error("Decrypted data doesn't match original")
	}
}

func TestDESEncryption(t *testing.T) {
	desKey := []byte("12345678") // 8-byte key for DES

	encrypted, err := EncryptDES(testData, desKey)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := DecryptDES(encrypted, desKey)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(testData, decrypted) {
		t.Error("Decrypted data doesn't match original")
	}
}

func TestRC4Encryption(t *testing.T) {
	encrypted, err := EncryptRC4(testData, testKey)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := DecryptRC4(encrypted, testKey)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(testData, decrypted) {
		t.Error("Decrypted data doesn't match original")
	}
}

func TestInvalidKeySize(t *testing.T) {
	invalidKey := []byte("tooshort")
	_, err := EncryptAES(testData, invalidKey)
	if err == nil {
		t.Error("Expected error for invalid AES key size")
	}
}
