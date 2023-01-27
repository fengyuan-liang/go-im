package router

import (
	"github.com/gin-gonic/gin"
	"go-im/service"
)

// @Description: 路由控制器
// @Version: 1.0.0
// @Date: 2023/01/27 22:04
// @Author: fengyuan-liang@foxmail.com

func Router() *gin.Engine {
	engine := gin.Default()
	engine.GET("/index", service.GetIndex)
	return engine
}
