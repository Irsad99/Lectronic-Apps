package input

type InputProduct struct {
	Name        string `json:"name" valid:"type(string) required"`
	Price       string `json:"price" valid:"type(string) required"`
	Stock       int `json:"stock" valid:"type(string) required"`
	Description string `json:"description" valid:"type(string) required"`
	Category    string `json:"category" valid:"type(string) required"`
	Image       string `json:"image"`
}
