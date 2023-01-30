package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

// @Description: 正则表达式工具类
// @Version: 1.0.0
// @Date: 2023/01/30 21:48
// @Author: fengyuan-liang@foxmail.com

func init() {
	// 注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("RegexPhone", RegexPhone); err != nil {
			fmt.Println("参数校验标签注册失败")
		}
		fmt.Println("参数校验标签注册成功")
	}
}

const (
	PHONE_ROLE = "1[3-9]{1}\\d{9}$"
)

var PHONE_REGEX, _ = regexp.Compile(PHONE_ROLE)

// RegexPhone 手机号码
func RegexPhone(level validator.FieldLevel) bool {
	if phoneNum, ok := level.Field().Interface().(string); ok {
		return PHONE_REGEX.MatchString(phoneNum)
	}
	return false
}

// ProcessErr go validator参数校验器自定义规则及提示
func ProcessErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}
	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() //获取是哪个字段不符合格式
		typeOf := reflect.TypeOf(u)
		// 如果是指针，获取其属性
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		field, ok := typeOf.FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("reg_error_info") // 获取field对应的reg_error_info tag值
			return fieldName + ":" + errorInfo           // 返回错误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}
