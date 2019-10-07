package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type ArrivalbillController struct {
	BaseController
}

func (c *ArrivalbillController) QueryArrivalBill() {
	var (
		param models.QueryArrivalBillStruct
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	pageNum := param.Pageno
	pageSize := param.Pagesize
	if pageNum > 0 {
		param.Pageno = pageNum - 1
	}
	if pageSize == 0 {
		param.Pagesize = 10
	}
	rets, total := models.QueryArrivalBill(param)
	util.RetContent.Code = util.SUCESSFUL
	util.PageResults.TotalCount = total
	util.PageResults.Result = rets
	util.RetContent.Data = util.PageResults
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *ArrivalbillController) GetArrivalbills() {
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
	arrivalbillcode := param["arrivalbillcode"]
	rets := models.GetArrivalbillBypage(arrivalbillcode)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *ArrivalbillController) GetArrivalbillById() {
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
	ret, _ := models.GetArrivalbillByUserID(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *ArrivalbillController) EditArrivalbillById() {
	var (
		param models.Arrivalbill
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditArrivalbillById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *ArrivalbillController) AddArrivalbill() {
	var (
		param models.Arrivalbill
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
	code, id := models.AddArrivalbill(param)
	util.RetContent.Code = code
	util.RetContent.Data = id
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *ArrivalbillController) DeleteArrivalbill() {
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
	code := models.DeleteArrivalbill(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
