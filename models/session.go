package models

import (
	"ApiTestApp/appUtil"
	"ApiTestApp/service"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
)

//構造体変数名パスカルケースしないとjsonにできない謎仕様なので要注意
type Session struct {
	SessionId          string `json:"session_id"`
	TemporaryCommonKey string `json:"temporary_common_key"`
	UserId             uint64 `json:"user_id"`
}

type ResponseTmp struct {
	ResultCode uint16 `json:"result_code"`
	TimeStamp  string `json:"time_stamp"`
}

type MakeSessionResponse struct {
	Session
	ResponseTmp
}

func (this *MakeSessionResponse) SetApiResponse() []byte {
	this.ResultCode = appUtil.RESULT_CODE_SUCCESS
	this.TimeStamp = service.GetTimeRFC3339()
	this.TemporaryCommonKey = ""
	this.SessionId = MakeSessionId()
	if this.SessionId == "" {
		// エラーコードを入れる
		this.ResultCode = appUtil.RESULT_CODE_ERROR
	}
	outputJson, err := json.Marshal(this)
	if err != nil {
		// エラーコードを入れる
		panic(err)
	}
	return outputJson
}

func MakeSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
