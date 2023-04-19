// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 10:12
// @Author: liangfengyuan@qiniu.com

package repository

import (
	"go-im/models"
	"go-im/pkg/orm"
)

var messageRepository *MessageRepository

type MessageRepository struct {
	orm.BaseRepository[models.Message]
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func GetMessageRepo() *MessageRepository {
	if messageRepository == nil {
		return NewMessageRepository()
	}
	return messageRepository
}
