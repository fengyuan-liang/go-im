// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/20 10:49
// @Author: liangfengyuan@qiniu.com

package orm

import (
	"fmt"
	"go-im/models"
	"reflect"
	"testing"
)

func TestMysql(t *testing.T) {
	userBasic := models.UserBasic{}
	v := reflect.ValueOf(&userBasic).FieldByName("ID")
	fmt.Println(v)
}
