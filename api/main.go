package main

import (
	user "github.com/ecash/router/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/u/api/register", user.UserRegisteration)

	log.Println("Starting eCASH API Server at : http://localhost:8081...")
	log.Fatal( http.ListenAndServe(":8081", router) )
	log.Println("Started...")

}

