package cbc

import (
	"fmt"
	"testing"
)

var (
	pt = []byte("Some plain text")
	// Key size for AES is either: 16 bytes (128 bits), 24 bytes (192 bits) or 32 bytes (256 bits)
	key = []byte("secretkey16bytes")
	iv  = []byte("1111111111111111")
)

// AES encryption with CBC and PKCS7 padding
func TestEncrypt(t *testing.T) {
	cipherText, err := Encrypt(pt, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cipherText: %X\n", cipherText)

	plainText, err := Decrypt(cipherText, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plainText: %s\n", plainText)
}

func TestHex(t *testing.T) {
	cipherText, err := EncryptToHex(pt, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cipherText: %s\n", cipherText)
	plainText, err := DecryptFromHex(cipherText, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plainText: %s\n", plainText)
}

func TestBase64(t *testing.T) {
	cipherText, err := EncryptToBase64(pt, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cipherText: %s\n", cipherText)
	plainText, err := DecryptFromBase64(cipherText, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plainText: %s\n", plainText)
}
