package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/model"
	"github.com/nansuri/gp-server/util"
)

// List all of User API
func ListAllCipherAPI(router *mux.Router) {
	router.HandleFunc("/encryptData", EncryptData).Methods("POST")
	router.HandleFunc("/decryptData", DecryptData).Methods("POST")
}

func EncryptData(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.GeneralRequest
	var response model.GeneralResponse

	// JSON Body decoder
	decodeRequest(w, r, &request)

	// Assemble the response
	response.EncryptedData = util.Encrypt(request.Data)

	// Send response
	EncodeResponse(w, r, response)
}

func DecryptData(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.GeneralRequest
	var response model.GeneralResponse

	// JSON Body decoder
	decodeRequest(w, r, &request)
	token := decodeToken(w, r)

	data := util.Decrypt(request.DataByte)

	// Assemble the response
	response.Data = string(data)
	response.ExtendInfo = token

	// Send response
	EncodeResponse(w, r, response)
}
