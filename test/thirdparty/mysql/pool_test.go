package mysql

import (
	"fmt"
	"github.com/VoluntaryFillingSys/common"
	"github.com/VoluntaryFillingSys/thirdparty/mysql_pool"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewConn(t *testing.T) {
	wg := sync.WaitGroup{}
	pool, err := mysql_pool.NewPool(10)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(20)
	fmt.Println("开启20个协程获取连接")
	for i := 0; i < 20; i++ {
		go ReleaseConn(pool, wg)
	}
	wg.Wait()
}

func ReleaseConn(pool *mysql_pool.MysqlPool, wg sync.WaitGroup) {
	s := rand.Int63n(2)
	time.Sleep(time.Duration(s) * time.Second)
	conn, err := pool.GetConnFromPool()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("连接池连接数",len(pool.ConnChan))
	time.Sleep(time.Duration(s) * time.Second)
	pool.ReleaseConn(conn)
	wg.Done()
}

func TestSqlDriver(t *testing.T) {

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

