package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fanyangyang/eam/helper"
	"github.com/fanyangyang/eam/models"
	"strings"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	var users []models.User
	pageNum, err := c.GetInt("p")
	if err != nil {
		pageNum = models.DEFAULT_PAGE_NUM
	}
	pageSize, err := c.GetInt("s")
	if err != nil {
		pageSize = models.DEFAULT_PAGE_SIZE
	}

	status, err := c.GetInt("fs")
	if err != nil {
		status = models.USER_STATUS_ON
	}
	n, err := models.Ormer.QueryTable("user").Filter("status",status).Limit(pageSize, (pageNum-1)*pageSize).OrderBy("id").All(&users) //分页查询,需要结合路由进行传参数
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["json"] = &models.UserRet{
		TotalNum: n,
		PageNum:  pageNum,
		PageSize: pageSize,
		Users:    users,
	}

	c.ServeJSON()
}

//对应增加操作
func (c *UserController) Post() {
	fmt.Println("user post")
	var newOne models.User
	var err error
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &newOne); err != nil {
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	}

	if newOne.DepartmentId == models.DEFAULT_ID {
		// TODO 添加TODO，说明未分配部门
		helper.AddTODO(models.TODO{Title: helper.DEPARTMENT_NOT_SET})
		resp.Message = "" // TODO message 扩展
	}

	if newOne.PositionId == models.DEFAULT_ID {
		// TODO 添加TODO，说明未分配至味
		helper.AddTODO(models.TODO{Title: helper.POSITION_NOT_SET})
		resp.Message = "" // TODO message 扩展，支持多条消息，提示待办事件已经添加到待办事件列表中
	}

	newOne.Status = models.USER_STATUS_ON

	if _, err = models.Ormer.Insert(&newOne); err != nil { //返回的n是表中的数据总数
		resp.Success = false
		resp.Desc = err.Error()
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

//对应修改操作
func (c *UserController) Put() {
	//修改操作必须携带ID作为主键
	fmt.Println("user put")
	var target models.User
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}
	err := update(target)
	if err != nil {
		fmt.Println("update user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func (c *UserController) Delete() {
	fmt.Println("user delete")
	//必须携带ID
	var target models.User
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}

	_, err := models.Ormer.Delete(&target)
	if err != nil {
		fmt.Println("update user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

//批量入职，接收excel表格进行批量倒入操作
//单个入职采用新增接口POST
func (c *UserController) MultiInput() {
	fmt.Println("批量入职")
	// TODO 抽象出post函数，确定入职离职等更新用户的操作
	fmt.Println("user post")
	var newOne models.User
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &newOne); err != nil {
		fmt.Println("parse user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	}

	if newOne.DepartmentId == models.DEFAULT_ID {
		// TODO 添加TODO，说明未分配部门
		helper.AddTODO(models.TODO{Title: helper.DEPARTMENT_NOT_SET})
		resp.Message = "" // TODO message 扩展
	}

	if newOne.PositionId == models.DEFAULT_ID {
		// TODO 添加TODO，说明未分配至味
		helper.AddTODO(models.TODO{Title: helper.POSITION_NOT_SET})
		resp.Message = "" // TODO message 扩展，支持多条消息，提示待办事件已经添加到待办事件列表中
	}

	_, err := models.Ormer.Insert(&newOne) //返回的n是表中的数据总数
	if err != nil {
		resp.Success = false
		resp.Desc = err.Error()
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func (c *UserController) Leave() {
	fmt.Println("离职")
	fmt.Println("user put")
	var target models.User
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}
	err := update(target)
	if err != nil {
		fmt.Println("update user error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func update(target models.User) error {
	cols := make([]string, 0)

	if target.PositionId != models.DEFAULT_ID {
		cols = append(cols, "position_id")
	}

	if target.DepartmentId != models.DEFAULT_ID {
		cols = append(cols, "department_id")
	}

	if target.Status != 0 {
		cols = append(cols, "status")
	}
	target.Name = strings.TrimSpace(target.Name)
	if target.Name != "" {
		cols = append(cols, "name")
	}

	if target.Age != 0 {
		cols = append(cols, "age")
	}
	target.BoardDate = strings.TrimSpace(target.BoardDate)
	if target.BoardDate != "" {
		cols = append(cols, "board_date")
	}

	if target.Sex != 0 {
		cols = append(cols, "sex")
	}

	_, err := models.Ormer.Update(&target, cols...)
	return err
}

/*
TODO 用户相关报表
比如：每月入职数量
	相比上月入职率
 */


