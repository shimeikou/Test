package tests

import (
	"ApiTestApp/apputil"
	"ApiTestApp/service"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	result := apputil.ResultCodeError
	if result != 500 {
		t.Fatalf("failed test %#v", result)
	}
}
func TestEncrypt(t *testing.T) {
	nonce := service.GCMEncrypt("abcvsdtewrwrgt")
	res2 := service.GCMDecrypt(string(nonce))

	if string(res2) != "abcvsdtewrwrgt" {
		t.Fatalf("failed test %#v", res2)
	}
}
