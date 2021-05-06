package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/model"
	service "github.com/nansuri/gp-server/service"
	util "github.com/nansuri/gp-server/util"
)

// List all of User API
func ListAllUserAPI(router *mux.Router, prefix string) {
	router.HandleFunc("/"+prefix+"/ping", TestParseAndReturn).Methods("POST")
	router.HandleFunc("/"+prefix+"/token", GetToken).Methods("POST")
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
	var isSuccess bool
	var errorContext string

	// JSON Body decoder
	decodeRequest(w, r, &getTokenRequest)

	// Some Logic Here
	decryptedUserInfo := service.Decrypt(getTokenRequest.EncryptedUserInfo)
	if decryptedUserInfo == "" {
		isSuccess = false
		errorContext = "Invalid Encryption Info"
	} else {
		isSuccess = true
	}

	switch getTokenRequest.Scope {
	case "TESTRAILEXPORTER":
		util.InfoLogger.Println("Scope is TESTRAILEXPORTER with " + decryptedUserInfo)
	case "GENERAL":
		util.InfoLogger.Println("Scope is GENERAL with " + decryptedUserInfo)
	default:
		isSuccess = false
		errorContext = "Invalid Scope"
	}

	token := service.QueryTokenByUserInfoAndScope(getTokenRequest.EncryptedUserInfo, getTokenRequest.Scope)
	if token == "" && getTokenRequest.Scope != "" && isSuccess {
		util.InfoLogger.Println("Generating Token")
		token = service.GenerateTokenAndStore("userId", getTokenRequest.EncryptedUserInfo, getTokenRequest.Scope)
	}

	// send response
	getTokenResponse.Token = token
	getTokenResponse.Status = isSuccess
	getTokenResponse.Error = errorContext

	EncodeResponse(w, r, getTokenResponse)
}
