package controllers

import (
	"ApiTestApp/apputil"
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
func (c *SignupController) Post() {
	sessionID := c.Ctx.Input.Header("session-id")

	dbConn := service.RedisConnectionPool.Get()
	defer dbConn.Close()
	sessionInfoBytes, err := redis.Bytes(dbConn.Do("Get", sessionID))

	if sessionInfoBytes == nil || err != nil {
		logs.Error("[signup] get session info failed. maybe session Expired :", sessionID)
		c.Data["json"] = "[signup] session expired!"
		c.ServeJSON()
		return
	}

	sessionInfo := new(models.MakeSessionResponse)
	if err = json.Unmarshal(sessionInfoBytes, sessionInfo); err != nil {
		logs.Error("[signup] unmarshal cache failed!!")
		c.Data["json"] = "[signup] unmarshal failed!"
		c.ServeJSON()
		return
	}

	if sessionInfo.UserId != 0 {
		logs.Error("[signup] session'userID is setted!!")
		c.Data["json"] = "[signup] session'userID is setted!!"
		c.ServeJSON()
		return
	}

	userID, json := SetupSignUpResponse()

	redisConn := service.RedisConnectionPool.Get()
	defer redisConn.Close()
	_, err = redisConn.Do("SET", sessionInfo.SessionId, userID, "EX", 60*60*6)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = string(json)
	c.ServeJSON()
}

//SetupSignUpResponse レスポンスセット
func SetupSignUpResponse() (uint64, []byte) {
	db := service.GetMysqlConnection("user_data")
	defer db.Close()

	trans, err := db.Begin()
	if err != nil {
		logs.Error(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if err := trans.Rollback(); err != nil {
				panic(err.Error())
			}
			logs.Error("Rollbacked")
		}
	}()

	var count uint64
	db.QueryRow(`select count(id) from users`).Scan(&count)

	shardID := (count + 1) % apputil.DataBaseShardMax

	now := service.GetTimeDefault()
	//UUID
	UUIDHash := service.EncodeUUID(count + 1)

	result, err := db.Exec(
		`INSERT INTO users(shard_id,register_date,uuid_hash) VALUES(?,?,?)`,
		shardID,
		now,
		UUIDHash,
	)
	if err != nil {
		panic(err.Error())
	}
	if err = trans.Commit(); err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()
	res := new(models.SignUpResponse)
	res.ID = uint64(id)
	res.RegisterDate = now
	res.UUIDHash = UUIDHash

	res.ResultCode = apputil.ResultCodeSuccess
	res.Time = service.GetTimeRFC3339()

	outputJSON, err := json.Marshal(res)
	return res.ID, outputJSON
}
