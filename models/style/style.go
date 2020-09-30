package style

import "cd-catalog-backend-go/common"

//Style is a representation of a music style
type Style struct {
	common.Base
	Name string `json:"name" gorm:"unique"`
}
