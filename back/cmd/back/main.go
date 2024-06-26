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

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI(os.Getenv("MONGODB_URI")).
			SetServerAPIOptions(serverAPI),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to MongoDB server: %+v", err))
	}

	defer client.Disconnect(context.Background())

	db := client.Database(os.Getenv("MONGO_DATABASE"))
	ctrl := controller.New(db)

	app.Use(adaptor.HTTPMiddleware(middleware.NewAuth0()))

	router.SetRoutes(app, ctrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
