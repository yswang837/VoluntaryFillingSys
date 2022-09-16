package main

import (
	"github.com/spf13/cobra"
	"github.com/yswang837/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid             string         `gorm:"type:varchar(50);primary_key;not null;" json:"uid"`    // 用户id，雪花算法生成的
	Email           string         `gorm:"type:varchar(64);unique_index;not null;" json:"email"` // 邮箱
	Password        string         `gorm:"type:varchar(20);not null" json:"password"`            // 密码
	Phone           string         `gorm:"type:varchar(11);" json:"phone"`                       // 电话号码
	NickName        string         `gorm:"type:varchar(20);" json:"nick_name"`                   // 昵称
	Role            int            `gorm:"type int;default:1" json:"role"`                       // 角色,用于权限管理,1普通用户,2管理员,3超级管理员
	Grade           int            `gorm:"type int;default:1" json:"grade"`                      // 账号等级
	Summary         string         `gorm:"type varchar(200);" json:"summary"`                    // 个人简介
	Sex             int            `gorm:"type int" json:"sex"`                                  // 1男,2女,3不详
	Age             int            `gorm:"type int" json:"age"`                                  // 年龄
	IsCertified     string         `gorm:"type varchar(5);default:'no'" json:"is_certified"`     //是否认证
	DataIntegrity   int            `gorm:"type int" json:"data_integrity"`                       //资料完整度
	CreateAt        time.Time      `gorm:"type datetime;not null" json:"create_at"`              // 创建时间
	ActiveAt        time.Time      `gorm:"type datetime;not null" json:"active_at"`              // 活跃时间
	UpdateAt        time.Time      `gorm:"type datetime;not null" json:"update_at"`              // 更新时间
	DeleteAt        gorm.DeletedAt `gorm:"type datetime;" json:"delete_at"`                      // 软删除时间
	CustomTableName string         `gorm:"-"`
}

func NewUser(tableName string) *User {
	return &User{CustomTableName: tableName}
}

func (u *User) TableName() string {
	return u.CustomTableName
}

// genMysqlTableCmd represents the genMysqlTable command
var genMysqlTableCmd = &cobra.Command{
	Use:   "gen-mysql-table",
	Short: "生成 mysql 表",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := mysql.NewClient("user")
		if err != nil {
			return
		}
		c.Master().AutoMigrate(NewUser("user_0"), NewUser("user_1"))
		c.Slave().AutoMigrate(NewUser("user_0"), NewUser("user_1"))
	},
}
