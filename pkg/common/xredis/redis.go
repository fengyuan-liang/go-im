package xredis

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"go-im/config"
)

var (
	Cli *RedisClient
)

type RedisClient struct {
	Client   *redis.Client
	RedsSync *redsync.Redsync
	Prefix   string
}

func NewRedisClient(cfg *config.RedisStruct) *redis.Client {
	var (
		client   *redis.Client
		pool     redsyncredis.Pool
		redsSync *redsync.Redsync
		err      error
	)
	// 单机redis
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.URL + ":" + cfg.PORT,
		Password: cfg.PASSWORD,
		DB:       cfg.DB,
	})
	// 判断是否能够链接到redis
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	// redis 锁
	pool = goredis.NewPool(client)
	redsSync = redsync.New(pool)
	Cli = &RedisClient{client, redsSync, cfg.Prefix}

	return client
}
