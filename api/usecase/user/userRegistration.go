package user

import (
	"encoding/json"
	"errors"
	"github.com/ecash/domain/entity"
	utls "github.com/ecash/usecase"
	"log"
	"time"
)


func UserRegistrationValidator(data []byte) (entity.RegisteredUser, error) {

	ur := entity.RegisteredUser{}
	// need to add option flag for debug and if on then only print this log
	// log.Println(string(data))
	err := json.Unmarshal([]byte(data), &ur)
	if err != nil {
		return ur, errors.New(utls.SrvErrorMap["SRVER0001"] + " " + err.Error())
	}

	isEmptyFirstName := utls.IsEmpty(ur.Firstname)
	isEmptyEmail := utls.IsEmpty(ur.Email)
	isEmptyMob := utls.IsEmpty(ur.Mobile)
	isEmptyPw := utls.IsEmpty(ur.Password)
	if isEmptyFirstName || isEmptyEmail || isEmptyMob || isEmptyPw {
		return ur, errors.New(utls.SrvErrorMap["SRVER0002"])
	}

	isValidEmail := utls.ValidateEmail(ur.Email)
	if !isValidEmail {
		return ur, errors.New(utls.SrvErrorMap["SRVER0003"])
	}
	ur.Password, _ = utls.HashActivationLink(ur.Password)
	ur.EmailCode = utls.RandomString(6)
	ur.MobileCode = utls.RandomString(6)
	ur.EnableFlag = false
	ur.MobilVerified = false
	ur.EmailVerified = false
	ur.ActiveFlag = true
	ur.ActivationURL, _ = utls.HashActivationLink(ur.Email + ur.Mobile)
	ur.ActivationFlag = false
	ur.DateCreatedOn = time.Now().Format("2006-01-02 15:04:05")
	ur.DateLastUpdatedOn = time.Now().Format("2006-01-02 15:04:05")

	log.Println("Completed : Data Validation & Setter :\n", ur)
	return ur, nil
}


// UserLoginValidator:
func UserLoginValidator(data []byte) (entity.RegisteredUser, error) {
	ur := entity.RegisteredUser{}
	err := json.Unmarshal([]byte(data), &ur)
	if err != nil {
		return ur, errors.New(utls.SrvErrorMap["SRVER0001"] + " " + err.Error())
	}

	log.Println("Login Data Packet :",ur)
	isEmptyEmail := utls.IsEmpty(ur.Email)
	isEmptyPw := utls.IsEmpty(ur.Password)
	log.Println("Flags : isEmpty email & password ? ", isEmptyEmail, isEmptyPw)
	if isEmptyEmail || isEmptyPw {
		return ur, errors.New(utls.SrvErrorMap["SRVER0002"])
	}
	log.Println("Login Data Packet Validated :\n", ur)
	return ur, nil
}
//
//func UserAuthenticationRules(usr *entity.RegisteredUser) (bool){
//
//}