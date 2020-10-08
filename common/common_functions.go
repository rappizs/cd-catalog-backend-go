package common

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

//Paginate ....
func Paginate(r *http.Request) (func(db *gorm.DB) *gorm.DB, error) {
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		return nil, err
	}
	perPage, err := strconv.Atoi(r.FormValue("per-page"))
	if err != nil {
		return nil, err
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * perPage).Limit(perPage)
	}, nil
}

//OrderBy sets ordering
func OrderBy(by string, orderType string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(by + " " + orderType)
	}
}

//CreatePaginatedResponse creates paginated response
func CreatePaginatedResponse(r *http.Request, db *gorm.DB, table string, data interface{}) (map[string]interface{}, error) {
	var recordCount int64
	db.Table(table).Count(&recordCount)
	perPage, err := strconv.Atoi(r.FormValue("per-page"))
	if err != nil {
		return nil, err
	}
	var pageCount int64
	if recordCount%int64(perPage) == 0 {
		pageCount = recordCount / int64(perPage)
	} else {
		pageCount = recordCount/int64(perPage) + 1
	}

	currPage, err := strconv.Atoi(r.FormValue("page"))

	return map[string]interface{}{
		"current_page": currPage,
		"last_page":    pageCount,
		"data":         data,
	}, nil
}

//CreateErrorStruct creates a struct containing the validation errors
func CreateErrorStruct(err error) ValidationError {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Field()+":"+err.Tag()+":"+err.Param())
	}

	validErr := ValidationError{}
	validErr.Errors = errors
	return validErr
}
