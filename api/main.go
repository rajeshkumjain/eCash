package main

import (
	"fmt"
	"log"
	"net/http"

	ctrl "./controller"
)

// type Database struct {
// 	*sql.DB
// }

func main() {

	APIserver := http.NewServeMux()
	APIserver.HandleFunc("/u/api/register", ctrl.Register)

	log.Println("starting api server at 8081...")
	// log.Println("Connecting database...")
	// Database, err := repo.NewDB()
	// if err != nil {
	// 	log.Println("DB connection failed")
	// }
	// fmt.Println("Connected :", Database)

	server := &http.Server{
		Addr:    ":8081",
		Handler: APIserver,
	}
	server.ListenAndServe()
	fmt.Println(APIserver)

}

// func startServer() {

// }
