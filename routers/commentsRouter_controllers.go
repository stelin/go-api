package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
