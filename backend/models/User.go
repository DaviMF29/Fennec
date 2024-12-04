package models

import (
	"context"
	"errors"
	"time"

	"github.com/DaviMF29/wombat/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	BirthDate string             `bson:"birth_date" json:"birth_date"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var ErrNoRows = errors.New("mongo: no documents in result set")

func InsertUser(user User) (id string, err error) {
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


func GetUserById(id string) (user User, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return User{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("wombat").Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, errors.New("ID inválido")
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return User{}, ErrNoRows
	}

	return user, nil
}

func GetUserByEmail(email string) (user User, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return User{}, err
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("wombat").Collection("users")
	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, ErrNoRows
	}
	return user, nil
}

func GetUserByUsername(username string) (user User, err error) {
	client, err := db.OpenConnection()
	if err != nil {
		return User{}, err
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("wombat").Collection("users")
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return User{}, ErrNoRows
	}
	return user, nil
}


func encryptPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}
