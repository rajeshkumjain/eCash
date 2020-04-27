package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	router "github.com/ecash/router"
	srv "github.com/ecash/services/user"
	)

// Register : Controller to manage the user registraion process
func UserRegisteration(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	log.Println("*** Inside Registration ***")

	header := r.Header.Get("Content-Type")
	if header != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		res := router.Response{}
		res.Error = "ERROR"
		res.ErrorCode = "E001"
		res.Message = router.ErrorMap["E001"]
		json.NewEncoder(w).Encode(res)
		log.Println("*** ERROR *** \n", router.ErrorMap["E001"])
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusPartialContent)
		w.Header().Set("Content-Type", "application/json")
		res := router.Response{}
		res.Error = "ERROR"
		res.ErrorCode = "E002"
		res.Message = router.ErrorMap["E002"]
		json.NewEncoder(w).Encode(res)
		log.Println("*** ERROR *** \n", router.ErrorMap["E002"])
		return
	}

	var id int64
	id, err = srv.UserRegistrationService(b)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		res := router.Response{}
		res.Error = "ERROR"
		res.ErrorCode = "SRVERR"
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		//log.Println("*** ERROR *** \n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := router.Response{}
	res.Error = "NONE"
	res.ErrorCode = "S001"
	res.Message = router.ErrorMap["S001"]
	json.NewEncoder(w).Encode(res)
	log.Println("*** Success *** \n", router.ErrorMap["S001"], " NEW ID", id)
	return

}

