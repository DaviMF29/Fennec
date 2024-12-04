package repository

import (
	"context"
	"errors"
	"time"

	"github.com/DaviMF29/wombat/db"
	"github.com/DaviMF29/wombat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var ErrNoRows = errors.New("mongo: no documents in result set")

func InsertUser(user models.User) (id string, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return "", err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("users")

	encryptedPassword, err := encryptPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = encryptedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func GetUserById(id string) (models.User, error) {
	client, err := db.OpenConnection()
	if err != nil {
		return models.User{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, errors.New("ID inválido")
	}

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return models.User{}, ErrNoRows
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	client, err := db.OpenConnection()
	if err != nil {
		return models.User{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, ErrNoRows
	}
	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	client, err := db.OpenConnection()
	if err != nil {
		return models.User{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, ErrNoRows
	}
	return user, nil
}

func encryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
