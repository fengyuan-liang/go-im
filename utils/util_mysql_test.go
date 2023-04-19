// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 16:39
// @Author: liangfengyuan@qiniu.com

package utils

import (
	"fmt"
	"testing"
)

type TestStruct struct{}

func (t *TestStruct) TableName() string {
	return "TestStruct..."
}

func TestGetTableName(t *testing.T) {
	fmt.Println(GetTableName(&TestStruct{}))
}
