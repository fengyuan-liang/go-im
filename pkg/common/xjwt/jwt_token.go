package xjwt

import "go-im/pkg/constant"

type JwtToken struct {
	Token     string
	SessionId string
	Expire    int64
	Uid       int64
	Platform  constant.PlatformType
}
