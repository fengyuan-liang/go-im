package service

import (
	"github.com/gin-gonic/gin"
	"go-im/common/entity/response"
	"go-im/models"
	"strconv"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/28 14:43
// @Author: fengyuan-liang@foxmail.com

// PageQueryUserList 分页查询
// @Tags 用户相关
// @BasePath /user
// @Summary 分页查询
// @param pageNo query integer true "第几页" default(1)
// @param pageSize query integer true "每页多少条" default(10)
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/pageQuery [get]
func PageQueryUserList(c *gin.Context) {
	// 拿到参数
	page := c.Query("pageNo")
	size := c.Query("pageSize")
	if page == "" || size == "" {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	// 转化类型
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	c.JSON(200, response.Ok.WithData(models.PageQueryUserList(pageInt, sizeInt)))
}

// CreateUser 创建一个用户
// @Tags 创建一个用户
// @BasePath /user
// @Summary 用于用户注册
// @Param param body models.UserBasic true "上传的JSON"
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/register [post]
func CreateUser(c *gin.Context) {
	userBasic := models.UserBasic{}
	// 第二次输入的密码
	rePassword := c.Query("rePassword")
	if err := c.BindJSON(&userBasic); err != nil || userBasic.Name == "" || userBasic.Age <= 0 || userBasic.PassWord == "" {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	if userBasic.PassWord != rePassword {
		c.JSON(-1, response.AppErr.WithMsg("两次密码不一致"))
		return
	}
	// 插入数据
	models.InsetOne(userBasic)
	c.JSON(200, response.Ok)
}
