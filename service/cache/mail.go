package cache

import (
	"context"
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
	err := global.Rdb.Set(ctx, c.key+":"+mail, value, 3*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheMail) Get(ctx context.Context, mail string) (any, error) {
	res, err := global.Rdb.Get(ctx, c.key+":"+mail).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
