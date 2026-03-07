package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jerry-dev-start/infra/config"
	"github.com/jerry-dev-start/infra/route"
)

type Server struct {
	Engine     *gin.Engine
	cnf        *config.Config
	httpServer *http.Server
	Close      func()
}

// NewServer 创建Server
func NewServer(conf *config.Config) *Server {
	if conf.Server == nil {
		panic("Server configuration not found.")
	}
	gin.SetMode(conf.Server.GetModel())

	ginEngine := gin.Default()

	return &Server{
		Engine: ginEngine,
	}
}

// StartWeb 启动Web服务
func (s *Server) StartWeb() {
	// 拼接地址
	address := fmt.Sprintf("%s:%d", s.cnf.Server.GetHost(), s.cnf.Server.GetPort())

	s.httpServer = &http.Server{
		Addr:    address,
		Handler: s.Engine,
	}

	go func() {
		log.Printf("Start web server at %s", address)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// ---- 优雅启停逻辑 ----
	quit := make(chan os.Signal, 1)
	// 如果程序捕获到 Ctrl+C 或者是系统关闭指令，则把他发送到quit Channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在没有把 syscall.SIGINT,syscall.SIGTERM 信息发送到quit 之前会一直阻塞到这里
	<-quit
	log.Printf("Shutting down server...")
	if s.Close != nil {
		s.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server exited gracefully.")
}

// RegisterRouter 批量注册路由
// 参数 route.IRouter 接口的实列
func (s *Server) RegisterRouter(routes ...route.IRouter) {
	PublicGroup := s.Engine.Group(s.cnf.Server.GetRouterPrefix())
	PrivateGroup := s.Engine.Group(s.cnf.Server.GetRouterPrefix())

	//这里要给 PrivateGroup 加入认证的中间件
	for _, r := range routes {
		r.Register(s.Engine, PublicGroup, PrivateGroup)
	}

	log.Println("Routers registered successfully.")
}
