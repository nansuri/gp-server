package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
	service "github.com/nansuri/gp-server/service"
	dbservice "github.com/nansuri/gp-server/service/database"
	logger "github.com/sirupsen/logrus"
)

// List all of User API
func JiraBridgeAPI(router *mux.Router, prefix string) {
	router.HandleFunc("/"+prefix+"/create", CreateJiraIssue).Methods("POST")
	router.HandleFunc("/"+prefix+"/account", GetAccountIdByEmailAPI).Methods("GET")
	router.HandleFunc("/"+prefix+"/fields", GetFieldInfo).Methods("GET")
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
	logger.WithFields(logger.Fields{"Project": request.Project, "Summary": request.Summary, "Reporter": request.Reporter}).Info("CreateJiraIssue")

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
		logger.WithFields(logger.Fields{"Project": request.Project}).Warn("Ding token not provided")
	}

	response.Key, response.Error = service.CreateJiraIssue(request, assignee)
	if request.Priority == "Blocker" || request.Priority == "Critical" || request.IsUrgent == "true" {
		service.SendNotification(token, request, response.Key)
	}

	// Assemble the response
	if response.Error == "" {
		response.Status = true
		logger.WithFields(logger.Fields{"Key": response.Key}).Info("Jira Issue Created")
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

func GetFieldInfo(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var response model.JiraFieldResult

	// Query param decoder
	field := getQueryParam(w, r, "name")

	// Logic
	if field == "JIRA_BOARD" || field == "TICKET_TYPE" || field == "TICKET_PRIORITY" {
		response.FieldName = field
		response.FieldItems = dbservice.DbQueryTicketInfo(field)

		if response.FieldItems != nil {
			response.Status = true
		}
	} else {
		response.Error = "Invalid field name"
	}

	// Send response
	EncodeResponse(w, r, response)
}
