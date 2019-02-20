package controllers

import (
	"ApiTestApp/models"

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
	res := models.ServerInfo{}
	json := res.SetApiResponse()
	c.Data["json"] = string(json)
	c.ServeJSON()
}
