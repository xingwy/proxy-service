package http

import (
	"proxy-service/api/handler"

	"github.com/gin-gonic/gin"
)

// Init 路由
func Reginster(instance *gin.Engine) *gin.Engine {
	api := instance.Group("/proxy-service")
	top := api.Group("/top")
	{
		top.Any("/send", handler.Instance.TopHandler().Proxy)
	}

	return instance
}
