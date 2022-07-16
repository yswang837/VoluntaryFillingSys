package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string

	StartTime  string
	MachinedId string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiNiuServer string
)

func init() {
	f, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadServer(f)
	LoadData(f)
	LoadSnowFlake(f)
	LoadQiNiu(f)
}
func LoadServer(f *ini.File) {
	AppMode = f.Section("server").Key("AppMode").MustString("debug")
	HttpPort = f.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = f.Section("server").Key("JwtKey").MustString("3247fifths33")
}

func LoadData(f *ini.File) {
	Db = f.Section("database").Key("Db").MustString("mysql")
	DbHost = f.Section("database").Key("DbHost").MustString(":localhost")
	DbPort = f.Section("database").Key("DbPort").MustString("10000")
	DbName = f.Section("database").Key("DbName").MustString(":gopher")
	DbUser = f.Section("database").Key("DbUser").MustString("root")
	DbPassword = f.Section("database").Key("DbPassword").MustString("root")
}

func LoadSnowFlake(f *ini.File) {
	StartTime = f.Section("snowflake").Key("StartTime").MustString("2022-07-16")
	MachinedId = f.Section("snowflake").Key("MachinedId").MustString("1")
}

func LoadQiNiu(f *ini.File) {
	AccessKey = f.Section("qiniu").Key("AccessKey").String()
	SecretKey = f.Section("qiniu").Key("SecretKey").String()
	Bucket = f.Section("qiniu").Key("Bucket").String()
	QiNiuServer = f.Section("qiniu").Key("QiNiuServer").String()

}
