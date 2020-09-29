package router

import (
	"cd-catalog-backend-go/config"
	"cd-catalog-backend-go/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//StartRouter starts the http router
func StartRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/discs", controller.GetAllDisc).Methods("GET")
	r.HandleFunc("/api/discs/{id}", controller.GetDisc).Methods("GET")
	r.HandleFunc("/api/discs", controller.CreateDisc).Methods("POST")
	r.HandleFunc("/api/discs/{id}", controller.UpdateDisc).Methods("PATCH")
	r.HandleFunc("/api/discs/{id}", controller.DeleteDisc).Methods("DELETE")

	fmt.Printf("Application listening on localhost, port: %v\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, r))
}
