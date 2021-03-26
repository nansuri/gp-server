package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/config"
	"github.com/nansuri/gp-server/model"
	"github.com/nansuri/gp-server/util"
)

// List all of User API
func ListAllUserAPI(router *mux.Router) {
	router.HandleFunc("/getUserInfo", GetAllUserInfo).Methods("GET")
	router.HandleFunc("/pingUser", TestParseAndReturn).Methods("POST")
	router.HandleFunc("/getToken", GetToken).Methods("POST")
}

// Get all user Info
func GetAllUserInfo(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.Response
	var arrUser []model.User

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.LastLogin)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrUser = append(arrUser, user)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrUser

	EncodeResponse(w, r, response)
}

// Test json request body
func TestParseAndReturn(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var user model.User
	var response model.UserResponse

	// JSON Body decoder
	err := decodeJSONBody(w, r, &user)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.errorMessage, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Assemble the response
	response.Status = 200
	response.FirstName = user.FirstName + "_test"

	// Send response
	EncodeResponse(w, r, response)
}

func GetToken(w http.ResponseWriter, r *http.Request) {

	// Declared the request and response struct
	var getTokenRequest model.GetTokenRequest
	var getTokenResponse model.GetTokenResponse
	var 

	// JSON Body decoder
	decodeRequest(w, r, &getTokenRequest)

	// Some Logic Here
	decryptedUserInfo := util.Decrypt(getTokenRequest.EncryptedUserInfo)

	switch getTokenRequest.Scope {
	case "TESTRAILEXPORTER":
		fmt.Println("Scope is TESTRAILEXPORTER with " + decryptedUserInfo)
	default:
		fmt.Println("Scope is Other")
	}

	token := util.QueryTokenByUserInfoAndScope(getTokenRequest.EncryptedUserInfo, getTokenRequest.Scope)
	if token == "" {
		fmt.Println("Generate token now")
		token = util.GenerateTokenAndStore("userId", getTokenRequest.EncryptedUserInfo, getTokenRequest.Scope)
	}

	// send response
	getTokenResponse.Token = token
	EncodeResponse(w, r, getTokenResponse)
}
