package common

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func init()  {
	InitFileLoad()
}

var Ini *ini.File

func InitFileLoad() (err error) {
	Ini, err = ini.Load("/Users/wang/Test/VoluntaryFillingSys/config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
		return err
	}
	return nil
}

func InitMysqlConfig(conf *mysql.Config) {

	conf.DBName = Ini.Section("mysql").Key("DbName").MustString("vfs")
	conf.Addr = Ini.Section("mysql").Key("DbHost").MustString(":localhost")
	conf.Net = Ini.Section("mysql").Key("DbPort").MustString("3306")
	conf.User = Ini.Section("mysql").Key("DbUser").MustString("root")
	conf.Passwd = Ini.Section("mysql").Key("DbPassword").MustString("root")

}