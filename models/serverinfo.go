package models

import (
	"ApiTestApp/AppUtil"
	"ApiTestApp/Service"
	"encoding/json"

	"github.com/astaxie/beego"
)

type ServerInfo struct {
	ServerVersion    string `json:"version"`
	MasterHash       string `json:"master_hash"`
	AssertHash       string `json:"assert_hash"`
	MaintenanceState int    `json:"maintenance_state"`
	ResponseTmp
}

func (this *ServerInfo) SetApiResponse() []byte {
	this.AssertHash = beego.AppConfig.String("AssertHash")
	this.MasterHash = beego.AppConfig.String("MasterHash")
	this.ServerVersion = beego.AppConfig.String("ServerVersion")
	this.MaintenanceState, _ = beego.AppConfig.Int("MaintenanceState!")
	this.ResultCode = AppUtil.RESULT_CODE_SUCCESS
	this.TimeStamp = Service.GetTimeRFC3339()

	outputJson, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return outputJson
}
