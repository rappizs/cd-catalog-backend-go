package common

import uuid "github.com/satori/go.uuid"

//Base contains common columns for all structs
type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
}
