package main

import (
	"go-im/router"
)

// @Description: 启动类
// @Version: 1.0.0
// @Date: 2023/01/27 20:58
// @Author: fengyuan-liang@foxmail.com

// @title
func main() {
	// 路由
	r := router.Router()
	// 监听端口
	r.Run(":8080")
}
