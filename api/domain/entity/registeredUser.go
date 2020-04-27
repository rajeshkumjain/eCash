package entity

// RegisteredUser : Basic user information capture at the time of registration process
type RegisteredUser struct {
	ID int  `json:"id, omitempty"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename, omitempty"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"` // encrypted password
	EmailCode  string `json:"emailcode, omitempty"`
	MobileCode string `json:"mobilecode, omitempty"`
	EmailVerified bool `json:"email_verified, omitempty"`
	MobilVerified bool `json:"mobile_verified, omitempty"`
	AgreeTermsCondition  bool `json:"iagree"`
	AgreeSendNewsletters bool `json:"iagreenewsletter"`
	RegistrationSource string `json:"source, omitempty"` // should be listed one
	SocialPluginID    int    `json:"plugid, omitempty"`
	SocialLoginPlugin string `json:"socialsigninname, omitempty"`
	// common fields to be part of all entities
	EnableFlag bool `json:"enable, omitempty"`
	ActiveFlag bool `json:"activeflag, omitempty"` // default to be active flag
	DateCreatedOn     string `json:"createdon, omitempty"`
	DateLastUpdatedOn string `json:"lastupdatedon, omitempty"`
}

