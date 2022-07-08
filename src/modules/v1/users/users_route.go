package users

import (
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
	route.HandleFunc("/verify", ctrl.VerifyEmail).Methods("GET")
	route.HandleFunc("/avatar", ctrl.UploadAvatar).Methods("POST")

	route.HandleFunc("/user", ctrl.GetId).Methods("GET")

	route.HandleFunc("/email", ctrl.GetEmail).Methods("GET")

	route.HandleFunc("/delete/{id_user}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}
