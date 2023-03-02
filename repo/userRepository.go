package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model "mockingtest/model"
)

const (
	ColUser = "User"
)

type UserRepository interface {
	AddUser(ctx context.Context, user model.User) (userId string, err error)
	RemoveUser(ctx context.Context, userId string) error
	UpdateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, userId string) (model.User, error)
}

type userRepository struct {
	client   *mongo.Client
	database string
}

func NewUserRepository(client *mongo.Client, database string) UserRepository {
	return &userRepository{
		client:   client,
		database: database,
	}
}

func (repo *userRepository) AddUser(ctx context.Context, user model.User) (userId string, err error) {
	res, err := repo.client.Database(repo.database).Collection(ColUser).InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(string), err
}

func (repo *userRepository) RemoveUser(ctx context.Context, userId string) error {
	filter := bson.M{
		"_id": userId,
	}
	_, err := repo.client.Database(repo.database).Collection(ColUser).DeleteOne(ctx, filter)
	return err
}

func (repo *userRepository) UpdateUser(ctx context.Context, user model.User) error {
	_, err := repo.client.Database(repo.database).Collection(ColUser).UpdateByID(ctx, user.Id, user)
	return err
}

func (repo *userRepository) GetUser(ctx context.Context, userId string) (model.User, error) {
	var user model.User
	filter := bson.M{
		"_id": userId,
	}
	err := repo.client.Database(repo.database).Collection(ColUser).FindOne(ctx, filter).Decode(&user)
	return user, err
}
