package controller

import (
	"github.com/mopeneko/okusuri-memo/back/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	logService *service.Log
}

func New(db *mongo.Database) *Controller {
	return &Controller{
		logService: service.NewLog(db),
	}
}
