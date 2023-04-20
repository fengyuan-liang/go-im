package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model `json:"gorm.Model"`
	OwnerId    uint64      `bson:"ownerId;type:int" json:"owner_id"`
	TargetId   uint64      `bson:"targetId;type:int" json:"target_id"`
	Type       ContactType `bson:"type;type:int" json:"type"`
	Desc       string      `bson:"desc;type:varchar(200)" json:"desc"`
}

// TableName 用户表名
func (*Contact) TableName() string {
	return "contact"
}

type ContactType uint

const (
	FRIENDS ContactType = iota
	LOVER
)
