package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mopeneko/okusuri-memo/back/internal/pagination"
	"github.com/mopeneko/okusuri-memo/back/internal/service"
	"github.com/samber/lo"
)

type GetLogsRequest struct {
	SortKey   string `query:"sort_key"`
	SortOrder string `query:"sort_order"`
	Limit     int64  `query:"limit"`
	LastID    string `query:"last_id"`
}

func (ctrl *Controller) GetLogs(c fiber.Ctx) error {
	req := new(GetLogsRequest)
	if err := c.Bind().Query(req); err != nil {
		return fmt.Errorf("failed to bind query: %w", err)
	}

	var sortOrder *pagination.SortOrder
	switch req.SortOrder {
	case "asc":
		sortOrder = lo.ToPtr(pagination.SortOrderAsc)

	case "desc":
		sortOrder = lo.ToPtr(pagination.SortOrderDesc)
	}

	logs, err := ctrl.logService.GetMulti(
		c.UserContext(),
		service.GetMultiParams{
			SortKey:   lo.EmptyableToPtr(req.SortKey),
			SortOrder: sortOrder,
			Limit:     lo.EmptyableToPtr(req.Limit),
			LastID:    lo.EmptyableToPtr(req.LastID),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to get all: %w", err)
	}

	return c.JSON(logs)
}
