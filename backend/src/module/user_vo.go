package module

import (
	"time"
)

type UserVo struct {
	// 用户昵称
	NickName string `json:"nickName" form:"nickName"`
	// 用户名
	UserName string `json:"userName" form:"userName"`
	// 密码
	Password string `json:"password" form:"password"`
	// 邮件地址
	UserMail string `json:"userMail" form:"userMail"`
	// 令牌
	Token string `json:"token" form:"token"`
	// 新密码
	NewPassword string `json:"newPassword" form:"newPassword"`
	// 状态
	State uint `json:"state" form:"state"`
	// 上级
	ParentName string `json:"parentName"`
	// 创建时间
	CreatedAt time.Time `json:"createdAt"`
}

type SimpleUserVo struct {
	// 用户昵称
	NickName string `json:"nickName" form:"nickName"`
	// 用户名
	UserName string `json:"userName" form:"userName"`
	// 邮件地址
	UserMail string `json:"userMail" form:"userMail"`
	// 状态
	State uint `json:"state" form:"state"`
	// 上级
	ParentName string `json:"parentName"`
	// 创建时间
	CreatedAt time.Time `json:"createdAt"`
	// 自己
	Self bool `json:"self"`
}

type LoginResponse struct {
	SessionId  string `json:"sessionId"`
	Privileges []Menu `json:"privileges"`
	UserDto    User   `json:"userDto"`
}
