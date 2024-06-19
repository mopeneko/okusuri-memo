package service

import (
	"context"

	"github.com/mopeneko/okusuri-memo/back/internal/model"
	"github.com/mopeneko/okusuri-memo/back/internal/pagination"
	"github.com/mopeneko/okusuri-memo/back/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	repo *repository.Log
}

func NewLog(db *mongo.Database) *Log {
	return &Log{
		repo: repository.NewLog(db),
	}
}

type GetMultiParams struct {
	SortKey   *string
	SortOrder *pagination.SortOrder
	Limit     *int64
	LastID    *string
}

func (s *Log) GetMulti(ctx context.Context, params GetMultiParams) ([]*model.Log, error) {
	return s.repo.Find(ctx, nil, &pagination.Pagination{
		SortKey:   params.SortKey,
		SortOrder: params.SortOrder,
		Limit:     params.Limit,
		LastID:    params.LastID,
	})
}

func (s *Log) Insert(ctx context.Context, obj *model.Log) error {
	return s.repo.Insert(ctx, obj)
}
