package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fanyangyang/eam/models"
	"strings"
)

type TodoController struct {
	BaseController
}

func (c *TodoController) Get() {
	var todos []models.TODO
	pageNum, err := c.GetInt("p")
	if err != nil {
		pageNum = models.DEFAULT_PAGE_NUM
	}
	pageSize, err := c.GetInt("s")
	if err != nil {
		pageSize = models.DEFAULT_PAGE_SIZE
	}

	status, err := c.GetInt("fs")
	query := models.Ormer.QueryTable("t_o_d_o")
	if err == nil {
		query = query.Filter("status", status)
	}
	n, err := query.Limit(pageSize, (pageNum-1)*pageSize).OrderBy("id").All(&todos) //分页查询,需要结合路由进行传参数
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["json"] = &models.TODORet{
		TotalNum: n,
		PageNum:  pageNum,
		PageSize: pageSize,
		TODOs:    todos,
	}

	c.ServeJSON()
}

func (c *TodoController) Post() {
	var newTODO models.TODO
	var err error
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &newTODO); err != nil {
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	}

	newTODO.Status = models.TODO_STATUS_NEW
	if _, err = models.Ormer.Insert(&newTODO); err != nil { //返回的n是表中的数据总数
		resp.Success = false
		resp.Desc = err.Error()
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func (c *TodoController) Put() {
	//修改操作必须携带ID作为主键
	var target models.TODO
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse todo error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}
	cols := make([]string, 0)
	if target.Status != 0 {
		cols = append(cols, "status")
	}
	if target.Level != 0 {
		cols = append(cols, "level")
	}

	target.Desc = strings.TrimSpace(target.Desc)
	if target.Desc != "" {
		cols = append(cols, "desc")
	}

	target.Body = strings.TrimSpace(target.Body)
	if target.Body != "" {
		cols = append(cols, "body")
	}

	target.DeadLine = strings.TrimSpace(target.DeadLine)
	if target.DeadLine != "" {
		cols = append(cols, "dead_line")
	}

	target.Title = strings.TrimSpace(target.Title)
	if target.Title != "" {
		cols = append(cols, "title")
	}

	if len(cols) > 0 {
		if _, err := models.Ormer.Update(&target, cols...); err != nil {
			fmt.Println("update todo error: ", err)
			resp.Success = false
			resp.Desc = err.Error()
			c.ServeJSON()
			return
		}

	}
	resp.Success = true

	c.ServeJSON()
}

func (c *TodoController) Delete(){
	//必须携带ID
	var target models.TODO
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse todo error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}

	_, err := models.Ormer.Delete(&target)
	if err != nil {
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}
