package input

type InputProduct struct {
	Name        string `json:"name" validate:"type(string) required"`
	Price       string `json:"price" validate:"type(string) required"`
	Stock       int    `json:"stock" validate:"type(string) required"`
	Description string `json:"description" validate:"type(string) required"`
	Category    string `json:"category" validate:"type(string) required"`
	Image       string `json:"image" validate:"type(string) required"`
}
