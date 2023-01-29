package utils

import (
	"github.com/gin-gonic/gin"
)

// @Description: gin工具
// @Version: 1.0.0
// @Date: 2023/01/29 16:18
// @Author: fengyuan-liang@foxmail.com

// GetAllQueryParams 获取所有请求的参数，并封装为map返回
func GetAllQueryParams(c *gin.Context) map[string]interface{} {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]interface{}, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}
