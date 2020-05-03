package user

import (
	"encoding/json"
	"github.com/ecash/router"
	srv "github.com/ecash/services/user"
	"io/ioutil"
	"log"
	"net/http"
)
// UserLogin: User Authentication Function
func UserLogin(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"

	log.Println("*** Inside Login / Authentication *** \n")
	//log.Println(r.Header)
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

	log.Println(b)
	id,sk, err := srv.UserLoginService(b,r.Header.Get("User-Agent"))

	log.Println("user id", id)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		res := router.Response{}
		res.Error = "ERROR"
		res.ErrorCode = "SRVERR"
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := router.Response{}
	res.Error = "NONE"
	res.ErrorCode = "S003"
	res.Message = router.ErrorMap["S003"]
	res.SessionToken = sk
	json.NewEncoder(w).Encode(res)
	log.Println("*** Success *** \n", router.ErrorMap["S003"])
	return

}

