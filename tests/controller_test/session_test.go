package test

import (
	"ApiTestApp/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserController_Handler(t *testing.T) {
	// prepare http.Request and http.ResponseWriter
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/user", nil)

	// execute the target method
	c := controllers.UserController{}
	c.Get()
	res := w.Result()
	defer res.Body.Close()

	testutil.AssertResponse(t, res, http.StatusOK, "./testdata/user_controller/response.golden")
}
