package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mopeneko/okusuri-memo/back/internal/controller"
	"github.com/mopeneko/okusuri-memo/back/internal/router"
)

func main() {
	app := fiber.New()

	ctrl := controller.New()
	router.SetRoutes(app, ctrl)

	log.Fatal(app.Listen(":8080"))
}
