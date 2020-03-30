package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SergeyDavidenko/subscription/config"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// RedisConnector ...
type RedisConnector struct {
	c *redis.Client
}

var (
	Client = &RedisConnector{}
)

//RedisInitialize get the redis client
func RedisInitialize() *RedisConnector {
	log.Debug("try connect to redis")
	redisHost := fmt.Sprintf("%s:%d",
		config.Conf.Redis.Address,
		config.Conf.Redis.Port)
	c := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})
	log.Debug("initialize redis connect")
	if err := c.Ping().Err(); err != nil {
		log.Fatal("Unable to connect to redis " + err.Error())
	}
	log.Debug("try redis ping")
	log.Info("Redis success connect to ", redisHost)
	Client.c = c
	return Client
}

//GetKey get key
func (client *RedisConnector) GetKey(key string) ([]byte, error) {
	log.Debug("Redis get key")
	val, err := client.c.Get(key).Bytes()
	if err == redis.Nil || err != nil {
		log.Warn(err.Error())
		return nil, err
	}
	return val, nil
}

//SetKey set key
func (client *RedisConnector) SetKey(key string, value interface{}, expiration int) error {
	log.Debug("Start redis SetKey func")
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	exp := time.Duration(expiration) * time.Second
	err = client.c.Set(key, cacheEntry, exp).Err()

	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//CheckKey get key
func (client *RedisConnector) CheckKey(key string) error {
	log.Debug("Redis get key")
	err := client.c.Get(key).Err()
	if err == redis.Nil || err != nil {
		log.Warn(err.Error())
		return err
	}
	return nil
}
