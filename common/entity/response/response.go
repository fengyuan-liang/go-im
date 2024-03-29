package response

import (
	"encoding/json"
	"log"
)

// @Description: 返回参数封装
// @Version: 1.0.0
// @Date: 2023/01/28 17:34
// @Author: fengyuan-liang@foxmail.com

var (
	Ok     = New(200, "操作成功")
	Err    = New(500, "操作失败")
	AppErr = New(-1, "请求异常")
)

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// New reply 构造函数
func New(code int, msg string) *Reply {
	return &Reply{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// WithMsg 覆盖响应消息
func (t *Reply) WithMsg(msg string) Reply {
	return Reply{
		Code: t.Code,
		Msg:  msg,
		Data: t.Data,
	}
}

// AppendMsg 覆盖响应消息
func (t *Reply) AppendMsg(msg string) Reply {
	return Reply{
		Code: t.Code,
		Msg:  t.Msg + "，错误信息为：【" + msg + "】",
		Data: t.Data,
	}
}

// WithData 覆盖响应数据
func (t *Reply) WithData(data interface{}) Reply {
	return Reply{
		Code: t.Code,
		Msg:  t.Msg,
		Data: data,
	}
}

// Json 返回JSON格式的数据
func (t *Reply) Json() []byte {
	s := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: t.Code,
		Msg:  t.Msg,
		Data: t.Data,
	}
	log.Printf("%+v\n", s)
	raw, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	return raw
}
