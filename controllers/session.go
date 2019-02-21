package controllers

import (
	"ApiTestApp/apputil"
	"ApiTestApp/models"
	"ApiTestApp/service"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// SessionController operations for Session
type SessionController struct {
	beego.Controller
	APICommonPrameter
}

// URLMapping ...
func (s *SessionController) URLMapping() {
	s.Mapping("Post", s.Post)
}

//Prepare ...
func (s *SessionController) Prepare() {
	s.StateCode = apputil.ResultCodeSuccess
	s.CheckServerVitality()
}

// Post ...
// @Title Create
// @Description Make Session
// @Success 200 {object} models.MakeSessionResponse
// @Failure 403 body is empty
// @router / [post]
func (s *SessionController) Post() {
	if s.StateCode != apputil.ResultCodeSuccess {
		s.Data["json"] = s.ErrorReturn()
		s.ServeJSON()
		return
	}

	res := models.MakeSessionResponse{}
	Sessionid, json := res.SetAPIResponse()

	//redis connection取得
	conn := service.RedisConnectionPool.Get()
	defer conn.Close()

	//5分間だけ、id -1のセッション
	val, err := conn.Do("SET", Sessionid, models.UndecidedUserID, "NX", "EX", 60*5)

	if err != nil {
		logs.Error(err)
		s.StateCode = apputil.ResultCodeRedisError
		s.Data["json"] = s.ErrorReturn()
		s.ServeJSON()
		return
	}

	//toDo: retry create session if it existed
	if val == nil {
		logs.Error("session id is exist!!", Sessionid)
		s.Data["json"] = "failed to create session id "
		s.StateCode = apputil.ResultCodeRedisError
		s.Data["json"] = s.ErrorReturn()
		s.ServeJSON()
		return
	}

	if s.StateCode == apputil.ResultCodeSuccess {
		s.Data["json"] = string(json)
	}

	s.ServeJSON()
}
