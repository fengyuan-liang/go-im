package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

// @Description: 系统初始化相关
// @Version: 1.0.0
// @Date: 2023/01/27 22:25
// @Author: fengyuan-liang@foxmail.com

var vip *viper.Viper

// InitConfig 读取config.yml中的配置文件
func InitConfig() *viper.Viper {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	vip = viper.New()
	vip.AddConfigPath(path + "/config") //设置读取的文件路径
	vip.SetConfigName("application")    //设置读取的文件名
	vip.SetConfigType("yaml")           //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal("配置文件初始化失败，info：", err)
	}
	return vip
}

// GetOrDefaultViper 获取解析到的viper
func GetOrDefaultViper() *viper.Viper {
	if vip == nil {
		vip = InitConfig()
	}
	return vip
}
