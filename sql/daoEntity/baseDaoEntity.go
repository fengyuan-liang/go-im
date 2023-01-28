package daoEntity

import "gorm.io/gorm"

// @Description: 通用dao层实体类
// @Version: 1.0.0
// @Date: 2023/01/28 14:32
// @Author: fengyuan-liang@foxmail.com

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
