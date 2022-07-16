package user

import "time"

type User struct {
	Uid        string    `gorm:"type:varchar(20);not null;" json:"uid"`     // 用户id，雪花算法生成的
	Email      string    `gorm:"type:varchar(64);not null;" json:"email"`   // 邮箱
	Password   string    `gorm:"type:varchar(20);not null" json:"password"` // 密码
	Phone      string    `gorm:"type:varchar(11);" json:"phone"`            // 电话号码
	NickName   string    `gorm:"type:varchar(20);" json:"nick_name"`        // 昵称
	HeadImage  string    `gorm:"type:varchar(64);" json:"head_image"`       // 头像
	Role       int       `gorm:"type int;default:1" json:"role"`            // 角色,用于权限管理,1普通用户,2管理员,3超级管理员
	Grade      int       `gorm:"type int;default:1" json:"grade"`           // 账号等级
	Summary    string    `gorm:"type varchar(200);" json:"summary"`         // 个人简介
	Sex        int       `gorm:"type int" json:"sex"`                       // 1男,2女,3不详
	CreateTime time.Time `gorm:"type datetime;not null" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"type datetime;not null" json:"update_time"` // 更新时间
	DeleteTime time.Time `gorm:"type datetime;not null" json:"delete_time"` // 删除时间

	UserName string `gorm:"type:varchar(20);not null;" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	//Password string `gorm:"type:varchar(20);not null;" json:"password" validate:"required,min=6,max=20" label:"密码"`
	//Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

//func CheckUser(username string) (code int) {
//	var users User
//	db.Select("id").Where("user_name = ?", username).First(&users)
//	if users.ID > 0 {
//		return errmsg.ErrorUsernameUsed
//	}
//	return errmsg.Success
//}
//
//func CreateUser(data *User) (code int) {
//	data.Password = ScryptPw(data.Password)
//	if err = db.Create(data).Error; err != nil {
//		return errmsg.Error
//	}
//	return errmsg.Success
//}
//
//func ScryptPw(password string) string {
//	const keyLen = 10
//	salt := make([]byte, 8)
//	salt = []byte{3, 48, 209, 38, 48, 22, 29, 99} //随机给的
//	hashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return base64.StdEncoding.EncodeToString(hashPw)
//}
//
//func GetUsers(username string, pageSize int, pageNum int) ([]User, int) {
//	var users []User
//	count := 0
//	if username == "" {
//		if err = db.Find(&users).Count(&count).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error; err != nil && err != gorm.ErrRecordNotFound {
//			return nil, 0
//		}
//	} else {
//		if err = db.Where("user_name like ?", username+"%").Find(&users).Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
//			return nil, 0
//		}
//	}
//	return users, count
//}
//
//func DeleteUser(id int) int {
//	if err = db.Where("id=?", id).Delete(&User{}).Error; err != nil {
//		return errmsg.Error
//	}
//	return errmsg.Success
//
//}
//
//func EditUser(id int, data *User) int {
//	var m = make(map[string]interface{})
//	m["user_name"] = data.UserName
//	m["role"] = data.Role
//	if err = db.Model(&User{}).Where("id=?", id).Updates(m).Error; err != nil {
//		return errmsg.Error
//	}
//	return errmsg.Success
//}
//
////登录验证
//func CheckLogin(username string, password string) int {
//	var user User
//	db.Where("user_name=?", username).First(&user)
//	if user.ID == 1 {
//		return errmsg.ErrorUserNotExist
//	}
//	if ScryptPw(password) != user.Password {
//		return errmsg.ErrorPasswordWrong
//	}
//	if user.Role != 0 {
//		return errmsg.ErrorUserNoRight
//	}
//	return errmsg.Success
//}
