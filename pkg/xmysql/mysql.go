package xmysql

import (
	"fmt"
	"go-im/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func InitMysql(configStruct *config.ConfigStruct) {
	if DB != nil {
		return
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		configStruct.Db.Mysql.USER_NAME,
		configStruct.Db.Mysql.PASSWORD,
		configStruct.Db.Mysql.URL,
		configStruct.Db.Mysql.PORT,
		configStruct.Db.Mysql.DB_NAME)
	// 自定义sql日志模版参数
	newLogConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: func() logger.LogLevel {
				switch configStruct.Db.Mysql.LogLevel {
				case "Info":
					fallthrough
				case "info":
					return logger.Info
				case "Warn":
					fallthrough
				case "warn":
					return logger.Warn
				case "Error":
					fallthrough
				case "error":
					return logger.Error
				case "Silent":
					fallthrough
				case "silent":
					return logger.Silent
				default:
					return logger.Info
				}
			}(),
			Colorful: configStruct.Db.Mysql.Colorful,
		},
	)
	var err error
	// 自定义日志模版
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogConfig})
	if err != nil {
		panic("数据库连接建立失败，错误信息为：" + err.Error())
	}
}
