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

	"github.com/astaxie/beego"
)

const (
	NS = "/api"

	SERVER_INFO = "/getserverinfo"
	USER        = "/user"
	OBJECT      = "/object"
	SESSION     = "/makesession"
)

func init() {
	ns := beego.NewNamespace(NS,
		beego.NSNamespace(OBJECT,
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace(USER,
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace(SESSION,
			beego.NSInclude(
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace(SERVER_INFO,
			beego.NSInclude(
				&controllers.ServerInfoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
