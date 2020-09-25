package models


type Product struct {
	Id int `json:"id"`
	Id_category int `json:"id_category"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	Image string `json:"image"`
}