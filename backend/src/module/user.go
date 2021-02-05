package module

import (
	"backend/src/global"
	"gorm.io/gorm"
	"time"
)

// 用户
type User struct {
	BaseTable
	// 用户名称
	UserName string `json:"userName" form:"userName"`
	// 密码
	Password string `json:"password" form:"password"`
	// 盐
	Salt string `json:"salt" form:"salt"`
	// 用户昵称
	NickName string `json:"nickName" form:"nickName"`
	// 邮件地址
	UserMail string `json:"userMail" form:"userMail"`
	// 权限
	Privileges string `json:"privileges" form:"privileges"`
	// 登录Token
	LoginToken string `json:"loginToken" form:"loginToken"`
	// 登录Token生成时间
	LoginTokenTime time.Time `json:"loginTokenTime" form:"loginTokenTime"`
	// 状态（0：初始，1：禁用，2：正常）
	State uint `json:"state" form:"state"`
	// 上级
	ParentName string `json:"parentName" form:"parentName"`
}

// 表名
func (user User) TableName() string {
	return "user"
}

// 表名
func (user User) Save(db *gorm.DB) {
	db.Save(user)
}

// 获取用户
func GetUser(userName string) *User {
	var dbUser User
	global.DataBase.Where(&User{UserName: userName}).First(&dbUser)
	return &dbUser
}
