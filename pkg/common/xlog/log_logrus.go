// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/20 15:52
// @Author: liangfengyuan@qiniu.com

package xlog

import "github.com/sirupsen/logrus"

var logger = logrus.New()

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
