package Service

import "time"

func GetTimeRFC3339() string {
	return time.Now().Format(time.RFC3339)
}
