package main

import (
	user "github.com/ecash/router/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/registration", user.UserRegisteration)
	router.HandleFunc("/api/user/activate/", user.UserActivate).Queries("key", "{key}")
	router.HandleFunc("/api/user/login", user.UserLogin)
	log.Println("Starting eCASH API Server at : http://localhost:8081...")
	log.Fatal( http.ListenAndServe(":8081", router) )

	log.Println("Started...")

}

