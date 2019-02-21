package models

import (
	"ApiTestApp/apputil"
	"ApiTestApp/service"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
)

//Session 構造体変数名パスカルケースしないとjsonにできない謎仕様なので要注意
type Session struct {
	SessionID          string `json:"session_id"`
	TemporaryCommonKey string `json:"temporary_common_key"`
	UserID             uint64 `json:"user_id"`
}

//ResponseTmp ...
type ResponseTmp struct {
	ResultCode uint16 `json:"result_code"`
	Time       string `json:"time"`
}

//MakeSessionResponse ...
type MakeSessionResponse struct {
	Session
	ResponseTmp
}

//SetAPIResponse ...
func (c *MakeSessionResponse) SetAPIResponse() (string, []byte) {
	c.ResultCode = apputil.ResultCodeSuccess
	c.Time = service.GetTimeRFC3339()
	c.TemporaryCommonKey = "abcdefg1234567"
	c.UserID = 0

	c.SessionID = MakeSessionID()
	if c.SessionID == "" {
		// エラーコードを入れる
		c.ResultCode = apputil.ResultCodeError
	}
	outputJSON, err := json.Marshal(c)
	if err != nil {
		// エラーコードを入れる
		panic(err)
	}
	return c.SessionID, outputJSON
}

//MakeSessionID ...
func MakeSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
