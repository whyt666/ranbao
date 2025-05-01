package ini

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ranbao/global"
)

func InitRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("127.0.0.1:%s", Conf.RedisConfig.Port),
	})
	fmt.Println(rdb)
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	global.Rdb = rdb
}
