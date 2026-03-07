package inits

import (
	"fmt"

	"github.com/jerry-dev-start/infra/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysqlInit 初始化Gorm数据库
func GormMysqlInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.VM_CNF.MysqlConfig.Username,
		global.VM_CNF.MysqlConfig.Password,
		global.VM_CNF.MysqlConfig.Host,
		global.VM_CNF.MysqlConfig.Port,
		global.VM_CNF.MysqlConfig.DbName)
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		return db
	}
}
