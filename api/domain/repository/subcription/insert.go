package subcription

import (
	"database/sql"
	"errors"
	"github.com/ecash/domain/entity"
	"context"
	"github.com/ecash/domain/infra"
	repo "github.com/ecash/domain/repository"
	"log"

	"time"
)

// InsertUser : Save Data in Entity RegisteredUser
func InsertFreeUserPlan(plans *entity.SubscriptionPlans) (int64, error) {

	log.Println("*** Inside Inserting User Plan Function ***")

	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		//		log.Println(repo.MessageMap["C001"])
		return 0, err
	}
	defer db.Close()

	tsql := `INSERT INTO UserSubscription ( sp_id, us_start_date, us_valid_upto, 
					us_free, us_paid, us_enable_flag, 
                    cf_active_flag, cf_date_created_on,  cf_date_last_updated_on) 
            VALUES (@SPID, @StartDate, @EndDate, 
                    @Free, @IsPaid, @EnableFlag, 
                    @ActiveFlag, @DateCreatedOn, @DateLastUpdatedOn );SELECT convert(bigint, SCOPE_IDENTITY())`

	stmt, err := db.Prepare(tsql)
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
		sql.Named("SPID", plans.ID),
		sql.Named("StartDate", time.Now().Format("2006-01-02 15:04:05")),
		sql.Named("EndDate", time.Now().AddDate(0, 0, plans.PeriodInDays)),
		sql.Named("Free", true),
		sql.Named("IsPaid", true),
		sql.Named("EnableFlag", true),
		sql.Named("ActiveFlag", true),
		sql.Named("DateCreatedOn", time.Now().Format("2006-01-02 15:04:05")),
		sql.Named("DateLastUpdatedOn", time.Now().Format("2006-01-02 15:04:05")),
	)
	var subcriptionid int64
	err = row.Scan(&subcriptionid)
	if err != nil {
		//log.Println(repo.MessageMap["C002"], err.Error())
		return 0, errors.New(repo.MessageMap["C002"] + " " + err.Error())
	}
	log.Println("New User ID : ", subcriptionid)
	return subcriptionid, nil
}
