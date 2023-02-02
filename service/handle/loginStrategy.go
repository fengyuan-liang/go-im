package handle

import (
	"go-im/common/bizError"
	"go-im/models"
	"go-im/utils"
)

var loginFactory LoginFactory

func init() {
	// 电话登录
	loginFactory.Register("PHONE_NUMBER_LOGIN_STRATEGY", &PhoneNumberLoginStrategy{})
}

// LoginBySign 根据登录策略进行登录
func LoginBySign(loginSign string, paramsMap map[string]interface{}) (*models.UserBasic, bizError.BizErrorer) {
	return loginFactory.GetLoginStrategy(loginSign).login(paramsMap)
}

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
	return nil
}
