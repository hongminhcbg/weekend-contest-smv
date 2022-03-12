package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/weekend-contest-smv/backend/cfg"
	"github.com/hongminhcbg/weekend-contest-smv/backend/store"
)

type UserVisitedService struct {
	s   store.UserVisitedStore
	cfg *cfg.Config
}

func NewUserVisitedService(s store.UserVisitedStore, cfg *cfg.Config) *UserVisitedService {
	return &UserVisitedService{s: s, cfg: cfg}
}

func (s *UserVisitedService) Register(ctx *gin.Context) {

}
