package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content   string             `bson:"content" json:"content"`
	UserId    string             `bson:"UserId" json:"UserId"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
