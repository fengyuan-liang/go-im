package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-im/common/bizError"
	"reflect"
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

// ParseMapFieldType
//
//	@Description: 根据bean的类型将jsonMap value类型转换
//	@args jsonMap 结构体对应泛化的map
//	@args bean 要转换的结构体类型
//	@args ignoreField 结构体内那些字段不需要加入到map中
//	@return map[string]interface{}
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

// ParseMap
//
//	@Description: 结构体转为Map[string]interface{}
//	@args in 传入的结构体
//	@args tagName 结构体内tag，用作map的key
//	@return map[string]interface{}
//	@return bizError.BizErrorer
func ParseMap(in interface{}, tagName string) (map[string]interface{}, bizError.BizErrorer) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, bizError.NewBizError(fmt.Sprintf("ToMap only accepts struct or struct pointer; got %T", v))
	}
	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}
