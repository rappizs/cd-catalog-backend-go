package user

import (
	"cd-catalog-backend-go/common"
	"cd-catalog-backend-go/database"
	"cd-catalog-backend-go/jwt"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"crypto/sha1"

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

	users := []User{}
	result := db.Find(&users)

	if result.Error != nil {
		//TODO error handling
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(users)
}

//GetByID returns a record by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := User{}
	id := params["id"]
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(404)
		} else {
			w.WriteHeader(400)
		}
		return
	}
	json.NewEncoder(w).Encode(user)
}

//Register creates a user
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	err := validate.Struct(user)
	if err != nil {
		validErr := common.CreateErrorStruct(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&validErr)
		return
	}

	h := sha1.New()
	h.Write([]byte(user.Password))
	user.Password = base64.URLEncoding.EncodeToString(h.Sum(nil))

	user.ID = uuid.NewV4()
	db.Create(user)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	userFromBody := User{}
	json.NewDecoder(r.Body).Decode(&userFromBody)

	err := validate.Struct(userFromBody)
	if err != nil {
		validErr := common.CreateErrorStruct(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&validErr)
		return
	}

	h := sha1.New()
	h.Write([]byte(userFromBody.Password))
	userFromBody.Password = base64.URLEncoding.EncodeToString(h.Sum(nil))

	user := User{}
	err = db.Where("user_name = ?", userFromBody.UserName).
		Where("password = ?", userFromBody.Password).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(401)
		} else {
			w.WriteHeader(400)
		}
		return
	}

	tokenString, err := jwt.Generate(user.ID)
	result := map[string]interface{}{"jwt": tokenString}
	json.NewEncoder(w).Encode(result)
}

/* //Update updates a record
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
} */

//Delete deletes a record
/* func Delete(w http.ResponseWriter, r *http.Request) {
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
} */
