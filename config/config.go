package config

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var conf = new(Config)

type Config struct {
	Server      *Server      `mapstructure:"server"`
	MysqlConfig *MysqlConfig `mapstructure:"mysql_config"`
}

// GetConfig 获取序列化的后配置
func GetConfig() *Config {
	if conf == nil {
		// 确保即使有人忘记调用 Init，程序也会以明确的方式失败
		panic("config not initialized, please call config.Init() first")
	}
	return conf
}

type Server struct {
	Port         *int    `mapstructure:"port"`
	Host         *string `mapstructure:"host"`
	Model        *string `mapstructure:"model"`
	DbType       *string `mapstructure:"db_type"`
	RouterPrefix *string `mapstructure:"prefix"`
}

func (s *Server) GetDbType() string {
	if s == nil || s.DbType == nil {
		return "mysql"
	}
	return *s.DbType
}

func (s *Server) GetModel() string {
	if s == nil || s.Model == nil {
		return "release"
	}
	return *s.Model
}

func (s *Server) GetRouterPrefix() string {
	if s == nil || s.RouterPrefix == nil {
		return ""
	}
	return *s.RouterPrefix
}

// GetHost 获取主机地址
func (s *Server) GetHost() string {
	if s == nil || s.Host == nil {
		return "127.0.0.1"
	}
	return *s.Host
}

// GetPort 获取端口号
func (s *Server) GetPort() int {
	if s == nil || s.Port == nil {
		return 8888
	}
	return *s.Port
}

// Init 初始化基础包的全局配置。
// 该方法会执行以下操作：
//  1. 加载配置文件，如果命令行指定了使用指定的，没有指定使用 ./config.yml
//
// 注意：如果配置文件不存在、格式错误或数据库无法连接，该方法将直接触发 panic。
// 在主服务启动初期（main 函数内）应当首先调用此方法。
func Init() *Config {
	// 读取命令行中的配置文件，如果有的话使用命令行指定的
	var configFilePath = pflag.StringP("config", "c", "./config.yml", "config file path")
	pflag.Parse()
	// 初始化 Viper
	v := viper.New()
	v.SetDefault("application.name", "VmApp")
	v.SetDefault("server.port", 8888)
	v.SetDefault("server.host", "127.0.0.1")

	// 如果命令行读取配置文件信息，就使用指定的
	if *configFilePath != "" {
		v.SetConfigFile(*configFilePath)
		log.Printf("Use the configuration file specified by command-line arguments: %s", *configFilePath)
	} else {
		// 如果没有指定，则按下面的方式查找
		v.AddConfigPath(".")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		log.Println("Search for config.yml within the current working directory.")
	}

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将配置序存放到 conf 变量中
	if err := v.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("Failed to unmarshal configuration into Config struct.: %s \n", err))
	}
	return conf
}
