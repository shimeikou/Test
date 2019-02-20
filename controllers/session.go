package controllers

import (
	"ApiTestApp/models"
	"ApiTestApp/service"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// SessionController operations for Session
type SessionController struct {
	beego.Controller
}

// URLMapping ...
func (s *SessionController) URLMapping() {
	s.Mapping("Post", s.Post)
}

// Post ...
// @Title Create
// @Description Make Session
// @Success 200 {object} models.MakeSessionResponse
// @Failure 403 body is empty
// @router / [post]
func (s *SessionController) Post() {
	res := models.MakeSessionResponse{}
	Sessionid, json := res.SetApiResponse()

	conn := service.RedisConnectionPool.Get()
	defer conn.Close()

	//5分間だけ、id -1のセッション
	val, err := conn.Do("SET", Sessionid, models.UndecidedUserID, "NX", "EX", 60*5)

	if err != nil {
		logs.Error(err)
	}

	//toDo: retry create session if it existed
	if val == nil {
		logs.Error("session id is exist!!", Sessionid)
		s.Data["json"] = "failed to create session id "
	} else {
		s.Data["json"] = string(json)
	}

	s.ServeJSON()
}
