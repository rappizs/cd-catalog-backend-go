package migrate

import (
	"cd-catalog-backend-go/database"
	"cd-catalog-backend-go/models/artist"
	"cd-catalog-backend-go/models/disc"
	"cd-catalog-backend-go/models/style"
)

//All migrates all structs to the db
func All() {
	db := database.GetDBInterface()
	db.AutoMigrate(&disc.Disc{})
	db.AutoMigrate(&style.Style{})
	db.AutoMigrate(&artist.Artist{})
}
