package models

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

type Session struct {
	session_id string
	//temporary_common_Key string
}

type SessionResponse struct {
	result_code int
	session     *Session
}

func MakeSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
