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

//contractcode=xxx& vehiclecode=xxx,两个参数必带
func (c *RepairitemController) GetRepairitems() {
	contractcode := c.GetString("contractcode")
	vehiclecode := c.GetString("vehiclecode")
	if contractcode == "" || vehiclecode == "" {
		util.RetContent.Code = util.FAILED
		util.RetContent.Message = "参数为空"
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}
	rets := models.GetRepairitemBCode(contractcode, vehiclecode)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *RepairitemController) GetRepairitemById() {
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

func (c *RepairitemController) EditRepairitemById() {
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

func (c *RepairitemController) AddRepairitem() {
	var (
		param models.Repairitem
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.FAILED
		util.RetContent.Message = "参数错误"
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

func (c *RepairitemController) DeleteRepairitem() {
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
	code := models.DeleteRepairitem(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
