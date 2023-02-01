package models

import (
	"go-im/common/driverUtil"
	"go-im/sql/daoEntity"
	"go-im/utils"
	"gorm.io/gorm"
	"strconv"
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
	Name          string `validate:"required" reg_error_info:"姓名不能为空"`
	Age           uint8  `validate:"gt=0,lt=200" reg_error_info:"年龄不合法"`
	Password      string `gorm:"column:password" json:"password"`
	PhoneNum      string `validate:"RegexPhone" reg_error_info:"手机号格式不正确"`
	Email         string `validate:"email" reg_error_info:"email为空或格式不正确"`
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

var localDB *gorm.DB

// getDB 这里可以改成根据配置文件选择数据源
func getDB() *gorm.DB {
	if localDB == nil {
		localDB, _ = driverUtil.GetGormDriver(&driverUtil.GormFormMySQLDriver{})
	}
	return localDB
}

//===================== 查询相关 ==========================

// FindUserByName 根据名字查用户
func FindUserByName(name string) *UserBasic {
	userBasic := &UserBasic{}
	getDB().Where("name = ?", name).First(userBasic)
	return userBasic
}

// FindUserByPhone 根据名字查用户
func FindUserByPhone(phone string) *UserBasic {
	userBasic := &UserBasic{}
	getDB().Where("phone_num = ?", phone).First(userBasic)
	return userBasic
}

// PageQueryUserList 分页查询
func PageQueryUserList(pageNo int, pageSize int) []*UserBasic {
	db := getDB()
	// 初始化容器
	userBasicList := make([]*UserBasic, pageSize)
	db.Scopes(daoEntity.Paginate(pageNo, pageSize)).Find(&userBasicList)
	return userBasicList
}

// PageQueryByFilter
//
//	 @Description: 根据过滤条件进行分页查询。可以使用
//	 @param pageNo 第几页
//	 @param pageSize 每页多少条数据
//	 @param filter 回调函数，ex：
//	 func(tx *gorm.DB) {
//			for k, v := range queryMap {
//				if k == "pageNo" || k == "pageSize" {
//					continue
//				}
//				tx.Where(k, v)
//			}
//	 @return []*UserBasic 返回查询到的分页数据
func PageQueryByFilter(pageNo int, pageSize int, filter func(*gorm.DB)) []*UserBasic {
	db := getDB()
	// 初始化容器
	tx := db.Scopes(daoEntity.Paginate(pageNo, pageSize))
	// 执行回调函数
	if filter != nil {
		filter(tx)
	}
	// 初始化容器
	userBasicList := make([]*UserBasic, pageSize)
	tx.Find(&userBasicList)
	return userBasicList
}

//===================== 插入相关 ==========================

// InsetOne
//
//	@Description: 插入相关，需要防止并发情况和集群情况插入多次的问题TODO
//	@param basic 用户结构体，请传入指针
//	@return tx 返回tx
func InsetOne(basic *UserBasic) (tx *gorm.DB) {
	db := getDB()
	// 检查名字是否已经有了
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
END2:
	// 生成`userNumber` 全局唯一
	var userNumber = strconv.FormatUint(utils.GetRandomUserId(), 10)
	db.Raw("select count(1) from ? where user_number = ? limit 1", basic.TableName(), userNumber).Scan(&cnt)
	// 已经存在了，重新生成
	if cnt >= 1 {
		goto END2
	}
	basic.UserNumber = userNumber
	return db.Create(basic)
}

// ===================== 更新相关相关 ======================

// Update 更新
func Update(userId uint64, callback func(tx *gorm.DB)) {
	db := getDB()
	tx := db.Model(&UserBasic{}).Where("user_id = ?", userId)
	// 执行回调，传入更新的值
	callback(tx)
}

//===================== 删除相关 ==========================

// LogicDelOne
//
//	@Description: 逻辑删除
//	@param userId 用户id
func LogicDelOne(userId uint64) {
	DelOneByUserId(userId, true)
}

func RealDelOne(userId uint64) {
	DelOneByUserId(userId, false)
}

// DelOneByUserId
//
//	@Description: 根据userId删除用户
//	@param userId 用户id
//	@param isLogicDel 是否逻辑删除
func DelOneByUserId(userId uint64, isLogicDel bool) {
	db := getDB()
	if isLogicDel {
		// 逻辑删除，将 DeletedAt字段标记为现在的时间
		db.Where("user_id = ?", userId).Delete(&UserBasic{})
	} else {
		// 物理删除
		db.Unscoped().Where("user_id = ?", userId).Delete(&UserBasic{})
	}
}
