package controller

import "github.com/gofiber/fiber/v3"

type HealthResponse struct {
	Status string `json:"status"`
}

func (ctrl *Controller) Health(c fiber.Ctx) error {
	return c.JSON(&HealthResponse{
		Status: "OK",
	})
}
