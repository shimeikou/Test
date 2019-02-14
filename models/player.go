package models

type SignUpResponse struct {
	Player
	ResponseTmp
	uuid string `json:"uuid"`
	salt string `json:"salt"`
}

type Player struct {
	PlayerId     int64  `json:"user_id"`
	RegisterDate string `json:"register_date"`
}

func (this *SignUpResponse) SetApiResponse() (string, []byte) {

}
