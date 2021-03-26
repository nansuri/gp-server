package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

type RSAKeyPair struct {
	Private rsa.PrivateKey
	Public  rsa.PublicKey
}

// Key
var key = "0123456789012345"

const NONCESIZE = 12

func Encrypt(plainString string) []byte {

	textByte := []byte(plainString)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, NONCESIZE)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, textByte, nil)
	return ciphertext
}

func Decrypt(encryptedString string) (decryptedString string) {

	decodeString, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, NONCESIZE)

	fmt.Println("hi" + string(nonce))

	plaintext, err := aesgcm.Open(nil, nonce, decodeString, nil)
	if plaintext != nil {
		fmt.Println("- Decryption success")
	}
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}
