package handle

import (
	"go-im/common/bizError"
	"go-im/models"
	"log"
)

// LoginHandle
// @Description: 登录接口
type LoginHandle interface {
	// login
	//  @Description: 登录接口
	//  @param map[string]interface{} 登录时传递的参数
	//
	login(paramsMap map[string]interface{}) (*models.UserBasic, bizError.BizErrorer)
}

type LoginFactory struct {
	// 登录工厂
	loginStrategyFactoryMap map[string]LoginHandle
}

func (loginFactory *LoginFactory) GetLoginStrategy(loginSign string) LoginHandle {
	return loginFactory.loginStrategyFactoryMap[loginSign]
}

func (loginFactory *LoginFactory) Register(loginSign string, loginHandle LoginHandle) {
	if loginSign == "" || loginHandle == nil {
		log.Fatal("fail register loginHandle，the args is illegal")
	}
	loginFactory.loginStrategyFactoryMap[loginSign] = loginHandle
}
