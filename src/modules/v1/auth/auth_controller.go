package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type controller struct {
	service interfaces.AuthService
}

func NewController(service interfaces.AuthService) *controller {
	return &controller{service}
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	var input input.AuthInput

	if err := helpers.ValidationError(input); err != nil {
		helpers.New(err.Error(), 401, true)
		return
	}

	json.NewDecoder(r.Body).Decode(&input)

	c.service.Login(input).Send(w)
}
