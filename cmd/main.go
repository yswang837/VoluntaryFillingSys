package main

import (
	"fmt"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/yswang837/email"
	"github.com/yswang837/redis/config"
	"github.com/yswang837/snowflake"
)

func main() {
	config.NewConfigByFileName("redis")
	if err := snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
		return
	}
	fmt.Println(snowflake.GenID("AAB"))
	code, err := email.SendEmail("1714113169@qq.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("code", code)

}
