package auth

import (
	"github.com/Irsad99/LectronicApp/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/auth").Subrouter()

	repository := users.NewRepo(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.Login).Methods("POST")
}
