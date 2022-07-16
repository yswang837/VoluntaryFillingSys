package main

import (
	"fmt"
	"github.com/VoluntaryFillingSys/pkg/snowflake"
	"github.com/VoluntaryFillingSys/utils"
)

func main() {
	if err := snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
		return
	}
	fmt.Println(snowflake.Node.Generate())
}
