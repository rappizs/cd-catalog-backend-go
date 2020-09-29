package controllers

import (
	"cd-catalog-backend-go/database"
	"cd-catalog-backend-go/structs"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//GetAllDisc returns every disc record
func GetAllDisc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	discs := []structs.Disc{}
	result := database.DB.Find(&discs)
	if result.Error != nil {
		//TODO error handling
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(discs)
}

//GetDisc returns a disc by id
func GetDisc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	disc := structs.Disc{}
	id := params["id"]
	err := database.DB.Where("id = ?", id).First(&disc).Error
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

//CreateDisc creates a disc
func CreateDisc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	disc := structs.Disc{}
	json.NewDecoder(r.Body).Decode(&disc)
	//TODO validation
	disc.ID = uuid.NewV4()
	database.DB.Create(disc)
	json.NewEncoder(w).Encode(disc)
}

//UpdateDisc updates a disc by id
func UpdateDisc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	disc := structs.Disc{}
	params := mux.Vars(r)
	id := params["id"]
	err := database.DB.Where("id = ?", id).First(&disc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	//TODO validations
	updatedDisc := structs.Disc{}
	json.NewDecoder(r.Body).Decode(&updatedDisc)
	database.DB.Model(&disc).Updates(&updatedDisc)
	database.DB.Where("id = ?", id).First(&disc)
	json.NewEncoder(w).Encode(disc)
}

//DeleteDisc deletes a disc by  id
func DeleteDisc(w http.ResponseWriter, r *http.Request) {
	disc := structs.Disc{}
	params := mux.Vars(r)
	id := params["id"]
	err := database.DB.Where("id = ?", id).First(&disc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	database.DB.Delete(&disc)
	w.WriteHeader(204)
}
