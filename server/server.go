package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jerry-dev-start/infra/config"
)

func NewServer(conf *config.Config) {
	if conf.Server == nil {
		panic("Server configuration not found.")
	}
	gin.SetMode(conf.Server.GetModel())

}
