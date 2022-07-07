package products

import (

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.FindAll).Methods("GET")
	route.HandleFunc("/product", ctrl.FindByID).Methods("GET")
	route.HandleFunc("/search", ctrl.SearchByName).Methods("GET")
	route.HandleFunc("/sort", ctrl.SortByCategory).Methods("GET")
	route.HandleFunc("/add", ctrl.AddData).Methods("POST")
	route.HandleFunc("/delete/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}