package user

import (
	"errors"
	"github.com/ecash/domain/entity"
	repo "github.com/ecash/domain/repository/user"
	reposub "github.com/ecash/domain/repository/subcription"
	uc "github.com/ecash/usecase/user"
	ucutils "github.com/ecash/usecase/utils"
	"log"
)

// UserRegistrationService : Service Component to complete the registration process
func UserRegistrationService(data []byte) (int64, error) {

	ur := entity.RegisteredUser{}

	// 1. Call Data Validator Use Case
	ur, err := uc.UserRegistrationValidator(data)
	if err != nil {
		//log.Println(err.Error())
		return 0, err
	}

	// 2. Write to the database
	id, err := repo.InsertUser(&ur)
	if err != nil {
		//log.Println("Insert Error : ",err.Error())
		return 0, err
	}

	// 3. Send welcome email
	log.Println("Sending Welcome Email")
	err = ucutils.SendEmail(ur.Email, ur.Firstname+" "+ur.Surname, ur.ActivationURL, ur.EmailCode)
	if err != nil {
		return id, err
	}

	// Write to log also

	return id, nil

}

// UserActivationService: Service to Activate the user
func UserActivationService(e string) (int64, error) {
	// Step 1: What is the current Status of User against Activation Key
	id, enable, activation, err := repo.FindByActivationURL(e)
	if err != nil {
		return 0, err
	}
	log.Println("ID :", id, " already enabled ? ", enable, " already activated ? ", activation)
	if enable || activation {
		return 0, errors.New("Account Already Activated or Disabled")
	}

	// Step 2: If Not activated Yet - Activate the user
	err = repo.ActivateUser(id, e)
	if err != nil {
		return 0, err
	}

	// Step 3: Fetch Free Plan Information
	fp := entity.SubscriptionPlans{}
	fp, err = reposub.FindByPlan("free")
	if err != nil {

	}
	log.Println(fp)

	// Step 4: AssignFreeTrailPlan - after successful activation, system to assign free plan to the user
	var subid int64
	subid, err = reposub.InsertFreeUserPlan(&fp)
	if err != nil {

	}
	log.Println("Subcription Activated Successfully : id is : ", subid)

	return id, nil
}
