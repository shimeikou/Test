// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ApiTestApp/controllers"
	"log"
	"os"

	"github.com/astaxie/beego"
)

const (
	apiNameSpace = "/api/"
	user         = "user"
)

func init() {
	/*ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)*/

	logfile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	defer logfile.Close()

	log.Printf("trace: router.go pass.")

	beego.Router(apiNameSpace+user, &controllers.UserController{}, "Get:Get")
	//beego.Router(apiNameSpace+user, &controllers.UserController{}, "Get:Get")

}
