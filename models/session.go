package models

import (
	"ApiTestApp/AppUtil"
	SessionService "ApiTestApp/service"
	"encoding/json"
)

//構造体変数名パスカルケースしないとjsonにできない謎仕様なので要注意
type Session struct {
	SessionId string `json:"session_id"`
	//temporary_common_Key string
}

type ResponseTmp struct {
	ResultCode int   `json:"result_code"`
	TtimeStamp int64 `json:"time_stamp"`
}

type MakeSessionResponse struct {
	Session
	ResponseTmp
}

func CreateNewSessionResponse() []byte {
	var sessionResponse = MakeSessionResponse{}
	sessionResponse.ResultCode = AppUtil.RESULT_CODE_SUCCESS
	sessionResponse.TtimeStamp = 0
	sessionResponse.SessionId = SessionService.MakeSessionId()
	if sessionResponse.SessionId == "" {
		// エラー時の処理
		sessionResponse.ResultCode = AppUtil.RESULT_CODE_ERROR
	}
	outputJson, _ := json.Marshal(&sessionResponse)
	return outputJson
}
