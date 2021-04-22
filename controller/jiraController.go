package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
	util "github.com/nansuri/gp-server/service"
)

// List all of User API
func JiraBridgeAPI(router *mux.Router, prefix string) {
	router.HandleFunc("/"+prefix+"/create", CreateJiraIssue).Methods("POST")
}

// Test json request body
func CreateJiraIssue(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var token string
	var request model.JiraRequest
	var response model.JiraResult

	// JSON Body decoder
	decodeRequest(w, r, &request)

	// Logic
	response.Key, response.Error = util.CreateJiraIssue(request)

	switch request.Project {
	case "MEMO":
		token = config.DingMember
	default:
		fmt.Println("\nInvalid ding token")
	}

	util.SendNotification(token, request, response.Key)

	// Assemble the response
	if response.Error == "" {
		response.Status = true
	} else {
		response.Status = false
	}

	// Send response
	EncodeResponse(w, r, response)
}
