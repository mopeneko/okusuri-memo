package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/mopeneko/okusuri-memo/back/internal/controller"
	"github.com/mopeneko/okusuri-memo/back/internal/middleware"
	"github.com/mopeneko/okusuri-memo/back/internal/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	client, err := mongo.Connect(
		context.Background(),
		options.Client().
			SetHosts([]string{os.Getenv("MONGO_SERVER") + ":27017"}).
			SetAuth(options.Credential{
				Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
				Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
			}),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to MongoDB server: %+v", err))
	}

	db := client.Database(os.Getenv("MONGO_DATABASE"))
	ctrl := controller.New(db)

	app.Use(adaptor.HTTPMiddleware(middleware.NewAuth0()))

	router.SetRoutes(app, ctrl)

	log.Fatal(app.Listen(":8080"))
}
