package products

import (
	"github.com/Irsad99/LectronicApp/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.Do(ctrl.FindAll, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/item", middleware.Do(ctrl.FindByID, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/search", middleware.Do(ctrl.SearchByName, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/sort", middleware.Do(ctrl.SortByCategory, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/add", middleware.Do(ctrl.AddData, middleware.CheckRoleAdmin)).Methods("POST")
	route.HandleFunc("/upload", middleware.Do(ctrl.UploadAvatar, middleware.CheckRoleAdmin)).Methods("POST")
	route.HandleFunc("/delete/{id}", middleware.Do(ctrl.Delete, middleware.CheckRoleAdmin)).Methods("DELETE")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, middleware.CheckRoleAdmin)).Methods("PUT")
}
