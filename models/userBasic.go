package models

import (
	"go-im/common/driverUtil"
	"go-im/sql/daoEntity"
	"gorm.io/gorm"
)

// @Description: 相关实体类
// @Version: 1.0.0
// @Date: 2023/01/27 21:03
// @Author: fengyuan-liang@foxmail.com

// UserBasic 用户基础属性表
type UserBasic struct {
	gorm.Model
	Name          string
	Age           uint8
	PassWord      string
	PhoneNum      string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartBeatTime uint64
	LogOutTime    uint64 `gorm:"column:logout_time" json:"logout_time"`
	isLogin       bool
	DeviceInfo    string // 登陆设备信息
}

// TableName 用户表名
func (user *UserBasic) TableName() string {
	return "user_basic"
}

// PageQueryUserList 分页查询
func PageQueryUserList(pageNo int, pageSize int) []*UserBasic {
	db, _ := driverUtil.GetGormDriver(&driverUtil.GormFormMySQLDriver{})
	// 初始化容器
	userBasicList := make([]*UserBasic, pageSize)
	db.Scopes(daoEntity.Paginate(pageNo, pageSize)).Find(&userBasicList)
	return userBasicList
}
