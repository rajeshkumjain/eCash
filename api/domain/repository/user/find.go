package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ecash/domain/infra"
	repo "github.com/ecash/domain/repository"
	"github.com/ecash/domain/entity"
	"log"
)

// FindBy : find by email / mobile in RegisteredUser table
func FindBy(findby string, s string) (bool, error) {
	var ErrorMessage string
	//log.Println("findby :", findby, " string to find :",s)
	if s == "" {
		ErrorMessage = "Search String Can't be blank"
		log.Println(ErrorMessage)
		return true, nil
	}
	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		//log.Println(repo.MessageMap["C001"])
		return false, errors.New(repo.MessageMap["C001"])
	}
	defer db.Close()
	var rid int64
	if findby == "email" {
		tsql := "SELECT ru_id FROM RegisteredUser where ru_email = @Email"
		rows, _ := db.QueryContext(ctx, tsql, sql.Named("Email", s))
		if rows.Next() {
			rows.Scan(&rid)
			//log.Println("Found Record : ", rid)
		}
		ErrorMessage = repo.MessageMap["V001"]
		defer rows.Close()
	} else {
		if findby == "mobile" {
			tsql := "SELECT ru_id FROM RegisteredUser where ru_mobile = @Mobile"
			rows, _ := db.QueryContext(ctx, tsql, sql.Named("Mobile", s))
			if rows.Next() {
				rows.Scan(&rid)
				//	log.Println("Found Record : ", rid)
			}
			ErrorMessage = repo.MessageMap["V002"]
			defer rows.Close()
		}
	}

	if rid > 0 {
		log.Println("ROW Found with ID", rid, ErrorMessage)
		return true, nil
	}
	return false, nil
}

// FindByActivationURL : Find User Inforamtion based on Activation Key
func FindByActivationURL(key string) (int, bool, bool, error) {
	log.Println("Inside FindByActivationURL")
	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		return 0, false, false, errors.New(repo.MessageMap["C001"])
	}
	defer db.Close()
	var rid int
	var enable, activation bool
	tsql := `SELECT ru_id, ru_enable_flag, ru_activation_flag 
               FROM RegisteredUser 
               WHERE ru_activationURL = @ActivationKey`

	rows, _ := db.QueryContext(ctx, tsql, sql.Named("ActivationKey", key))
	if rows.Next() {
		rows.Scan(&rid, &enable, &activation)
	}
	if rid > 0 { // success
		return rid, enable, activation, nil
	}
	return rid, false, false, errors.New(repo.MessageMap["V003"])
}

// FindAuthentication : Authenticate the user by email & password. Fetching record
func FindAuthentication(e string, pw string) (entity.RegisteredUser, error) {

	var usr = entity.RegisteredUser{}
	log.Println("Inside User Authentication")
	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		return usr, errors.New(repo.MessageMap["C001"])
	}
	defer db.Close()

	tsql := `SELECT r.ru_id, ru_email, r.ru_enable_flag, r.cf_active_flag , 
                        r.ru_activation_flag, r.ru_is_mobile_verified, r.ru_is_email_verified 
                 FROM RegisteredUser r 
                WHERE r.ru_email = @Email AND r.ru_password=@Password`

	rows, _ := db.QueryContext(ctx,
		tsql,
		sql.Named("Email", e),
		sql.Named("Password", pw))

	if rows.Next() {
		rows.Scan(&usr.ID, &usr.Email, &usr.EnableFlag, &usr.ActiveFlag, &usr.ActivationFlag, &usr.MobilVerified, &usr.EmailVerified)
	}

	if usr.ID > 0 { // success
		return usr, nil
	} else {
		return usr, errors.New(repo.MessageMap["V005"])
	}

}
