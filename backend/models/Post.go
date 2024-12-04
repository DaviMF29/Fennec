package models

import (
	"context"
	"time"

	"github.com/DaviMF29/wombat/db"
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

func InsertPost(post Post) (id string, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("posts")

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	result, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return "", err
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func GetPostById(id string) (post Post, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return Post{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("posts")

	post = Post{}
	err = collection.FindOne(context.Background(), primitive.M{"_id": id}).Decode(&post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}