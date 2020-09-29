package database

import (
	"cd-catalog-backend-go/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//Init connects to the db
func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBname)
	var err error
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected to %v\n", config.DBname)
}

//GetDBInterface returns the db interface
func GetDBInterface() *gorm.DB {
	return db
}
