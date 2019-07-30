package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type SaledetailController struct {
	BaseController
}

func (c *SaledetailController) GetSaledetails() {
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
	rets := models.GetSaledetailBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *SaledetailController) GetSaledetailById() {
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

	id := param["contractid"]

	logs.FileLogs.Info("%v ---id = ", id)
	ret, _ := models.GetSaledetailById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *SaledetailController) EditSaledetailById() {
	var (
		param models.Saledetail
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditSaledetailById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *SaledetailController) AddSaledetail() {
	var (
		param models.Saledetail
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
	code := models.AddSaledetail(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *SaledetailController) DeleteSaledetail() {
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

	id := param["contractid"]

	logs.FileLogs.Info("%v ---", id)
	code := models.DeleteSaledetail(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
