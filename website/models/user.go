package models

// - let me 1st get things working with what was working - then would make changes in it

// RegisteredUser : User Login Table - created at the time of registration
type RegisteredUser struct {
	ID int

	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`

	Password    string `json:"password"`
	OldPassword string `json:"oldpassword"`
	EmailCode   string `json:"emailcode"`
	MobileCode  string `json:"mobilecode"`

	AgreeTermsCondition  bool `json:"iagree"`
	AgreeSendNewsletters bool `json:"iagreenewsletter"`

	RegistrationSource string `json:"source"`

	SocialPluginID int `json:"plugid"`

	SocialLoginPlugin string `json:"socialsigninname"`

	EnableFlag bool `json:"enable"`
	ActiveFlag bool `json:"activeflag"`
}
