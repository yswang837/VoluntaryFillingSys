package user

import (
	"fmt"
	"github.com/VoluntaryFillingSys/data/user"
	"github.com/gin-gonic/gin"
)

type Consumer struct {
	userRedis *user.ClientRedis
	userMysql *user.ClientMysql
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Init() error {
	var err error
	if c.userMysql, err = user.NewMysqlClient(); err != nil {
		return err
	}
	if c.userRedis, err = user.NewRedisClient(); err != nil {
		return err
	}
	return nil
}

func (c *Consumer) AddUser(ctx *gin.Context) {
	fmt.Println("Add......")
	var u user.User
	_ = ctx.ShouldBindJSON(&u)
}
