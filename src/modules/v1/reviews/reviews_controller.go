package reviews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type review_ctrl struct {
	svc interfaces.ReviewService
}

func NewCtrl(ctrl interfaces.ReviewService) *review_ctrl {
	return &review_ctrl{ctrl}
}

func (ctrl *review_ctrl) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataID = r.URL.Query()
	id, err := strconv.Atoi(dataID["id"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data, err := ctrl.svc.FindByID(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *review_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getID := r.Header.Get("user_id")
	id, err := strconv.Atoi(getID)
	if err != nil {
		fmt.Println("error")
	}

	var data input.InputReview

	json.NewDecoder(r.Body).Decode(&data)

	if err := helpers.ValidationError(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ctrl.svc.Add(id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}
