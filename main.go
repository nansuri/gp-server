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
	var port string = "3025"
	router := mux.NewRouter()

	controller.ListAllUserAPI(router)
	controller.ListAllCipherAPI(router)

	http.Handle("/", router)
	fmt.Println("==== Server Listen on port " + port + " ===")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
