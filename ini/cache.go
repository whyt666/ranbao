package ini

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ranbao/global"
)

func InitRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	fmt.Println(rdb)
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	global.Rdb = rdb
}
