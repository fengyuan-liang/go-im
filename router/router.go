package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-im/docs"
	"go-im/pkg/common/middleware"
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
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
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
	//================== 不需要验证token ===============
	defaultGroup.POST("/login", service.Login)
	defaultGroup.POST("/createUser", service.CreateUser)
	//================== 页面跳转 =====================
	// 页面跳转
	{
		r.GET("/toRegister", service.ToRegister)
		r.GET("/toChat", service.ToChat)
	}
	//================== 用户相关 =====================
	// 用户路由组 ，处理用户相关请求
	u1 := r.Group("/api")
	// 需要验证token
	u1.Use(middleware.JwtAuth())
	{
		// 查询相关
		{
			u1.GET("/pageQuery", service.PageQueryUserList)
			u1.GET("/pageQueryByFilter", service.PageQueryByFilter)
			u1.POST("/searchFriends", service.SearchFriends)
		}
		// 修改相关
		{
			//u1.POST("/createUser", service.CreateUser)
			u1.POST("/delOne", service.DelOne)
			u1.POST("/update", service.Update)
		}
		// 业务功能
		{
			//u1.POST("/login", service.Login)
		}
		// ws相关
		{
			u1.GET("/sendMsg", service.WsSendMsg)
			u1.GET("/chat", service.Chat)
		}
		// 用户关系
		contactGroup := u1.Group("/contact")
		{
			contactGroup.POST("/addfriend", service.AddFriend)
			contactGroup.POST("/searchFriends", service.SearchFriends)
		}
	}
	return r
}
