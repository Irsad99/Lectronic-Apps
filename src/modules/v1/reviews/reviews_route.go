package reviews

import (

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/review").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/review", ctrl.FindByID).Methods("GET")
	route.HandleFunc("/add", ctrl.AddData).Methods("POST")
}