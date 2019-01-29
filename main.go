package main

import (
	RedisUtil "ApiTestApp/Util"
	_ "ApiTestApp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else if beego.BConfig.RunMode == "prod" {
		beego.BConfig.Listen.EnableFcgi = true
		//beego.BConfig.Listen.EnableStdIo = true
	}

	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log","level":7,"maxlines":10000,"maxsize":256,"daily":true,"maxdays":7,"color":true}`)
	logs.EnableFuncCallDepth(true)
	//ogs.SetLogFuncCallDepth(1)
	//logs.SetLogger(logs.AdapterSlack, `{"webhookurl":"https://slack.com/xxx","level":1}`)

	RedisUtil.Init()
	beego.Run()
}
