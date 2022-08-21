package test

import (
	"fmt"
	"github.com/yswang837/snowflake"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	startTime := "2022-08-12"
	machinedId := "1"
	if err := snowflake.Init(startTime, machinedId); err != nil {
		return
	}
	id := snowflake.GenID("aa_")
	fmt.Println(id)

}
