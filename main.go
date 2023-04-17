package main

import (
	"flag"
	"go-im/common/driverHelper"
	"go-im/router"
	"go-im/utils"
)

// @Description: 启动类
// @Version: 1.0.0
// @Date: 2023/01/27 20:58
// @Author: fengyuan-liang@foxmail.com

var (
	confFile = flag.String("f", "./config/application.yml", "config file")
	port     = flag.Int("p", 8080, "please input port")
)

func init() {
	flag.Parse()
	utils.InitConfig(*confFile)
	driverHelper.GetOrDefaultGormDriver(&driverHelper.GormFormMySQLDriver{})
	driverHelper.GetOrDefaultRedis()
}

// @title
func main() {
	// 路由
	r := router.Router()
	// 监听端口
	r.Run(":" + utils.ParseString(*port))
}
