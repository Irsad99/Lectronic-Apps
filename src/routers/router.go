package routers

import (
	"net/http"

	database "github.com/Irsad99/LectronicApp/src/database/gorm"
	"github.com/Irsad99/LectronicApp/src/modules/v1/auth"
	"github.com/Irsad99/LectronicApp/src/modules/v1/orders"
	"github.com/Irsad99/LectronicApp/src/modules/v1/products"
	"github.com/Irsad99/LectronicApp/src/modules/v1/reviews"
	"github.com/Irsad99/LectronicApp/src/modules/v1/users"
	"github.com/rs/cors"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New() (http.Handler, error) {
	mainRoute := mux.NewRouter()

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(mainRoute)

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	products.NewRouter(mainRoute, db)
	users.NewRouter(mainRoute, db)
	auth.NewRouter(mainRoute, db)
	orders.NewRouter(mainRoute, db)
	reviews.NewRouter(mainRoute, db)

	return c, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"hello\": \"world\"}"))
}
