package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const key = "thisis32bitlongpassphraseimusing"

func encrypt(plainText string) (string, error) {
	// Create a new AES cipher using the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Convert plain text to byte slice
	plainTextBytes := []byte(plainText)

	// Create a byte array for the ciphertext with space for the IV
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))

	// Generate a random IV and store it in the beginning of the ciphertext slice
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the plain text using AES in CFB mode
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextBytes)

	// Return the ciphertext as a hex string
	return hex.EncodeToString(cipherText), nil
}

func decrypt(cipherTextHex string) (string, error) {
	// Convert the hex string to byte slice
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher using the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Ensure that the length of cipherText is at least the block size
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract the IV from the beginning of the ciphertext
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Create a buffer to hold the decrypted plaintext
	plainText := make([]byte, len(cipherText))

	// Decrypt the ciphertext using AES in CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainText, cipherText)

	// Return the decrypted plaintext as a string
	return string(plainText), nil
}
