package cache

import (
	"context"
	"sync"

	"github.com/gookit/config/v2"
	libWechatCache "github.com/silenceper/wechat/v2/cache"
)

var cache *libWechatCache.Redis
var once sync.Once

// RedisCache 获得一个微信RedisCache对象
func RedisCache() *libWechatCache.Redis {
	once.Do(func() {
		redisDNS := config.String("Redis.DNS")
		redisPwd := config.String("Redis.Password")

		options := &libWechatCache.RedisOpts{
			Host:        redisDNS,
			Password:    redisPwd,
			Database:    0,
			MaxActive:   1,
			MaxIdle:     1,
			IdleTimeout: 60,
		}

		cache = libWechatCache.NewRedis(context.Background(), options)
	})
	return cache
}
