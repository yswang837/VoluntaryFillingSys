package user

import (
	"fmt"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/jinzhu/gorm"
	"github.com/yswang837/mysql"
	"time"
)

type ClientMysql struct {
	Client *mysql.Client
}

func NewMysqlClient() (*ClientMysql, error) {
	var err error
	c := &ClientMysql{}
	c.Client, err = mysql.NewClient("user")
	if err != nil {
		return nil, err
	}
	return &ClientMysql{Client: c.Client}, nil
}

func (c *ClientMysql) master(uid string) *gorm.DB {
	return c.Client.Master().Model(&User{}).Scopes(shardTable(uid))
}

func (c *ClientMysql) slave(uid string) *gorm.DB {
	return c.Client.Slave().Model(&User{}).Scopes(shardTable(uid))
}

func whereEmail(email string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("email = ?", email)
	}
}

func shardTable(uid string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(fmt.Sprintf("user_%d", crc([]byte(uid))%utils.TableNumOfUser))
	}
}

// todo mysql的基本操作
func (c *ClientMysql) Add(user *User) error {
	timeNow := time.Now()
	user.CreateAt = timeNow
	user.ActiveAt = timeNow
	user.UpdateAt = timeNow
	if err := c.master(user.Uid).Create(user); err.Error != nil {
		return err.Error
	}
	return nil
}

func (c *ClientMysql) Delete(uid string) {
	c.master("aaa").Delete(&User{}, "uid = ?", uid)
}

func (c *ClientMysql) Update() {

}

func (c *ClientMysql) CheckEmailExist(email string) bool {
	u := &User{}
	uids := []string{utils.UidOfUser0, utils.UidOfUser1}
	for _, uid := range uids {
		if err := c.slave(uid).Scopes(whereEmail(email)).First(u).Error; err != nil {
			return false
		}
		if u.Email != "" {
			return true
		}
	}
	return false

}

func (c *ClientMysql) GetUser() {

}

func (c *ClientMysql) GetUsers() {

}
