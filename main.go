package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	userController "github.com/nansuri/gp-server/controller"
)

func main() {
	router := mux.NewRouter()
	userController.ListAllUserAPI(router)
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
