package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	models "../models"
	repo "../repository"
)

// Register : save new user information into db
func Register(w http.ResponseWriter, r *http.Request) {
	res := models.ResponseResult{}
	h := r.Header.Get("Content-Type")

	if h != "application/json" {
		w.Header().Set("Content-Type", "application/json")
		res.Error = "ERROR"
		res.Message = "ERROR: Error in the header"
		json.NewEncoder(w).Encode(res)
		return

	} else {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
			w.Header().Set("Content-Type", "application/json")
			res.Error = "ERROR"
			res.Message = "ERROR : Unable to decode the body of content"
			json.NewEncoder(w).Encode(res)
			return

		} else {
			// decode json packet &
			user := models.RegisteredUser{}
			json.Unmarshal([]byte(b), &user)
			ID, ierr := repo.InsertUser(&user)

			if ierr != nil {
				log.Fatalln(err)
				w.Header().Set("Content-Type", "application/json")
				res.Error = "ERROR"
				res.Message = "ERROR : Unable to register user"
				json.NewEncoder(w).Encode(res)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				res.Error = "NONE"
				res.Message = "SUCCESS : User Registration (user id (" + string(ID) + ")"
				json.NewEncoder(w).Encode(res)
				return
			}
		}
	}
}
