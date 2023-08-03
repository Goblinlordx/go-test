package accounting

import (
	"strconv"
	"strings"
	"sync"
	"time"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/goblinlordx/go-test/currency"
	"github.com/google/uuid"
)

type TestDataset struct {
	Accounts     []Account
	Transactions []Transaction
	Categories   []Category
}

var payees = []string{}
var payeeOnce = sync.Once{}

func GetTestDatasetPayee() []string {
	payeeOnce.Do(func() {
		faker := gofakeit.New(901)
		for i := 0; i < 20; i++ {
			payees = append(payees, faker.Name())
		}
	})
	return payees
}

var accounts = []Account{}
var accountOnce = sync.Once{}

func GetTestDatasetAccounts() []Account {
	accountOnce.Do(func() {
		faker := gofakeit.New(101)
		for i := 0; i < 10; i++ {
			accounts = append(accounts, Account{
				Id:             uuid.MustParse(faker.UUID()),
				Name:           faker.Name(),
				CurrencySymbol: currency.CurrencySymbol("USD"),
			})
		}
	})
	return accounts
}

var cats = []Category{}
var catOnce = sync.Once{}

func GetTestDatasetCategories() []Category {
	catOnce.Do(func() {
		exist := make(map[string]bool)
		faker := gofakeit.New(102)
		for i := 0; i <= 10; i++ {
			name := faker.Noun()
			i := -1
			k := strings.Join([]string{name}, "-")
			for exist[k] {
				i++
				k = strings.Join([]string{name, strconv.Itoa(i)}, "-")
			}
			exist[k] = true
			cats = append(cats, Category{
				Id:    uuid.MustParse(faker.UUID()),
				Value: k,
			})
		}
	})
	return cats
}

var tx = []Transaction{}
var txOnce = sync.Once{}

func GetTestDatasetTransactions() []Transaction {
	txOnce.Do(func() {
		accounts := GetTestDatasetAccounts()
		payees := GetTestDatasetPayee()
		cats := GetTestDatasetCategories()
		faker := gofakeit.New(102)
		for i := 0; i < 10000; i++ {
			tx = append(tx, Transaction{
				Id:        uuid.MustParse(faker.UUID()),
				AccountId: accounts[faker.IntRange(0, len(accounts)-1)].Id,
				At: faker.DateRange(
					time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 07, 19, 23, 59, 0, 0, time.UTC),
				),
				Payee:      payees[faker.IntRange(0, len(payees)-1)],
				Amount:     faker.IntRange(-1000, 1100),
				Note:       faker.LoremIpsumSentence(4),
				CategoryId: cats[faker.IntRange(0, len(cats)-1)].Id,
			})
		}
	})
	return tx
}

func GetTestDataset() *TestDataset {
	return &TestDataset{
		Accounts:     GetTestDatasetAccounts(),
		Transactions: GetTestDatasetTransactions(),
		Categories:   GetTestDatasetCategories(),
	}
}
