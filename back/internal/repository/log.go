package repository

import (
	"context"
	"fmt"

	"github.com/mopeneko/okusuri-memo/back/internal/filter"
	"github.com/mopeneko/okusuri-memo/back/internal/model"
	"github.com/mopeneko/okusuri-memo/back/internal/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	logDefaultSortKey   = "created_at"
	logDefaultSortOrder = pagination.SortOrderAsc
)

type Log struct {
	collection *mongo.Collection
}

func NewLog(db *mongo.Database) *Log {
	collection := db.Collection("logs")

	return &Log{
		collection: collection,
	}
}

func (repo *Log) Find(ctx context.Context, filter *filter.Log, p *pagination.Pagination) ([]*model.Log, error) {
	bsonFilter := bson.M{}
	opts := options.Find()
	opts = opts.SetSort(bson.D{{Key: logDefaultSortKey, Value: -1}})

	if filter != nil {
	}

	if p != nil {
		sortOrder := logDefaultSortOrder
		if p.SortOrder != nil {
			sortOrder = *p.SortOrder
		}

		sortKey := logDefaultSortKey
		if p.SortKey != nil {
			sortKey = *p.SortKey
		}

		opts = opts.SetSort(bson.D{{Key: sortKey, Value: sortOrder}})

		if p.Limit != nil {
			opts = opts.SetLimit(*p.Limit)
		}

		if p.LastID != nil {
			lastID, err := primitive.ObjectIDFromHex(*p.LastID)
			if err != nil {
				return nil, fmt.Errorf("failed to create a new object ID: %w", err)
			}
			bsonFilter["_id"] = bson.M{"$gt": lastID}
		}
	}

	cur, err := repo.collection.Find(ctx, bsonFilter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return make([]*model.Log, 0), nil
		}
		return nil, fmt.Errorf("failed to find logs: %w", err)
	}

	defer cur.Close(ctx)

	res := make([]*model.Log, 0)
	if err = cur.All(ctx, &res); err != nil {
		return nil, fmt.Errorf("failed to iterate and decode documents: %w", err)
	}

	return res, nil
}

func (repo *Log) Insert(ctx context.Context, obj *model.Log) error {
	res, err := repo.collection.InsertOne(ctx, obj)
	if err != nil {
		return fmt.Errorf("failed to insert log: %w", err)
	}

	objectID, ok := (res.InsertedID).(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("failed to cast to ObjectID: %w", err)
	}

	obj.ID = objectID

	return nil
}
