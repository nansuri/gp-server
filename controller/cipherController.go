package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/model"
	"github.com/nansuri/gp-server/util"
)

// All variable
var key = "0123456789012345"

// List all of User API
func ListAllCipherAPI(router *mux.Router) {
	router.HandleFunc("/getKeyPair", GetPublicKey).Methods("GET")
	router.HandleFunc("/encryptData", EncryptData).Methods("POST")
	router.HandleFunc("/decryptData", DecryptData).Methods("POST")
}

type RSAKeyPair struct {
}

func GetPublicKey(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var response model.GeneralResponse
	var keyPair util.RSAKeyPair

	// Logic
	// keyPair.Private, keyPair.Public = util.GenerateKeyPair(128)

	// Assemble the response
	response.Data = keyPair.Public.N.String()

	// Send response
	EncodeResponse(w, r, response)
}

func EncryptData(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.GeneralRequest
	var response model.GeneralResponse

	// JSON Body decoder
	decodeRequest(w, r, &request)

	// Assemble the response
	response.EncryptedData = util.Encrypt([]byte(request.Data), key)
	response.Data = string(util.Encrypt([]byte(request.Data), key))

	// Send response
	EncodeResponse(w, r, response)
}

func DecryptData(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.GeneralRequest
	var response model.GeneralResponse

	// JSON Body decoder
	decodeRequest(w, r, &request)

	data := util.Decrypt([]byte(request.DataByte), key)

	// Assemble the response
	response.Data = string(data)

	// Send response
	EncodeResponse(w, r, response)
}
