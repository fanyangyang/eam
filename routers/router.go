package routers

import (
	"github.com/astaxie/beego"
	"github.com/fanyangyang/eam/controllers"
)

func init() {
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/entry", &controllers.UserController{}, "post:Entry")
	beego.Router("/user/leave", &controllers.UserController{}, "post:Leave")
	beego.Router("/type", &controllers.ITypeController{})
	beego.Router("/position", &controllers.PositioinController{})
	beego.Router("/item", &controllers.ItemController{})
	beego.Router("/department", &controllers.DepartmentController{})
	beego.Router("asks", &controllers.AskListController{})
}
