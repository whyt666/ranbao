package cache

import (
	"context"
	"go.uber.org/zap"
	"ranbao/global"
	"time"
)

var CacheMailCtl *CacheMail

func init() {
	CacheMailCtl = &CacheMail{
		key: "mailCode",
	}
}

type CacheMail struct {
	key string
}

func (c *CacheMail) Set(ctx context.Context, mail string, value any) error {
	_, err := global.Rdb.Set(ctx, c.key+":"+mail, value, 30*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheMail) Get(ctx context.Context, mail string) (any, error) {
	res, err := global.Rdb.Get(ctx, c.key+":"+mail).Result()
	zap.S().Info(c.key + ":" + mail)
	if err != nil {
		return res, err
	}
	return res, nil
}
