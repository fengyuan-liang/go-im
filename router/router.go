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
	//================== 中间件 =====================
	// jwt鉴权
	//r.Use(middleware.JwtAuth())
	//================== 系统相关 =====================
	// 加载静态资源
	r.Static("/asset", "assert/")
	r.LoadHTMLGlob("views/**/*")
	defaultGroup := r.Group("/")
	{
		defaultGroup.GET("/", service.GetIndex)
		defaultGroup.GET("index", service.GetIndex)
	}
	//================== swagger相关 =====================
	// 设置docs文件相对路径
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//================== 用户相关 =====================
	// 用户路由组 ，处理用户相关请求
	u1 := r.Group("/user")
	{
		// 查询相关
		{
			u1.GET("/pageQuery", service.PageQueryUserList)
			u1.GET("/pageQueryByFilter", service.PageQueryByFilter)
		}
		// 修改相关
		{
			u1.POST("/register", service.CreateUser)
			u1.POST("/delOne", service.DelOne)
			u1.POST("/update", service.Update)
		}
		// 业务功能
		{
			u1.POST("/login", service.Login)
		}
		// ws相关
		{
			u1.GET("/sendMsg", service.WsSendMsg)
			u1.GET("/chat", service.Chat)
		}
	}
	return r
}
