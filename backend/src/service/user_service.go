package service

import (
	"backend/src/cache"
	"backend/src/constants"
	"backend/src/global"
	"backend/src/module"
	"backend/src/repository"
	"backend/src/utils"
	"errors"
	"fmt"
	"gopkg.in/ldap.v2"
	"log"
	"reflect"
	"strings"
)

const (
	// 默认密码
	DefaultPassword string = "123456"
	// 登录token长度
	LoginTokenLength int = 20
)

// 登录
func Login(loginUser module.UserVo) ResponseBody {
	dbUser := module.GetUser(loginUser.UserName)
	if reflect.DeepEqual(dbUser, module.User{}) {
		return NewCustomErrorResponseBody("用户名不存在")
	}
	if dbUser.State == constants.UserStateDisable {
		return NewCustomErrorResponseBody("用户已禁用")
	}

	// admin账号登录
	if constants.Admin == loginUser.UserName {
		if dbUser.Password != utils.MD5Encryption(loginUser.Password, dbUser.Salt) {
			return NewCustomErrorResponseBody("密码不正确")
		}
		dbUser.LoginToken = utils.GetRandomString(LoginTokenLength)
		repository.Save(&dbUser)
		sessionId := cache.Join(*dbUser)
		return NewSuccessResponseBody(module.LoginResponse{
			UserDto:    *dbUser,
			Privileges: buildMenuTree(loginUser.UserName),
			SessionId:  sessionId,
		})
	} else { // 其他账号登录
		e := ldapCheckPwd(loginUser)
		if e != nil {
			return NewCustomErrorResponseBody(fmt.Sprintf("登陆失败!%s", e.Error()))
		}

		if dbUser.State == constants.UserStateInit {
			dbUser.State = constants.UserStateEnable
		}
		repository.Save(&dbUser)

		dbUser.LoginToken = utils.GetRandomString(LoginTokenLength)
		repository.Save(&dbUser)
		sessionId := cache.Join(*dbUser)

		return NewSuccessResponseBody(module.LoginResponse{UserDto: *dbUser, Privileges: buildMenuTree(loginUser.UserName), SessionId: sessionId})
	}
}

func LoginByLADPToken(token, sessionId string) ResponseBody {
	user, hasUser := cache.GetUser(sessionId)
	if !hasUser {
		return NewCustomErrorResponseBody("token失效")
	}
	if user.LoginToken != token {
		return NewCustomErrorResponseBody("token不正确")
	}
	user = *module.GetUser(user.UserName)
	if user.State == constants.UserStateDisable {
		return NewCustomErrorResponseBody("用户已禁用")
	}
	return NewSuccessResponseBody(module.LoginResponse{UserDto: user, Privileges: buildMenuTree(user.UserName), SessionId: sessionId})
}

// 登出
func Logout(userName, sessionId string) ResponseBody {
	dbUser := module.GetUser(userName)
	if reflect.DeepEqual(dbUser, module.User{}) {
		return NewCustomErrorResponseBody("用户名不存在")
	}
	cache.Invalid(sessionId)
	return NewSimpleSuccessResponseBody()
}

// 禁用账号
func DisableUser(userName string) ResponseBody {
	dbUser := module.GetUser(userName)
	if reflect.DeepEqual(dbUser, module.User{}) {
		return NewCustomErrorResponseBody("用户名不存在")
	}
	dbUser.State = constants.UserStateDisable
	repository.Save(&dbUser)
	return NewSimpleSuccessResponseBody()
}

// 启用账号
func EnableUser(userName string) ResponseBody {
	dbUser := module.GetUser(userName)
	if reflect.DeepEqual(dbUser, module.User{}) {
		return NewCustomErrorResponseBody("用户名不存在")
	}
	dbUser.State = constants.UserStateEnable
	repository.Save(&dbUser)
	return NewSimpleSuccessResponseBody()
}

// 域账户验证账号密码
func ldapCheckPwd(loginUser module.UserVo) error {
	domainName := loginUser.UserName
	if !strings.Contains(domainName, global.Config.Ldap.EmailSuffix) {
		domainName = domainName + global.Config.Ldap.EmailSuffix
	}
	con, e := ldap.Dial("tcp", global.Config.Ldap.Host+":"+global.Config.Ldap.Port)
	defer con.Close()
	if e != nil {
		log.Printf("LDAP dial fail %s", e.Error())
		return errors.New("连接域账户失败")
	}
	e = con.Bind(domainName, loginUser.Password)
	if e != nil {
		log.Printf("LDAP bind fail %s", e.Error())
		return errors.New("域账户校验失败")
	}
	return nil
}

// 修改ldap登陆类型的admin账号的密码
func UpdateLdapAdminPassword(param module.UserVo) ResponseBody {
	if global.Config.System.LoginType != constants.LoginTypeLDAP {
		return NewCustomErrorResponseBody("只支持修改 ldap登陆类型的 admin的密码")
	}
	if param.UserName != constants.Admin {
		return NewCustomErrorResponseBody("只支持修改 ldap登陆类型的 admin的密码")
	}
	if len(param.Password) < 6 || len(param.Password) > 18 {
		return NewCustomErrorResponseBody("密码长度6到18位")
	}
	user := module.GetUser(param.UserName)
	user.Salt = utils.GetPasswordSalt()
	user.Password = utils.MD5Encryption(param.Password, user.Salt)
	repository.Save(&user)
	return NewSimpleSuccessResponseBody()
}

func getCheckTokenUrl() string {
	sso := global.Config.SSO
	return sso.BaseUrl + sso.CheckTokenUrl
}

// 是否有这个子账户
func hasSubAccount(parentUserName, userName string) bool {
	subs := make([]module.User, 0)
	getSubAccountList(parentUserName, &subs)
	for _, sub := range subs {
		if sub.UserName == userName {
			return true
		}
	}
	return userName == parentUserName
}

// 获取子账号
func getSubAccountList(userName string, list *[]module.User) {
	currentLen := len(*list)
	subAccountList := repository.GetSubAccount(userName)
	for _, u := range subAccountList {
		*list = append(*list, u)
		getSubAccountList(u.UserName, list)
	}
	if currentLen == len(*list) {
		return
	}
}
