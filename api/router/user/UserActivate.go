package user

import (
	"fmt"
	"github.com/ecash/router"
	srv "github.com/ecash/services/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// UserActivate: Activate the user
func UserActivate(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	var vars = mux.Vars(r)
	var encrypKey = vars["key"]
	log.Println("*** Inside User Activation  *** \n", encrypKey)

	if encrypKey == "" || len(encrypKey) < 5 {
		log.Println(router.ErrorMap["E005"])
		fmt.Fprintf(w, "\n**************************************************\n")
		fmt.Fprintf(w, router.ErrorMap["E005"])
		fmt.Fprintf(w, "\n**************************************************")
		return
	}
	_, err := srv.UserActivationService(encrypKey)
	if err != nil {
		log.Println(router.ErrorMap["E005"])
		fmt.Fprintf(w, "\n**************************************************\n")
		fmt.Fprintf(w, router.ErrorMap["E005"])
		fmt.Fprintf(w, "\n**************************************************")
	} else {
		log.Println("Success")
		fmt.Fprintf(w, "\n**************************************************\n")
		fmt.Fprintf(w, router.ErrorMap["S002"])
		fmt.Fprintf(w, "\n**************************************************")
	}

}
