package main

import (
	"fmt"
	"go-im/common/driverUtil"
	"go-im/models"
)

// @Description: 启动类
// @Version: 1.0.0
// @Date: 2023/01/27 20:58
// @Author: fengyuan-liang@foxmail.com
func main() {
	driver, _ := driverUtil.GetGormDriver(&driverUtil.GormFormMySQLDriver{})
	user := models.UserBasic{}
	driver.Find(&user)
	fmt.Println(user)
	//// 路由
	//r := router.Router()
	//// 监听端口
	//r.Run(":8080")
}
