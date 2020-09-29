package disc

import (
	"cd-catalog-backend-go/database"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

//Init gets the db interface from database package
func Init() {
	db = database.GetDBInterface()
}

//GetAll returns every record
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	discs := []Disc{}
	result := db.Find(&discs)
	if result.Error != nil {
		//TODO error handling
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(discs)
}

//GetByID returns a record by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	disc := Disc{}
	id := params["id"]
	err := db.Where("id = ?", id).First(&disc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	json.NewEncoder(w).Encode(disc)
}

//Create creates a record
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	disc := Disc{}
	json.NewDecoder(r.Body).Decode(&disc)
	//TODO validation
	disc.ID = uuid.NewV4()
	db.Create(disc)
	json.NewEncoder(w).Encode(disc)
}

//Update updates a record
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	disc := Disc{}
	params := mux.Vars(r)
	id := params["id"]
	err := db.Where("id = ?", id).First(&disc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	//TODO validations
	updatedDisc := Disc{}
	json.NewDecoder(r.Body).Decode(&updatedDisc)
	db.Model(&disc).Updates(&updatedDisc)
	db.Where("id = ?", id).First(&disc)
	json.NewEncoder(w).Encode(disc)
}

//Delete deletes a record
func Delete(w http.ResponseWriter, r *http.Request) {
	disc := Disc{}
	params := mux.Vars(r)
	id := params["id"]
	err := db.Where("id = ?", id).First(&disc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	db.Delete(&disc)
	w.WriteHeader(204)
}
