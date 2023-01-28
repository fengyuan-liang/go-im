package driverUtil

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"go-im/common/bizError"
	"go-im/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// @Description: 驱动设置
// @File:  config
// @Version: 1.0.0
// @Date: 2023/01/19 15:59
// @Author: fengyuan-liang@foxmail.com

var db *sql.DB
var client *mongo.Client
var gormDb *gorm.DB

//=========================== 原生sql操作 ==================================

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

//=========================== gorm操作 ==================================

type InitGormDriverInterface interface {
	InitGormDriver() (*gorm.DB, bizError.BizErrorer)
}

// GetGormDriver 多态方法
func GetGormDriver(gormDriver InitGormDriverInterface) (*gorm.DB, bizError.BizErrorer) {
	return gormDriver.InitGormDriver()
}

type GormDriverBasic struct {
	DSN string
}

type GormFormMySQLDriver struct {
	basic GormDriverBasic
}

func (d *GormFormMySQLDriver) InitGormDriver() (*gorm.DB, bizError.BizErrorer) {
	if gormDb != nil {
		return gormDb, nil
	}
	// 获取配置
	var configStruct config.ConfigStruct
	configStruct = configStruct.GetDbInfo()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		configStruct.Db.Mysql.USER_NAME,
		configStruct.Db.Mysql.PASSWORD,
		configStruct.Db.Mysql.URL,
		configStruct.Db.Mysql.PORT,
		configStruct.Db.Mysql.DB_NAME)
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, bizError.NewBizError("数据库连接建立失败，错误信息为：", err.Error())
	}
	return gormDb, nil
}

// TODO gorm现在好像不支持mongodb，因为原生操作就已经很方便了

type GormFormMongoDbDriver struct {
	basic GormDriverBasic
}

func (d *GormFormMongoDbDriver) InitGormDriver() (*gorm.DB, bizError.BizErrorer) {
	if gormDb != nil {
		return gormDb, nil
	}
	// 获取配置
	var configStruct config.ConfigStruct
	configStruct = configStruct.GetDbInfo()
	dsn := fmt.Sprintf("mongodb://%v:%v",
		configStruct.Db.MongoDb.URL,
		configStruct.Db.MongoDb.PORT)
	fmt.Println(dsn)
	var err error
	gormDb, err = gorm.Open(nil, &gorm.Config{})
	if err != nil {
		return nil, bizError.NewBizError("数据库连接建立失败，错误信息为：", err.Error())
	}
	return gormDb, nil
}
