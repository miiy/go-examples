package ecb

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

var (
	pt = []byte("Some plain text")
	// Key size for AES is either: 16 bytes (128 bits), 24 bytes (192 bits) or 32 bytes (256 bits)
	key = []byte("secretkey16bytes")
)

// AES encryption with ECB and PKCS7 padding
func TestAesEcb(t *testing.T) {
	cipherText, err := Encrypt(pt, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ciphertext:", fmt.Sprintf("%X", cipherText))

	recoveredPt, err := Decrypt(cipherText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Recovered plaintext:", fmt.Sprintf("%s", recoveredPt))

	h := sha1.New()
	h.Write(pt)
	fmt.Println("Sha1:", fmt.Sprintf("%x", h.Sum(nil)))
}

func TestHex(t *testing.T) {
	cipherText, err := EncryptToHex(pt, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cipherText: %s\n", cipherText)
	plainText, err := DecryptFromHex(cipherText, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plainText: %s\n", plainText)
}

func TestSha1ToHex(t *testing.T) {
	fmt.Printf("Sha1: %s\n", Sha1ToHex(pt))
}

// Output:
// cipherText: AF3B0173EAE9DD013A649F4EAABA1376
// plainText plaintext: Some plain text
// Sha1: d85e382c4a48731d850ec5956a20e5b3ccaa0e7d
//
// MySQL
// select hex(aes_encrypt('Some plain text', 'secretkey16bytes'))
// -- AF3B0173EAE9DD013A649F4EAABA1376
// select sha1('Some plain text')
// -- d85e382c4a48731d850ec5956a20e5b3ccaa0e7d
