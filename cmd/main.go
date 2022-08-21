package main

import (
	routes "github.com/VoluntaryFillingSys/router"
	"github.com/VoluntaryFillingSys/utils"
)

func main() {
	//if err := snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
	//	return
	//}
	//fmt.Println(snowflake.GenID("AAB"))
	//code, err := email.SendEmail("1714113169@qq.com")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("code", code)
	r := routes.InitRouter()
	panic(r.Run(utils.HttpPort))

}
