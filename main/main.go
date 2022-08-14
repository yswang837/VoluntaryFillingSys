package main

import (
	"fmt"
	"github.com/VoluntaryFillingSys/common"
	"github.com/VoluntaryFillingSys/thirdparty/mysql_pool"
	"github.com/go-sql-driver/mysql"
)

func main()  {
	var conf mysql.Config

	common.InitMysqlConfig(&conf)

	db, _ := mysql_pool.NewMysqlClient(conf)


	sqlStr := "select email, password from users"
	rows , _:= db.Query(sqlStr)
	fmt.Println(rows)
	var email, password string

	for rows.Next() {
		rows.Scan(&email, &password)
		fmt.Println(email, password)
	}

}