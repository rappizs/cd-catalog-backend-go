package artist

import (
	"cd-catalog-backend-go/structs"
)

//Artist is a representation of a Disc's artist
type Artist struct {
	structs.Base
	Name string `json:"name" gorm:"unique"`
}
