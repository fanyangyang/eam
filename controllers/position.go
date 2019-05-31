package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fanyangyang/eam/models"
	"strings"
)

type PositionController struct {
	BaseController
}

func (c *PositionController) Get() {
	var positions []models.Position
	pageNum, err := c.GetInt("p")
	if err != nil {
		pageNum = models.DEFAULT_PAGE_NUM
	}
	pageSize, err := c.GetInt("s")
	if err != nil {
		pageSize = models.DEFAULT_PAGE_SIZE
	}

	status, err := c.GetInt("fs")
	query := models.Ormer.QueryTable("position")
	if err == nil {
		query = query.Filter("status", status)
	}
	n, err := query.Limit(pageSize, (pageNum-1)*pageSize).OrderBy("id").All(&positions) //分页查询,需要结合路由进行传参数
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, _ := range positions {
		models.Ormer.LoadRelated(&positions[i],"users")
	}
	c.Data["json"] = &models.PositionRet{
		TotalNum:    n,
		PageNum:     pageNum,
		PageSize:    pageSize,
		Positions: positions,
	}

	c.ServeJSON()
}

func (c *PositionController) Post() {
	var newPosition models.Position
	var err error
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &newPosition); err != nil {
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	}

	newPosition.Status = models.POSITION_STATUS_ON
	if _, err = models.Ormer.Insert(&newPosition); err != nil { //返回的n是表中的数据总数
		resp.Success = false
		resp.Desc = err.Error()
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func (c *PositionController) Put() {
	//修改操作必须携带ID作为主键
	var target models.Position
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse Position error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}
	cols := make([]string, 0)
	if target.Status != 0 {
		cols = append(cols, "status")
	}

	if target.Name != "" {
		cols = append(cols, "name")
	}

	target.Desc = strings.TrimSpace(target.Desc)
	if target.Desc != "" {
		cols = append(cols, "desc")
	}

	if target.Status != 0 {
		cols = append(cols, "status")
	}

	if len(cols) > 0 {
		if _, err := models.Ormer.Update(&target, cols...); err != nil {
			fmt.Println("update Position error: ", err)
			resp.Success = false
			resp.Desc = err.Error()
			c.ServeJSON()
			return
		}

	}
	resp.Success = true

	c.ServeJSON()
}

func (c *PositionController) Delete() {
	//必须携带ID
	var target models.Position
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse Position error")
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

