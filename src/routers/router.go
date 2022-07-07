package routers

import (
	"net/http"

	// database "github.com/Irsad99/LectronicApp/src/database/gorm"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	// db, err := database.New()
	// if err != nil {
	// 	return nil, err
	// }

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"hello\": \"world\"}"))
}
