package driverUtil

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"go-im/common/bizError"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// @Author: fengyuan-liang@foxmail.com
// @Description:
// @File:  config
// @Version: 1.0.0
// @Date: 2023/01/19 15:59

// DBConfig 映射yml配置文件结构体
type DBConfig struct {
	Labels map[string]*BaseDBStruct `yaml:"db"`
}

type BaseDBStruct struct {
	DRIVER_NAME string `yaml:"DRIVER_NAME"`
	USER_NAME   string `yaml:"USER_NAME"`
	PASSWORD    string `yaml:"PASSWORD"`
	URL         string `yaml:"URL"`
	PORT        string `yaml:"PORT"`
	DB_NAME     string `yaml:"DB_NAME"`
}

var db *sql.DB
var client *mongo.Client
var gormDb *gorm.DB

// InitMySQLDB 初始化mysql数据库连接
func InitMySQLDB() (*sql.DB, bizError.BizErrorer) {
	if db != nil {
		return db, nil
	}
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		viper.GetString("db.mysql.USER_NAME"),
		viper.GetString("db.mysql.PASSWORD"),
		viper.GetString("db.mysql.URL"),
		viper.GetString("db.mysql.PORT"),
		viper.GetString("db.mysql.DB_NAME"))
	fmt.Printf("dsb读取成功，dsb【%v】\n", dsn)
	// 不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		bizError.NewBizError(fmt.Sprintf("数据库连接建立失败，错误信息为：%v", err))
		return db, bizError.NewBizError(fmt.Sprintf("数据库连接建立失败，错误信息为：%v", err))
	}
	err = db.Ping()
	// 尝试和数据库建立连接，会校验dsn的正确性
	if err != nil {
		return db, bizError.NewBizError(fmt.Sprintf("dsn错误，错误信息为：%v", err))
	}
	return db, nil
}

// InitMongoDB 初始化mongodb连接
func InitMongoDB() (*mongo.Client, bizError.BizErrorer) {
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dsn := fmt.Sprintf("mongodb://%v:%v", viper.GetString("db.mongodb.URL"),
		viper.GetString("db.mongodb.MOGO_PORT"))
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil && client.Ping(ctx, nil) != nil {
		return nil, bizError.NewBizError(fmt.Sprintf("数据库连接失败，错误信息为：%v", err))
	}
	fmt.Printf("初始化数据库连接成功，dsb【%v】\n", dsn)
	return client, nil
}

// InitGormDriver 初始化gorm对象
func InitGormDriver() (*gorm.DB, bizError.BizErrorer) {
	if gormDb != nil {
		return gormDb, nil
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		viper.GetString("db.mysql.USER_NAME"),
		viper.GetString("db.mysql.PASSWORD"),
		viper.GetString("db.mysql.URL"),
		viper.GetString("db.mysql.PORT"),
		viper.GetString("db.mysql.DB_NAME"))
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, bizError.NewBizError("数据库连接建立失败，错误信息为：", err.Error())
	}
	return gormDb, nil
}
