package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

// @Description:
// @Version: 1.0.0
// @Date: 2023/01/27 22:06
// @Author: fengyuan-liang@foxmail.com

func GetIndex(c *gin.Context) {
	files, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	files.Execute(c.Writer, "index")
}
