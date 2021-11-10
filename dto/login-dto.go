package dto

type LoginDTO struct {
	Email    string `bson:"email" json: "email" form: "email" binding: "required" validate:"required,email"`
	Password string `bson:"password" json: "password" form: "password" binding: "required" validate:"required,min:8"`
}
