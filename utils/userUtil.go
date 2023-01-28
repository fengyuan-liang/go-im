package utils

import "math/rand"

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/28 23:13
// @Author: fengyuan-liang@foxmail.com

func GetRandomUserId() uint64 {
	return uint64(rand.Int63())
}
