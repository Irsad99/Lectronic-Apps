package auth

import (
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
	return &service{repository}
}

func (s *service) Login(input input.AuthInput) *helpers.Response {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return helpers.New("email/password incorrect, please correct this", 401, false)
	}

	if !helpers.CheckPassword(user.Password, input.Password) {
		return helpers.New("email/password incorrect, please correct this", 401, false)
	}

	if !user.Verified {
		return helpers.New("please check your email for verification", 401, false)
	}

	new := helpers.NewToken(user.ID, user.Email, user.Role)
	token, err := new.Create()
	if err != nil {
		return helpers.New("failed to create token", 401, false)
	}

	return helpers.New(token, 200, true)
}
