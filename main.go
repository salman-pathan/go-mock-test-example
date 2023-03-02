package main

import (
	"context"
	"time"

	"mockingtest/route"
	"mockingtest/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"mockingtest/repo"
)

const (
	database = "MockingTest"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	var userRepo repo.UserRepository
	userRepo = repo.NewUserRepository(client, database)

	router := gin.Default()
	userRouterGroup := router.Group("/user")

	userService := service.NewUserService(userRepo)
	userRoute := route.NewUserRoute(userRouterGroup, &userService)

	userRoute.HandleRoutes()

	router.Run(":8000")
}
