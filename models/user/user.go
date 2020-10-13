package user

import "cd-catalog-backend-go/common"

//User is a representation of a user in the app
type User struct {
	common.Base
	UserName string `json:"user_name" gorm:"unique" validate:"required,min=3,max=20"`
	Password string `json:"password" gorm:"not null" validate:"required,min=5"`
}
