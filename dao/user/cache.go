package user

import (
	"github.com/yswang837/redis"
)

type Client struct {
	client *redis.Client
}

func NewClient() (*Client, error) {
	c := &Client{}
	var err error
	//conf, err := config.NewConfigByFileName("redis")
	//if err != nil {
	//	return nil, err
	//}

	c.client, err = redis.NewClient()
	if err != nil {
		return nil, err
	}
	return &Client{client: c.client}, nil
}

func (c *Client) SetTTL(uid string, val string, ttl string) bool {
	_, err := c.client.EXPIRE(uid, val, ttl)
	if err != nil {
		return false
	}
	return true
}
