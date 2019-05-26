package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	fmt.Println("user get")
}

func (c *UserController) Post() {
	fmt.Println("user post")
}

func (c *UserController) Entry() {
	fmt.Println("入职")
}

func (c *UserController) Leave() {
	fmt.Println("离职")
}
