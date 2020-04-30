package user


import (
	"context"
	"database/sql"
	"errors"
	"github.com/ecash/domain/infra"
	repo "github.com/ecash/domain/repository"
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
func FindByActivationURL(key string) (int64, bool, bool, error) {
	log.Println("Inside FindByActivationURL")
	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		return 0, false, false, errors.New(repo.MessageMap["C001"])
	}
	defer db.Close()
	var rid int64
	var enable,activation bool
	tsql := `SELECT ru_id, ru_enable_flag, ru_activation_flag 
               FROM RegisteredUser 
               WHERE ru_activationURL = @ActivationKey`

	rows, _ := db.QueryContext(ctx, tsql, sql.Named("ActivationKey", key))
	if rows.Next() {
		rows.Scan(&rid,&enable,&activation)
	}
	if rid > 0 { // success
		return rid,enable,activation, nil
	}
	return rid, false, false, errors.New(repo.MessageMap["V003"])
}
