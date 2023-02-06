package utils

import (
	"context"
	"go-im/common/driverUtil"
)

const (
	PublishKey = "websocket"
)

// Publish
//
//	@Description: 发送消息给订阅channel的接受者
//	@args ctx
//	@args channel
//	@args msg
//	@return bizError.BizErrorer
func Publish(ctx context.Context, channel string, msg string) error {
	redis, _ := driverUtil.InitRedis()
	result := redis.Publish(ctx, channel, msg)
	return result.Err()
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	redis, _ := driverUtil.InitRedis()
	subscribe := redis.Subscribe(ctx, channel)
	message, err := subscribe.ReceiveMessage(ctx)
	return message.Payload, err
}
