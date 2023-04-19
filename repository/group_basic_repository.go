// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 14:07
// @Author: liangfengyuan@qiniu.com

package repository

import "go-im/pkg/orm"

type GroupBasicRepository struct {
	*orm.BaseRepository[GroupBasicRepository]
}

func NewGroupBasicRepository() *GroupBasicRepository {
	return &GroupBasicRepository{}
}
