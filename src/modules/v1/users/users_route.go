package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/user", ctrl.GetId).Methods("GET")
	route.HandleFunc("/email", ctrl.GetEmail).Methods("GET")
	route.HandleFunc("/add", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/delete/{id_user}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}