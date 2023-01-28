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
