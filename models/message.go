package models

import (
	"gorm.io/gorm"
)

type MessageType int

const (
	SINGLE_CHAT MessageType = iota
	GROUP_CHAT
	BROADCASE_CHAT
)

type Message struct {
	gorm.Model
	FormId      uint64      `gorm:"column:form_id;type:varchar(50)" json:"form_id"`
	TargetId    uint64      `gorm:"column:target_id;type:varchar(50)" json:"target_id"`
	Type        MessageType `gorm:"column:type;type:varchar(10)" json:"type"`        // 群聊 私聊 广播
	Content     string      `gorm:"column:content;type:varchar(200)" json:"content"` // 消息内容
	ContentType int         `gorm:"column:content_type" json:"content_type"`         // 消息类型 图片 广播 文字
	Avatar      string      `gorm:"column:avatar;type:varchar(100)" json:"avatar"`
	Desc        string      `gorm:"column:desc;type:varchar(200)" json:"desc"`
	Amount      int         `gorm:"column:amount;type:int" json:"amount"` // 其他数字统计
}

// TableName 用户表名
func (message *Message) TableName() string {
	return "message"
}
