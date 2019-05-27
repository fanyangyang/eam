package controllers

import (
	"github.com/astaxie/beego"
	"github.com/fanyangyang/eam/models"
)

type BaseController struct {
	beego.Controller
	login   bool
	Account models.Account
}

func (c *BaseController) Prepare() {
}
