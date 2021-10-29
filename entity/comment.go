package entity

import "time"

type Comment struct {
	ID         string                 `bson:"_id,omitempty" json:"id"`
	ReviewerID string                 `bson:"reviwer_id" json:"reviewer_id" validate:"required"`
	PostID     string                 `bson:"post_id" json:"post_id" validate:"required"`
	ExtID      string                 `bson:"ext_id" json:"ext_id" validate:"required"`
	Comment    map[string]interface{} `bson:"comment" json:"comment" validate:"required"`
	CreatedAt  time.Time              `bson:"created_at" json:"created_at" validate:"required"`
	UpdatedAt  time.Time              `bson:"updated_at" json:"updated_at" validate:"required"`
}
