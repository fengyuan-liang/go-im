// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/04/19 15:13
// @Author: liangfengyuan@qiniu.com

package xgin

import (
	"fmt"
	"go-im/config"
	"go-im/router"
	"go-im/utils"
)

type GinServer struct{}

func (g GinServer) Initialize() error {
	fmt.Println("Initialize callback...")
	return nil
}

func (g GinServer) RunLoop() {
	router.Router().Run(":" + utils.ParseString(config.GetConfig().Server.PORT))
}

func (g GinServer) Destroy() {
	fmt.Println("Destroy callback...")
}

func NewGinServer() *GinServer {
	return &GinServer{}
}
