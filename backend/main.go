package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hongminhcbg/weekend-contest-smv/backend/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hongminhcbg/weekend-contest-smv/backend/cfg"
	"github.com/hongminhcbg/weekend-contest-smv/backend/services"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, api-key")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
func main() {
	fmt.Println("Hello world")
	conf, err := cfg.LoadConfig()
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(conf, "", "\t")
	fmt.Println("start server with config: ", string(b))

	db := mustConnectMysql(conf)
	cache := mustConnectRedis(conf)
	userTrackingStore := store.NewUserVisitedStore(db)
	service := services.NewUserVisitedService(userTrackingStore, conf, cache)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", service.Register)
	r.Run()
}

func mustConnectMysql(c *cfg.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.MysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// force a connection and test that it worked
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func mustConnectRedis(c *cfg.Config) *redis.Client {
	opts, err := redis.ParseURL(c.RedisUrl)
	if err != nil {
		log.Fatal("parse redis url error", err)
	}

	cli := redis.NewClient(opts)
	err = cli.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("ping redis error", err)
	}

	return cli
}
