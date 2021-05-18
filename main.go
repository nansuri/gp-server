package main

import (
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/nansuri/gp-server/controller"
	logger "github.com/sirupsen/logrus"
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
	logger.Info("==== Server Listen on port " + port + " ===")
	logger.Fatal(http.ListenAndServe(":"+port, router))
}
