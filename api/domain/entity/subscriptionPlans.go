package entity

// SubscriptionPlans :
type SubscriptionPlans struct {
	ID           int     `json:"sp_id"`
	Name         string  `json:"name"`
	Desc         string  `json:"description, omitempty"`
	PeriodInDays int  `json:"duration_in_days"`
	Currency     string  `json:"currency"`
	Amount       float64 `json:"amount, omitempty"`
	GST          float64 `json:"gst_amount, omitempty"`
	GrossAmount  float64 `json:"gross_amount, omitempty"`
	Free         bool    `json:"free, omitempty"`
	// common fields to be part of all entities
	EnableFlag        bool   `json:"enable, omitempty"`
	ActiveFlag        bool   `json:"activeflag, omitempty"` // default to be active flag
	DateCreatedOn     string `json:"createdon, omitempty"`
	DateLastUpdatedOn string `json:"lastupdatedon, omitempty"`
}

/*
MS SQL
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
	sp_free              bit    ,
	cf_date_created_on   varchar(100)    ,
	cf_active_flag       bit    ,
	cf_date_last_updated_on varchar(100)    ,
	CONSTRAINT Pk_SubcriptionPlans_sp_id PRIMARY KEY  ( sp_id )
 );

*/
