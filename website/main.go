package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	uc "./usecases"
)

// const - definition
const (
	ConstTemplateFiles = "pages/*.gohtml"
	LogPath            = "./logs/"
)

var Tpl *template.Template
var Logfile os.File

// init: initialise the code
func init() {
	Tpl = template.Must(template.ParseGlob(ConstTemplateFiles))
	Logfile, err := os.OpenFile(LogPath+"logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ERROR: Error in opening file: logger.log ", err)
	}
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(Logfile)
	//defer Logfile.Close()
	log.Println("initialise logs")
}

// HomePage : call homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "Index", nil)
}

// SignUp : call User Registration Page
func SignUp(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "Registration", nil)
}

// main : main calling program
func main() {
	// - need to think about the controller
	http.Handle("/pages/", http.StripPrefix("/pages", http.FileServer(http.Dir("./pages"))))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/registration", SignUp)
	http.HandleFunc("/SignUp", uc.RegisterNewUser)
	startServer()
}

func startServer() {
	log.Println("Server started on : http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
