package main

import (
	"cd-catalog-backend-go/database"
	"cd-catalog-backend-go/router"
)

func main() {
	database.InitDB()
	router.StartRouter()
}
