package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize the database
func InitializeUserDB(postgresDBUrl, userDB string) *gorm.DB {
	dsn := fmt.Sprintf("%s%s", postgresDBUrl, userDB)
	userDatabase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error initializing user database")
	}
	return userDatabase
}
