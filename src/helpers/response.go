package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status      string      `json:"status"`
	IsError     bool        `json:"isError"`
	Data        interface{} `json:"data,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func New(data interface{}, code int, isError bool) *Response {

	if isError {
		return &Response{
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}

	}
	return &Response{
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}

func ValidationError(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

func Validation(dataString ...string) error {
	for _, value := range dataString {
		splitValue := strings.Split(value, " ")
		cekValues := strings.Join(splitValue, "")
		if boolName := govalidator.IsAlpha(cekValues); !boolName {
			return errors.New("value : {"+ value +"} / please fill in the column with string data type /")
		}
	}
	return nil
}
