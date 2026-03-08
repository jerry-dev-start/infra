package inits

import (
	"github.com/jerry-dev-start/infra/global"
	"github.com/jerry-dev-start/infra/logs"
)

// InitializeComponents 初始化各部件的方法
func InitializeComponents() {
	//打印 Banner
	logs.PrintBanner()
	// 初始化日志
	global.VM_LOG = logs.InitLogger()
	// 初始化 Gorm
	global.VM_DB = InitGorm()
}
