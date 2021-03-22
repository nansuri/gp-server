package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	config "github.com/nansuri/gp-server/config"
	model "github.com/nansuri/gp-server/model"
)

// List all of User API
func ListAllUserAPI(router *mux.Router) {
	router.HandleFunc("/getUserInfo", GetAllUserInfo).Methods("GET")
	router.HandleFunc("/insertUser", InsertUserInfo).Methods("POST")
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// InsertUserInfo = Insert User API
func InsertUserInfo(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	loc, _ := time.LoadLocation("Asia/Jakarta")

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	login_date := time.Now().UTC().In(loc)

	_, err = db.Exec("INSERT INTO user(first_name, last_name, login_date) VALUES(?, ?, ?)", first_name, last_name, login_date)

	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "System Error"
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
