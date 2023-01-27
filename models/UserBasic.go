package models

import "gorm.io/gorm"

// @Description: 相关实体类
// @Version: 1.0.0
// @Date: 2023/01/27 21:03
// @Author: fengyuan-liang@foxmail.com

// UserBasic 用户基础属性表
type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	PhoneNum      string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartBeatTime uint64
	LoginOutTime  uint64 `gorm:"column:login_out_tim" json:"login_out_time"`
	isLogin       bool
	DeviceInfo    string // 登陆设备信息
}

// TableName 用户表名
func (table *UserBasic) TableName() string {
	return "user_basic"
}
