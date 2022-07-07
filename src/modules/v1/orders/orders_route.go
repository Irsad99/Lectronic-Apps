package orders

import (
	"github.com/Irsad99/LectronicApp/src/modules/v1/payments"
	"github.com/Irsad99/LectronicApp/src/modules/v1/products"
	"github.com/Irsad99/LectronicApp/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/order").Subrouter()

	repository := NewRepository(db)
	userRepo := users.NewRepo(db)
	productRepo := products.NewRepo(db)
	paymentService := payments.NewService(repository, productRepo)
	service := NewService(repository, userRepo, paymentService)
	controller := NewController(service, paymentService)

	route.HandleFunc("/new", controller.NewOrder).Methods("POST")
	route.HandleFunc("/:id", controller.GetOrderDetail).Methods("GET")
	route.HandleFunc("/me", controller.MyOrder).Methods("GET")
	route.HandleFunc("/notification", controller.GetNotificationOrder).Methods("POST")

	// ADMIN ACCESS
	route.HandleFunc("/all", controller.FindAll).Methods("GET")
	route.HandleFunc("/update/:id", controller.UpdateOrder).Methods("POST")
	route.HandleFunc("/update/:id", controller.DeleteOrder).Methods("DELETE")
}
