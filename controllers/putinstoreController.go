package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type PutinstoreController struct {
	BaseController
}

func (c *PutinstoreController) PutinStore() {
	var (
		param models.Putinstore
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
	code, id := models.AddPutinstore(param)
	util.RetContent.Code = code
	util.RetContent.Data = id
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PutinstoreController) GetPutinstores() {
	var (
		param models.QueryPutistoreStruct
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	if param.Warehouseid == 0 || param.Pageno == 0 || param.Pagesize == 0 ||
		(param.Incode == "" && param.Datebegin == "") {
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
	rets, total := models.GetPutinstoreBypage(param)
	util.RetContent.Code = util.SUCESSFUL
	util.PageResults.TotalCount = total
	util.PageResults.Result = rets
	util.RetContent.Data = util.PageResults
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PutinstoreController) GetPutinstoreById() {
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
	ret, _ := models.GetPutinstoreById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PutinstoreController) EditPutinstoreById() {
	var (
		param models.Putinstore
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditPutinstoreById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PutinstoreController) AddPutinstore() {
	var (
		param models.Putinstore
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
	code, id := models.AddPutinstore(param)
	util.RetContent.Code = code
	util.RetContent.Data = id
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PutinstoreController) DeletePutinstore() {
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
	code := models.DeletePutinstore(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
