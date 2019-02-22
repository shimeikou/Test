package controllers

import (
	"ApiTestApp/apputil"
	"ApiTestApp/models"
	"ApiTestApp/service"
	"encoding/json"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

//APICommonPrameter ...
type APICommonPrameter struct {
	userID    uint64
	StateCode uint16
}

//GetUserIDFromSessionID ...
func (c *APICommonPrameter) GetUserIDFromSessionID(sessionID string) {
	dbConn := service.RedisConnectionPool.Get()
	defer dbConn.Close()
	ID, err := redis.Int64(dbConn.Do("Get", sessionID))

	if err != nil {
		logs.Error(err)
		c.StateCode = apputil.ResultCodeSessionError
	}
	if ID <= 0 {
		logs.Error("this session hasn't user ID! :", sessionID)
		c.StateCode = apputil.ResultCodeCantGetIDFromSession
	}
	c.userID = uint64(ID)
}

//CheckUserVitality ...
func (c *APICommonPrameter) CheckUserVitality() {
	/*if userId == banId {
		c.StateCode = ResultCodeBANUser
	}
	*/
}

//CheckServerVitality サーバ側状態確認(requestHash & masterHashの確認も含む)
func (c *APICommonPrameter) CheckServerVitality() {
	db := service.GetMysqlConnection(service.MasterDataBaseName)
	defer db.Close()
	var state int8
	db.QueryRow("select state from mantenance_infos").Scan(&state)
	if state != apputil.ServerStateOnline {
		c.StateCode = apputil.ResultCodeMantenance
	}
}

//ErrorReturn エラー発生により、エラーのレスポンスを作成
func (c *APICommonPrameter) ErrorReturn() string {

	if c.StateCode == apputil.ResultCodeSuccess {
		logs.Warning("ResultCodeSuccess can't ErrorReturn")
		return ""
	}

	resStruct := models.ResponseTmp{}
	resStruct.ResultCode = c.StateCode
	resStruct.Time = service.GetTimeRFC3339()

	json, err := json.Marshal(resStruct)
	if err != nil {
		logs.Error("json Marshal failed in ErrorReturn!")
		return ""
	}
	return string(json)
}
