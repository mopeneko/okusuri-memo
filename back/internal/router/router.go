package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mopeneko/okusuri-memo/back/internal/controller"
)

func SetRoutes(app *fiber.App, ctrl *controller.Controller) {
	app.Get("/health", ctrl.Health)
}
