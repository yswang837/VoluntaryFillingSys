package main

import (
	"fmt"
	"github.com/VoluntaryFillingSys/pkg/snowflake"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/yswang837/email"
)

func main() {
	if err := snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
		return
	}
	fmt.Println(snowflake.Node.Generate())
	code, err := email.SendEmail("1714113169@qq.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("code", code)

}

//bxvjwuyyeoqfdced
