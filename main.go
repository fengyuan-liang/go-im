package main

import (
	"flag"
	"go-im/config"
	"go-im/pkg/commands"
	"go-im/pkg/common/xredis"
	"go-im/pkg/xgin"
	"go-im/pkg/xmysql"
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
	conf := config.InitConfig(confFile, port)
	xmysql.InitMysql(conf)
	xredis.NewRedisClient(&conf.Db.Redis)
}

// @title
func main() {
	commands.Run(xgin.GinServer{})
}
