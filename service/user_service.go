package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-im/common/entity/response"
	"go-im/models"
	"go-im/pkg/common/xjwt"
	"go-im/pkg/constant"
	"go-im/repository"
	"go-im/service/handle/loginHanle"
	"go-im/utils"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/28 14:43
// @Author: fengyuan-liang@foxmail.com

var (
	userValidate = validator.New()
)

func init() {
	userValidate.RegisterValidation("RegexPhone", utils.RegexPhone)
}

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

type UserParams struct {
	Name       string `json:"name" validate:"required" reg_error_info:"姓名不能为空"`
	Password   string `json:"password" gorm:"column:password" json:"password"`
	RePassword string `json:"rePassword" validate:"eqfield=Password" reg_error_info:"两次密码不一样"`
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
	userParams := &UserParams{}
	if err := c.BindJSON(userParams); err != nil {
		c.JSON(-1, response.AppErr.WithMsg("参数有误，err:"+err.Error()))
		return
	}
	if err := userValidate.Struct(userParams); err != nil {
		c.JSON(-1, response.AppErr.WithMsg(utils.ProcessErr(userParams, err)))
		return
	}
	// 检查用户名 电话是否已经存在
	userByName := models.FindUserByName(userParams.Name)
	if userByName.Name != "" {
		c.JSON(-1, response.AppErr.WithMsg("用户名已注册"))
		return
	}
	//userByPhone := models.FindUserByPhone(userParams.PhoneNum)
	//if userByPhone.PhoneNum != "" {
	//	c.JSON(-1, response.AppErr.WithMsg("手机号已注册"))
	//	return
	//}
	userBasic := &models.UserBasic{}
	err := utils.Copy(userParams).To(userBasic)
	if err != nil {
		fmt.Println("拷贝失败")
		c.JSON(500, response.AppErr.AppendMsg(err.Error()))
		return
	}
	// 加盐用来存储密码
	salt := fmt.Sprintf("%06d", rand.Int31())
	userBasic.Salt = salt
	userBasic.Password = utils.EncodeBySalt(userBasic.Password, salt)
	// 插入数据
	repository.GetUserRepo().AddOrModify(userBasic)
	c.JSON(200, response.Ok)
}

type DelOneParams struct {
	UserId     uint64 `json:"userId"`
	IsLogicDel bool   `json:"isLogicDel"`
}

// DelOne 删除一个用户
// @Tags 删除一个用户
// @Summary 参数类型：{"userId":123456,"isLogicDel":true}
// @Param param body models.UserBasic true "上传的JSON"
// @Produce json
// @Success 200 {UserBasic} []*UserBasic
// @Router /user/delOne [post]
func DelOne(c *gin.Context) {
	params := &DelOneParams{}
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

type loginVO struct {
	models.UserBasic
	Token string `json:"token"`
}

type LoginParams struct {
	Name      string `validate:"required" reg_error_info:"姓名不能为空" json:"name"`
	Password  string `validate:"required" reg_error_info:"密码不能为空" json:"password"`
	LoginSign string `validate:"required" reg_error_info:"登录标识不能为空" json:"loginSign"`
}

// Login 通用登录接口
// @Tags 通用登录接口
// @Param param body models.UserBasic true "上传的JSON"
// @Produce json
// @Router /user/login [post]
func Login(c *gin.Context) {
	params := &LoginParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(-1, response.Err.WithMsg("参数有误，err:"+err.Error()))
		return
	}
	if err := userValidate.Struct(params); err != nil {
		c.JSON(-1, response.AppErr.WithMsg(utils.ProcessErr(params, err)))
		return
	}
	parseMap, _ := utils.ParseMap(params, "json")
	// 登录
	if userBasic, err := loginHanle.LoginBySign(params.LoginSign, parseMap); err != nil {
		c.JSON(-1, response.AppErr.WithMsg(err.ErrorMsg()))
	} else {
		vo := &loginVO{}
		utils.Copy(userBasic).To(vo)
		// 生成token
		jwtToken, _ := xjwt.CreateToken(int64(userBasic.UserId), constant.ANDROID, false, constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
		vo.Token = jwtToken.Token
		c.JSON(200, response.Ok.WithData(vo))
	}
}
