package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-im/docs"
	"go-im/service"
)

// @Description: 路由控制器
// @Version: 1.0.0
// @Date: 2023/01/27 22:04
// @Author: fengyuan-liang@foxmail.com

func Router() *gin.Engine {
	r := gin.Default()
	//================== 系统相关 =====================
	defaultGroup := r.Group("/")
	{
		defaultGroup.GET("index", service.GetIndex)
	}
	//================== swagger相关 =====================
	// 设置docs文件相对路径
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//================== 用户相关 =====================
	// 路由组1 ，处理用户相关GET请求
	u1 := r.Group("user/")
	{
		u1.GET("pageQuery", service.PageQueryUserList)
	}
	return r
}
