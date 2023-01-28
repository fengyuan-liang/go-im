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
	//================== 系统相关 =====================
	defaultGroup := engine.Group("/")
	{
		defaultGroup.GET("index", service.GetIndex)
	}
	//================== 用户相关 =====================
	// 路由组1 ，处理用户相关GET请求
	u1 := engine.Group("user/")
	{
		u1.GET("pageQuery", service.PageQueryUserList)
	}
	return engine
}
