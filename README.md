# GP Base Package
** GP Base Package ** 是一个面向 Go 语言 Web 开发的高性能基础库。它深度集成并二次封装了 Gin、GORM 和 Redis，旨在通过标准化的配置和开箱即用的核心组件，帮助开发者快速构建稳定、规范的 Web 服务。

核心特性：
- ⚡ 极速启动: 预集成 Gin 引擎，支持标准化的中间件管理。
- 🗄️ ORM 增强: 基于 GORM 的数据库连接池自动化管理，支持主从配置。
- 🚀 Redis 封装: 内置常用缓存操作、分布式锁等实用工具。
- 🛠️ 统一配置: 采用结构化配置管理（支持 YAML/JSON/Env），一处定义，全局引用。
- 📊 规范响应: 内置标准化的 JSON 响应格式与错误码处理机制。

项目结构：
```Plaintext
gp-base/
├── core/           # 核心逻辑 (启动器、全局上下文)
├── config/         # 配置解析与全局 Config 结构体
├── database/       # GORM 封装与初始化
├── cache/          # Redis 封装与连接池管理
├── middleware/     # 常用 Gin 中间件 (日志、跨域、鉴权)
├── response/       # 统一返回格式封装
└── pkg/            # 通用工具函数 (加密、时间、转换)
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
    "github.com/yourname/gp-base/core"
    "github.com/yourname/gp-base/config"
)

func main() {
    // 1. 加载配置
    cfg := config.MustLoad("config.yaml")

    // 2. 初始化底层基础组件 (DB, Redis 等)
    app := core.NewApp(cfg)
    
    // 3. 注册路由并启动
    app.Router.GET("/ping", func(c *gin.Context) {
        response.Success(c, "pong")
    })
    
    app.Run(":8080")
}
```