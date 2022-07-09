package users

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

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(ctrl interfaces.UserService) *user_ctrl {
	return &user_ctrl{ctrl}
}

// GET ALL USERS
func (ctrl *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.GetAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

// GET USER BY ID
func (ctrl *user_ctrl) GetId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataID = r.URL.Query()
	id, err := strconv.Atoi(dataID["id"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data, err := ctrl.svc.GetId(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

// GET USER BY EMAIL
func (ctrl *user_ctrl) GetEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getdata = r.URL.Query()
	email := string(getdata["email"][0])

	result, err := ctrl.svc.GetEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// ADD USER aka CREATE USER aka REGISTER
func (ctrl *user_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input input.RegisterInput
	json.NewDecoder(r.Body).Decode(&input)

	if err := helpers.ValidationError(input); err != nil {
		helpers.New(err.Error(), 401, true)
		return
	}

	result, err := ctrl.svc.AddUser(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}

// UPDATE USER
func (ctrl *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.User

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result, err := ctrl.svc.Update(uint(id), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// DELETE USER
func (ctrl *user_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result, err := ctrl.svc.Delete(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(result)
}

func (ctrl *user_ctrl) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("i")

	result, err := ctrl.svc.Verify(token)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result.Send(w)
}

func (ctrl *user_ctrl) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	i := r.Header.Get("user_id")
	id, err := strconv.Atoi(i)
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

	result.Send(w)
}

func (ctrl *user_ctrl) MyProfile(w http.ResponseWriter, r *http.Request) {
	i := r.Header.Get("user_id")
	id, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(id)

	result, err := ctrl.svc.GetId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}
