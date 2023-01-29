package utils

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// @Description: 通用工具类
// @Version: 1.0.0
// @Date: 2023/01/29 17:22
// @Author: fengyuan-liang@foxmail.com

//================  number相关  ====================

func ParseString(args interface{}) string {
	if args == nil {
		log.Panic("NPE")
	}
	// 判断类型
	switch reflect.TypeOf(args).String() {
	// 关键是float转换有问题
	case "float32":
		return strconv.Itoa(int(args.(float32)))
	case "float64":
		return strconv.Itoa(int(args.(float64)))
	case "string":
		return args.(string)
	case "int":
		return strconv.Itoa(args.(int))
	default:
		return fmt.Sprintf("%s", args)
	}
}

func ParseInt(args interface{}) int {
	if args == nil {
		log.Panic("NPE")
	}
	intArgs, err := strconv.Atoi(ParseString(args))
	if err != nil {
		log.Panicf("类型转换失败，错误信息：%v", err)
	}
	return intArgs
}

//================  string相关  ====================

// UpperTitle 首字母大写
func UpperTitle(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

// LowerTitle 首字母小写
func LowerTitle(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}
