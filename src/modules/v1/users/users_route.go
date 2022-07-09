package users

import (
	"github.com/Irsad99/LectronicApp/src/middleware"
	"github.com/Irsad99/LectronicApp/src/modules/v1/verifyemail"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	verifyReppo := verifyemail.NewRepo(db)
	svc := NewService(repo, verifyReppo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/register", ctrl.AddUser).Methods("POST")
	route.HandleFunc("/me", middleware.Do(ctrl.MyProfile, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/verify", ctrl.VerifyEmail).Methods("GET")
	route.HandleFunc("/avatar", middleware.Do(ctrl.UploadAvatar, middleware.CheckAuth)).Methods("POST")

	route.HandleFunc("/user", ctrl.GetId).Methods("GET")

	route.HandleFunc("/email", ctrl.GetEmail).Methods("GET")

	route.HandleFunc("/delete/{id_user}", middleware.Do(ctrl.Delete, middleware.CheckAuth, middleware.CheckRoleAdmin)).Methods("DELETE")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, middleware.CheckAuth)).Methods("PUT")
}
