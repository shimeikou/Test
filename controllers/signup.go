package controllers

import (
	"ApiTestApp/models"
	"ApiTestApp/service"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
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
func (this *SignupController) Post() {
	sessionId := this.Ctx.Input.Header("session-id")

	conn := service.RedisConnectionPool.Get()
	defer conn.Close()
	sessionInfoBytes, err := redis.Bytes(conn.Do("Get", sessionId))

	if sessionInfoBytes == nil || err != nil {
		logs.Error("[signup] get session info failed. maybe session Expired :", sessionId)
		panic(err)
	}

	sessionInfo := new(models.MakeSessionResponse)
	if err = json.Unmarshal(sessionInfoBytes, sessionInfo); err != nil {
		logs.Error("[signup] unmarshal json failed!!")
		panic(err)
	}

	logs.Debug(sessionInfo)

	this.Data["json"] = sessionInfo.SessionId
	this.ServeJSON()
}
