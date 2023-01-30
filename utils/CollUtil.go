package utils

// @Description: 集合工具类
// @Version: 1.0.0
// @Date: 2023/01/30 13:34
// @Author: fengyuan-liang@foxmail.com

func ContainsValue(targetValue interface{}, args ...string) bool {
	for _, v := range args {
		if targetValue == v {
			return true
		}
	}
	return false
}

func ArrayContainsValue(targetValue *interface{}, args []*interface{}) bool {
	for _, v := range args {
		if targetValue == v {
			return true
		}
	}
	return false
}
