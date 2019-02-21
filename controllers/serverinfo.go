package controllers

import (
	"ApiTestApp/apputil"
	"ApiTestApp/models"
	"ApiTestApp/service"
	"encoding/json"

	"github.com/astaxie/beego"
)

//ServerInfoController operations for ServerInfo
type ServerInfoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ServerInfoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create ServerInfo
// @Param	body		body 	models.ServerInfo	true		"body for ServerInfo content"
// @Success 201 {int} models.ServerInfo
// @Failure 403 body is empty
// @router / [post]
func (c *ServerInfoController) Post() {
	json := SetAPIResponse()
	c.Data["json"] = string(json)
	c.ServeJSON()
}

//SetAPIResponse ...
func SetAPIResponse() []byte {

	res := models.ServerInfo{}
	res.AssertHash = beego.AppConfig.String("AssertHash")
	res.MasterHash = beego.AppConfig.String("MasterHash")
	res.ServerVersion = beego.AppConfig.String("ServerVersion")
	res.MaintenanceState, _ = beego.AppConfig.Int("MaintenanceState!")
	res.ResultCode = apputil.ResultCodeSuccess
	res.Time = service.GetTimeRFC3339()

	outputJSON, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return outputJSON
}
