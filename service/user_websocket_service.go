package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-im/common/driverHelper/redisHelper"
	"go-im/common/entity/constant"
	"net/http"
	"time"
)

//====================== websocket相关 ==============================

// 防止跨域伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsSendMsg
//
//	@Description: 通过ws发送消息
//	@args c
//	@return bizError.BizErrorer
func WsSendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("ws消息发送失败，err Info：", err.Error())
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println("ws消息发送失败，err Info：", err.Error())
		}
	}(ws)
	MsgHandle(ws, c)
}

func MsgHandle(ws *websocket.Conn, c *gin.Context) {
	for {
		// 接收redis消息队列里的消息
		msg, err := redisHelper.Subscribe(c, redisHelper.PublishKey)
		tm := time.Now().Format(constant.TIME_PATTERN)
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		fmt.Printf("msg[%v]\n", m)
		if err != nil {
			panic(err)
		}
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			panic(err)
		}
	}
}
