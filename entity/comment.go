package entity

import "time"

type Comment struct {
	ID         string                 `bson:"_id,omitempty" json:"id"`
	ReviewerID string                 `bson:"reviwer_id" json:"reviewer_id"`
	PostID     string                 `bson:"post_id" json:"post_id"`
	ExtID      string                 `bson:"ext_id" json:"ext_id"`
	Comment    map[string]interface{} `bson:"comment" json:"comment"`
	CreatedAt  time.Time              `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time              `bson:"updated_at" json:"updated_at"`
}
