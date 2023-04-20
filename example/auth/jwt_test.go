// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/20 11:24
// @Author: liangfengyuan@qiniu.com

package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-im/pkg/common/xjwt"
	"go-im/pkg/constant"
	"go-im/utils"
	"testing"
)

func TestJWT(t *testing.T) {
	jwtToken, err := xjwt.CreateToken(1000, constant.ANDROID, false, constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwtToken.Token)
	token, _ := xjwt.ParseFromToken(jwtToken.Token)
	fmt.Println(token.Claims)
	fmt.Println(utils.ParseInt64(token.Claims.(jwt.MapClaims)["exp"]))
}
