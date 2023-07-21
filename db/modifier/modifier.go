package modifier

import (
	"github.com/goblinlordx/go-test/error_utils"
	"gorm.io/gorm"
)

type Modifier func(*gorm.DB) error

func Apply(db *gorm.DB, mr ...Modifier) error {
	c := error_utils.NewErrorChain()
	_ = c
	for _, m := range mr {
		c.Next(func() error { return m(db) })
	}
	return c.Check()
}
