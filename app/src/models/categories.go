package models


type Category struct {
	Id int `json:"id"`
	Main bool `json:"main"`
	Name string `json:"name"`
	Icon_name string `json:"icon_name"`
	Sub_categories []Category `json:"subcategories"`
}