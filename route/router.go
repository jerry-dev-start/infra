package route

import "github.com/gin-gonic/gin"

type IRouter interface {
	Register(engine *gin.Engine, PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup)
}
