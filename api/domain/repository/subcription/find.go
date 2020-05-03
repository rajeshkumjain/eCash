package subcription

import (
	"context"
	"errors"
	"github.com/ecash/domain/infra"
	repo "github.com/ecash/domain/repository"
	"github.com/ecash/domain/entity"
	"log"
)

// FindByPlan : Search Plans from the Subcription Information Table & return object
func FindByPlan(fb string) (entity.SubscriptionPlans, error) {
	sp := entity.SubscriptionPlans{}
	if fb == "" {
		log.Println(repo.MessageMap["V004"])
		return sp, errors.New(repo.MessageMap["V004"])
	}

	db, err := infra.NewDB()
	ctx := context.Background()
	if err != nil {
		return sp, errors.New(repo.MessageMap["C001"])
	}
	defer db.Close()

	if fb == "free" {
		tsql := `SELECT TOP 1 sp_id,  sp_duration_in_days, sp_enable, sp_free, cf_active_flag
                  FROM [SubscriptionPlans] 
                 WHERE cf_active_flag=1 AND sp_free=1 AND sp_enable=1
                 ORDER BY  [cf_date_created_on] DESC`
		rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
			return sp, errors.New(repo.MessageMap["C003"])
		}
		if rows.Next() {
			rows.Scan(&sp.ID, &sp.PeriodInDays, &sp.EnableFlag, &sp.Free, &sp.ActiveFlag)
		}
		defer rows.Close()
	}
	return sp, nil
}
