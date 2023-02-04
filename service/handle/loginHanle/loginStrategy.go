package loginHanle

import (
	"go-im/common/bizError"
	"go-im/models"
	"go-im/utils"
)

var loginFactory = LoginFactory{loginStrategyFactoryMap: map[string]LoginHandle{}}

func init() {
	// 账号密码登录
	AddLoginStrategy("ACCOUNT_PWD_LOGIN_STRATEGY", &NamePwdLoginStrategy{})
}

// LoginBySign
//
//	@Description: 根据登录策略进行登录
//	@args loginSign 具体的登录方式，ex：账号密码登录，手机号码登录
//	@args paramsMap 登录的参数
//	@return *models.UserBasic 返回用户信息
//	@return bizError.BizErrorer
func LoginBySign(loginSign string, paramsMap map[string]interface{}) (*models.UserBasic, bizError.BizErrorer) {
	return loginFactory.GetLoginStrategy(loginSign).login(paramsMap)
}

// AddLoginStrategy
//
//	@Description: 增加登录的逻辑。符合开闭原则
//	@args loginSign 登录方式标识
//	@args loginHandle 具体的handle
func AddLoginStrategy(loginSign string, loginHandle LoginHandle) {
	loginFactory.Register(loginSign, loginHandle)
}

type (
	NamePwdLoginStrategy     struct{}
	PhoneNumberLoginStrategy struct{}
)

func (p *NamePwdLoginStrategy) login(paramsMap map[string]interface{}) (*models.UserBasic, bizError.BizErrorer) {
	// 用户名和密码
	name := utils.ParseString(paramsMap["name"])
	password := utils.ParseString(paramsMap["password"])
	userBasic := models.FindUserByName(name)
	if userBasic.Name == "" {
		return nil, bizError.NewBizError("用户不存在，请注册")
	}
	if !utils.CheckBySalt(password, userBasic.Salt, userBasic.Password) {
		return nil, bizError.NewBizError("密码不正确")
	}
	return userBasic, nil
}

func (p *PhoneNumberLoginStrategy) login(params map[string]interface{}) (*models.UserBasic, bizError.BizErrorer) {
	// p.login(nil)
	return nil, nil
}
