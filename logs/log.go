package logs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化日志框架
// 日志的框架是使用的 Zap
func InitLogger() *zap.Logger {
	zapConfig := zap.NewDevelopmentEncoderConfig()
	// 自定义时间格式，默认是浮点数，改为人类可读格式
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 关键字大写
	zapConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig),   // 编码器配置
		zapcore.AddSync(os.Stdout),          // 打印到控制台，生产环境可改为文件
		zap.NewAtomicLevelAt(zap.InfoLevel), // 日志级别
	)
	// zap.AddCaller() 会在日志中显示调用函数的文件名和行号
	log := zap.New(core, zap.AddCaller())
	return log
}
