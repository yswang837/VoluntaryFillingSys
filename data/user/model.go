package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid       string `gorm:"type:varchar(20);not null;" json:"uid"`     // 用户id，雪花算法生成的
	Email     string `gorm:"type:varchar(64);not null;" json:"email"`   // 邮箱
	Password  string `gorm:"type:varchar(20);not null" json:"password"` // 密码
	Phone     string `gorm:"type:varchar(11);" json:"phone"`            // 电话号码
	NickName  string `gorm:"type:varchar(20);" json:"nick_name"`        // 昵称
	HeadImage string `gorm:"type:varchar(64);" json:"head_image"`       // 头像
	Role      int    `gorm:"type int;default:1" json:"role"`            // 角色,用于权限管理,1普通用户,2管理员,3超级管理员
	Grade     int    `gorm:"type int;default:1" json:"grade"`           // 账号等级
	Summary   string `gorm:"type varchar(200);" json:"summary"`         // 个人简介
	Sex       int    `gorm:"type int" json:"sex"`                       // 1男,2女,3不详
	Age       uint   `gorm:"type unit" json:"age"`                      // 年龄
}
