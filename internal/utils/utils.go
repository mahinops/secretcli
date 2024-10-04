package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

const key = "thisis32bitlongpassphraseimusing"

func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextBytes)

	return hex.EncodeToString(cipherText), nil
}

func Decrypt(cipherTextHex string) (string, error) {
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	plainText := make([]byte, len(cipherText))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

// hashPassword hashes the password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
