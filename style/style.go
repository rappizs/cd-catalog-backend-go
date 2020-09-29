package style

import "cd-catalog-backend-go/structs"

//Style is a representation of a music style
type Style struct {
	structs.Base
	Name string `json:"name" gorm:"unique"`
}
