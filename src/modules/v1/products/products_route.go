package products

import (

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.FindAll).Methods("GET")
	route.HandleFunc("/item", ctrl.FindByID).Methods("GET")
	route.HandleFunc("/search", ctrl.SearchByName).Methods("GET")
	route.HandleFunc("/sort", ctrl.SortByCategory).Methods("GET")
	route.HandleFunc("/add", ctrl.AddData).Methods("POST")
	route.HandleFunc("/upload", ctrl.UploadAvatar).Methods("POST")
	route.HandleFunc("/delete/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}