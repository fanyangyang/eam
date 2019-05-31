package routers

import (
	"github.com/astaxie/beego"
	"github.com/fanyangyang/eam/controllers"
)

func init() {
	/*
	列出所有员工，分两组：在职和离职两种
	get：select
	 */

	beego.Router("/user", &controllers.UserController{})
	/*
	入职，新增操作，提供员工基础信息，提供单个操作和批量操作（excel）
	post：insert
	 */
	beego.Router("/user/getin", &controllers.UserController{}, "post:MultiInput")
	/*
	离职，修改操作，提供员工姓名，提交前先进行确认，需要先查询出该员工信息进行确认
	delete：update
	 */
	beego.Router("/user/leave", &controllers.UserController{}, "post:Leave")
	/*
	针对资产分类的增删改查
	 */
	beego.Router("/type", &controllers.ITypeController{})
	/*
	针对岗位的增删改查
	 */
	beego.Router("/position", &controllers.PositionController{})
	/*
	针对资产的增删改查
	 */
	beego.Router("/item", &controllers.ItemController{})
	/*
	针对部门的增删改查
	 */
	beego.Router("/department", &controllers.DepartmentController{})
	/*
	TODO
	针对申请的增删改查
	 */
	beego.Router("/apply", &controllers.ApplyController{})

	/*
	针对待办任务的增删改查
	 */
	 beego.Router("/todo",&controllers.TodoController{})
}
