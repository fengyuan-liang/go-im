// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/20 16:01
// @Author: liangfengyuan@qiniu.com

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-im/common/entity/response"
	"go-im/models"
	"go-im/repository"
	"go-im/utils"
	"strconv"
)

var (
	validate = validator.New()
)

func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	friends, err := repository.GetContactRepo().SearchFriends(uint64(id))
	if err != nil {
		c.JSON(-1, response.AppErr.WithData(fmt.Sprintf("查询失败，%v", err.Error())))
		return
	}
	c.JSON(200, response.Ok.WithData(friends))
}

func AddFriend(c *gin.Context) {
	var (
		err error
	)
	contact := &models.Contact{}
	if err = c.BindJSON(contact); err != nil {
		c.JSON(-1, response.Err.WithMsg("参数有误，err:"+err.Error()))
	}
	if err = validate.Struct(contact); err != nil {
		c.JSON(-1, response.AppErr.WithMsg(utils.ProcessErr(contact, err)))
		return
	}
	if err = repository.GetContactRepo().AddFriend(contact); err != nil {
		c.JSON(-1, response.AppErr.WithData(err.Error()))
		return
	}
	c.JSON(200, response.Ok)
}
