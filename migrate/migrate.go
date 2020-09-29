package migrate

import (
	"cd-catalog-backend-go/artist"
	"cd-catalog-backend-go/database"
	"cd-catalog-backend-go/disc"
	"cd-catalog-backend-go/style"
)

//All migrates all structs to the db
func All() {
	db := database.GetDBInterface()
	db.AutoMigrate(&disc.Disc{})
	db.AutoMigrate(&style.Style{})
	db.AutoMigrate(&artist.Artist{})
}
