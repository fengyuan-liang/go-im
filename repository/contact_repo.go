// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/20 16:22
// @Author: liangfengyuan@qiniu.com

package repository

import (
	"go-im/models"
	"go-im/pkg/common/xmysql"
	"go-im/pkg/orm"
)

var contactRepo *ContactRepository

type ContactRepository struct {
	*orm.BaseRepository[models.Contact]
}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

func GetContactRepo() *ContactRepository {
	if groupBasicRepo == nil {
		contactRepo = NewContactRepository()
	}
	return contactRepo
}

func (repo *ContactRepository) AddFriend(contact *models.Contact) error {
	return xmysql.GetDB().Create(contact).Error
}

func (repo *ContactRepository) SearchFriends(ownerId uint64) (friends *[]models.Contact, err error) {
	err = xmysql.GetDB().Where("owner_id = ?", ownerId).Find(friends).Error
	return
}
