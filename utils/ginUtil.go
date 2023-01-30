package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
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

// ParseMapFieldType 根据bean的类型将jsonMap value类型转换
//
// @params filter 过滤器，那些字段不用添加
func ParseMapFieldType(jsonMap map[string]interface{}, bean interface{}, ignoreField ...string) map[string]interface{} {
	// 拿到结构体所有属性的类型
	stringTypeMap := GetLowerTitleFieldStringType(bean)
	// 新的容器
	m := make(map[string]interface{})
	for k, v := range jsonMap {
		for fieldName, fieldStringType := range stringTypeMap {
			// 字段是否需要忽略
			if ignoreField != nil && ContainsValue(k, ignoreField...) {
				break
			}
			// 忽略大小写进行比较
			if strings.EqualFold(k, fieldName) {
				m[k] = ParseType(v, fieldStringType)
			}
		}
	}
	return m
}
