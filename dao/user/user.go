package user

//
//import (
//	"fmt"
//	"github.com/VoluntaryFillingSys/entity/user"
//
//	"gorm.io/gorm"
//)
//
//type IUser interface {
//	AddUser(user *user.User)
//	ListUser() []user.User
//}
//
//type Dao struct {
//	Db *gorm.DB
//}
//
//var MyDao IUser
//
//// var DB *gorm.DB
//
////func init() {
////	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
////	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
////	if err != nil {
////		log.Fatal("fail to init mysql db, err: ", err)
////	}
////	MyDao = &Dao{Db: db}
////	db.AutoMigrate(&user.User{})
////}
//
//func (dao *Dao) AddUser(user *user.User) {
//	dao.Db.Create(user)
//}
//
//func (dao *Dao) ListUser() []user.User {
//	var users []user.User
//	res := dao.Db.Find(&users)
//	fmt.Println(res)
//	return users
//}
