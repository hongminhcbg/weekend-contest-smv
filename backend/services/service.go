package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

func (s *UserVisitedService) ListUsersVisited(ctx *gin.Context) {
	lastId := utils.GetQueryInt(ctx, "last_id", 9e10)
	itemPerPage := utils.GetQueryInt(ctx, "item_per_page", 10)
	rCtx := ctx.Request.Context()

	users, err := s.store.GetUsersVisited(rCtx, lastId, itemPerPage)
	if err != nil {
		log.Println(err, "ListUsersVisited get in db error")
		utils.ResponseError(ctx, erp.ERR_INTENAL_SERVER)
		return
	}

	result := models.GetVisitedUsersResponse{
		HasMoreData: !(len(users) < itemPerPage),
		Users:       users,
	}

	utils.ResponseData(ctx, result)
}

func (s *UserVisitedService) MostVisitors(ctx *gin.Context) {
	countVisitUsers := make([]*models.CntVisitTimesByIp, 0)
	rCtx := ctx.Request.Context()

	rawResp, err := s.cache.Get(rCtx, erp.MOST_VISITED_USERS).Result()
	if err == nil {
		log.Println("[INFOR] MostVisitors response from cache")
		err = json.NewDecoder(strings.NewReader(rawResp)).Decode(&countVisitUsers)
		if err == nil {
			utils.ResponseData(ctx, &models.MostVisitedUserResponse{Users: countVisitUsers})
			return
		}
		return
	}

	countVisitUsers, err = s.store.GetMostVisitedByIp(rCtx)
	if err != nil {
		log.Println(err, "MostVisitors get in db error")
		utils.ResponseError(ctx, erp.ERR_INTENAL_SERVER)
		return
	}

	go func() {
		b, _ := json.Marshal(countVisitUsers)
		s.cache.Set(context.Background(), erp.MOST_VISITED_USERS, string(b), 40*time.Second)
	}()

	utils.ResponseData(ctx, &models.MostVisitedUserResponse{Users: countVisitUsers})
}
