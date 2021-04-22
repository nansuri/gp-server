package util

import (
	"database/sql"
	"log"

	"github.com/nansuri/gp-server/config"
)

func StoreUserInfoAndToken(userId string, encryptedUserInfo string, scope string, token string) sql.Result {

	db := config.Connect()
	defer db.Close()

	// db Execution
	result, err := db.Exec("INSERT INTO user_token(user_id, encrypted_user_info, scope, token) VALUES(?, ?, ?, ?)", userId, encryptedUserInfo, scope, token)
	if err != nil {
		log.Panic(err.Error)
	}
	return result
}

func QueryTokenByUserInfoAndScope(encryptedUserInfo string, scope string) string {

	var token string

	db := config.Connect()
	defer db.Close()

	err := db.QueryRow("SELECT token FROM user_token WHERE encrypted_user_info=? AND scope=?", encryptedUserInfo, scope)
	switch err := err.Scan(&token); err {
	case sql.ErrNoRows:
		token = ""
	case nil:
		return token
	default:
		panic(err)
	}

	return token
}

func QueryUserInfoByTokenAndScope(token string, scope string) string {

	var encryptedUserInfo string

	db := config.Connect()
	defer db.Close()

	err := db.QueryRow("SELECT encryptedUserInfo FROM user_token WHERE token=? AND scope=?", token, scope)
	switch err := err.Scan(&token); err {
	case sql.ErrNoRows:
		encryptedUserInfo = ""
	case nil:
		return encryptedUserInfo
	default:
		panic(err)
	}

	return encryptedUserInfo
}
