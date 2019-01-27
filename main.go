package main

import (
	_ "ApiTestApp/routers"
	"io"
	"log"
	"os"

	"github.com/astaxie/beego"
)

func main() {

	logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	defer logfile.Close()

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else if beego.BConfig.RunMode == "prod" {
		beego.BConfig.Listen.EnableFcgi = true
		//beego.BConfig.Listen.EnableStdIo = true
	}

	log.Printf("trace: main.go pass.")

	beego.Run()
}
