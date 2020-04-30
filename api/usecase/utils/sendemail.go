package utils

import (
	"errors"
	utls "github.com/ecash/usecase"
	"net/smtp"
)

// smtpServer: data to smtp server.
type smtpServer struct {
	host string
	port string
}

// Address :  URI to smtp server.
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// SendEmail : Send the welcome email with activation code and security code
func SendEmail(e string, user string, hash string, scode string) error {
	// need to change my own email id
	from := "pointsofinterest2019@gmail.com"
	password := "Poi@123#"
	to := []string{e}

	message := []byte("Subject: Welcome - Activate Your Account\r\n" + "Dear " + user + `, 

Thank you for signing up to the App.  

To complete your sign up please click the link below and enter the following security code to activate the account.

Step 1 Click the link below 

http://localhost:8081/api/user/activate/?key=` + hash + `

Step 2 : Your security code is   [ ` + scode + ` ] 

Enter this code into the sign up page of the app

Thank you 
Rajesh Jain
`)

	// smtp server configuration
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message, Authentication
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// sending email
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		//log.Println("Error: in sending out the email :", err)
		return errors.New(utls.SrvErrorMap["SRVER0004"] + " " + err.Error())
	}
	return nil
}
