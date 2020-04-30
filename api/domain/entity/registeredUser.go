package entity

// RegisteredUser : Basic user information capture at the time of registration process
type RegisteredUser struct {
	ID                   int    `json:"id, omitempty"`
	Firstname            string `json:"firstname"`
	Middlename           string `json:"middlename, omitempty"`
	Surname              string `json:"surname"`
	Email                string `json:"email"`
	Mobile               string `json:"mobile"`
	Password             string `json:"password"` // encrypted password
	ActivationURL        string `json:"activationURL"`
	ActivationFlag       bool `json:"activationflag"`
	EmailCode            string `json:"emailcode, omitempty"`
	MobileCode           string `json:"mobilecode, omitempty"`
	EmailVerified        bool   `json:"email_verified, omitempty"`
	MobilVerified        bool   `json:"mobile_verified, omitempty"`
	AgreeTermsCondition  bool   `json:"iagree"`
	AgreeSendNewsletters bool   `json:"iagreenewsletter"`
	RegistrationSource   string `json:"source, omitempty"` // should be listed one
	SocialPluginID       int    `json:"plugid, omitempty"`
	SocialLoginPlugin    string `json:"socialsigninname, omitempty"`
	// common fields to be part of all entities
	EnableFlag        bool   `json:"enable, omitempty"`
	ActiveFlag        bool   `json:"activeflag, omitempty"` // default to be active flag
	DateCreatedOn     string `json:"createdon, omitempty"`
	DateLastUpdatedOn string `json:"lastupdatedon, omitempty"`
}

/*
MS SQL
CREATE TABLE dbo.RegisteredUser (
	ru_id                int NOT NULL   IDENTITY,
	ru_firstname         varchar(150) NOT NULL   ,
	ru_middlename        varchar(150)    ,
	ru_surname           varchar(150)    ,
	ru_email             varchar(200) NOT NULL   ,
	ru_emailVerificationCode varchar(10)    ,
	ru_mobile            varchar(50)    ,
	ru_mobileVerificationCode varchar(10)    ,
	ru_password          varchar(250)    ,
	ru_enable_flag       bit NOT NULL   ,
	ru_i_agree_t_n_c     bit NOT NULL   ,
	ru_i_agree_newsletter bit NOT NULL   ,
	ru_third_party_signin_name varchar(100)    ,
	ru_third_party_signin_id varchar(100)    ,
	ru_registration_source varchar(100)    ,
	ru_activationURL     varchar(250)    ,
	ru_is_mobile_verified bit    ,
	ru_activation_flag   bit    ,
	ru_is_email_verified bit    ,
	cf_date_created_on   varchar(100)    ,
	cf_last_updated_by_id int    ,
	cf_created_by_id     int    ,
	cf_date_last_updated_on varchar(100)    ,
	cf_active_flag       bit    ,
	CONSTRAINT Pk_RegisteredUser_ru_id PRIMARY KEY  ( ru_id )
 );


CREATE TABLE dbo.SubscriptionPlans (
	sp_id                int NOT NULL   IDENTITY,
	sp_name              varchar(100)    ,
	sp_description       varchar(500)    ,
	sp_duration_in_days  int    ,
	sp_currency          char(2)    ,
	sp_amount            money    ,
	sp_gst_amount        money    ,
	sp_gross_amount      money    ,
	sp_enable            bit    ,
	sp_free_trail        bit    ,
	cf_date_created_on   varchar(100)    ,
	cf_created_by_id     int    ,
	cf_last_updated_by_id int    ,
	cf_active_flag       bit    ,
	cf_date_last_updated_on varchar(100)    ,
	CONSTRAINT Pk_SubcriptionPlans_sp_id PRIMARY KEY  ( sp_id )
 );


CREATE TABLE dbo.UserSubscription (
	s_id                 int NOT NULL   IDENTITY,
	ru_id                int NOT NULL   ,
	sp_id                int    ,
	us_start_date        date    ,
	us_valid_upto        date    ,
	us_free_trail        bit    ,
	us_gross_amount      money    ,
	us_paid_on           date    ,
	us_paid_through      varchar(25)    ,
	us_payment_ID        varchar(100)    ,
	us_paid_flag         bit    ,
	us_enable_flag       bit    ,
	cf_date_created_on   varchar(100)    ,
	cf_created_by_id     int    ,
	cf_last_updated_by_id int    ,
	cf_active_flag       bit    ,
	cf_date_last_updated_on varchar(100)    ,
	CONSTRAINT Pk_Subcription_s_id PRIMARY KEY  ( s_id ),
	CONSTRAINT Unq_Subcription_sp_id UNIQUE ( sp_id ) ,
	CONSTRAINT Unq_Subcription_ru_id UNIQUE ( ru_id )
 );

 */