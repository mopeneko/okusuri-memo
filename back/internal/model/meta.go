package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meta struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at"`
	DeletedBy *string            `json:"deleted_by" bson:"deleted_by"`
}
