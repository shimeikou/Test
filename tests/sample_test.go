package tests

import (
	"ApiTestApp/apputil"
	"ApiTestApp/service"
	"fmt"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	result := apputil.ResultCodeError
	if result != 500 {
		t.Fatalf("failed test %#v", result)
	}
}
func TestCrypt(t *testing.T) {
	answer := "abcvsdtewrwrgt"
	nonce := service.GCMEncrypt(answer)
	res2 := service.GCMDecrypt(string(nonce))

	if string(res2) != answer {
		t.Fatalf("failed test %#v", res2)
	}
}

func TestUUID(t *testing.T) {
	UUID := service.EncodeUUID(2)
	fmt.Print(UUID)
	UUIDHash := service.UUIDToHash(UUID)
	fmt.Print("\n")
	fmt.Print(UUIDHash)
	err := service.VerifyUUID(UUID, UUIDHash)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
