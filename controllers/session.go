package controllers

import (
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
// @Success 200 {object} models.SessionResponse
// @Failure 403 body is empty
// @router / [post]
func (this *SessionController) Post() {

	this.Data["json"] = `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`
	this.ServeJSON()

}
