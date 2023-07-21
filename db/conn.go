package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

func getDB() (*gorm.DB, error) {
	if conn != nil {
		return conn, nil
	}
	dbp := os.Getenv("DB")
	conn, err := gorm.Open(sqlite.Open(dbp), &gorm.Config{})
	return conn, err
}
