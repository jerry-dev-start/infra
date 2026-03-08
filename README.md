<div align="center">
  <h1 align="center">GP Base Package</h1>
  <p align="center">
    🚀 一个为快速开发而生的 Go 语言基础设施包
  <p>
</div>

**GP Base Package** 是一个面向 Go 语言 Web 开发的高性能基础库。它深度集成并二次封装了 Gin、GORM 和 Redis，旨在通过标准化的配置和开箱即用的核心组件，帮助开发者快速构建稳定、规范的 Web 服务。

核心特性：
- ⚡ 极速启动: 预集成 Gin 引擎，支持标准化的中间件管理。
- 🗄️ ORM 增强: 基于 GORM 的数据库连接池自动化管理，支持主从配置。
- 🚀 Redis 封装: 内置常用缓存操作、分布式锁等实用工具。
- 🛠️ 统一配置: 采用结构化配置管理（支持 YAML/JSON/Env），一处定义，全局引用。
- 📊 规范响应: 内置标准化的 JSON 响应格式与错误码处理机制。

项目结构：
```Plaintext
gp-base/
├── config/         # 配置文件解析并序列化
├── global/         # 全局使用的组件，如：日志、数据库等
├── inits/          # 所有组件都在这里初始化
├── logs/           # 日志Zap初始化
├── route/          # 路由相关
├── server/         # 服务启动相关
└── utils/          # 一直可以直接使用的工具包
```

## 快速上手
1. 安装
在你的 Web 项目中，使用 go get 引入该基础包：
```bash
go get github.com/jerry-dev-start/infra
```

2.初始化示例
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jerry-dev-start/infra/config"
	"github.com/jerry-dev-start/infra/global"
	bp "github.com/jerry-dev-start/infra/inits"
	"github.com/jerry-dev-start/infra/server"
)

func main() {
	global.VM_CNF = config.Init()
	s := server.NewServer(global.VM_CNF, func(context *gin.Context) {

	})
	// 初始化基础包中的部件
	bp.InitializeComponents()
	//初始化路由
	s.RegisterRouter(
		&route.AuthRouter{},
	)
	global.VM_LOG.Info("初始化成功，下面启动服务...")
	s.StartWeb()
}
```