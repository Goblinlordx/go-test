package accounting

import (
	"strconv"
	"strings"
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

func GetTestDatasetPayee() []string {
	faker := gofakeit.New(901)
	p := make([]string, 20)
	for i, _ := range p {
		p[i] = faker.Name()
	}
	return p
}

func GetTestDatasetAccounts() []Account {
	faker := gofakeit.New(101)
	a := make([]Account, 10)
	for i, _ := range a {
		a[i] = Account{
			Id:             uuid.MustParse(faker.UUID()),
			Name:           faker.Name(),
			CurrencySymbol: currency.CurrencySymbol("USD"),
		}
	}
	return a
}

func GetTestDatasetCategories() []Category {
	cats := []Category{}
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
	return cats
}

func GetTestDatasetTransactions() []Transaction {
	accounts := GetTestDatasetAccounts()
	payees := GetTestDatasetPayee()
	cats := GetTestDatasetCategories()
	faker := gofakeit.New(102)
	t := make([]Transaction, 10000)
	_ = t
	for i, _ := range t {
		t[i] = Transaction{
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
		}
	}
	return t
}

func GetTestDataset() *TestDataset {
	return &TestDataset{
		Accounts:     GetTestDatasetAccounts(),
		Transactions: GetTestDatasetTransactions(),
		Categories:   GetTestDatasetCategories(),
	}
}
