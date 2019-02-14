package models

//SignUpResponse はsignUpApiのレスポンス
type SignUpResponse struct {
	Player
	ResponseTmp
}

//Player ユーザの基本データ
type Player struct {
	ID           uint64 `json:"id"`
	ShardID      uint8  `json:"shard_id"`
	RegisterDate string `json:"register_date"`
	UUIDHash     string `json:"uuid_hash"`
}
