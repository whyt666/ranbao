package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Rdb *redis.Client
)
