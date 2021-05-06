package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	controller "github.com/nansuri/gp-server/controller"
	util "github.com/nansuri/gp-server/util"
)

func main() {
	// Initialization
	var port string = "3025"
	router := mux.NewRouter()

	// List of all registered API
	controller.ListAllUserAPI(router, "user")
	controller.ListAllCipherAPI(router, "cipher")
	controller.JiraBridgeAPI(router, "jira")

	// Handler
	http.Handle("/", router)
	fmt.Println("==== Server Listen on port " + port + " ===")
	util.InfoLogger.Println("==== Server Listen on port " + port + " ===")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
