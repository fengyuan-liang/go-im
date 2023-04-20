// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/20 11:37
// @Author: liangfengyuan@qiniu.com

package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-im/pkg/common/xjwt"
	"go-im/utils"
)

func GetToken(ctx *gin.Context) (token string, expire int64) {
	var (
		tk *jwt.Token
		//key   string
		//value interface{}
		//exp   int64 // jwt过期时间
		err error
	)
	token, err = ctx.Cookie("jwt")
	if err != nil {
		return
	}
	tk, err = xjwt.ParseFromCookie(ctx)
	if err != nil {
		return
	}
	if exp, ok := tk.Claims.(jwt.MapClaims)["exp"]; ok {
		expire = utils.ParseInt64(exp)
		return
	}
	//for key, value = range tk.Claims.(jwt.MapClaims) {
	//	if key == "exp" {
	//		exp = utils.ParseInt(value)
	//	}
	//}
	//expire = int64(exp)
	return
}
