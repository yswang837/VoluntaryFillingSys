package user

import (
	"github.com/VoluntaryFillingSys/data/user"
	"github.com/gin-gonic/gin"
	"github.com/yswang837/snowflake"
	"net/http"
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
	var u user.User
	_ = ctx.ShouldBindJSON(&u)
	u.Uid = snowflake.GenID("AA")
	if err := c.userMysql.Add(&u); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "10001",
			"msg":  err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": "10000",
		"msg":  "success",
	})
}
