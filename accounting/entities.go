package accounting

import (
	"time"

	"github.com/goblinlordx/go-test/currency"
	"github.com/google/uuid"
)

type Account struct {
	Id             uuid.UUID `gorm:"primaryKey"`
	CurrencySymbol currency.CurrencySymbol
	Name           string
}

type Transaction struct {
	Id         uuid.UUID `gorm:"primaryKey,index:idx_tx_at"`
	AccountId  uuid.UUID
	Account    Account
	At         time.Time `gorm:"index:idx_tx_at,priority:1,sort:desc"`
	Payee      string
	Amount     int
	Note       string
	CategoryId uuid.UUID
	Category   Category
}

type Category struct {
	Id    uuid.UUID `gorm:"primaryKey"`
	Value string    `gorm:"index:unq_tx_tag,unique"`
}
