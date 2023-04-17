package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model `json:"gorm.Model"`
	OwnerId    uint   `bson:"ownerId;type:int" json:"owner_id"`
	TargetId   uint   `bson:"targetId;type:int" json:"target_id"`
	Type       int    `bson:"type;type:int" json:"type"`
	Desc       string `bson:"desc;type:varchar(200)" json:"desc"`
}

// TableName 用户表名
func (*Contact) TableName() string {
	return "contact"
}
