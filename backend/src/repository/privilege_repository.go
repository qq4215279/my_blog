package repository

import (
	"backend/src/constants"
	"backend/src/global"
	"backend/src/module"
)

// 获取父节点权限
func GetEnableParentPrivilege(allEnableParent interface{}) {
	global.DataBase.Where(&module.Privilege{IsForbidden: constants.PrivilegeEnable, IsLeaf: constants.NotLeaf}).Order("parent_path, label").Find(allEnableParent)
}

// 获取子节点权限
func GetEnableLeafPrivilege(allEnableLeaf interface{}) {
	global.DataBase.Where(&module.Privilege{IsForbidden: constants.PrivilegeEnable, IsLeaf: constants.IsLeaf}).Order("parent_path, label").Find(allEnableLeaf)
}
