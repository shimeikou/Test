package main

import (
	"ApiTestApp/Sevice"
	_ "ApiTestApp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	Init()
	beego.Run()
}

func Init() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		logs.SetLogger(logs.AdapterConsole)
	} else if beego.BConfig.RunMode == "prod" {
		beego.BConfig.Listen.EnableFcgi = true
		logs.SetLogger(logs.AdapterFile, `{"filename":"test.log","level":7,"maxlines":10000,"maxsize":256,"daily":true,"maxdays":7,"color":true}`)
		logs.EnableFuncCallDepth(true)
	}

	//logs.SetLogFuncCallDepth(1)
	//logs.SetLogger(logs.AdapterSlack, `{"webhookurl":"https://slack.com/xxx","level":1}`)

	Sevice.InitRedisPool()

	/*res := Service.RedisGet("123")
	logs.Debug(res)
	*/
}
