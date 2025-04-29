package cache

import (
	"context"
	"ranbao/global"
	"ranbao/ini"
	"testing"
	"time"
)

func TestCacheMail_Set(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		ctx   context.Context
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "redis基础测试",
			fields: fields{
				key: "why",
			},
			args: args{
				ctx:   context.Background(),
				value: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ini.InitRDB()

			_, err := global.Rdb.Set(tt.args.ctx, tt.fields.key, tt.args.value, 5*time.Minute).Result()
			if err != nil {
				t.Error(err)
			}
			res, err := global.Rdb.Get(tt.args.ctx, tt.fields.key).Result()
			if err != nil {
				t.Error(err)
			}
			t.Log(res)
			if res != "123" {
				t.Error("no")
			}
		})
	}
}
