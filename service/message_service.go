// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 18:53
// @Author: liangfengyuan@qiniu.com

package service

import (
	"encoding/json"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go-im/models"
	"go-im/pkg/entity/vo"
	"io"
	"net"
	"net/http"
	"sync"
)

func init() {
	go udpSendProc()
	go udpRecvProc()
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets mapset.Set
}

var (
	clientMap   = make(map[uint64]*Node)
	rwLock      = sync.RWMutex{}
	sendMsgLock = sync.RWMutex{}
	isvalida    = true
	udpSendChan = make(chan []byte, 1024)
)

func Chat(ctx *gin.Context) {
	var (
		node *Node
		conn *websocket.Conn
		err  error
	)
	// 1. 参数校验
	messageVO := &vo.MessageVO{}
	err = ctx.BindQuery(messageVO)
	if err != nil {
		panic(err)
	}
	conn, err = (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Panicf("chat websocket panic， err[%v]", err.Error())
	}
	// 2. 获取连接
	node = &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: mapset.NewSet(),
	}
	// 3. 用户关系

	// 4. userId 与 Node进行绑定
	rwLock.Lock()
	clientMap[messageVO.UserId] = node
	rwLock.Unlock()
	// 5. 完成发送逻辑
	go sendMessageTask(node)
	// 6. 完成接收逻辑
	go recvMessageTask(node)
	// 欢迎
	sendMsg(messageVO.TargetId, []byte("欢迎进入聊天室"))
}

func sendMessageTask(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}

func recvMessageTask(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			logrus.Error(err)
			return
		}
		broadMsg(data)
		logrus.Infof("[ws] receive message[%v]", data)
	}
}

func broadMsg(data []byte) {
	udpSendChan <- data
}

func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	defer CloseConn(conn)
	if err != nil {
		logrus.Error(err)
	}
	for {
		select {
		case data := <-udpSendChan:
			_, err = conn.Write(data)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func udpRecvProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		// 所有id都可以连接
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		logrus.Error(err)
	}
	defer CloseConn(udpConn)
	for {
		buffer := make([]byte, 512)
		offset, err := udpConn.Read(buffer[0:])
		if err != nil {
			logrus.Error(err)
			return
		}
		dispatch(buffer[0:offset])
	}
}

// dispatch
func dispatch(data []byte) {
	msg := models.Message{}
	json.Unmarshal(data, &msg)
	switch msg.Type {
	case models.SINGLE_CHAT:
		sendSingleChat(msg)
	case models.GROUP_CHAT:
		sendGroupChat(msg)
	case models.BROADCASE_CHAT:
		sendBroadCastChat(msg)
	default:
		logrus.Errorf("no match type[%v]", msg.Type)
	}

}

func sendSingleChat(msg models.Message) {

}

func sendGroupChat(msg models.Message) {

}

func sendBroadCastChat(msg models.Message) {

}

func sendMsg(targetId uint64, msg []byte) {
	sendMsgLock.Lock()
	// 获取到websocket连接
	if node, ok := clientMap[targetId]; ok {
		node.DataQueue <- msg
	} else {
		logrus.Errorf("have no websocket connection, targetId[%v]", targetId)
	}
	sendMsgLock.Unlock()
}

func CloseConn(closer io.Closer) {
	if closer == nil {
		logrus.Error("resource is close")
		return
	}
	err := closer.Close()
	if err != nil {
		logrus.Errorf("failed to close resource, err[%v]", err.Error())
	}
}
