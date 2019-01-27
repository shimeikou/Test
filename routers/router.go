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
	apiNameSpace = "/api/"

	user   = "user"
	object = "object"
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

	beego.Router(apiNameSpace+user, &controllers.UserController{}, "Get:Get")
	beego.Router(apiNameSpace+object, &controllers.ObjectController{}, "Get:GetAll")

}
