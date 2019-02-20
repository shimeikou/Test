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

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

//APICommonMethod ...
type APICommonMethod struct {
	userID uint64
}

//GetUserIDFromSessionID ...
func (c *APICommonMethod) GetUserIDFromSessionID(sessionID string) interface{} {
	dbConn := service.RedisConnectionPool.Get()
	defer dbConn.Close()
	userID, err := redis.Int64(dbConn.Do("Get", sessionID))

	if err != nil {
		logs.Error(err)
		return -1
	}
	return userID
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
}

//LoginRequest ...
type LoginRequest struct {
	ID   uint64 `json:"id"`
	UUID string `json:"uuid"`
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
	sessionID := c.Ctx.Input.Header("session-id")

	dbConn := service.RedisConnectionPool.Get()
	defer dbConn.Close()
	_, err := redis.Int(dbConn.Do("Get", sessionID))

	if err != nil {
		logs.Error(err)
		c.Data["json"] = "session expired!"
		c.ServeJSON()
		return
	}

	jsonBytes := c.Ctx.Input.RequestBody
	request := new(LoginRequest)
	err = json.Unmarshal(jsonBytes, request)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	db := service.GetMysqlConnection(service.UserEntryDataBaseName)
	defer db.Close()

	var UUIDHash string
	err = db.QueryRow("SELECT uuid_hash FROM users WHERE id = ?", request.ID).Scan(&UUIDHash)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if err = service.VerifyUUID(request.UUID, UUIDHash); err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error
		c.ServeJSON()
		return
	}

	//パス認証終わったのでトランザクションでデータ更新
	tx, err := db.Begin()
	if err != nil {
		logs.Error(err)
	}
	defer func() {
		if err := recover(); err != nil {
			if err := tx.Rollback(); err != nil {
				panic(err.Error())
			}
			logs.Error("Rollbacked")
		}
	}()

	_, err = tx.Exec(
		`UPDATE users SET login_at = ? WHERE id = ?  `,
		service.GetTimeDefault(),
		request.ID,
	)
	if err != nil {
		logs.Error(err)
	}

	if err = tx.Commit(); err != nil {
		logs.Error(err)
	}

	var res models.ResponseTmp
	res.ResultCode = apputil.ResultCodeSuccess
	res.Time = service.GetTimeRFC3339()

	outputJSON, err := json.Marshal(res)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = string(outputJSON)
	}
	c.ServeJSON()
}
