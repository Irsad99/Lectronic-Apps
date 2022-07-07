package orders

import (
	"net/http"

	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type controller struct {
	service interfaces.OrderService
}

func NewController(service interfaces.OrderService) *controller {
	return &controller{service: service}
}

func (c *controller) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}
