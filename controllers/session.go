package controllers

import (
	"ApiTestApp/models"

	"github.com/astaxie/beego"
)

// SessionController operations for Session
type SessionController struct {
	beego.Controller
}

// URLMapping ...
func (s *SessionController) URLMapping() {
	s.Mapping("Post", s.Post)
	s.Mapping("Get", s.Get)
}

// @Title Post
// @Description Make Session
// @Success 200 {object} models.MakeSessionResponse
// @Failure 403 body is empty
// @router / [post]
func (this *SessionController) Post() {
	json := models.CreateNewSessionResponse()
	this.Data["json"] = string(json)
	this.ServeJSON()
}
