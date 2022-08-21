package user

import (
	"github.com/yswang837/mysql"
)

type ClientMysql struct {
	Client *mysql.Client
}

func (c *ClientMysql) NewMysqlClient() (*ClientMysql, error) {
	var err error
	c.Client, err = mysql.NewClient()
	if err != nil {
		return nil, err
	}
	return &ClientMysql{Client: c.Client}, nil
}

// todo mysql的基本操作
func (c *ClientMysql) AddUser(user *User) {

	c.Client.Db.Create(user)
}
