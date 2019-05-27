package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fanyangyang/eam/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	fmt.Println("user get")
	var users []models.User
	//分页：o.QueryTable("user").Limit(pageSize,(ps-1)*pageSize).All(user)
	n, err := models.Ormer.QueryTable("user").All(&users) //查询全部
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("n is : ", n)
}

//对应增加操作
func (c *UserController) Post() {
	fmt.Println("user post")
	var newOne models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &newOne); err != nil {
		fmt.Println("parse user error")
	}
	fmt.Printf("newOne: %#v\n", newOne)
}

//对应修改操作
func (c *UserController) Put() {
	fmt.Println("user put")
}

func (c *UserController) Delete() {
	fmt.Println("user delete")
}

func (c *UserController) GetIn() {
	fmt.Println("入职")
}

func (c *UserController) Leave() {
	fmt.Println("离职")
}
