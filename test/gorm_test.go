package test

import (
	"fmt"
	dao "github.com/VoluntaryFillingSys/dao/user"
	"testing"
)


func TestGorm(t *testing.T) {
	users := dao.MyDao.ListUser()
	fmt.Println(users)
}
