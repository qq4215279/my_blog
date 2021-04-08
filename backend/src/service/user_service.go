package service

import (
	"backend/src/cache"
	"backend/src/constants"
	"backend/src/global"
	"backend/src/module"
	"backend/src/repository"
	"backend/src/utils"
	"errors"
	"gopkg.in/ldap.v2"
	"log"
	"reflect"
	"strings"
	"time"
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

		accessToken, _ := utils.CreateAccessIdToken(dbUser.UserName)
		refreshToken, _ := utils.CreateRefreshIdToken(dbUser.UserName)

		cache.Join(*dbUser, refreshToken)
		return NewSuccessResponseBody(module.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Privileges:   buildMenuTree(loginUser.UserName),
			UserDto:      *dbUser,
		})
	} else { // 其他账号登录
		return NewCustomErrorResponseBody("当前项目只能admin用户登录")
	}
}

func RefreshAccessToken(refreshToken string) ResponseBody {
	mc, err := utils.ParseToken(refreshToken)
	if err != nil {
		return NewRefreshTokenExpireResponseBody()
	}
	userName := mc.Username

	// 创建新的accessToken
	accessToken, err := utils.CreateAccessIdToken(userName)

	//刷新refreshToken
	user, hasUser := cache.GetUser(refreshToken)
	if !hasUser {
		user = *module.GetUser(user.UserName)
	}

	// 判断是否刷新 refreshToken，如果refreshToken 快过期了 需要重新生成一个替换掉
	// refreshToken 有效时长是应该为accessToken有效时长的2倍
	minTimeOfRefreshToken := utils.AccessTokenExpirationTime * 2
	now := time.Duration(time.Now().UnixNano())
	refreshTokenStartTime, _ := cache.GetStartTime(refreshToken)
	// (refreshToken上次创建的时间点 + refreshToken的有效时长 - 当前时间点) 表示refreshToken还剩余的有效时长，如果小于2倍accessToken时长 ，则刷新 refreshToken
	if refreshTokenStartTime == 0 || (refreshTokenStartTime+utils.RefreshTokenExpirationTime)-now <= minTimeOfRefreshToken {
		cache.Invalid(refreshToken)

		refreshToken, _ = utils.CreateRefreshIdToken(userName)
		cache.Join(user, refreshToken)
	}

	return NewSuccessResponseBody(module.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken, Privileges: buildMenuTree(user.UserName), UserDto: user})
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
	return NewSuccessResponseBody(module.LoginResponse{UserDto: user, Privileges: buildMenuTree(user.UserName)})
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
