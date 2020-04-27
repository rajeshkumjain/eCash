package user

import (
	"github.com/ecash/domain/entity"
	repo "github.com/ecash/domain/repository/user"
	uc "github.com/ecash/usecase/user"
)

// UserRegistrationService : Service Component to complete the registration process
func UserRegistrationService(b []byte) (int64, error) {
	ur := entity.RegisteredUser{}

	// 1. Call Data Validator Use Case
	ur,err := uc.UserRegistrationValidator(b)
	if err != nil {
		//log.Println(err.Error())
		return 0,err
	}

	// 2. Write to the database
	id, err:= repo.InsertUser(&ur)
	if err != nil {
		//log.Println("Insert Error : ",err.Error())
		return 0,err
	}

	// 3. Send welcome email
	return id,nil

}
