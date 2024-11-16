package main

import (
	"fmt"
	"log"
	"order_inventory_management/models"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("Failed to load .env :", err)
	}
	postgresDbUrl := os.Getenv("POSTGRES_DB_URL")
	database := os.Getenv("DB")

	fmt.Print("1 : Create Database\n", "2 : Run Migration\n", "Choose the option : ")
	var Option int
	fmt.Scan(&Option)

	if Option == 1 {
		postgresDb, err := gorm.Open(postgres.Open(postgresDbUrl), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to PostgreSQL server")
		}
		CheckAndCreateDatabase(postgresDb, database)
	} else if Option == 2 {
		dsn := fmt.Sprintf("%s%s", postgresDbUrl, database)
		Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database")
		}
		if err := MigrateDatabase(Db).Migrate(); err != nil {
			log.Fatal("Failed to migrate database")
		}
		fmt.Println("Database migrated successfully")

	} else {
		log.Println("Enter the valid option.")
		return
	}
}

func CheckAndCreateDatabase(initialDB *gorm.DB, dbName string) {
	var exists bool
	err := initialDB.Raw("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		log.Fatal("failed to check if database exists:", err)
	}

	if !exists {
		err = initialDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			log.Fatal("failed to create database:", err)
		}
		log.Println("Database created successfully:", dbName)
	} else {
		log.Println("Database already exists:", dbName)
	}
}

func MigrateDatabase(db *gorm.DB) *gormigrate.Gormigrate {
	migrate := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "added_users_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.User{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable(&models.User{})
			},
		},
		{
			ID: "added_productss_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Product{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable(&models.Product{})
			},
		},
		{
			ID: "added_inventory_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Inventory{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable(&models.Inventory{})
			},
		},
		{
			ID: "added_orders_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Order{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable(&models.Order{})
			},
		},
	})

	if err := migrate.Migrate(); err != nil {
		log.Println("Failed to migrate")
	}
	return migrate
}
