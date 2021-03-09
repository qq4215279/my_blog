package repository

import (
	"backend/src/global"
	"backend/src/module"
)

func GetUser(userName string) *module.User {
	var dbUser module.User
	global.DataBase.Where(&module.User{UserName: userName}).First(&dbUser)
	return &dbUser
}

// 获取下一级账号
func GetSubAccount(userNames string) []module.User {
	var users []module.User
	global.DataBase.Where("parent_name = ?", userNames).Find(&users)
	return users
}

// 获取被赋了Api权限的用户
func GetHasPrivilegeUser() []module.User {
	var users []module.User
	global.DataBase.Where("high_privileges is not null").Find(&users)
	return users
}
