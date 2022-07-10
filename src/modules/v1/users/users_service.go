package users

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/Irsad99/LectronicApp/src/database/gorm/models"
	help "github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type user_service struct {
	rep        interfaces.UserRepo
	verifyRepo interfaces.TokenRepository
}

func NewService(svc interfaces.UserRepo, verifyRepo interfaces.TokenRepository) *user_service {
	return &user_service{svc, verifyRepo}
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
func (svc *user_service) AddUser(input *input.RegisterInput) (*help.Response, error) {
	var user models.User
	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Role = "user"

	hashPwd, err := help.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPwd
	data, err := svc.rep.AddUser(&user)
	if err != nil {
		return nil, err
	}

	var verify models.Token
	token := help.GenToken(24)

	verify.UserID = data.IdUser
	verify.Token = token

	_, err = svc.verifyRepo.InsertToken(&verify)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://localhost:8080/user/verify?i=%s", token)

	email := []string{user.Email}
	cc := []string{user.Email}

	help.SendMail(email, cc, "Activation Your Account", url)

	res := help.New(data, 200, false)
	return res, nil
}

// UPDATE USER
func (svc *user_service) Update(id uint, data *models.User) (*help.Response, error) {
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

// Verify Account
func (svc *user_service) Verify(token string) (*help.Response, error) {
	tok, err := svc.verifyRepo.FindByToken(token)
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	_, err = svc.verifyRepo.DeleteToken(token)
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	user, err := svc.rep.GetId(int(tok.UserID))
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	user.Updated_at = time.Now()
	user.Verified = true

	_, err = svc.rep.Update(user.IdUser, user)
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	res := help.New("Your account has been verified", 200, false)
	return res, nil
}

func (svc *user_service) Upload(id int, file multipart.File, handle *multipart.FileHeader) (*help.Response, error) {
	user, err := svc.rep.GetId(id)
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	images, err := help.UploadImages("avatar", file, handle)
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	user.Image = images.URL

	_, err = svc.rep.Update(uint(id), user)
	if err != nil {
		res := help.New(err.Error(), 400, true)
		return res, nil
	}

	res := help.New("Upload file successfully", 200, false)
	return res, nil
}
