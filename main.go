package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initDB()
	startRouter()
}

func startRouter() {

}

var db *gorm.DB

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Disc{})
	db.AutoMigrate(&Style{})
	db.AutoMigrate(&Artist{})

	fmt.Printf("Successfully connected to %v\n", dbname)
}
