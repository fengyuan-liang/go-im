package service

import (
	"github.com/gin-gonic/gin"
	"go-im/models"
	"strconv"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/28 14:43
// @Author: fengyuan-liang@foxmail.com

func PageQueryUserList(c *gin.Context) {
	// 拿到参数
	page := c.Query("pageNo")
	size := c.Query("pageSize")
	if page == "" || size == "" {
		c.JSON(501, gin.H{
			"message": "请求参数缺失",
		})
		return
	}
	// 转化类型
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    models.PageQueryUserList(pageInt, sizeInt),
	})
}
