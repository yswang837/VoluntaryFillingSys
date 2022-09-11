package user

import (
	"github.com/VoluntaryFillingSys/data/user"
	"github.com/VoluntaryFillingSys/utils"
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
	if err = snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
		return err
	}
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
	u.Uid = snowflake.GenID(utils.PrefixUser)
	if c.CheckEmailExist(u.Uid, u.Email) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": utils.ErrorEmailUsed,
			"msg":  "邮箱已注册",
		})
		return
	}
	if err := c.userMysql.Add(&u); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": utils.MysqlErr,
			"msg":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "用户注册成功",
	})
	return
}

func (c *Consumer) CheckEmailExist(uid, email string) bool {
	return c.userMysql.CheckEmailExist(uid, email)
}
