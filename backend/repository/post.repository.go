package repository

import (
	"context"
	"errors"
	"time"
	"github.com/DaviMF29/fennec/db"
	"github.com/DaviMF29/fennec/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertPost(post models.Post) (id string, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("fennec").Collection("posts")

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	result, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return "", err
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func GetPostById(id string) (models.Post, error) {
	client, err := db.OpenConnection()
	if err != nil {
		return models.Post{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("fennec").Collection("posts")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Post{}, errors.New("ID inválido")
	}

	var post models.Post
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func DeletePostById(id string) error {
	client, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("fennec").Collection("posts")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(id string,post models.Post) error {
	client, err := db.OpenConnection()
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	collection := client.Database("fennec").Collection("posts")

	objectID, err := primitive.ObjectIDFromHex(id)

	post.UpdatedAt = time.Now()

	_,err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": post})
	if err != nil {
		return err
	}

	return nil
}