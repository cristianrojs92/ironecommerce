package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string	`json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	}

//Se la estructura de payload
type PayloadUser struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Exp int64  `json:"exp"`
	}

	type Auth struct {
		UserId int `json:id`
		Token string `json:token`
	}