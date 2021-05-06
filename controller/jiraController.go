package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
	service "github.com/nansuri/gp-server/service"
	util "github.com/nansuri/gp-server/util"
)

// List all of User API
func JiraBridgeAPI(router *mux.Router, prefix string) {
	router.HandleFunc("/"+prefix+"/create", CreateJiraIssue).Methods("POST")
	router.HandleFunc("/"+prefix+"/account", GetAccountIdByEmailAPI).Methods("GET")
}

// Test json request body
func CreateJiraIssue(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var token string
	var assignee string
	var request model.JiraRequest
	var response model.JiraResult

	// JSON Body decoder
	decodeRequest(w, r, &request)

	// Logic
	switch request.Project {
	case "MEMO":
		token = config.DingMember
	case "RSO":
		token = ""
	case "ACO":
		token = ""
	case "MPO":
		token = config.DingMerchantPortal
	default:
		util.WarningLogger.Println("Ding token not provided")
	}

	response.Key, response.Error = service.CreateJiraIssue(request, assignee)
	if request.Priority == "Blocker" || request.Priority == "Critical" {
		service.SendNotification(token, request, response.Key)
	}

	// Assemble the response
	if response.Error == "" {
		response.Status = true
	} else {
		response.Status = false
	}

	// Send response
	EncodeResponse(w, r, response)
}

// Test json request body
func GetAccountIdByEmailAPI(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var request model.GeneralRequest
	var response model.GeneralResponse

	// JSON Body decoder
	decodeRequest(w, r, &request)

	// Logic
	response.DataOutput = service.GetAccountIdByEmail(request.DataInput)

	// Send response
	EncodeResponse(w, r, response)
}
