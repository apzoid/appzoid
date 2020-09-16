package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseInit() {
	dsn := "host=lallah.db.elephantsql.com user=zupcuvyz password=x_broQwnrHZW_lGVNbN74qQlhzpsAW-T dbname=zupcuvyz port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database not found")
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
