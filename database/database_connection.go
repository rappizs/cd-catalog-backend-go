package database

import (
	"cd-catalog-backend-go/config"
	structs "cd-catalog-backend-go/structs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB is the interface to communicate with the db
var DB *gorm.DB

//InitDB connects to the db
func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBname)
	var err error
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&structs.Disc{})
	DB.AutoMigrate(&structs.Style{})
	DB.AutoMigrate(&structs.Artist{})

	fmt.Printf("Successfully connected to %v\n", config.DBname)
}
