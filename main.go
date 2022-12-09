package main

import (
	"workflow_http/container"
)

var app *container.Container

func main() {

	// 加载配置文件
	app = container.GetContainer("conf/config.yaml")
	r := SetupRouter()
	err := r.Run(app.Config.HttpConfig.ListenAddress)
	if err != nil {
		panic("server startup failed, err:" + err.Error())
	}
}
