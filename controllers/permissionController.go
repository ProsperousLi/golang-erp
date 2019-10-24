package controllers

import (
	"encoding/json"

	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type PermissionController struct {
	BaseController
}

func (c *PermissionController) GetPermissions() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}
	pageNum := param["pageNum"]
	pageSize := param["pageSize"]
	if pageNum > 0 {
		pageNum = pageNum - 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	rets := models.GetPermissionBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PermissionController) GetPermissionById() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	beego.Info("id = ", id)
	ret, _ := models.GetPermissionByID(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PermissionController) EditPermissionById() {
	var (
		param models.Permission
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error(err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditPermissionById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PermissionController) AddPermission() {
	var (
		param models.Permission
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error(err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	} else {
		beego.Info(param)
	}
	code := models.AddPermission(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PermissionController) DeletePermission() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	beego.Info(id)
	code := models.DeletePermission(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{cardid: 1, read:[1,2,3], write:[4,5,6]}
func (c *PermissionController) SetPermission() {
	var param models.AddPersionStruct

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}
	code, msg := models.SetPermission(param)
	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//cardid=’xxx’
func (c *PermissionController) QueryPermission() {
	cardid := c.GetString("cardid")
	ret, err := models.QueryPermission(cardid)
	if err != nil {
		beego.Error("%s", err)
		util.RetContent.Code = util.FAILED
		util.RetContent.Message = "未查询到该用户"
	}
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
