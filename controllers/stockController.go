package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type StockController struct {
	BaseController
}

func (c *StockController) GetStocks() {
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
	rets := models.GetStockBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *StockController) GetStockById() {
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

	warehouseid := param["warehouseid"]

	logs.FileLogs.Info("%v ---warehouseid = ", warehouseid)
	ret, _ := models.GetStockById(warehouseid)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *StockController) EditStockById() {
	var (
		param models.Stock
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditStockById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *StockController) AddStock() {
	var (
		param models.Stock
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
	code := models.AddStock(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *StockController) DeleteStock() {
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

	id := param["warehouseid"]

	logs.FileLogs.Info("%v ---", id)
	code := models.DeleteStock(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
