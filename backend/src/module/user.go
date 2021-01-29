package module

import (
	"gorm.io/gorm"
	"time"
)

// 用户
type User struct {
	BaseTable
	// 用户名称
	UserName string `json:"userName"`
	// 密码
	Password string `json:"password"`
	// 盐
	Salt string `json:"salt"`
	// 用户昵称
	NickName string `json:"nickName"`
	// 邮件地址
	UserMail string `json:"userMail"`
	// 权限
	Privileges string `json:"privileges"`
	// 登录Token
	LoginToken string `json:"loginToken"`
	// 登录Token生成时间
	LoginTokenTime time.Time `json:"loginTokenTime"`
	// 状态（0：初始，1：禁用，2：正常）
	State uint `json:"state"`
	// 上级
	ParentName string `json:"parentName"`
}

// 表名
func (user User) TableName() string {
	return "user"
}

// 表名
func (user User) Save(db *gorm.DB) {
	db.Save(user)
}
