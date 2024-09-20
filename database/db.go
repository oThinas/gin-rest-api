package database

import (
	"gin-rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=postgres password=root dbname=school port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic("Could not connect to the database")
	}

	DB.AutoMigrate(&models.Student{})
}
