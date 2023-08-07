package accounting

import (
	"context"
	"fmt"
	"time"

	"github.com/goblinlordx/go-test/db"
	"github.com/goblinlordx/go-test/db/modifier"
	db_tx "github.com/goblinlordx/go-test/db/tx"
	"github.com/goblinlordx/go-test/error_utils"
	"github.com/goblinlordx/go-test/repository"
	"gorm.io/gorm"
)

type AccountRepository struct {
	repository.BasicRepository[Account]
}

func ProvideAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (ar *AccountRepository) Search(ctx context.Context) ([]Account, error) {
	res := make([]Account, 0)
	err := db.WithDb(ctx, func(db *gorm.DB) error {
		q := db.Model(&Account{})
		m := modifier.GetModifiers(ctx)
		err := modifier.Apply(q, m...)
		if err != nil {
			return err
		}

		return q.Scan(&res).Error
	})

	return res, err
}

type CategoryRepository struct {
	repository.BasicRepository[Category]
}

func ProvideCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

type TransactionRepository struct {
	repository.BasicRepository[Transaction]
}

func ProvideTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

type aggregateRangeType string

const (
	RANGE_YEARLY  = aggregateRangeType("YEARLY")
	RANGE_MONTHLY = aggregateRangeType("MONTHLY")
	RANGE_DAILY   = aggregateRangeType("DAILY")
	RANGE_HOURLY  = aggregateRangeType("HOURLY")
)

type aggregateByRangeResult struct {
	At     string
	Amount int
}

func (r *TransactionRepository) AggregateByRange(ctx context.Context, art aggregateRangeType, start time.Time, end time.Time, mr ...modifier.Modifier) ([]aggregateByRangeResult, error) {
	results := []aggregateByRangeResult{}
	var err error
	db_tx.WithTx(ctx, func(db *gorm.DB) error {
		df := "%Y-%m-%dT%h:00:00Z"
		if art == RANGE_DAILY {
			df = "%Y-%m-%dT00:00:00Z"
		} else if art == RANGE_MONTHLY {
			df = "%Y-%m-00T00:00:00Z"
		} else if art == RANGE_YEARLY {
			df = "%Y-%m-%dT00:00:00Z"
		}

		q := db.Model(&Transaction{}).
			Select(fmt.Sprintf("strftime(\"%s\", at) as at, sum(amount) as amount", df)).
			Where("at >= ? AND at < ?", start, end).
			Group("strftime(\"%Y%m\", at)").
			Order("at")

		c := error_utils.NewErrorChain()
		c.Next(func() error { return modifier.Apply(db, mr...) })
		c.Next(func() error { return q.Scan(&results).Error })

		return c.Check()
	})

	return results, err
}
