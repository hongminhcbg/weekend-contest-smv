package store

import (
	"context"
	"github.com/hongminhcbg/weekend-contest-smv/backend/models"
)

type UserVisitedStore interface {
	Save(ctx context.Context, r *models.UserVisited) error
	GetMostVisitedByIp(ctx context.Context, limit, offset int) ([]*models.UserVisited, error)
}
