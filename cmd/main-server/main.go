package main

import (
	"github.com/VoluntaryFillingSys/router"
)

func main() {
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
