package models

import (

	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string 				 `bson:"user_id" json:"user_id"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Likes     int                `bson:"likes" json:"likes"`
	Comments  []Comment          `bson:"comments" json:"comments"`
	Saves     int                `bson:"saves" json:"saves"`
}

