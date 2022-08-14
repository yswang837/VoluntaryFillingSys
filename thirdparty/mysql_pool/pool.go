package mysql_pool

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"sync"
)

// 连接池定义
type MysqlPool struct {
	sync.Mutex                // 保证连接池的线程安全
	Size       int            // 连接池的链接数量
	ConnChan   chan io.Closer // 存储链接的通道
	IsClose	   bool           // 连接池是否关闭
	ctx        context.Context
}

// NewPool 初始化连接池
func NewPool(size int) (*MysqlPool, error) {
	if size <= 0 {
		return nil, errors.New("mysql_pool pool size is invalid")
	}
	pool := &MysqlPool{
		ConnChan: make(chan io.Closer, size),
		ctx:      context.Background(),
	}
	return pool, nil
}

// GetConnFromPool 从连接池里获取连接
func (pool *MysqlPool) GetConnFromPool() (io.Closer, error) {
	if pool.IsClose {
		return nil, errors.New("mysql_pool pool has closed")
	}
	select {
	case conn, ok := <- pool.ConnChan:
		if !ok {
			return nil, errors.New("mysql_pool pool has closed")
		}
		return conn, nil
	default:
		// 连接池中没有连接， 新建连接
		return pool.NewConn(pool.ctx)
	}
}

// NewConn 构建新连接
func (pool *MysqlPool) NewConn(ctx context.Context) (io.Closer, error) {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/vfs?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("连接数据库失败")
	}
	conn, _ := db.Conn(ctx)
	select {
	case pool.ConnChan <- conn:
		log.Println("将连接放入连接池")
	default:
		conn.Close()
		log.Println("连接池满，连接丢弃")
	}
	return conn, nil
}

// ReleaseConn 释放连接
func (pool *MysqlPool) ReleaseConn(conn io.Closer) error {
	pool.Lock()
	defer pool.Unlock()

	if pool.IsClose {
		return errors.New("mysql_pool pool has closed")
	}

	select {
	case pool.ConnChan <- conn:
		log.Println("将连接放入连接池")
	default:
		conn.Close()
		log.Println("连接池满，连接丢弃")
	}
	return nil
}

// ClosePool 关闭连接池
func (pool *MysqlPool) ClosePool() error {
	pool.Lock()
	defer pool.Unlock()

	if pool.IsClose {
		return errors.New("mysql_pool pool has closed")
	}
	pool.IsClose = true

	// 先关闭通道里的连接
	for conn := range pool.ConnChan {
		conn.Close()
	}
	// 再关闭通道
	close(pool.ConnChan)
	return nil
}


// go-sql-driver/mysql_pool 类已实现mysql连接池
func NewMysqlClient(c mysql.Config) (*sql.DB, error) {

	//default_dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Passwd, c.Addr, c.Net, c.DBName)

	db, _ := sql.Open("mysql", dsn)
	// 设置最大连接数
	db.SetMaxOpenConns(10)

	db.SetMaxIdleConns(2)//连接数据库查询

	if err := db.Ping(); err != nil {
		fmt.Println("连接失败")
		return nil, errors.New("数据库连接失败")
	}
	return db, nil
}
