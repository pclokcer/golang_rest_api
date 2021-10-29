package entity

import "time"

type LoginWith struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Type      string    `bson:"type" json:"type" validate:"required"`
	UserID    string    `bson:"user_id" json:"user_id" validate:"required"`
	Email     string    `bson:"email" json:"email" validate:"required"`
	Password  string    `bson:"password" json:"password" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at" validate:"required"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at" validate:"required"`
}
