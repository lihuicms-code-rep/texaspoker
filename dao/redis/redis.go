package redis

import (
	"github.com/go-redis/redis"
	"github.com/lihuicms-code-rep/texaspoker/log"
)

var redisClient *redis.Client

func InitRedisClient(addr, password string, db int) error {
	log.Console.Infof("redis client init, addr:%s, password:%s, dbnumber:%d", addr, password, db)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Console.Errorf("redis client ping error:%+v", err)
		return err
	}

	log.Console.Info("redis client ok")
	return nil
}
