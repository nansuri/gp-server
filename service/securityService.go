package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"

	"github.com/nansuri/gp-server/config"
)

type RSAKeyPair struct {
	Private rsa.PrivateKey
	Public  rsa.PublicKey
}

// Key
var key = config.RSAKey

const NONCESIZE = 12

// const defaultTokenLength = 12

// Use this to generate a token
func GenerateSecureToken() string {
	b := make([]byte, config.DefaultTokenLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GenerateTokenAndStore(userId string, encryptedUserInfo string, scope string) string {
	var token = GenerateSecureToken()
	StoreUserInfoAndToken(userId, encryptedUserInfo, scope, token)
	return token
}

/**
/ Used for encrypting a data
**/

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

/**
/ Used for decrypting a data
**/

func Decrypt(encryptedString string) (decryptedString string) {

	decodeString, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		// fmt.Println("!! - Decryption Failed")
		return ""
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
		panic(err.Error())
	}

	nonce := make([]byte, NONCESIZE)

	plaintext, err := aesgcm.Open(nil, nonce, decodeString, nil)
	if plaintext != nil {
		// fmt.Println("OK - Decryption Success")
	}
	if err != nil {
		return ""
		panic(err.Error())
	}
	return string(plaintext)
}
