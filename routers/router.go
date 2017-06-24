package routers

import (
	"api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSInclude(
			&controllers.UserController{},
		),
	)

	beego.AddNamespace(ns)

	//	ns := beego.NewNamespace("/v1",
	//		beego.NSNamespace("/user",
	//			beego.NSInclude(
	//				&controllers.UserController{},
	//			),
	//		),
	//	)

	//	beego.AddNamespace(ns)

	//	beego.Router("/user", &controllers.UserController{})
}
