package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func NewUser(name string, email string, password string) User {
	return User{
		Id:       primitive.NewObjectID().Hex(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}
