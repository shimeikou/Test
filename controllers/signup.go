package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// SignupController operations for Signup
type SignupController struct {
	beego.Controller
}

// URLMapping ...
func (c *SignupController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Signup
// @Param	body		body 	models.Signup	true		"body for Signup content"
// @Success 201 {object} models.Signup
// @Failure 403 body is empty
// @router / [post]
func (c *SignupController) Post() {
	h := c.Ctx.Input.Header("session-id")
	logs.Debug(h)
}
