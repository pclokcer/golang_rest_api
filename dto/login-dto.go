package dto

type LoginDTO struct {
	email string `json: "email" form: "email" binding: "required" validate: "email"`
	password string `json: "password" form: "password" binding: "required" validate: "min:8"`
}