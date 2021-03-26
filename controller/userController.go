package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	config "github.com/nansuri/gp-server/config"
	userModel "github.com/nansuri/gp-server/model"
)

// List all of User API
func ListAllUserAPI(router *mux.Router) {
	router.HandleFunc("/getUserInfo", GetAllUserInfo).Methods("GET")
	router.HandleFunc("/insertUser", InsertUserInfo).Methods("POST")
	router.HandleFunc("/pingUser", TestParseAndReturn).Methods("POST")
	router.HandleFunc("/getToken", GetToken).Methods("POST")
}

// Get all user Info
func GetAllUserInfo(w http.ResponseWriter, r *http.Request) {
	var user userModel.User
	var response userModel.Response
	var arrUser []userModel.User

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

// InsertUserInfo = Insert User API
func InsertUserInfo(w http.ResponseWriter, r *http.Request) {

	var request userModel.User
	var response userModel.ResponseInsert

	// Init db
	db := config.Connect()
	defer db.Close()

	// decode request
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

	// set time
	loc, _ := time.LoadLocation("Asia/Jakarta")
	login_date := time.Now().UTC().In(loc)

	// db Execution
	_, err = db.Exec("INSERT INTO user(first_name, last_name, login_date) VALUES(?, ?, ?)", request.FirstName, request.LastName, login_date)

	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "System Error"
		EncodeResponse(w, r, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database of " + request.FirstName + "\n")

	EncodeResponse(w, r, response)
}

// Test json request body
func TestParseAndReturn(w http.ResponseWriter, r *http.Request) {

	// Define your request and response data struct here
	var user userModel.User
	var response userModel.UserResponse

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
	var getTokenRequest userModel.GetTokenRequest
	var getTokenResponse userModel.GetTokenResponse
	var decryptedMessage string

	// JSON Body decoder
	err := decodeJSONBody(w, r, &getTokenRequest)
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

	// securityUtil.GenerateKeyPair()
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// send response
	getTokenResponse.Token = decryptedMessage
	EncodeResponse(w, r, getTokenResponse)
}
