package tests

import "testing"
import AppUtil "ApiTestApp/AppUtil"

func TestExampleSuccess(t *testing.T) {
	result := AppUtil.RESULT_CODE_ERROR
	if result != 500 {
		t.Fatalf("failed test %#v", result)
	}
}
