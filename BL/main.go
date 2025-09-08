package main

import (
	"blog-go/core"
	"blog-go/global"
)

func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()

	core.RunServer()
}
