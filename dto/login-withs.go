package dto

type LoginWithsDTO struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Email    string `bson:"email" json: "email" form: "email" binding: "required" validate: "email"`
	Password string `bson:"password" json: "password" form: "password" binding: "required" validate: "min:8"`
}
