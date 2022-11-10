package ecb

import (
	"crypto/aes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func Encrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	plainText, err = padder.Pad(plainText) // pad last block of plaintext if block size less than block cipher size
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(plainText))
	mode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

func Decrypt(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBDecrypter(block)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	plainText, err = padder.Unpad(plainText) // unpad plaintext after decryption
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func EncryptToHex(plainText, key []byte) (string, error) {
	b, err := Encrypt(plainText, key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}

func DecryptFromHex(plainText string, key []byte) (string, error) {
	b, err := hex.DecodeString(plainText)
	if err != nil {
		return "", err
	}
	pb, err := Decrypt(b, key)
	if err != nil {
		return "", err
	}
	return string(pb), nil
}

func Sha1ToHex(pt []byte) string {
	h := sha1.New()
	h.Write(pt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
