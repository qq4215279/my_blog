/**
 * config
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:57
 */
package config

var (
	ConfPath = "conf"
)

// 系统配置内容
type systemInfo struct {
	// 版本号
	Version string `yaml:"version"`
	// 端口号
	Addr string `yaml:"addr"`
	// 登陆类型 sso\ldap
	LoginType string `yaml:"login-type"`
}

// mysql配置内容
type mysqlInfo struct {
	// 用户名
	Username string `yaml:"username"`
	// 密码
	Password string `yaml:"password"`
	// 连接URL
	URL string `yaml:"url"`
	// 数据库名
	DbName string `yaml:"db-name"`
	// 数据库配置
	Config string `yaml:"config"`
	// 最大空闲连接数
	MaxIdleConns string `yaml:"max-idle-conns"`
	// 最大连接数
	MaxOpenConns string `yaml:"max-open-conns"`
	// 日志模式
	LogMode string `yaml:"log-mode"`
}

// 日志配置内容
type logInfo struct {
	// 日志前缀
	Prefix string `yaml:"prefix"`
	// 日志根目录
	LogBasePath string `yaml:"log-base-path"`
	// gin日志目录
	GinPath string `yaml:"gin-path"`
	// info日志输出目录
	InfoPath string `yaml:"info-path"`
	// 持久层日志目录
	OrmPath string `yaml:"orm-path"`
	// 打包日志目录
	BuildJobLogPath string `yaml:"build-job-log-path"`
}

// 打包配置
type build struct {
	// 打包结果路径
	FilePath string `yaml:"file-path"`
}

// 打包配置
type client struct {
	// 负载策略
	BalanceStrategy string `yaml:"balance-strategy"`
	RootUrl         string `yaml:"root-url"`
}

// 定时任务配置
type job struct {
	// 清理任务
	Clear                      string `yaml:"clear"`
	ClearLogRetentionDay       int    `yaml:"clear-log-retention-day"`
	ClearBuildFileRetentionDay int    `yaml:"clear-build-file-retention-day"`
}

// 域账户配置
type ldap struct {
	// 清理任务
	Host string `yaml:"host"`
	// 端口
	Port string `yaml:"port"`
	// 搜索地址
	SearchBase string `yaml:"search-base"`
	// 搜索过滤
	SearchFilter string `yaml:"search-filter"`
	// 域名
	Domain string `yaml:"domain"`
	// 邮箱后缀
	EmailSuffix string `yaml:"email-suffix"`
}

// 域账户配置
type crypto struct {
	// 清理任务
	DecodeKeySuffix string `yaml:"decode-key-suffix"`
	// 端口
	CipherTextSuffix string `yaml:"cipher-text-suffix"`
}

// sso配置
type sso struct {
	// sso 中 本应用Id
	AppId string `yaml:"app-id"`
	// 根路径
	BaseUrl string `yaml:"base-url"`
	// 校验token
	CheckTokenUrl string `yaml:"check-token-url"`
}

// 配置根节点
type Config struct {
	// 系统配置节点
	System systemInfo `yaml:"system"`
	// mysql配置节点
	Mysql mysqlInfo `yaml:"mysql"`
	// 日志配置节点
	Log logInfo `yaml:"log"`
	// 客户端
	Client client `yaml:"client"`
	// 构建相关
	Build build `yaml:"build"`
	// 定时任务
	Job job `yaml:"job"`
	// 域账户配置
	Ldap ldap `yaml:"ldap"`
	// 密码加密
	Crypto crypto `yaml:"crypto"`
	// sso 配置
	SSO sso `yaml:"sso"`
}
