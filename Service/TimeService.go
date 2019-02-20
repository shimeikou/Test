package service

import "time"

//ExcuteTime ...
var ExcuteTime time.Time

//GetTimeRFC3339 現在時刻のRFC3339を取得
func GetTimeRFC3339() string {
	if ExcuteTime.IsZero() {
		ExcuteTime = time.Now()
	}
	return ExcuteTime.Format(time.RFC3339)
}

//GetTimeDefault 現在時刻のyyyy/mm/dd h:m:sを取得
func GetTimeDefault() string {
	if ExcuteTime.IsZero() {
		ExcuteTime = time.Now()
	}
	return ExcuteTime.Format("2006/01/02 15:04:05")
}
