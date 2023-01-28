package models

import (
	"go-im/common/driverUtil"
	"go-im/sql/daoEntity"
	"go-im/utils"
	"gorm.io/gorm"
)

// @Description: 相关实体类
// @Version: 1.0.0
// @Date: 2023/01/27 21:03
// @Author: fengyuan-liang@foxmail.com

// UserBasic 用户基础属性表
type UserBasic struct {
	gorm.Model
	UserId        uint64 `gorm:"column:user_id" json:"userId"`
	UserNumber    string `gorm:"column:user_number" json:"userNumber"`
	Name          string
	Age           uint8
	PassWord      string `gorm:"column:password" json:"password"`
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

//===================== 查询相关 ==========================

// PageQueryUserList 分页查询
func PageQueryUserList(pageNo int, pageSize int) []*UserBasic {
	db, _ := driverUtil.GetGormDriver(&driverUtil.GormFormMySQLDriver{})
	// 初始化容器
	userBasicList := make([]*UserBasic, pageSize)
	db.Scopes(daoEntity.Paginate(pageNo, pageSize)).Find(&userBasicList)
	return userBasicList
}

//===================== 插入相关 ==========================

// InsetOne 插入相关，需要防止并发情况和集群情况插入多次的问题TODO
func InsetOne(basic UserBasic) (tx *gorm.DB) {
	db, _ := driverUtil.GetGormDriver(&driverUtil.GormFormMySQLDriver{})
END:
	// 生成`userId`，必须全局唯一
	var cnt = 0
	var userId = utils.GetRandomUserId()
	db.Raw("select count(1) from ? where user_id = ? limit 1", basic.TableName(), userId).Scan(&cnt)
	// 已经存在了，重新生成
	if cnt >= 1 {
		goto END
	}
	basic.UserId = userId
	return db.Create(&basic)
}
