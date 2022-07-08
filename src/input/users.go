package input

type RegisterInput struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=6,required"`
}
