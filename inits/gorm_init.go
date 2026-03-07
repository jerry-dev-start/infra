package inits

import (
	"strings"

	"github.com/jerry-dev-start/infra/global"
	"gorm.io/gorm"
)

// InitGorm 初始化Gorm
func InitGorm() *gorm.DB {
	// 判断db_type的值来初始化对应数据库
	switch strings.ToLower(global.VM_CNF.Server.GetDbType()) {
	case "mysql":
		return GormMysqlInit()
	default:
		panic("database type not supported")
	}
}
