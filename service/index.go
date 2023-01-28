package service

import (
	"github.com/gin-gonic/gin"
	"time"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/27 22:06
// @Author: fengyuan-liang@foxmail.com

func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello,time is " + time.Now().String(),
	})
}
