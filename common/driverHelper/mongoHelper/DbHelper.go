package mongdbHelper

import (
	"container/list"
	"context"
	"fmt"
	"go-im/common/bizError"
	"go-im/common/driverHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// @Author: fengyuan-liang@foxmail.com
// @Description: 封装mongodb原生操作
// @File:  config
// @Version: 1.0.0
// @Date: 2023/01/19 15:59

type MongoDbHelper struct {
	Client          *mongo.Client
	collection      *mongo.Collection // 连接对象
	DB_NAME         string
	COLLECTION_NAME string
}

// InitMongoDbHelper 初始化数据库连接
func InitMongoDbHelper(dbName string, collectionName string) (MongoDbHelper, bizError.BizErrorer) {
	var dbConnectionError bizError.BizErrorer
	var mongodbHelper MongoDbHelper
	// 检查参数
	if collectionName == "" || dbName == "" {
		return mongodbHelper, bizError.NewBizError("传参不能为空！")
	}
	// 初始化数据库连接
	start := time.Now()
	mongodbHelper.Client, dbConnectionError = driverHelper.InitMongoDB()
	if dbConnectionError != nil {
		return mongodbHelper, bizError.NewBizError("数据库连接建立失败，错误信息为：", dbConnectionError.BizError())
	}
	end := time.Now()
	mongodbHelper.DB_NAME = dbName
	mongodbHelper.COLLECTION_NAME = collectionName
	// 连接对象
	mongodbHelper.collection = mongodbHelper.Client.Database(dbName).Collection(collectionName)
	fmt.Printf("数据库连接建立成功，耗时%v\n", end.Sub(start).String())
	return mongodbHelper, nil
}

//========================= 插入相关 ====================================

// InsertOne 插入一条数据
func (dbHelper *MongoDbHelper) InsertOne(bean interface{}) (any, bizError.BizErrorer) {
	// 1. 如果连接未建立，提示用户需要调用`InitMongoDbHelper`主动建立连接
	if dbHelper.Client == nil {
		return nil, bizError.NewBizError("连接未建立！请调用InitMongoDbHelper建立连接")
	}
	// 2. 插入一个文档
	oneResult, err := dbHelper.collection.InsertOne(context.TODO(), bean)
	if err != nil {
		return nil, bizError.NewBizError("插入文档失败，错误原因：", err.Error())
	}
	// 3. 返回insertId
	return oneResult, nil
}

// InsertMany 插入多条数据
func (dbHelper *MongoDbHelper) InsertMany(documents ...interface{}) (*mongo.InsertManyResult, bizError.BizErrorer) {
	insertManyResult, err := dbHelper.collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return nil, bizError.NewBizError("批量插入文档失败，错误原因：", err.Error())
	}
	return insertManyResult, nil
}

//========================= 查询相关 ====================================

// FindAll 查询所有文档，一般不允许使用
func (dbHelper *MongoDbHelper) FindAll() (*list.List, bizError.BizErrorer) {
	// 没有过滤条件，即查询全部
	return dbHelper.FindByFilter(bson.D{})
}

// FindOne 只查询一条数据
func (dbHelper *MongoDbHelper) FindOne() (*list.List, bizError.BizErrorer) {
	// 没有过滤条件，即查询全部（这里还不对，需要TODO一下）
	return dbHelper.FindByFilter(bson.D{{"skip", 0}, {"limit", 1}})
}

// FindByFilter 根据查询条件进行查询
func (dbHelper *MongoDbHelper) FindByFilter(filter bson.D) (*list.List, bizError.BizErrorer) {
	// 默认查询超时时间20s
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// 存储结果集合
	resultList := list.New()
	defer cancel()
	// 需要过滤条件将`bson.D{}`抽象出来即可，例如`bson.D{{"name", "tom"}}`
	cur, err := dbHelper.collection.Find(ctx, filter)
	if err != nil {
		return resultList, bizError.NewBizError("查询失败：", err.Error())
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Panic("关闭连接失败")
		}
	}(cur, ctx)
	for cur.Next(ctx) {
		var result bson.D
		err2 := cur.Decode(&result)
		if err2 != nil {
			return nil, bizError.NewBizError("解码失败：", err.Error())
		}
		resultList.PushFront(result)
	}
	return resultList, nil
}

// ========================= 更新相关 ====================================

// UpdateByPrams 根据参数更新，不用写$set，后续将优化成链式编程
func (dbHelper *MongoDbHelper) UpdateByPrams(query bson.D, update bson.D) (int64, bizError.BizErrorer) {
	return dbHelper.Update(query, bson.D{{"$set", update}})
}

// Update 根据查询和更新条件进行更新
func (dbHelper *MongoDbHelper) Update(query bson.D, update bson.D) (int64, bizError.BizErrorer) {
	// 默认查询超时时间20ms
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	updateResult, err := dbHelper.collection.UpdateMany(ctx, query, update)
	if err != nil {
		return -1, bizError.NewBizError("解码失败：", err.Error())
	}
	return updateResult.ModifiedCount, nil
}

// ========================= 删除相关 ====================================

// Del 删除
func (dbHelper *MongoDbHelper) Del(delFilter bson.D) (int64, bizError.BizErrorer) {
	// 默认查询超时时间20ms
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	delResult, err := dbHelper.collection.DeleteMany(ctx, delFilter)
	if err != nil {
		return -1, bizError.NewBizError("删除失败：", err.Error())
	}
	return delResult.DeletedCount, nil
}
