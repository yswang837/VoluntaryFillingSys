package user

import (
	"github.com/yswang837/redis"
)

type ClientRedis struct {
	client *redis.Client
}

func NewRedisClient() (*ClientRedis, error) {
	c := &ClientRedis{}
	var err error
	//conf, err := config.NewConfigByFileName("redis")
	//if err != nil {
	//	return nil, err
	//}

	c.client, err = redis.NewClient()
	if err != nil {
		return nil, err
	}
	return &ClientRedis{client: c.client}, nil
}

func (c *ClientRedis) SetTTL(uid string, val string, ttl string) bool {
	_, err := c.client.EXPIRE(uid, val, ttl)
	if err != nil {
		return false
	}
	return true
}
