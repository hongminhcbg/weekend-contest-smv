package store

import (
	"context"
	"gorm.io/gorm"

	"github.com/hongminhcbg/weekend-contest-smv/backend/models"
)

type UserVisitedStore interface {
	Save(ctx context.Context, r *models.UserVisited) error
	GetMostVisitedByIp(ctx context.Context, limit, offset int) ([]*models.UserVisited, error)
}

type userVisitedStoreImpl struct {
	*gorm.DB
}

func NewUserVisitedStore(db *gorm.DB) UserVisitedStore {
	return &userVisitedStoreImpl{db}
}

func (s *userVisitedStoreImpl) Save(ctx context.Context, r *models.UserVisited) error {
	return s.WithContext(ctx).Save(r).Error
}

func (s *userVisitedStoreImpl) GetMostVisitedByIp(ctx context.Context, limit, offset int) ([]*models.UserVisited, error) {
	return nil, nil
}
