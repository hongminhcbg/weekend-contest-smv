package store

import (
	"context"
	"gorm.io/gorm"

	"github.com/hongminhcbg/weekend-contest-smv/backend/models"
)

type UserVisitedStore interface {
	Save(ctx context.Context, r *models.UserVisited) error
	GetMostVisitedByIp(ctx context.Context) ([]*models.CntVisitTimesByIp, error)
	GetUsersVisited(ctx context.Context, lastId, limit int) ([]*models.UserVisited, error)
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

func (s *userVisitedStoreImpl) GetMostVisitedByIp(ctx context.Context) ([]*models.CntVisitTimesByIp, error) {
	result := make([]*models.CntVisitTimesByIp, 0)
	err := s.WithContext(ctx).
		Raw("select count(*) as num, ip from user_visited group by ip order by num desc LIMIT 100").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userVisitedStoreImpl) GetUsersVisited(ctx context.Context, lastId, limit int) ([]*models.UserVisited, error) {
	users := make([]*models.UserVisited, 0, limit)
	err := s.WithContext(ctx).Where("id < ?", lastId).Order("id DESC").Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

