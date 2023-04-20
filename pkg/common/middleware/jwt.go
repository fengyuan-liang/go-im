// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 14:23
// @Author: liangfengyuan@qiniu.com

package middleware

import (
	"github.com/gin-gonic/gin"
	"go-im/common/entity/response"
	"go-im/pkg/common/xjwt"
	"go-im/pkg/common/xlog"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := xjwt.ParseFromHeader(ctx)
		if err != nil {
			ctx.Abort()
			ctx.SecureJSON(-1, response.AppErr.WithMsg("请登录"))
			return
		}
		xlog.Infof("验证成功, token[%v]", token)
	}
}
