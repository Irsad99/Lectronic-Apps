package models

import "time"

type Product struct {
	Id_Product  uint      `gorm:"primaryKey" json:"id_product"`
	Name        string    `json:"name" valid:"type(string), required"`
	Price       string    `json:"price" valid:"type(string), required"`
	Category    string    `json:"category" valid:"type(string), required"`
	Description string    `json:"description" valid:"type(string), required"`
	Stock       string    `json:"stock" valid:"type(int), required"`
	Image       string    `json:"image" valid:"type(string), required"`
	Sold        string    `json:"sold" valid:"type(int), required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Products []Product
