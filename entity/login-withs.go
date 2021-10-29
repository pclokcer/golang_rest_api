package entity

import "time"

type LoginWith struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Type      string    `bson:"type" json:"type"`
	UserID    string    `bson:"user_id" json:"user_id"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
