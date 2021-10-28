package dto

type LoginWithsDTO struct {
	Email    string `bson:"email" json: "email" form: "email" binding: "required" validate: "email"`
	Password string `bson:"password" json: "password" form: "password" binding: "required" validate: "min:8"`
}
