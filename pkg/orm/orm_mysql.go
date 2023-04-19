// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 10:00
// @Author: liangfengyuan@qiniu.com

package orm

import (
	"go-im/pkg/xmysql"
	"go-im/utils"
	"gorm.io/gorm"
	"reflect"
)

type BaseEntity struct {
	Id uint64
}

type IBaseRepository[T any] interface {
	AddOrModify(t *T) error
	FindById(id interface{}) (t *T, err error)
	List(pageNo int, pageSize int) (*[]T, error)
	ListByFilter(pageNo int, pageSize int, filter func(*gorm.DB)) (*[]T, error)
}

type BaseRepository[T any] struct {
	TableName string
}

func (repo *BaseRepository[T]) FindById(id interface{}) (t *T, err error) {
	err = xmysql.DB.Where("id = ?", id).First(&t).Error
	return
}

func (repo *BaseRepository[T]) List(pageNo int, pageSize int) (*[]T, error) {
	return repo.ListByFilter(pageNo, pageSize, nil)
}

// ListByFilter
//
//	 @Description: 根据过滤条件进行分页查询。可以使用
//	 @args pageNo 第几页
//	 @args pageSize 每页多少条数据
//	 @args filter 回调函数，ex：
//	 func(tx *gorm.DB) {
//			for k, v := range queryMap {
//				if k == "pageNo" || k == "pageSize" {
//					continue
//				}
//				tx.Where(k, v)
//			}
//	 @return []*UserBasic 返回查询到的分页数据
func (repo *BaseRepository[T]) ListByFilter(pageNo int, pageSize int, filter func(*gorm.DB)) (*[]T, error) {
	// 初始化容器
	tx := xmysql.DB.Scopes(utils.Paginate(pageNo, pageSize))
	// 执行回调函数
	if filter != nil {
		filter(tx)
	}
	// 初始化容器
	t := make([]T, pageSize)
	err := tx.Find(&t).Error
	return &t, err
}

func (repo *BaseRepository[T]) AddOrModify(t *T) error {
	// 好像不能通过泛型获取字段，只能反射了
	v := reflect.ValueOf(t).FieldByName("Id")
	if t, err := repo.FindById(v.Int()); err == nil {
		// 修改，拿到改变了的字段
		return xmysql.DB.Updates(t).Error
	}
	// 没有查到， 新增
	return xmysql.DB.Create(t).Error
}

func (repo *BaseRepository[T]) DelOneById(id int64) error {
	t := new(T)
	err := xmysql.DB.Where("id = ?", id).Delete(t).Error
	return err
}
