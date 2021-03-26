package model

type User struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	LastLogin string `form:"last_login" json:"last_login"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}

type UserResponse struct {
	Status    int    `json:"status"`
	FirstName string `json:"first_name"`
}

type ResponseInsert struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Get Token API
type GetTokenRequest struct {
	EncryptedUserInfo string `json:"encrypted_user_info"`
	Scope             string `json:"scope"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}

// Encryptor API
