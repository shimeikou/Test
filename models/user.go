package models

//SignUpResponse signUpApiのレスポンス
type SignUpResponse struct {
	User
	ResponseTmp
}

//UndecidedUserID ...
const UndecidedUserID = -999

//User ユーザの基本データ
type User struct {
	ID           uint64 `json:"id"`
	ShardID      uint8  `json:"shard_id"`
	UUID         string `json:"uuid"`
	RegisterDate string `json:"register_date"`
	LoginAt      string `json:"login_at"`
}
