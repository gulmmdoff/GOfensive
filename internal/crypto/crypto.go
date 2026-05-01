package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Change your key
var MasterKey = []byte("X-C2-Maaster-Key-32-Bytes-Long!!") // 32 characters

func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(MasterKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(cryptoText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %w", err)
	}

	block, err := aes.NewCipher(MasterKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(plainText), nil
}
