package utils

import (
	"fmt"
	"log"
	"strconv"
)

// @Description: 通用工具类
// @Version: 1.0.0
// @Date: 2023/01/29 17:22
// @Author: fengyuan-liang@foxmail.com

func ParseString(args interface{}) string {
	if args == nil {
		log.Panic("NPE")
	}
	return fmt.Sprintf("%v", args)
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
