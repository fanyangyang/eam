package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fanyangyang/eam/models"
	"strings"
)

type ItemController struct {
	BaseController
}


func (c *ItemController) Get(){
	var items []models.Item
	pageNum, err := c.GetInt("p")
	if err != nil {
		pageNum = models.DEFAULT_PAGE_NUM
	}
	pageSize, err := c.GetInt("s")
	if err != nil {
		pageSize = models.DEFAULT_PAGE_SIZE
	}

	query := models.Ormer.QueryTable("item")
	status, err := c.GetInt("fs")
	if err == nil {
		query = query.Filter("status", status)
	}
	n, err := query.Limit(pageSize, (pageNum-1)*pageSize).OrderBy("id").All(&items) //分页查询,需要结合路由进行传参数
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["json"] = &models.ItemRet{
		TotalNum: n,
		PageNum:  pageNum,
		PageSize: pageSize,
		Items:    items,
	}

	c.ServeJSON()
}

func (c *ItemController) Post(){
	var newItem models.Item
	var err error
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &newItem); err != nil {
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	}

	newItem.Status = models.ITEM_STATUS_STORAGE
	if _, err = models.Ormer.Insert(&newItem); err != nil { //返回的n是表中的数据总数
		resp.Success = false
		resp.Desc = err.Error()
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}

func (c *ItemController) Put(){
	//修改操作必须携带ID作为主键
	var target models.Item
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse Item error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}
	cols := make([]string, 0)
	if target.Status != 0 {
		cols = append(cols, "status")
	}
	target.Number = strings.TrimSpace(target.Number)
	if target.Number != ""{
		cols = append(cols,"number")
	}

	if target.User.Id > 0 {
		cols = append(cols, "user_id")
	}

	target.SerialCode = strings.TrimSpace(target.SerialCode)
	if target.SerialCode != ""{
		cols = append(cols,"serial_code")
	}

	target.ShoppingCode = strings.TrimSpace(target.ShoppingCode)
	if target.ShoppingCode != ""{
		cols = append(cols,"shopping_code")
	}

	target.SourcePlateForm = strings.TrimSpace(target.SourcePlateForm)
	if target.SourcePlateForm != ""{
		cols = append(cols,"source_plate_form")
	}

	target.ShopDate = strings.TrimSpace(target.ShopDate)
	if target.ShopDate != ""{
		cols = append(cols,"shop_date")
	}

	target.Desc = strings.TrimSpace(target.Desc)
	if target.Desc != "" {
		cols = append(cols, "desc")
	}

	target.Name = strings.TrimSpace(target.Name)
	if target.Name != ""{
		cols = append(cols,"name")
	}

	if target.IType.Id > 0 {
		cols = append(cols,"i_type_id")
	}

	if len(cols) > 0 {
		if _, err := models.Ormer.Update(&target, cols...); err != nil {
			fmt.Println("update Item error: ", err)
			resp.Success = false
			resp.Desc = err.Error()
			c.ServeJSON()
			return
		}

	}
	resp.Success = true

	c.ServeJSON()
}

func (c *ItemController) Delete(){
	// TODO 子类是否也一并删除，或者给出提示，是否删除子类
	fmt.Println("Item delete")
	//必须携带ID
	var target models.Item
	resp := models.RespSuccess{}
	c.Data["json"] = &resp

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &target); err != nil {
		fmt.Println("parse Item error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()

		return
	}

	_, err := models.Ormer.Delete(&target)
	if err != nil {
		fmt.Println("update Item error")
		resp.Success = false
		resp.Desc = err.Error()
		c.ServeJSON()
		return
	} else {
		resp.Success = true
	}

	c.ServeJSON()
}