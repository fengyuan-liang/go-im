package main

import (
	"go-im/utils"
)

// @Description: 启动类
// @Version: 1.0.0
// @Date: 2023/01/27 20:58
// @Author: fengyuan-liang@foxmail.com
func main() {
	// 读取配置文件
	utils.InitConfig()
	//driver, _ := driverUtil.InitGormDriver()
	//user := models.UserBasic{}
	//driver.Find(&user)
	//fmt.Println(user)
	//// 路由
	//r := router.Router()
	//// 监听端口
	//r.Run(":8080")
}
