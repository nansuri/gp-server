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
		// util.InfoLogger.Println("Encryption `" + plainString + "` FAILED\n" + err.Error())
	}
	nonce := make([]byte, NONCESIZE)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		// util.InfoLogger.Println("Encryption `" + plainString + "` FAILED\n" + err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, textByte, nil)
	if ciphertext != nil {
		// util.InfoLogger.Println("Encryption `" + plainString + "` SUCCESS")
	}
	return ciphertext
}

/**
/ Used for decrypting a data
**/

func Decrypt(encryptedString string) (decryptedString string) {

	decodeString, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		// util.ErrorLogger.Println("Decryption `" + encryptedString + "` FAILED\n" + err.Error())
		return ""
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		// util.ErrorLogger.Println("Decryption `" + encryptedString + "` FAILED\n" + err.Error())
		return ""
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		// util.ErrorLogger.Println("Decryption " + encryptedString + " FAILED\n" + err.Error())
		return ""
	}

	nonce := make([]byte, NONCESIZE)

	plaintext, err := aesgcm.Open(nil, nonce, decodeString, nil)
	if plaintext != nil {
		// util.InfoLogger.Println("Decryption `" + encryptedString + "` SUCCESS")
	}
	if err != nil {
		// util.ErrorLogger.Println("Decryption `" + encryptedString + "` FAILED\n" + err.Error())
		return ""
	}
	return string(plaintext)
}
