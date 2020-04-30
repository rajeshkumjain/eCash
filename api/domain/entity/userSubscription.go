package entity

import "time"

// UserSubscription :
type UserSubscription struct {
	ID             int       `json:"s_id"`
	UserID         int       `json:"ru_id"`
	SPID           int       `json:"sp_id"`
	StartDate      string    `json:"start_date"`
	ValidUpto      string    `json:"valid_upto, omitempty"`
	Free           bool      `json:"free, omitempty"`
	GrossAmount    float64   `json:"gross_amount, omitempty"`
	PaidOn         time.Time `json:"paid_on, omitempty"`
	PaidThrough    string    `json:"paid_through, omitempty"`
	PaymentTransID string    `json:"payment_trans_id, omitempty"`
	PaidFlag       bool      `json:"paid"`
	// common fields to be part of all entities
	EnableFlag        bool   `json:"enable, omitempty"`
	ActiveFlag        bool   `json:"activeflag, omitempty"` // default to be active flag
	DateCreatedOn     string `json:"createdon, omitempty"`
	DateLastUpdatedOn string `json:"lastupdatedon, omitempty"`
}

/*
MS SQL
CREATE TABLE dbo.UserSubscription (
	s_id                 int NOT NULL   IDENTITY,
	ru_id                int NOT NULL   ,
	sp_id                int    ,
	us_start_date        date    ,
	us_valid_upto        date    ,
	us_free              bit    ,
	us_gross_amount      money    ,
	us_paid_on           date    ,
	us_paid_through      varchar(25)    ,
	us_payment_ID        varchar(100)    ,
	us_paid              bit    ,
	us_enable_flag       bit    ,
	cf_date_created_on   varchar(100)    ,
	cf_active_flag       bit    ,
	cf_date_last_updated_on varchar(100)    ,
*/
