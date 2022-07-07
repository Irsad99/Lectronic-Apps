package auth

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/auth").Subrouter()

	repository := users.NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.Login).Methods("POST")
}
