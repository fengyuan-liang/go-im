// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 18:15
// @Author: liangfengyuan@qiniu.com

package repository

import (
	"go-im/models"
	"go-im/pkg/orm"
)

var userRepository *UserRepository

type UserRepository struct {
	orm.BaseRepository[models.UserBasic]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func GetUserRepo() *UserRepository {
	if userRepository == nil {
		userRepository = NewUserRepository()
	}
	return userRepository
}
