package test

import (
	"fmt"
	"testing"
	"github.com/yswang837/snowflake"
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