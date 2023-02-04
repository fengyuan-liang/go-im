package loginHanle

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
	login(paramsMap map[string]interface{}) (*models.UserBasic, bizError.BizErrorer)
}

type LoginFactory struct {
	// 登录工厂
	loginStrategyFactoryMap map[string]LoginHandle
}

// GetLoginStrategy
//
//	@Description: 获取具体登录逻辑执行的handle
//	@receiver loginFactory 登录工厂
//	@args loginSign 登录标识，由前端提供
//	@return LoginHandle 返回具体执行handle
func (loginFactory *LoginFactory) GetLoginStrategy(loginSign string) LoginHandle {
	return loginFactory.loginStrategyFactoryMap[loginSign]
}

// Register
//
//	@Description: 将登录逻辑注册到登录工厂中
//	@receiver loginFactory 登录工厂
//	@args loginSign 登录标识，由前端提供
//	@args loginHandle 具体执行handle
func (loginFactory *LoginFactory) Register(loginSign string, loginHandle LoginHandle) {
	if loginSign == "" || loginHandle == nil {
		log.Fatal("fail register loginHandle，the args is illegal")
	}
	loginFactory.loginStrategyFactoryMap[loginSign] = loginHandle
}
