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
	UserID, err := redis.Int(dbConn.Do("Get", sessionID))

	if err != nil {
		logs.Error("[signup] get session info failed. maybe session Expired :", sessionID)
		c.Data["json"] = "[signup] session expired!"
		c.ServeJSON()
		return
	}

	if UserID != models.UndecidedUserID {
		logs.Error("[signup] session'userID is setted!!")
		c.Data["json"] = "[signup] session'userID is setted!!"
		c.ServeJSON()
		return
	}

	newUserID, json := SetupSignUpResponse()

	redisConn := service.RedisConnectionPool.Get()
	defer redisConn.Close()

	//セッションIDに正式のuserIDを付けて、さらに5分の命を与える
	_, err = redisConn.Do("SET", sessionID, newUserID, "EX", 60*5)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = string(json)
	}
	c.ServeJSON()
}

//SetupSignUpResponse レスポンスセット
func SetupSignUpResponse() (uint64, []byte) {
	db := service.GetMysqlConnection(service.UserEntryDataBaseName)
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
	trans.QueryRow(`select count(id) from users`).Scan(&count)

	shardID := (count + 1) % service.DataBaseShardMax

	now := service.GetTimeDefault()
	//UUID
	UUID := service.EncodeUUID(count + 1)
	UUIDHash := service.UUIDToHash(UUID)

	result, err := trans.Exec(
		`INSERT INTO users(shard_id,uuid_hash,register_date,login_at) VALUES(?,?,?,?)`,
		shardID,
		UUIDHash,
		now,
		now,
	)
	if err != nil {
		return 0, []byte(err.Error())
	}
	if err = trans.Commit(); err != nil {
		return 0, []byte(err.Error())
	}

	id, _ := result.LastInsertId()
	res := new(models.SignUpResponse)
	res.ID = uint64(id)
	res.RegisterDate = now
	res.LoginAt = now
	res.UUID = UUID

	res.ResultCode = apputil.ResultCodeSuccess
	res.Time = service.GetTimeRFC3339()

	outputJSON, err := json.Marshal(res)
	return res.ID, outputJSON
}
