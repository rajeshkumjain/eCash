package models

import "time"

// - let me 1st get things working with what was working - then would make changes in it

// RegisteredUser : User Login Table - created at the time of registration
type RegisteredUser struct {
	ID int

	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Surname    string `json:"surname"`

	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	EmailCode  string `json:"emailcode"`
	MobileCode string `json:"mobilecode"`

	AgreeTermsCondition  bool `json:"iagree"`
	AgreeSendNewsletters bool `json:"iagreenewsletter"`

	RegistrationSource string `json:"source"`

	SocialPluginID    int    `json:"plugid"`
	SocialLoginPlugin string `json:"socialsigninname"`

	EnableFlag bool `json:"enable"`
	ActiveFlag bool `json:"activeflag"`

	DateCreatedOn     time.Time `json:"createdon"`
	DateLastUpdatedOn time.Time `json:"lastupdatedon"`
}

/* SQL Structure
CREATE TABLE [dbo].[RegisteredUser](

	*	[ru_id] [int] IDENTITY(1,1) NOT NULL,

*	[ru_firstname] [varchar](150) NOT NULL,
*	[ru_middlename] [varchar](150) NULL,
*	[ru_surname] [varchar](150) NULL,

*	[ru_email] [varchar](200) NOT NULL,
*	[ru_mobile] [varchar](50) NULL,
*	[ru_password] [varchar](250) NULL,
*	[ru_emailVerificationCode] [varchar](10) NULL,
*	[ru_mobileVerificationCode] [varchar](10) NULL,

*	[ru_i_agree_t_n_c] [bit] NOT NULL,
*	[ru_i_agree_newsletter] [bit] NOT NULL,

*	[ru_registration_source] [varchar](100) NULL,

*	[ru_third_party_signin_id] [varchar](100) NULL,
*	[ru_third_party_signin_name] [varchar](100) NULL,

*	[ru_enable_flag] [bit] NOT NULL,
*	[cf_active_flag] [bit] NULL,

*	[cf_date_created_on] [varchar](100) NULL,
*	[cf_date_last_updated_on] [varchar](100) NULL,

	[cf_created_by_id] [int] NULL,
	[cf_last_updated_by_id] [int] NULL,
*/
