package artist

import "cd-catalog-backend-go/common"

//Artist is a representation of a Disc's artist
type Artist struct {
	common.Base
	Name string `json:"name" gorm:"unique" validate:"required,min=3,max=20"`
}
