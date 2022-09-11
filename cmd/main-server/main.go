package main

import (
	"github.com/VoluntaryFillingSys/router"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/yswang837/snowflake"
)

func main() {
	if err := snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
		return
	}
	if err := routes.StartGinService(); err != nil {
		return
	}

	//fmt.Println(snowflake.GenID("AAB"))
	//code, err := email.SendEmail("1714113169@qq.com")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("code", code)
}
