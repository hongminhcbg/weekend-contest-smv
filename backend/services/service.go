package services

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/hongminhcbg/weekend-contest-smv/backend/cfg"
	"github.com/hongminhcbg/weekend-contest-smv/backend/erp"
	"github.com/hongminhcbg/weekend-contest-smv/backend/models"
	"github.com/hongminhcbg/weekend-contest-smv/backend/store"
	"github.com/hongminhcbg/weekend-contest-smv/backend/utils"
)

type UserVisitedService struct {
	store store.UserVisitedStore
	cfg   *cfg.Config
	cache *redis.Client
}

func NewUserVisitedService(s store.UserVisitedStore, cfg *cfg.Config, cache *redis.Client) *UserVisitedService {
	return &UserVisitedService{store: s, cfg: cfg, cache: cache}
}

func (s *UserVisitedService) Register(ctx *gin.Context) {
	var r models.UserVisited
	sCtx := ctx.Request.Context()
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		log.Println(err, "bind json error")
		utils.ResponseError(ctx, erp.ERR_BAD_REQUEST)
		return
	}

	if len(r.Ip) == 0 || len(r.Location) == 0 {
		log.Println("[ERROR], ip or location empty")
		utils.ResponseError(ctx, erp.ERR_BAD_REQUEST)
		return
	}

	key := fmt.Sprintf(erp.FORMAT_RATELIMIT_BY_IP, r.Ip)
	err = s.cache.Get(sCtx, key).Err()
	if err == nil {
		log.Println("[ERROR], rate limit")
		utils.ResponseError(ctx, erp.ERR_RATE_LIMIT)
		return
	}

	go func() {
		s.cache.Set(sCtx, key, "1", 3*time.Second)
	}()

	err = s.store.Save(sCtx, &r)
	if err != nil {
		log.Println(err, "internal server error")
		utils.ResponseError(ctx, erp.ERR_INTENAL_SERVER)
		return
	}

	utils.ResponseData(ctx, r)
}
