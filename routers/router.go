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
	//NS NameSpace
	NS = "/api"

	//ServerInfo サーバ状態取得API
	ServerInfo = "/getserverinfo"
	//Session ...
	Session = "/makesession"

	//SignUp ...
	SignUp = "/signup"

	//Login ...
	Login = "/login"
)

func init() {
	ns := beego.NewNamespace(NS,
		beego.NSNamespace(Session,
			beego.NSInclude(
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace(ServerInfo,
			beego.NSInclude(
				&controllers.ServerInfoController{},
			),
		),
		beego.NSNamespace(SignUp,
			beego.NSInclude(
				&controllers.SignupController{},
			),
		),
		beego.NSNamespace(Login,
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
