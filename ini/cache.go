package ini

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"ranbao/global"
)

func InitRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	fmt.Println(rdb)

	global.Rdb = rdb
}
