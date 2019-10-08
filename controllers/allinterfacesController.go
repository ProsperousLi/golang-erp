package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type AllinterfacesController struct {
	BaseController
}

func (c *AllinterfacesController) QueryTimeStamp() {
	queryType := c.GetString("type")
	rets := models.QueryTimeStamp(queryType)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *AllinterfacesController) QueryTimeStampDays() {
	queryType := c.GetString("type")
	rets := models.QueryTimeStampDays(queryType)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//supplierid=XXX
func (c *AllinterfacesController) QueryMattersOfSupplier() {
	supplierid, err := c.GetInt64("supplierid")
	if err != nil {
		util.RetContent.Code = util.FAILED
		c.Data["json"] = util.RetContent
	}

	supliers, err := models.GetSupplyrelationBySupplierid(supplierid)
	if err != nil {
		util.RetContent.Code = util.FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	var matters []models.Matter
	for _, suplier := range supliers {
		var matter models.Matter
		matter, err := models.GetMatterById(suplier.Matterid)
		if err != nil {
			continue
		}

		matters = append(matters, matter)
	}

	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = matters
	c.Data["json"] = util.RetContent
	c.ServeJSON()

}

//matterid=xxx
func (c *AllinterfacesController) QuerySuppliersOfMatter() {
	matterid, err := c.GetInt64("matterid")
	if err != nil {
		util.RetContent.Code = util.FAILED
		c.Data["json"] = util.RetContent
	}

	matters, err := models.GetSupplyrelationBySupplierid(matterid)
	if err != nil {
		util.RetContent.Code = util.FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	var supliers []models.Supplier
	for _, matter := range matters {
		suplier, err := models.GetSupplierById(matter.Supplierid)
		if err != nil {
			continue
		}
		supliers = append(supliers, suplier)
	}
	if err != nil {
		util.RetContent.Code = util.FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
	}

	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = supliers
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

type UpdateMatterStruct struct {
	Supplierid int64
	MatterList []int64
}

//{supplierid: 1, matterList[1, 2, 3]}
//需要先把该供货商的供货信息删除，再插入
func (c *AllinterfacesController) UpdateMatterListOfSupplier() {
	var (
		param UpdateMatterStruct
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	code, msg := models.UpdateRelationBySupplierid(param.Supplierid, param.MatterList)

	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

type UpdateSupplierStruct struct {
	Matterid     int64
	SupplierList []int64
}

func (c *AllinterfacesController) UpdateSupplierListOfMatter() {
	var (
		param UpdateSupplierStruct
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	code, msg := models.UpdateSupplierListOfMatter(param.Matterid, param.SupplierList)

	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

// 3.根据采购单号查询到货情况
// 采购单号查询所有物料列表，入库单查询该采购单号的入库物料数量。
