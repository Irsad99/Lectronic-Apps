package input

type InputReview struct {
	ProductID int `json:"product_id" validate:"type(int) required"`
	Comment string `json:"comment"  validate:"type(string) required"`
	Rating int `json:"rating" validate:"type(int) required"`
}