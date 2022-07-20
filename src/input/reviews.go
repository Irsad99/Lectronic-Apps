package input

type InputReview struct {
	ProductID int `json:"product_id"`
	Comment string `json:"comment"`
	Rating int `json:"rating"`
}