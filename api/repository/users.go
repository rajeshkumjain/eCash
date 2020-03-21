package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"../models"
	//repo "../repository"
	"golang.org/x/crypto/bcrypt"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// InsertUser : create new user
func InsertUser(u *models.RegisteredUser) (int64, error) {
	var tsql string

	hash, _ := HashPassword(u.Password)
	u.Password = hash
	u.EmailCode = RandomString(6)
	u.DateLastUpdatedOn = time.Now()
	u.DateCreatedOn = time.Now()
	u.ActiveFlag = true

	db, err := NewDB()
	ctx := context.Background()
	if err != nil {
		log.Println("DB connection failed")
		return 0, err
	}
	defer db.Close()

	// checks - duplicate email id ,  mobil #
	emFound, err := FindBy("email", strings.ToLower(strings.TrimSpace(u.Email)))
	mFound, err := FindBy("mobile", u.Mobile)
	if emFound || mFound {
		if emFound {
			fmt.Println("Email Already Registered - duplicate ")
		}
		if mFound {
			fmt.Println("Mobil # Already Registered - duplicate")
		}
		return 0, err
	}

	tsql = `INSERT INTO RegisteredUser ( ru_firstname, ru_middlename, ru_surname, ru_email, ru_mobile, ru_password, ru_emailVerificationCode, 
		ru_mobileVerificationCode, ru_i_agree_t_n_c, ru_i_agree_newsletter, ru_registration_source, 
		ru_third_party_signin_id, ru_third_party_signin_name, ru_enable_flag, cf_active_flag, 
		cf_date_created_on, cf_date_last_updated_on ) VALUES (@Firstname, @Middlename, @Surname, @Email, @Mobile, @Password, @EmailCode,
			@MobileCode, @AgreeTermsCondition, @AgreeSendNewsletters, @RegistrationSource,
			@SocialPluginID, @SocialLoginPlugin, @EnableFlag, @ActiveFlag,
			@DateCreatedOn, @DateLastUpdatedOn);SELECT convert(bigint, SCOPE_IDENTITY())`

	stmt, err := db.Prepare(tsql)
	defer stmt.Close()

	fmt.Println(strings.ToLower(strings.TrimSpace(u.Email)))
	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Firstname", u.Firstname),
		sql.Named("Surname", u.Surname),
		sql.Named("Middlename", u.Middlename),
		sql.Named("Email", strings.ToLower(strings.TrimSpace(u.Email))),
		sql.Named("Mobile", u.Mobile),
		sql.Named("Password", u.Password),
		sql.Named("EmailCode", u.EmailCode),
		sql.Named("MobileCode", u.MobileCode),
		sql.Named("AgreeTermsCondition", u.AgreeTermsCondition),
		sql.Named("AgreeSendNewsletters", u.AgreeSendNewsletters),
		sql.Named("RegistrationSource", u.RegistrationSource),
		sql.Named("SocialPluginID", u.SocialPluginID),
		sql.Named("SocialLoginPlugin", u.SocialLoginPlugin),
		sql.Named("EnableFlag", u.EnableFlag),
		sql.Named("ActiveFlag", u.ActiveFlag),
		sql.Named("DateCreatedOn", u.DateCreatedOn),
		sql.Named("DateLastUpdatedOn", u.DateLastUpdatedOn))

	var uID int64
	err = row.Scan(&uID)
	if err != nil {
		log.Println("DB Write Failed : ", err.Error())
		return 0, err
	}
	fmt.Println("New User ID : ", uID)
	return 0, nil
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

// FindBy : find by email / mobile in RegisteredUser table
func FindBy(f string, v string) (bool, error) {

	//f fieldname, v fieldvalue
	var ErrorMessage string
	db, err := NewDB()
	ctx := context.Background()
	var Count int

	if err != nil {
		log.Println("DB connection failed")
		return false, err
	}
	defer db.Close()

	if f == "email" {
		tsql := "SELECT ru_id FROM RegisteredUser where ru_email = @Email"
		rows, _ := db.QueryContext(ctx, tsql, sql.Named("Email", v))
		ErrorMessage = "Dupliate email"
		defer rows.Close()
		for rows.Next() {
			Count++
		}
	} else {
		if f == "mobile" {
			tsql := "SELECT ru_id FROM RegisteredUser where ru_mobile = @Mobile"
			rows, _ := db.QueryContext(ctx, tsql, sql.Named("Mobile", v))
			ErrorMessage = "Dupliate Mobile"
			defer rows.Close()
			for rows.Next() {
				Count++
			}
		}
	}

	if Count > 0 {
		log.Println(ErrorMessage)
		return true, nil
	}
	return false, nil
}
