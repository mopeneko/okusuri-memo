package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mopeneko/okusuri-memo/back/internal/model"
)

type PostLogsRequest struct {
	Medicines []string `json:"medicines"`
}

func (ctrl Controller) PostLogs(c fiber.Ctx) error {
	req := new(PostLogsRequest)
	if err := c.Bind().JSON(req); err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}

	obj := &model.Log{
		Medicines: req.Medicines,
		Meta:      model.NewMeta(),
	}

	if err := ctrl.logService.Insert(c.UserContext(), obj); err != nil {
		return fmt.Errorf("failed to insert log: %w", err)
	}

	return c.JSON(obj)
}
