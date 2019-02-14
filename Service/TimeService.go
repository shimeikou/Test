package service

import "time"

//GetTimeRFC3339 現在時刻のRFC3339を取得
func GetTimeRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

//GetTimeDefault 現在時刻のyyyy/mm/dd h:m:sを取得
func GetTimeDefault() string {
	return time.Now().Format("2006/01/02 15:04:05")
}
