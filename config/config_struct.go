package config

import (
	"go-im/utils"
)

// @Description: 配置文件映射结构体
// @Version: 1.0.0
// @Date: 2023/01/27 23:22
// @Author: fengyuan-liang@foxmail.com

// ConfigStruct 配置文件映射实体类
type ConfigStruct struct {
	Db DBConfig `yaml:"db"`
}

// DBConfig 映射yml配置文件结构体
type DBConfig struct {
	// 如果
	Mysql   BaseDBStruct `yaml:"mysql"`
	MongoDb BaseDBStruct `yaml:"mongodb"`
	Redis   RedisStruct  `yaml:"redis"`
}

type BaseDBStruct struct {
	DRIVER_NAME   string `yaml:"DRIVER_NAME"`
	USER_NAME     string `yaml:"USER_NAME"`
	PASSWORD      string `yaml:"PASSWORD"`
	URL           string `yaml:"URL"`
	PORT          string `yaml:"PORT"`
	DB_NAME       string `yaml:"DB_NAME"`
	SlowThreshold int    `yaml:"SlowThreshold"` // 慢sql阈值，单位毫秒
	LogLevel      string `yaml:"LogLevel"`      // 日志打印级别 例如info error
	Colorful      bool   `yaml:"Colorful"`      // 是否彩色打印sql
}

// RedisStruct
// @Description: redis数据源对应配置
type RedisStruct struct {
	URL          string `yaml:"URL"`
	PORT         string `yaml:"PORT"`
	PASSWORD     string `yaml:"PASSWORD"`
	DB           string `yaml:"DB"`
	POOL_SIZE    int    `yaml:"POOL_SIZE"`
	MinIdleConns int    `yaml:"MinIdleConns"`
	PoolTimeout  int    `yaml:"PoolTimeout"`
}

// GetDbInfo 获取config下`db`的配置
func (dbInfo ConfigStruct) GetDbInfo() ConfigStruct {
	vip := utils.GetOrDefaultViper()
	// 读取db配置
	if err := vip.Unmarshal(&dbInfo); err != nil {
		panic("解析db配置异常, info:" + err.Error())
	}
	return dbInfo
}
