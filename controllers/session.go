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
func (this *SessionController) Post() {
	res := models.MakeSessionResponse{}
	Sessionid, json := res.SetApiResponse()
	setCache(Sessionid, json)
	this.Data["json"] = string(json)
	this.ServeJSON()
}

func setCache(key string, sessionResponse []byte) {
	conn := service.RedisConnectionPool.Get()
	defer conn.Close()
	val, err := conn.Do("SET", key, sessionResponse, "NX", "EX", 60*5)
	if val == nil {
		logs.Error("session id is exist!!", key)
		panic(err)
	}
}
