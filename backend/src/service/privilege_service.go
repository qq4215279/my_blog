package service

import (
	"backend/src/constants"
	"backend/src/module"
	"backend/src/repository"
	"backend/src/utils"
	"strings"
)

func buildMenuTree(userName string) []module.Menu {
	menuTree := make([]module.Menu, 0, 2)

	// 读取库里全部的权限
	var allEnableParent []module.Privilege
	repository.GetEnableParentPrivilege(&allEnableParent)
	for _, privilege := range allEnableParent {
		if privilege.IsLeaf == constants.IsLeaf {
			continue
		}

		// TODO 用户权限管理
		//if !HasModulePrivilege(userName, constants.BaseUrl+privilege.Name) {
		//	continue
		//}

		var menu module.Menu
		menu.Label = privilege.Label
		menu.Path = privilege.Path
		menu.Name = privilege.Icon
		menu.Children = make([]module.Menu, 0, 0)

		menuTree = append(menuTree, menu)
	}

	var allEnableLeaf []module.Privilege
	repository.GetEnableLeafPrivilege(&allEnableLeaf)
	for _, privilege := range allEnableLeaf {

		needContinue := true
		index := 0
		for i, module := range menuTree {
			if privilege.ParentPath == module.Path {
				needContinue = false
				index = i
				break
			}
		}
		if needContinue {
			continue
		}

		subMenu := new(module.Menu)
		subMenu.Label = privilege.Label
		subMenu.Path = privilege.Path
		subMenu.Name = privilege.Name
		subMenu.Icon = privilege.Icon
		subMenu.Children = make([]module.Menu, 0, 0)

		menuTree[index].Children = append(menuTree[index].Children, *subMenu)
	}
	return menuTree
}

// 是否有这个模块的权限
func HasModulePrivilege(userName, fullPath string) bool {
	if userName == constants.Admin {
		return true
	}
	user := repository.GetUser(userName)
	userPrivileges := strings.Split(user.Privileges, ",")
	for _, userModulePrivilege := range userPrivileges {
		if utils.IsBlank(userModulePrivilege) {
			continue
		}
		if strings.HasPrefix(fullPath, constants.BaseUrl+userModulePrivilege) {
			return true
		}
	}
	return false
}
