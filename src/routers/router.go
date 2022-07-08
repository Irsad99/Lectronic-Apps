package routers

import (
	"net/http"

	database "github.com/Irsad99/LectronicApp/src/database/gorm"
	"github.com/Irsad99/LectronicApp/src/modules/v1/auth"
	"github.com/Irsad99/LectronicApp/src/modules/v1/orders"
	"github.com/Irsad99/LectronicApp/src/modules/v1/products"
	"github.com/Irsad99/LectronicApp/src/modules/v1/reviews"
	"github.com/Irsad99/LectronicApp/src/modules/v1/users"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	products.NewRouter(mainRoute, db)
	users.NewRouter(mainRoute, db)
	auth.NewRouter(mainRoute, db)
	orders.NewRouter(mainRoute, db)
	reviews.NewRouter(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"hello\": \"world\"}"))
}
