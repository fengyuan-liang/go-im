// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 14:07
// @Author: liangfengyuan@qiniu.com

package repository

import "go-im/pkg/orm"

var groupBasicRepo *GroupBasicRepository

type GroupBasicRepository struct {
	*orm.BaseRepository[GroupBasicRepository]
}

func NewGroupBasicRepository() *GroupBasicRepository {
	return &GroupBasicRepository{}
}

func GetGroupRepo() *GroupBasicRepository {
	if groupBasicRepo == nil {
		groupBasicRepo = NewGroupBasicRepository()
	}
	return groupBasicRepo
}
