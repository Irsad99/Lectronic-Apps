package orders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
	"github.com/gorilla/mux"
)

type controller struct {
	service        interfaces.OrderService
	paymentService interfaces.PaymentService
}

func NewController(service interfaces.OrderService, paymentService interfaces.PaymentService) *controller {
	return &controller{service, paymentService}
}

func (c *controller) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := c.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	oid := params["id"]
	id, err := strconv.Atoi(oid)
	if err != nil {
		fmt.Println("error")
	}

	res, err := c.service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) MyOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h := r.Header.Get("user_id")
	id, err := strconv.Atoi(h)
	if err != nil {
		fmt.Println("error")
	}

	res, err := c.service.FindByUserID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) NewOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h := r.Header.Get("user_id")
	id, err := strconv.Atoi(h)
	if err != nil {
		fmt.Println("error")
	}

	var input input.OrderInput

	json.NewDecoder(r.Body).Decode(&input)

	if err := helpers.ValidationError(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := c.service.Create(uint64(id), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	var order models.Order

	json.NewDecoder(r.Body).Decode(&order)
	if err := helpers.ValidationError(order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := c.service.Update(id, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	res, err := c.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) GetNotificationOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input input.OrderNotificationInput

	json.NewDecoder(r.Body).Decode(&input)
	if err := helpers.ValidationError(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := c.paymentService.ProcessPayment(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}
