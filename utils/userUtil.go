package utils

import (
	"math/rand"
	"strconv"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/28 23:13
// @Author: fengyuan-liang@foxmail.com

// GetRandomUserId userId十位
func GetRandomUserId() uint64 {
END:
	uid := uint64(rand.Int31())
	if len(strconv.FormatUint(uid, 10)) < 10 {
		goto END
	}
	return uid
}
