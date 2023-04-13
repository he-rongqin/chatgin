package config

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"org.chatgin/src/common"
	"org.chatgin/src/modules/module"
)

// 全局配置
type AppConfig struct {
	ServerPort  string // 服务端口
	ContextPath string // 服务上下文路径
	AppName     string // 服务名称

}

// redis 配置
type RedisConfig struct {
	Addr         string // 连接地址
	Database     int    // 数据库
	Password     string
	PoolSize     int // 最大连接数
	MinIdleConns int // 最小空闲连接数
}

func InitConfig() {
	viper.AddConfigPath("./evn")
	viper.SetConfigFile("application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalln("加载配置文件出错，程序退出: %v\n", err)
	}
	// 初始化 server配置
	initServerConfig()
	// 初始化mysql
	initMysql(viper.GetString(common.CONFIG_PREFIX + "mysql.dsn"))
	// 初始化redis client
	initRedis(&RedisConfig{
		Addr:         viper.GetString(common.CONFIG_PREFIX + "redis.addr"),
		Database:     viper.GetInt(common.CONFIG_PREFIX + "redis.db"),
		Password:     viper.GetString(common.CONFIG_PREFIX + "redis.password"),
		PoolSize:     viper.GetInt(common.CONFIG_PREFIX + "redis.poolSize"),
		MinIdleConns: viper.GetInt(common.CONFIG_PREFIX + "redis.minIdleConns"),
	})

}

var (
	MysqlDB      *gorm.DB
	RedisClient  *redis.Client
	ServerConfig *AppConfig
)

// 初始化server 配置
func initServerConfig() {
	port := viper.GetString(common.CONFIG_PREFIX + ".server.port")
	path := viper.GetString(common.CONFIG_PREFIX + ".server.context.path")
	appName := viper.GetString(common.CONFIG_PREFIX + ".server.appName")

	if "" == port {
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
	db, error := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if error != nil {
		logrus.Fatalln("初始化mysql数据库连接错误。", error)
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

// 初始化数据库表
func InitTables() {
	// 创建用户表
	user := &module.UserBasic{}
	user.CreateTable()

}
