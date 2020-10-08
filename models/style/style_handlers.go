package style

import (
	"cd-catalog-backend-go/common"
	"cd-catalog-backend-go/database"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var db *gorm.DB
var validate *validator.Validate

//Init gets the db interface from database package
func Init() {
	db = database.GetDBInterface()
	validate = validator.New()
}

//GetAll returns every record
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	searchValue := r.FormValue("search-value")
	orderType := r.FormValue("type")
	order := common.OrderBy("name", orderType)

	styles := []Style{}
	result := db.Where("name like ?", "%"+searchValue+"%").
		Scopes(order).
		Find(&styles)

	if result.Error != nil {
		//TODO error handling
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(styles)
}

//GetByID returns a record by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	style := Style{}
	id := params["id"]
	err := db.Where("id = ?", id).First(&style).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	json.NewEncoder(w).Encode(style)
}

//Create creates a record
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	style := Style{}
	json.NewDecoder(r.Body).Decode(&style)

	err := validate.Struct(style)
	if err != nil {
		validErr := common.CreateErrorStruct(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&validErr)
		return
	}

	style.ID = uuid.NewV4()
	db.Create(style)
	json.NewEncoder(w).Encode(style)
}

//Update updates a record
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	style := Style{}
	params := mux.Vars(r)
	id := params["id"]
	err := db.Where("id = ?", id).First(&style).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	
	json.NewDecoder(r.Body).Decode(&style)

	err = validate.Struct(style)
	if err != nil {
		validErr := common.CreateErrorStruct(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&validErr)
		return
	}

	db.Model(&style).Updates(&style)
	db.Where("id = ?", id).First(&style)
	json.NewEncoder(w).Encode(style)
}

//Delete deletes a record
func Delete(w http.ResponseWriter, r *http.Request) {
	style := Style{}
	params := mux.Vars(r)
	id := params["id"]
	err := db.Where("id = ?", id).First(&style).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	db.Delete(&style)
	w.WriteHeader(204)
}
