package dto

type Register struct {
	Email    string `bson:"email" json: "email" validate:"required,email"`
	Password string `bson:"password" json: "password" validate:"required,min:8"`
	Type     string `bson:"password" json: "type" validate:"required"`
}
