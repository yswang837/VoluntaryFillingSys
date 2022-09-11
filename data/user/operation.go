package user

import (
	"fmt"
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
	return c.Client.Master().Model(&User{}).Scopes(selectTable(uid))
}

func (c *ClientMysql) slave(uid string) *gorm.DB {
	return c.Client.Slave().Model(&User{}).Scopes(selectTable(uid))
}

func whereEmail(uid string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("email = ?", uid)
	}
}

func selectTable(uid string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(fmt.Sprintf("user_%d", crc([]byte(uid))%2))
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

func (c *ClientMysql) CheckEmailExist(uid, email string) bool {
	u := User{}
	c.slave(uid).Scopes(whereEmail(email)).First(&u)
	if u.Email != "" {
		return true
	}
	return false

}

func (c *ClientMysql) GetUser() {

}

func (c *ClientMysql) GetUsers() {

}
