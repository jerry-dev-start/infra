package config

type MysqlConfig struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}
