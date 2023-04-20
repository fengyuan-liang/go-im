package models

import (
	"gorm.io/gorm"
)

type GroupBasic struct {
	gorm.Model `json:"gorm.Model"`
	Name       string `bson:"name;type:varchar(100)" json:"name"`
	OwnerId    uint   `bson:"ownerId;type:bigint(20)" json:"owner_id"`
	Icon       string `bson:"icon;type:varchar(100)" json:"icon"`
	Type       int    `bson:"type;type:int" json:"type"`
	Desc       string `bson:"desc;type:varchar(200)" json:"desc"`
}

// TableName 用户表名
func (*GroupBasic) TableName() string {
	return "group_basic"
}
