package cbc

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/andreburgaud/crypt2go/padding"
)

func Encrypt(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	padder := padding.NewPkcs7Padding(blockSize)
	plainText, err = padder.Pad(plainText) // pad last block of plaintext if block size less than block cipher size
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypted, plainText)
	return crypted, nil
}

func Decrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	plainText, err = padder.Unpad(plainText) // unpad plaintext after decryption
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func EncryptToHex(plainText, key, iv []byte) (string, error) {
	b, err := Encrypt(plainText, key, iv)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}

func DecryptFromHex(plainText string, key, iv []byte) (string, error) {
	b, err := hex.DecodeString(plainText)
	if err != nil {
		return "", err
	}
	pb, err := Decrypt(b, key, iv)
	if err != nil {
		return "", err
	}
	return string(pb), nil
}

func EncryptToBase64(plainText, key, iv []byte) (string, error) {
	b, err := Encrypt(plainText, key, iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func DecryptFromBase64(cipherText string, key, iv []byte) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	pb, err := Decrypt(b, key, iv)
	if err != nil {
		return "", err
	}
	return string(pb), nil
}
