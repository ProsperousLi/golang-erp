package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type RepaircontractController struct {
	BaseController
}

func (c *RepaircontractController) GetRepaircontracts() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
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
	rets := models.GetRepaircontractBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepaircontractController) GetRepaircontractById() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	logs.FileLogs.Info("%v ---id = ", id)
	ret, _ := models.GetRepaircontractById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepaircontractController) EditRepaircontractById() {
	var (
		param models.Repaircontract
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditRepaircontractById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepaircontractController) AddRepaircontract() {
	var (
		param models.Repaircontract
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	} else {
		logs.FileLogs.Info("%v", param)
	}
	code := models.AddRepaircontract(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepaircontractController) DeleteRepaircontract() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	logs.FileLogs.Info("%v ---", id)
	code := models.DeleteRepaircontract(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
