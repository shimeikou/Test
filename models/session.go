package models

import (
	"ApiTestApp/apputil"
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
	Time       string `json:"time"`
}

type MakeSessionResponse struct {
	Session
	ResponseTmp
}

func (this *MakeSessionResponse) SetApiResponse() (string, []byte) {
	this.ResultCode = apputil.ResultCodeSuccess
	this.Time = service.GetTimeRFC3339()
	this.TemporaryCommonKey = "abcdefg1234567"
	this.UserId = 0

	this.SessionId = MakeSessionId()
	if this.SessionId == "" {
		// エラーコードを入れる
		this.ResultCode = apputil.ResultCodeError
	}
	outputJson, err := json.Marshal(this)
	if err != nil {
		// エラーコードを入れる
		panic(err)
	}
	return this.SessionId, outputJson
}

func MakeSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
