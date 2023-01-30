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

func ParseInt64(args interface{}) int64 {
	return ParseDefaultInt(args, func(str string) (interface{}, error) {
		return strconv.ParseInt(str, 10, 64)
	}).(int64)
}

func ParseInt32(args interface{}) int32 {
	return int32(ParseInt64(args))
}

func ParseUint8(args interface{}) uint8 {
	return uint8(ParseInt64(args))
}

func ParseUint32(args interface{}) uint32 {
	return uint32(ParseInt64(args))
}

func ParseUint64(args interface{}) uint64 {
	return uint64(ParseInt64(args))
}

func ParseInt(args interface{}) int {
	return ParseDefaultInt(args, func(str string) (interface{}, error) {
		return strconv.Atoi(str)
	}).(int)
}

func ParseDefaultInt(args interface{}, callBack func(str string) (interface{}, error)) interface{} {
	if args == nil {
		log.Panic("NPE")
	}
	intArgs, err := callBack(ParseString(args))
	if err != nil {
		log.Panicf("类型转换失败，错误信息：%v", err)
	}
	return intArgs
}

// ParseType 将args转为指定的类型【stringType】
func ParseType(args interface{}, stringType string) (newArgs interface{}) {
	switch stringType {
	case "int":
		return ParseInt(args)
	case "int32":
		return ParseInt32(args)
	case "int64":
		return ParseInt64(args)
	case "uint8":
		return ParseUint8(args)
	case "uint32":
		return ParseUint32(args)
	case "uint64":
		return ParseUint64(args)
	case "string":
		fallthrough
	default:
		return ParseString(args)
	}
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
