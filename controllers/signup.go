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
	APICommonPrameter
}

// URLMapping ...
func (c *SignupController) URLMapping() {
	c.Mapping("Post", c.Post)
}

//Prepare ...
func (c *SignupController) Prepare() {
	c.StateCode = apputil.ResultCodeSuccess
	c.CheckServerVitality()
}

// Post ...
// @Title Create
// @Description create Signup
// @Param	body		body 	models.Signup	true		"body for Signup content"
// @Success 201 {object} models.Signup
// @Failure 403 body is empty
// @router / [post]
func (c *SignupController) Post() {
	if c.StateCode != apputil.ResultCodeSuccess {
		c.Data["json"] = c.ErrorReturn()
		c.ServeJSON()
		return
	}

	sessionID := c.Ctx.Input.Header("session-id")

	dbConn := service.RedisConnectionPool.Get()
	defer dbConn.Close()
	UserID, err := redis.Int(dbConn.Do("Get", sessionID))

	if err != nil {
		logs.Error("[signup] get session info failed. maybe session Expired :", sessionID)
		c.StateCode = apputil.ResultCodeSessionError
		c.Data["json"] = c.ErrorReturn()
		c.ServeJSON()
		return
	}

	if UserID != models.UndecidedUserID {
		logs.Error("[signup] session'userID is setted!!", sessionID)
		c.StateCode = apputil.ResultCodeSessionError
		c.Data["json"] = c.ErrorReturn()
		c.ServeJSON()
		return
	}

	json := SetupSignUpResponse()
	c.Data["json"] = string(json)

	c.ServeJSON()
}

//SetupSignUpResponse レスポンスセット
func SetupSignUpResponse() []byte {
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
		return []byte(err.Error())
	}
	if err = trans.Commit(); err != nil {
		return []byte(err.Error())
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
	return outputJSON
}
