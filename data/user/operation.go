package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yswang837/mysql"
)

type ClientMysql struct {
	Client *mysql.Client
}

func NewMysqlClient() (*ClientMysql, error) {
	var err error
	c := &ClientMysql{}
	c.Client, err = mysql.NewClient()
	if err != nil {
		return nil, err
	}
	return &ClientMysql{Client: c.Client}, nil
}

func (c *ClientMysql) master(uid string) *gorm.DB {
	return c.Client.Db.Model(&User{}).Scopes(selectTable(uid))
}

func (c *ClientMysql) slave(uid string) *gorm.DB {
	return c.Client.Db.Model(&User{}).Scopes(selectTable(uid))
}

func selectTable(uid string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(fmt.Sprintf("user_%d", crc([]byte(uid))%2))
	}
}

// todo mysql的基本操作
func (c *ClientMysql) Add(user *User) {

	c.Client.Db.Create(user)
}

func (c *ClientMysql) Delete(uid string) {
	c.Client.Db.Delete(&User{}, "uid = ?", uid)
}

func (c *ClientMysql) Update() {

}

func (c *ClientMysql) GetUser() {

}

func (c *ClientMysql) GetUsers() {

}
