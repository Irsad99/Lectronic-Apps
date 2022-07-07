package users

import (
	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	help "github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type user_service struct {
	rep interfaces.UserRepo
}

func NewService(svc interfaces.UserRepo) *user_service {
	return &user_service{svc}
}

// GET ALL USER
func (svc *user_service) GetAll() (*help.Response, error) {
	result, err := svc.rep.GetAll()
	if err != nil {
		res := help.New(result, 400, true)
		return res, nil
	}

	res := help.New(result, 200, false)
	return res, nil
}

// GET USER BY ID
func (svc *user_service) GetId(id int) (*help.Response, error) {
	
	data, err := svc.rep.GetId(id)
	if err != nil {
		return nil, err
	}

	res := help.New(data, 200, false)
	return res, nil
}

// GET USER BY EMAIL
func (svc *user_service) GetEmail(email string) (*help.Response, error) {
	data, err := svc.rep.GetEmail(email)
	if err != nil {
		return nil, err
	}

	res := help.New(data, 200, false)
	return res, nil
}

// ADD USER aka CREATE USER aka REGISTER
func (svc *user_service) AddUser(usr *models.User) (*help.Response, error) {
	hashPwd, err := help.HashPassword(usr.Password)
	if err != nil {
		return nil, err
	}

	usr.Password = hashPwd
	data, err := svc.rep.AddUser(usr)
	if err != nil {
		return nil, err
	}

	res := help.New(data, 200, false)
	return res, nil
}

// UPDATE USER
func (svc *user_service) Update(id int, data *models.User) (*help.Response, error) {
	data, err := svc.rep.Update(id, data)
	if err != nil {
		return nil, err
	}

	res := help.New(data, 200, false)
	return res, nil
}

// DELETE USER
func (svc *user_service) Delete(id int) (*help.Response, error) {
	result, err := svc.rep.Delete(id)
	if err != nil {
		res := help.New(result, 400, true)
		return res, nil
	}

	res := help.New(result, 200, false)
	return res, nil
}