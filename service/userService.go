package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-im/common/entity/response"
	"go-im/models"
	"go-im/utils"
	"gorm.io/gorm"
	"strconv"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/28 14:43
// @Author: fengyuan-liang@foxmail.com

// PageQueryUserList 分页查询
// @Tags 用户相关
// @BasePath /user
// @Summary 分页查询
// @param pageNo query integer true "第几页" default(1)
// @param pageSize query integer true "每页多少条" default(10)
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/pageQuery [get]
func PageQueryUserList(c *gin.Context) {
	// 拿到参数
	page := c.Query("pageNo")
	size := c.Query("pageSize")
	if page == "" || size == "" {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	// 转化类型
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	c.JSON(200, response.Ok.WithData(models.PageQueryUserList(pageInt, sizeInt)))
}

// PageQueryByFilter 带筛选条件的分页查询
// @Tags 用户相关
// @BasePath /user
// @Summary 前端请求参数应为：http://xx:xx/pageQueryByFilter?pageSize=1&pageNo=1&name=1&age=2&email=xxx@xxx
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/pageQueryByFilter [get]
func PageQueryByFilter(c *gin.Context) {
	// 拿到所有的过滤条件
	queryMap := utils.GetAllQueryParams(c)
	// 拿到参数
	page := queryMap["pageNo"]
	size := queryMap["pageSize"]
	if page == "" || size == "" {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	// 转化类型
	pageInt := utils.ParseInt(page)
	sizeInt := utils.ParseInt(size)
	// 没有过滤条件
	if len(queryMap) <= 2 {
		c.JSON(200, response.Ok.WithData(models.PageQueryUserList(pageInt, sizeInt)))
		return
	}
	// 添加所有的过滤条件
	c.JSON(200, response.Ok.WithData(models.PageQueryByFilter(pageInt, sizeInt, func(tx *gorm.DB) {
		for k, v := range queryMap {
			if k == "pageNo" || k == "pageSize" {
				continue
			}
			tx.Where(k, v)
		}
	})))
}

// CreateUser 创建一个用户
// @Tags 创建一个用户
// @BasePath /user
// @Summary 用于用户注册
// @Param param body models.UserBasic true "上传的JSON"
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/register [post]
func CreateUser(c *gin.Context) {
	type UserParams struct {
		models.UserBasic
		RePassword string // 第二次输入的密码
	}
	userParams := &UserParams{}
	if err := c.BindJSON(userParams); err != nil || userParams.Name == "" || userParams.Age <= 0 || userParams.PassWord == "" {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	if userParams.PassWord != userParams.RePassword {
		c.JSON(-1, response.AppErr.WithMsg("两次密码不一致"))
		return
	}
	userBasic := &models.UserBasic{}
	err := utils.Copy(userParams).To(userBasic)
	if err != nil {
		fmt.Println("拷贝失败")
		c.JSON(500, response.AppErr.AppendMsg(err.Error()))
		return
	}
	// 插入数据
	models.InsetOne(userBasic)
	c.JSON(200, response.Ok)
}

// DelOne 删除一个用户
// @Tags 删除一个用户
// @Summary 参数类型：{"userId":123456,"isLogicDel":true}
// @Param param body models.UserBasic true "上传的JSON"
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/delOne [post]
func DelOne(c *gin.Context) {
	type Params struct {
		UserId     uint64 `json:"userId"`
		IsLogicDel bool   `json:"isLogicDel"`
	}
	params := &Params{}
	if err := c.BindJSON(params); err != nil || params.UserId <= 0 {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	models.DelOneByUserId(params.UserId, params.IsLogicDel)
	c.JSON(200, response.Ok)
}

// Update 更新用户信息
// @Tags 更新用户信息
// @Param param body models.UserBasic true "上传的JSON"
// @Produce json
// @Router /user/update [post]
func Update(c *gin.Context) {
	jsonMap := make(map[string]interface{}) // 注意该结构接受的内容
	if err := c.BindJSON(&jsonMap); err != nil || len(jsonMap) <= 0 || jsonMap["userId"] == nil {
		c.JSON(500, response.Err.WithMsg("参数缺失"))
		return
	}
	// 不能直接用 map[string]interface{}解析，因为int会范化为float，更新会失败，必须要确定类型
	parseMap := utils.ParseMapFieldType(jsonMap, models.UserBasic{}, "userId")
	// TODO 从token中拿到userId
	var userId = uint64(utils.ParseInt(jsonMap["userId"]))
	models.Update(userId, func(tx *gorm.DB) {
		tx.Updates(parseMap)
	})
	c.JSON(200, response.Ok)
}
