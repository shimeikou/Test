package service

import (
	"github.com/astaxie/beego/logs"
	hashids "github.com/speps/go-hashids"
	"golang.org/x/crypto/bcrypt"
)

//UUIDSalt uuid生成salt
const UUIDSalt = "May the Force be with you."

//UUIDMinLength uuid最短長さ
const UUIDMinLength = 32

//ExtraNum uuid生成時に足すエクストラ数値
const ExtraNum = 123456789

//EncodeUUID idを元にuuid作成
func EncodeUUID(id uint64) string {
	hd := hashids.NewData()
	hd.Salt = UUIDSalt
	hd.MinLength = UUIDMinLength
	h, _ := hashids.NewWithData(hd)

	uuid, _ := h.EncodeInt64([]int64{int64(id), ExtraNum})
	return uuid
}

//DecodeUUID UUIDからユーザid算出する
func DecodeUUID(UUID string) int {
	hd := hashids.NewData()
	hd.Salt = UUIDSalt
	hd.MinLength = UUIDMinLength
	h, _ := hashids.NewWithData(hd)

	numbers, _ := h.DecodeWithError(UUID)
	return numbers[0]
}

//UUIDToHash 生成したUUIDをpwと見立てて、そこからさらにDB保存用Hashを生成
func UUIDToHash(UUID string) string {
	//まあ本番運用ないのでmincostで生成するわ
	hash, err := bcrypt.GenerateFromPassword([]byte(UUID), bcrypt.MinCost)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return string(hash)
}

//VerifyUUID 生成したUUIDをpwと見立てて、そこからさらにDB保存用Hashを生成
func VerifyUUID(UUID string, UUIDHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(UUIDHash), []byte(UUID))
}
