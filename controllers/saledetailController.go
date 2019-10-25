package controllers

import (
	"encoding/json"

	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type SaledetailController struct {
	BaseController
}

// util.RetContent.Code = util.FAILED
// util.RetContent.Message = "参数错误"
// c.Data["json"] = util.RetContent
// c.ServeJSON()
// return

func (c *SaledetailController) AddOrUpdateSaleDetail() {
	var (
		param models.AddAndUpdateSaledetailStruct
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.FAILED
		util.RetContent.Message = "参数错误"
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	if param.Contractcode == "" {
		util.RetContent.Code = util.FAILED
		util.RetContent.Message = "参数错误"
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	code, msg := models.AddOrUpdateSaleDetail(param)

	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *SaledetailController) GetSaledetails() {
	contractcode := c.GetString("contractcode")
	if contractcode == "" {
		util.RetContent.Code = util.SUCESSFUL
		util.RetContent.Message = "contractcode is null"
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}
	rets := models.GetSaledetailByContractcode(contractcode)
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
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["contractid"]

	beego.Info("id = ", id)
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
		beego.Error(err)
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
		beego.Error(err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	} else {
		beego.Info(param)
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
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["contractid"]

	beego.Info(id)
	code := models.DeleteSaledetail(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
