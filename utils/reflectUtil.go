package utils

import (
	"reflect"
)

// @Description: 反射工具类
// @Version: 1.0.0
// @Date: 2023/01/29 14:31
// @Author: fengyuan-liang@foxmail.com

// CopyStruct 拷贝两个结构体内相同的属性
// dst 目标结构体，src 源结构体
// 必须传入指针，且不能为nil
// 它会把src与dst的相同字段名的值，复制到dst中
// Deprecated: 有bug，请使用 utils.deepcopier 进行深拷贝
func CopyStruct(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()
	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		srcName := srcValue.Type().Field(i).Name
		dstFieldByName := dstValue.FieldByName(srcName)
		if dstFieldByName.IsValid() {
			switch dstFieldByName.Kind() {
			case reflect.Ptr:
				switch srcField.Kind() {
				case reflect.Ptr:
					if srcField.IsNil() {
						dstFieldByName.Set(reflect.New(dstFieldByName.Type().Elem()))
					} else {
						dstFieldByName.Set(srcField)
					}
				default:
					dstFieldByName.Set(srcField.Addr())
				}
			default:
				switch srcField.Kind() {
				case reflect.Ptr:
					if srcField.IsNil() {
						dstFieldByName.Set(reflect.Zero(dstFieldByName.Type()))
					} else {
						dstFieldByName.Set(srcField.Elem())
					}
				default:
					dstFieldByName.Set(srcField)
				}
			}
		}
	}
}

// GetFieldStringType 反射获取结构体属性和其类型
// 注意这里返回的一般是大写，比较请使用`strings.EqualFold(x1, x2)`忽略大小写进行比较
func GetFieldStringType(bean *struct{}) map[string]string {
	typ := reflect.TypeOf(bean)
	m := make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		structFieldType := typ.Field(i)
		m[structFieldType.Name] = structFieldType.Type.String()
	}
	return m
}
