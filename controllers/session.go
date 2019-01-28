package controllers

import (
	"ApiTestApp/models"

	"github.com/astaxie/beego"
)

// SessionController operations for Session
type SessionController struct {
	beego.Controller
}

type SessionResponse struct {
	session_id           string
	temporary_common_Key string
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
	//var req models
	//if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
	//	log.Fatal(err)
	//}
	var responseObj models.Session
	responseObj.InitTmpContent()
	this.Data["json"] = responseObj
	this.ServeJSON()
}
