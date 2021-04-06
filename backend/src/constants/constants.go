/**
 * constants
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:56
 */
package constants

const (
	Admin                         = "admin"
	UserNameKey                   = "username"
	SessionIdKey                  = "sessionId"
	BuildFlowTypeCustomer         = "customer"
	BuildFlowTypeIncrementVersion = "incrementVersion"
	RepositoryTypeSdata           = "sdata"
	BaseUrl                       = "/root/"

	// 客户端状态 活跃
	ClientStateActive = 0
	// 客户端状态 下线
	ClientStateOffline = 1
	// 客户端状态 忙碌
	ClientStateBusy = 2
	// 客户端状态 空闲
	ClientStateIdle = 3
	// 客户端状态 满负荷
	ClientStateFullLoad = 4
	// 客户端初始化权重
	ClientInitWeight = 1

	// 使用sso登陆
	LoginTypeSSO = "sso"
	// 使用域账户登陆
	LoginTypeLDAP = "ldap"

	// 密码加密key前缀
	PasswordKeyPrefix string = "PcXS"
	// 静态salt
	StaticSalt string = "hario"

	// 并行执行
	BuildFlowIsConcurrent = 1
	// 不允许并行执行
	BuildFlowNotConcurrent = 2

	// 关注
	ProjectBranchStar = 1
	// 不关注
	ProjectBranchUnStar = 0

	// 初始化
	UserStateInit uint = 0
	// 禁用
	UserStateDisable uint = 1
	// 启用
	UserStateEnable uint = 2

	// 启用权限
	PrivilegeEnable uint = 0
	// 禁用权限
	PrivilegeDisable uint = 1

	// 是否是叶子节点（0：不是，1：是）
	NotLeaf uint = 0
	IsLeaf  uint = 1

	Development string = "dev"
	Test        string = "test"
	Production  string = "prod"

	// 用户session失效时间
	UserInvalidTime int64 = 60 * 60 * 12
	//UserInvalidTime int64  = 10
	SSHKeyDir string = "ssh_key"

	// 服务端创建
	BuildRecordStateInit = 1
	// 执行成功
	BuildRecordStateSuccess = 2
	// 执行失败
	BuildRecordStateFail = 3
	// 终止执行
	BuildRecordStateStop = 4
	// 被客户端放弃执行
	BuildRecordStateGiveUp = 5

	// 业务失败
	ResponseStateFail = 0
	// 成功
	ResponseStateSuccess = 1
	// accessToken 过期
	ResponseStateAccessExpire = 2
	// refreshToken 过期
	ResponseStateRefreshExpire = 3
)

var (
	WorkingState = []int{ClientStateActive, ClientStateBusy, ClientStateIdle}
)
