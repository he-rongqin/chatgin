package config

import (
	"os"

	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	common "org.chatgin/pkg/util"
)

// 全局配置
type AppConfig struct {
	ServerPort  string // 服务端口
	ContextPath string // 服务上下文路径
	AppName     string // 服务名称

}

// 会话配置
type SessionConfig struct {
	ExpiresTimes uint   // 会话有效期，单位：小时
	PrivateKey   string // jwt 加密key
}

// redis 配置
type RedisConfig struct {
	Addr         string // 连接地址
	Database     int    // 数据库
	Password     string
	PoolSize     int // 最大连接数
	MinIdleConns int // 最小空闲连接数
}

// 全局初始化,
// initMysql: 初始化mysql 数据连接;
// initRedis: 初始化 redis client 连接;
// initServerConfig: 初始化 server 配置;
func InitConfig() {
	// 获取项目的执行路径
	dir, _ := os.Getwd()
	configFile = viper.New()
	configFile.AddConfigPath(dir + "/.env")
	configFile.SetConfigName("application")
	configFile.SetConfigType("yaml")
	err := configFile.ReadInConfig()
	if err != nil {
		logrus.Fatalln("加载配置文件出错，程序退出: %v\n", err)
	}
	// 初始化 server配置
	initServerConfig()
	// 初始化mysql
	initMysql(configFile.GetString(common.CONFIG_PREFIX + "mysql.dsn"))
	// 初始化redis client
	initRedis(&RedisConfig{
		Addr:         configFile.GetString(common.CONFIG_PREFIX + "redis.addr"),
		Database:     configFile.GetInt(common.CONFIG_PREFIX + "redis.db"),
		Password:     configFile.GetString(common.CONFIG_PREFIX + "redis.password"),
		PoolSize:     configFile.GetInt(common.CONFIG_PREFIX + "redis.poolSize"),
		MinIdleConns: configFile.GetInt(common.CONFIG_PREFIX + "redis.minIdleConns"),
	})
	// 初始化toekn配置
	initSessionConfig()

}

var (
	MysqlDB      *gorm.DB
	RedisClient  *redis.Client
	ServerConfig *AppConfig
	configFile   *viper.Viper
	TokenConfig  *SessionConfig
)

func initSessionConfig() {
	expiresTimes := configFile.GetUint(common.CONFIG_PREFIX + "token.expiresTimes")
	privateKey := configFile.GetString(common.CONFIG_PREFIX + "token.privateKey")

	if expiresTimes == 0 {
		expiresTimes = common.DEFAULT_EXPIRES_TIMES
	}
	if privateKey == "" {
		privateKey = common.DEFAULT_PRIVATE_KEY
	}

	TokenConfig = &SessionConfig{
		ExpiresTimes: expiresTimes,
		PrivateKey:   privateKey,
	}
}

// 初始化server 配置
func initServerConfig() {
	port := configFile.GetString(common.CONFIG_PREFIX + ".server.port")
	path := configFile.GetString(common.CONFIG_PREFIX + ".server.context.path")
	appName := configFile.GetString(common.CONFIG_PREFIX + ".server.appName")

	if port == "" {
		port = common.DEFAULT_SERVER_PORT
	}
	if path == "" {
		path = common.DEFAULT_CONTEXT_PATH
	}
	if appName == "" {
		appName = common.DEFAULT_APP_NAME
	}
	ServerConfig = &AppConfig{
		ServerPort:  port,
		ContextPath: path,
		AppName:     appName,
	}

}

// 初始化mysql
func initMysql(dns string) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("初始化mysql数据库连接错误。%v", err)
	}
	MysqlDB = db
}

// 初始化redis
func initRedis(config *RedisConfig) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		DB:       config.Database,
		Password: config.Password,
	})
}
