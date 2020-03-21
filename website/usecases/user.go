package usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"../models"
	"golang.org/x/crypto/bcrypt"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

/*
	1. Gather the data submited from the frontend form
	2. format it properly
	3. Create as JSON Packet
	4. Call the API Server
	5. Process the Response from the API Service
*/
// RegisterNewUser : post user data for registration - What this function should do?
func RegisterNewUser(w http.ResponseWriter, r *http.Request) {

	ru := models.RegisteredUser{}
	if r.Method == "POST" {
		// name
		ru.Firstname = r.FormValue("FirstName")
		ru.Middlename = r.FormValue("MiddleName")
		ru.Surname = r.FormValue("Surname")

		ru.Email = r.FormValue("UserEmail")
		ru.Mobile = r.FormValue("Mobile")
		// todo : unique email validation
		// todo (DONE) : password encryption - used golang crypto/bcryp - simple - default settings
		//hash, _ := HashPassword(r.FormValue("UserPassword"))
		ru.Password = r.FormValue("UserPassword")
		ru.EnableFlag = true // default is Enabled
		if r.FormValue("IAgreeTC") == "on" {
			ru.AgreeTermsCondition = true
		}
		if r.FormValue("IAgreeNewsletter") == "on" {
			ru.AgreeSendNewsletters = true
		}
		// todo (DONE) : email code generation - || Random function used here to generate 6 digit code
		rand.Seed(time.Now().UnixNano())
		ru.EmailCode = "99999" //RandomString(6)
		ru.MobileCode = "99999"
		ru.SocialLoginPlugin = "Blank"
		ru.SocialPluginID = 99999
		ru.RegistrationSource = "Website"
	}

	j, err := json.Marshal(ru)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(j))
	fmt.Println(j)
	resp, err := http.Post("http://localhost:8081/u/api/register", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	http.Redirect(w, r, "/", 301)

}

// RandomString : Function to generate Randam String
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// HashPassword : used for encrypting the password with default settings - this is just start the process.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
