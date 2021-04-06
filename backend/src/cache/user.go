package cache

import (
	"backend/src/constants"
	"backend/src/module"
	"backend/src/utils"
	"log"
	"time"
)

var (
	// 在线用户
	sessionMap = make(map[string]*expiredUser)
)

type expiredUser struct {
	// 开始时间
	startTime int64
	// 过期时间
	expiredTime int64
	// 用户数据
	user module.User
}

// 获取用户
func GetUser(sessionId string) (module.User, bool) {
	user := sessionMap[sessionId]
	if user == nil {
		return module.User{}, false
	}
	return user.user, true
}

// 获取开始时间
func GetStartTime(sessionId string) (time.Duration, bool) {
	user := sessionMap[sessionId]
	if user == nil {
		return 0, false
	}
	return time.Duration(user.startTime), true
}

func GetOnlineUser() map[string]*expiredUser {
	return sessionMap
}

// 加入
func Join(user module.User, sessionId string) {
	sessionMap[sessionId] = &expiredUser{startTime: time.Now().Unix(), expiredTime: time.Now().Unix() + constants.UserInvalidTime, user: user}
	log.Println(sessionId + " join...")
}

// 加入
func Join2(user module.User) string {
	sessionId := utils.GetRandomString(20)
	sessionMap[sessionId] = &expiredUser{startTime: time.Now().Unix(), expiredTime: time.Now().Unix() + constants.UserInvalidTime, user: user}
	log.Println(sessionId + " join...")
	return sessionId
}

// 延迟失效时间
func (u *expiredUser) addExpiredTime() {
	u.expiredTime = time.Now().Unix() + constants.UserInvalidTime
}

// 延迟失效时间
func AddInvalidTime(sessionId string) {
	expiredUser := sessionMap[sessionId]
	expiredUser.addExpiredTime()
	log.Println(sessionId + " addInvalidTime...")
}

// 失效
func Invalid(sessionId string) {
	log.Println(sessionId + " invalid...")
	delete(sessionMap, sessionId)
}

// 清理应该失效的用户
func ClearUser() {
	for sessionId := range sessionMap {
		SessionIsInvalid(sessionId)
	}
}

// session是否过期
func SessionIsInvalid(sessionId string) bool {
	expiredUser, ok := sessionMap[sessionId]
	// 不存在 返回过期
	if !ok {
		return true
	}
	// 未过期 返回未过期
	if expiredUser.expiredTime > time.Now().Unix() {
		return false
	}
	// 过期 清掉session 并返回过期
	Invalid(sessionId)
	return true
}
