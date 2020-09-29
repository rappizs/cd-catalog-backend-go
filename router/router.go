package router

import (
	"cd-catalog-backend-go/config"
	"cd-catalog-backend-go/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//StartRouter starts the http router
func StartRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/discs", controllers.GetAllDisc).Methods("GET")
	r.HandleFunc("/api/discs/{id}", controllers.GetDisc).Methods("GET")
	r.HandleFunc("/api/discs", controllers.CreateDisc).Methods("POST")
	r.HandleFunc("/api/discs/{id}", controllers.UpdateDisc).Methods("PATCH")
	r.HandleFunc("/api/discs/{id}", controllers.DeleteDisc).Methods("DELETE")

	fmt.Printf("Application listening on localhost, port: %v\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, r))
}
