package artist

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
	artists := []Artist{}
	result := db.Find(&artists)
	if result.Error != nil {
		//TODO error handling
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(artists)
}

//GetByID returns a record by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	artist := Artist{}
	id := params["id"]
	err := db.Where("id = ?", id).First(&artist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	json.NewEncoder(w).Encode(artist)
}

//Create creates a record
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	artist := Artist{}
	json.NewDecoder(r.Body).Decode(&artist)
	//TODO validation
	artist.ID = uuid.NewV4()
	db.Create(artist)
	json.NewEncoder(w).Encode(artist)
}

//Update updates a record
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	artist := Artist{}
	params := mux.Vars(r)
	id := params["id"]
	err := db.Where("id = ?", id).First(&artist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	//TODO validations
	updatedArtist := Artist{}
	json.NewDecoder(r.Body).Decode(&updatedArtist)
	db.Model(&artist).Updates(&updatedArtist)
	db.Where("id = ?", id).First(&artist)
	json.NewEncoder(w).Encode(artist)
}

//Delete deletes a record
func Delete(w http.ResponseWriter, r *http.Request) {
	artist := Artist{}
	params := mux.Vars(r)
	id := params["id"]
	err := db.Where("id = ?", id).First(&artist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	db.Delete(&artist)
	w.WriteHeader(204)
}
