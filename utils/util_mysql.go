// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 16:25
// @Author: liangfengyuan@qiniu.com

package utils

import (
	"fmt"
	"gorm.io/gorm/schema"
	"reflect"
)

func GetTableName(dest interface{}) (tableName string) {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr && value.IsNil() {
		value = reflect.New(value.Type().Elem())
	}
	modelType := reflect.Indirect(value).Type()
	if modelType.Kind() == reflect.Interface {
		modelType = reflect.Indirect(reflect.ValueOf(dest)).Elem().Type()
	}
	// 获取到值为止
	for modelType.Kind() == reflect.Slice || modelType.Kind() == reflect.Array || modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	modelValue := reflect.New(modelType)
	if tabler, ok := modelValue.Interface().(schema.Tabler); ok {
		tableName = tabler.TableName()
		fmt.Println(tabler)
	}
	return
}
