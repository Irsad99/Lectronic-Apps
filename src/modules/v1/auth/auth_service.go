package auth

import (
	"strconv"

	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
	"github.com/Irsad99/LectronicApp/src/interfaces"
)

type service struct {
	repository interfaces.UserRepo
}

func NewService(repository interfaces.UserRepo) *service {
	return &service{repository}
}

func (s *service) Login(input input.AuthInput) *helpers.Response {
	if err := helpers.ValidationError(input); err != nil {
		return helpers.New(err.Error(), 401, true)
	}

	user, err := s.repository.GetEmail(input.Email)
	if err != nil {
		return helpers.New("email/password incorrect, please correct this", 401, true)
	}

	if !helpers.CheckPassword(user.Password, input.Password) {
		return helpers.New("email/password incorrect, please correct this", 401, true)
	}

	if !user.Verified {
		return helpers.New("please check your email for verification", 401, true)
	}

	uid := strconv.FormatUint(uint64(user.IdUser), 10)
	new := helpers.NewToken(uid, user.Email, user.Role)
	token, err := new.Create()
	if err != nil {
		return helpers.New("failed to create token", 401, true)
	}

	return helpers.New(token, 200, false)
}
