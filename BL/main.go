package main

import (
	"blog-go/core"
	"blog-go/global"
	"blog-go/initialize"
)

func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectES()
	initialize.OtherInit()
	defer global.Redis.Close()

	initialize.InitCron()

	core.RunServer()
}
