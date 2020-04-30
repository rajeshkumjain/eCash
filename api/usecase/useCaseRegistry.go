package usecase

import "github.com/ecash/domain/entity"
// NEED TO THINK AND COMPLETE THE WORK HERE
var UserRegistry interface{
	UserRegistrationValidator(data []byte) (entity.RegisteredUser, error)
}

