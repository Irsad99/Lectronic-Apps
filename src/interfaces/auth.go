package interfaces

import (
	"github.com/Irsad99/LectronicApp/src/helpers"
	"github.com/Irsad99/LectronicApp/src/input"
)

type AuthService interface {
	Login(input input.AuthInput) *helpers.Response
}
