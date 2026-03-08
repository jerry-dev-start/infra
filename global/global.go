package global

import (
	"github.com/jerry-dev-start/infra/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// VM_DB 是全局数据库对象
	VM_DB  *gorm.DB
	VM_CNF *config.Config
	VM_LOG *zap.Logger
)
