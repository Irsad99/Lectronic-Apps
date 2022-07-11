package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
	"github.com/asaskevich/govalidator"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type product_ctrl struct {
	svc interfaces.ProductService
}

func NewCtrl(ctrl interfaces.ProductService) *product_ctrl {
	return &product_ctrl{ctrl}
}

func (ctrl *product_ctrl) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.FindAll()
	if err != nil {
		helpers.New(data, 404, true).Send(w)
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *product_ctrl) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataID = r.URL.Query()
	id, err := strconv.Atoi(dataID["id"][0])
	if err != nil {
		helpers.New(id, 400, true).Send(w)
	}

	data, err := ctrl.svc.FindByID(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *product_ctrl) SearchByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getdata = r.URL.Query()
	name := string(getdata["name"][0])

	data, err := ctrl.svc.SearchByName(name)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *product_ctrl) SortByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getdata = r.URL.Query()
	category := string(getdata["category"][0])

	data, err := ctrl.svc.SortByCategory(category)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *product_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data input.InputProduct
	var decoder = schema.NewDecoder()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = govalidator.ToInt(data.Price)
	if err != nil {
		helpers.New(err.Error(), 400, true).Send(w)
		return 
	}

	err = helpers.Validation(data.Name, data.Category, data.Description)
	if err != nil {
		helpers.New(err.Error(), 400, true).Send(w)
		return 
	}

	// _, err = govalidator.ValidateStruct(data)
	// if err != nil {
	// 	helpers.New(err.Error(), 400, true).Send(w)
	// 	return 
	// }

	result, err := ctrl.svc.Add(&data, file, handler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}

func (ctrl *product_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])
	if err != nil {
		helpers.New(data, 400, true).Send(w)
	}

	result, err := ctrl.svc.Delete(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *product_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.Product

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result, err := ctrl.svc.Update(id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *product_ctrl) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var dataId = r.URL.Query()
	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	result, err := ctrl.svc.Upload(id, file, handler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&result)
}
