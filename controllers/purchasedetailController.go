package controllers

import (
	"encoding/json"

	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type PurchasedetailController struct {
	BaseController
}

func (c *PurchasedetailController) GetPurchasedetails() {
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
	contractcode := param["contractcode"]
	rets := models.GetPurchasedetailBypage(contractcode)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PurchasedetailController) GetPurchasedetailById() {
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

	id := param["contractcode"]

	beego.Info("contractcode = ", id)
	ret, _ := models.GetPurchasedetailById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PurchasedetailController) EditPurchasedetailById() {
	var (
		param models.Purchasedetail
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error(err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditPurchasedetailById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PurchasedetailController) AddPurchasedetail() {
	var (
		param models.Purchasedetail
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
	code := models.AddPurchasedetail(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *PurchasedetailController) DeletePurchasedetail() {
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

	id := param["contractcode"]

	beego.Info(id)
	code := models.DeletePurchasedetail(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
