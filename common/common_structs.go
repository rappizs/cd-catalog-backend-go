package common

import uuid "github.com/satori/go.uuid"

//Base contains common columns for all structs
type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
}

//ValidationError is a struct to be sent back as response body if a validation fails
type ValidationError struct {
	Errors []string `json:"errors"`
}
