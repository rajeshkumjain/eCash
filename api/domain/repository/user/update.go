package user

import (
	"database/sql"
	"context"
	"errors"
	"github.com/ecash/domain/infra"
	"log"
)

// InsertUser : Save Data in Entity RegisteredUser
func ActivateUser(id int, key string) error {

	log.Println("*** Activate User ***")
	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		log.Println("Error -Activate User")
		return  err
	}
	defer db.Close()

	tsql := `UPDATE RegisteredUser SET ru_activation_flag=1, ru_enable_flag=1
				WHERE ru_activationURL= @ActivationURL`
	result, err := db.ExecContext(ctx, tsql, sql.Named("ActivationURL", key))
	if err != nil {
		return errors.New("Unable to update the record")
	}
	log.Println("\nRows Affected : ")
	log.Println(result.RowsAffected())
	return nil

}
