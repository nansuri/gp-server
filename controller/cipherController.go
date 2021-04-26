package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/model"
	util "github.com/nansuri/gp-server/service"
)

// List all of User API
func ListAllCipherAPI(router *mux.Router, prefix string) {
	router.HandleFunc("/"+prefix+"/encrypt", EncryptData).Methods("POST")
	router.HandleFunc("/"+prefix+"/decrypt", DecryptData).Methods("POST")
}

func EncryptData(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.CipherRequest
	var response model.CipherResponse

	// JSON Body decoder
	decodeRequest(w, r, &request)

	// Assemble the response
	response.EncryptedData = util.Encrypt(request.Data)

	// Send response
	EncodeResponse(w, r, response)
}

func DecryptData(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.CipherRequest
	var response model.CipherResponse

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
