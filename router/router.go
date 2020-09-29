package router

import (
	"cd-catalog-backend-go/config"
	"cd-catalog-backend-go/models/artist"
	"cd-catalog-backend-go/models/disc"
	"cd-catalog-backend-go/models/style"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Start starts the http router
func Start() {
	initControllers()

	r := mux.NewRouter()

	r.HandleFunc("/api/discs", disc.GetAll).Methods("GET")
	r.HandleFunc("/api/discs/{id}", disc.GetByID).Methods("GET")
	r.HandleFunc("/api/discs", disc.Create).Methods("POST")
	r.HandleFunc("/api/discs/{id}", disc.Update).Methods("PATCH")
	r.HandleFunc("/api/discs/{id}", disc.Delete).Methods("DELETE")

	r.HandleFunc("/api/artists", artist.GetAll).Methods("GET")
	r.HandleFunc("/api/artists/{id}", artist.GetByID).Methods("GET")
	r.HandleFunc("/api/artists", artist.Create).Methods("POST")
	r.HandleFunc("/api/artists/{id}", artist.Update).Methods("PATCH")
	r.HandleFunc("/api/artists/{id}", artist.Delete).Methods("DELETE")

	r.HandleFunc("/api/styles", style.GetAll).Methods("GET")
	r.HandleFunc("/api/styles/{id}", style.GetByID).Methods("GET")
	r.HandleFunc("/api/styles", style.Create).Methods("POST")
	r.HandleFunc("/api/styles/{id}", style.Update).Methods("PATCH")
	r.HandleFunc("/api/styles/{id}", style.Delete).Methods("DELETE")

	fmt.Printf("Application listening on localhost, port: %v\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, r))
}

func initControllers() {
	disc.Init()
	artist.Init()
	style.Init()
}
