package user

import (
	"errors"
	"fmt"
	"github.com/ecash/domain/entity"
	reposub "github.com/ecash/domain/repository/subcription"
	repo "github.com/ecash/domain/repository/user"
	uc "github.com/ecash/usecase/user"
	ucutils "github.com/ecash/usecase/utils"

	"log"
	"crypto/rand"
)

// UserRegistrationService : Service Component to complete the registration process
func UserRegistrationService(data []byte) (int, error) {

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
func UserActivationService(e string) (int, error) {
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
	sp := entity.SubscriptionPlans{}
	sp, err = reposub.FindByPlan("free")
	if err != nil {
		return 0, err
	}

	// Step 4: AssignFreeTrailPlan - after successful activation, system to assign free plan to the user
	var subid int
	subid, err = reposub.InsertFreeUserPlan(&sp, id)
	if err != nil {
		return 0, err
	}
	log.Println("Subscription Activated Successfully : id is : ", subid)

	return id, nil
}

// UserLoginService: User Authentication Services
func UserLoginService(data []byte, agentinfo string) (int, string, error) {
	var sessionkey string
	log.Println("Inside Login Services")
//	log.Println("User Agent :", agentinfo)
	ur := entity.RegisteredUser{}
	cust := entity.RegisteredUser{}

	ur, err := uc.UserLoginValidator(data)
	if err != nil {
		return 0, sessionkey, err
	}

	// FindAuthentication(e string, pw string)
	cust, err = repo.FindAuthentication(ur.Email, ur.Password)
	if err != nil {
		/* log into failed cases
		ur.Email, ur.Password, Device Information, GeoLocation, TimeStamp
		 */
		log.Println("Login Error : ", ur.Email, ur.Password, agentinfo)
		return 0,sessionkey, err
	}

	log.Println("ID : ", cust.ID)
	log.Println("Email : ", cust.Email)
	log.Println("Enable ?  : ", cust.EnableFlag)
	log.Println("Active ? : ", cust.ActiveFlag)
	log.Println("Activation ? : ", cust.ActivationFlag)
	log.Println("Is Mobil Verified  ? : ", cust.MobilVerified)
	log.Println("Is Email Verified : ? ", cust.EmailVerified)

	// Now call user authentication rules before returning back.
	// authentication rules (basic)

	/* THESE NEEDS TO BE MOVED TO USE CASE SERVICES (BASIC)
	IF Authorized AND Activation flag is False - “Your account is not activated yet. Try again after activating your account”.
	If Authorized AND Account Enable is False : “Your is disabled, contact customer support for further details”

	**** still need to implement ***
	If Authorized AND Temporarily Blocked is True : “Your account is temporarily Suspended, try after 30 mins”
	If Authorized AND Subscription Activation Status is False : “Your Subscription has expired”

	ELSE Success
	*/

	if (cust.ActivationFlag) && !(cust.EnableFlag) && (cust.ActiveFlag) {
		return 0, sessionkey, errors.New("Your account has been disabled, contact customer support for further details ")
	}
	if !(cust.ActivationFlag && cust.EnableFlag && cust.ActiveFlag) {
		return 0, sessionkey, errors.New("Your account is not activated yet. Try again after activating your account ")
	}

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Println(err)
	}
	sessionkey = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
//	fmt.Println("session key :",sessionkey)

	/*  login history to recorded
		cust.ID, cust.Email, SessionToken, LoginID, Subscription Plan ID & Subscription Status
	 */

	return cust.ID, sessionkey, nil

}
