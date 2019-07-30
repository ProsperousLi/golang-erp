package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type RepairitemController struct {
	BaseController
}

func (c *RepairitemController) GetRepaircosts() {
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
	rets := models.GetRepairitemBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepairitemController) GetRepaircostById() {
	var (
		param = make(map[string]string)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	contractcode := param["contractcode"]

	logs.FileLogs.Info("%v ---contractcode = ", contractcode)
	ret, _ := models.GetRepairitemById(contractcode)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepairitemController) EditRepaircostById() {
	var (
		param models.Repairitem
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditRepairitemById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepairitemController) AddRepaircost() {
	var (
		param models.Repairitem
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
	code := models.AddRepairitem(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepairitemController) DeleteRepaircost() {
	var (
		param = make(map[string]string)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["contractcode"]

	logs.FileLogs.Info("%v ---", id)
	code := models.DeleteRepairitem(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
