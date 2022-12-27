package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)
var DEFAULT_KEY = "8v]CBx_d$ibrcA7t[SFgF.MRg,F+F=v{"
func Encrypt(stringToEncrypt , keyString string) (encryptedString string, errorS error) {
	// Create a new cipher block
	key, text := []byte(keyString), []byte(stringToEncrypt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given ciphertext using the given key
func Decrypt(encryptedString string, keyString string) (decryptedString string, errorS error) {
	// Create a new cipher block
	ciphertextBytes, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	key, text := []byte(keyString), ciphertextBytes
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	return string(data), nil
}



func TestEncrypt(stringToEncrypt , keyString string) (encryptedString string) {
	en, err := Encrypt(stringToEncrypt, keyString)
	fmt.Println(en, err)
	return en
}
func TestDecrypt(encryptedString string, keyString string) (decryptedString string) {
	en, err := Decrypt(encryptedString, keyString)
	fmt.Println(en, err)
	return en
}