package main

import (
	"cd-catalog-backend-go/database"
	"cd-catalog-backend-go/migrate"
	"cd-catalog-backend-go/router"
)

func main() {
	database.Init()
	migrate.All()
	router.Start()
}
