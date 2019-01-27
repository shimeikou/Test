package controllers

import (
	"github.com/astaxie/beego"
)

// SessionController operations for Session
type SessionController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create Session
// @Param	body		body 	models.Session	true		"body for Session content"
// @Success 201 {object} models.Session
// @Failure 403 body is empty
// @router / [post]
func (c *SessionController) Post() {

}
