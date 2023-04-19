package vo

import "go-im/models"

type MessageVO struct {
	UserId   uint64             `json:"userId"`
	Type     models.MessageType `json:"type"`
	TargetId uint64             `json:"targetId"`
	Content  string             `json:"content"`
}
