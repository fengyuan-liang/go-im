package models

import (
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

type messageType string

type Message struct {
	gorm.Model
	FormId      string      `gorm:"column:form_id;type:varchar(50)" json:"form_id"`
	TargetId    string      `gorm:"column:target_id;type:varchar(50)" json:"target_id"`
	Type        messageType `gorm:"column:type;type:varchar(10)" json:"type"`        // 群聊 私聊 广播
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

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	// GroupSets set.Interface
}

var (
	clientMap map[int64]*Node = make(map[int64]*Node)
	rwLock                    = sync.RWMutex{}
)

func Chat(writer http.ResponseWriter, request *http.Request) {

}
