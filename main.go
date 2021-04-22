package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nansuri/gp-server/controller"
)

func main() {
	// Initialization
	var port string = "3025"
	router := mux.NewRouter()

	// List of all registered API
	controller.ListAllUserAPI(router, "user")
	controller.ListAllCipherAPI(router, "cipher")
	controller.JiraBridgeAPI(router)

	// Handler
	http.Handle("/", router)
	fmt.Println("==== Server Listen on port " + port + " ===")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
