package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fanyangyang/eam/models"
	"strings"
)

type ITypeController struct {
	BaseController
}


func (c *ITypeController) Get(){
	var types []models.IType
	pageNum, err := c.GetInt("p")
	if err != nil {
		pageNum = models.DEFAULT_PAGE_NUM
	}
	pageSize, err := c.GetInt("s")
	if err != nil {
		pageSize = models.DEFAULT_PAGE_SIZE
	}

	query := models.Ormer.QueryTable("i_type")
	status, err := c.GetInt("fs")
	if err == nil {
		query = query.Filter("status", status)
	}
	n, err := query.Limit(pageSize, (pageNum-1)*pageSize).OrderBy("id").All(&types) //分页查询,需要结合路由进行传参数
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["json"] = &models.ITypeRet{
		TotalNum: n,
		PageNum:  pageNum,
		PageSize: pageSize,
		ITypes:    types,
	}

	c.ServeJSON()
}

func (c *ITypeController) Post(){
	var newIType models.IType
	var err error
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &newIType); err != nil {
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	}

	newIType.Status = models.ITYPE_STATUS_ON
	if _, err = models.Ormer.Insert(&newIType); err != nil { //返回的n是表中的数据总数
		resp.Success = false
		resp.Desc = err.Error()
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func (c *ITypeController) Put(){
	//修改操作必须携带ID作为主键
	var target models.IType
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse itype error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}
	cols := make([]string, 0)
	if target.Status != 0 {
		cols = append(cols, "status")
	}

	target.Desc = strings.TrimSpace(target.Desc)
	if target.Desc != "" {
		cols = append(cols, "desc")
	}

	target.Name = strings.TrimSpace(target.Name)
	if target.Name != ""{
		cols = append(cols,"name")
	}

	if target.ParentId != 0 {
		cols = append(cols,"parent_id")
	}


	if len(cols) > 0 {
		if _, err := models.Ormer.Update(&target, cols...); err != nil {
			fmt.Println("update itype error: ", err)
			resp.Success = false
			resp.Desc = err.Error()
			c.ServeJSON()
			return
		}

	}
	resp.Success = true

	c.ServeJSON()
}

func (c *ITypeController) Delete(){
	fmt.Println("itype delete")
	//必须携带ID
	var target models.IType
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse itype error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}

	_, err := models.Ormer.Delete(&target)
	if err != nil {
		fmt.Println("update itype error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}