package config

import "go-im/common/driverUtil"

// @Description: 配置文件映射结构体
// @Version: 1.0.0
// @Date: 2023/01/27 23:22
// @Author: fengyuan-liang@foxmail.com

type ConfigStruct struct {
	Db driverUtil.DBConfig `yaml:"db"`
}
