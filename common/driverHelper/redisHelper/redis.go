package redisHelper

import (
	"context"
	"fmt"
	"go-im/pkg/common/xredis"
)

const (
	PublishKey = "websocket"
)

// Publish
//
//	@Description: 通过redis发送消息给订阅channel的接受者
//	@args ctx
//	@args channel
//	@args msg
//	@return bizError.BizErrorer
func Publish(ctx context.Context, channel string, msg string) error {
	result := xredis.Cli.Client.Publish(ctx, channel, msg)
	fmt.Println("向redis推送消息：", msg)
	return result.Err()
}

// Subscribe
//
//	@Description: 获取redis中订阅的消息
//	@args ctx
//	@args channel
//	@return string
//	@return error
func Subscribe(ctx context.Context, channel string) (string, error) {
	subscribe := xredis.Cli.Client.Subscribe(ctx, channel)
	message, err := subscribe.ReceiveMessage(ctx)
	return message.Payload, err
}
