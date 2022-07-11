package reviews

import (
	"github.com/Irsad99/LectronicApp/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/review").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/review", middleware.Do(ctrl.FindByID, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/add", middleware.Do(ctrl.AddData, middleware.CheckAuth)).Methods("POST")
}
