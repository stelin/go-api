package controllers

import (
	"api/models/logic"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @router /get/:key [get]
func (this *UserController) Get() {
	logic := logic.UserLogic{}
	logic.GetUserInfos(1, 2)

	this.Ctx.WriteString("hello beego")
}
