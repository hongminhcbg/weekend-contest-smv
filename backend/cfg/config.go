package cfg

import (
	"fmt"
	"os"
)

type Config struct {
	MysqlDsn string
	RedisUrl string
}

func LoadConfig() (*Config, error) {
	mysqlDsn := os.Getenv("MYSQL_DSN")
	redisUrl := os.Getenv("REDIS_URL")
	if len(mysqlDsn) == 0 {
		return nil, fmt.Errorf("mysql config is required")
	}

	if len(redisUrl) == 0 {
		return nil, fmt.Errorf("redis config is required")
	}

	return &Config{
		MysqlDsn: mysqlDsn,
		RedisUrl: redisUrl,
	}, nil
}
