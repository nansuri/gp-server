package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/model"
	util "github.com/nansuri/gp-server/util"
)

// List all of User API
func JiraBridgeAPI(router *mux.Router) {
	router.HandleFunc("/createJiraIssue", CreateJiraIssue).Methods("GET")
}

// Test json request body
func CreateJiraIssue(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.JiraRequest
	var response model.JiraResult

	// JSON Body decoder
	err := decodeJSONBody(w, r, &request)
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

	// Logic
	response.Key = util.CreateIssue(request)

	// Assemble the response
	response.Status = true

	// Send response
	EncodeResponse(w, r, response)
}
