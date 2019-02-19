package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

//CryptSecretKey 暗号化キー(都度生成した方がいいけど)
const CryptSecretKey = "AES256Key-32Characters1234567890"
const NonceLength = 12

//GCMEncrypt AES GCMで暗号化
func GCMEncrypt(plainText string) []byte {
	key := []byte(CryptSecretKey)
	plaintextBytes := []byte(plainText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, NonceLength)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	fmt.Printf("AAA %x\n", nonce)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintextBytes, nil)
	return append(nonce, ciphertext...)
}

//GCMDecrypt AES+GCMで復号
func GCMDecrypt(Raw string) []byte {
	key := []byte(CryptSecretKey)

	rawBytes := []byte(Raw)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintextBytes, err := aesgcm.Open(nil, rawBytes[:NonceLength], rawBytes[NonceLength:], nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintextBytes
}
