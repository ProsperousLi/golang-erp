package controllers

import (
	"encoding/json"

	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type MarketcontractController struct {
	BaseController
}

//type=xxx&contractcode=xxx&custcode=xxx&handler=xxx&execstatus=1&pageno=1&pagesize=10
func (c *MarketcontractController) GetMarketcontracts() {
	var (
		marketType, execstatus, contractcode, custcode, handler string
		pageno, pagesize                                        int64
	)

	marketType = c.GetString("marketType")
	execstatus = c.GetString("execstatus")
	contractcode = c.GetString("contractcode")
	custcode = c.GetString("custcode")
	handler = c.GetString("handler")
	pageno, _ = c.GetInt64("pageno")
	pagesize, _ = c.GetInt64("pagesize")
	if pageno > 0 {
		pageno = pageno - 1
	}
	if pagesize == 0 {
		pagesize = 10
	}
	rets, total := models.GetMarketcontractBypage(marketType, execstatus,
		contractcode, custcode, handler, pageno, pagesize)
	util.RetContent.Code = util.SUCESSFUL
	util.PageResults.TotalCount = total
	util.PageResults.Result = rets
	util.RetContent.Data = util.PageResults
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MarketcontractController) GetMarketcontractById() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	id := param["id"]

	beego.Info("id = ", id)
	ret, _ := models.GetMarketcontractById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MarketcontractController) EditMarketcontractById() {
	var (
		param models.Marketcontract
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error(err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditMarketcontractById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{"cardid":"123","name":"小李","sex":0,"compID":10,"deptID":1,"dutyID":1,"health":"very good","height":"170","nativeplace":"安徽省","nation":"汉","maritalstatus":"未婚","education":"本科","university":"大连东软","major":"计算  机","qualification":"666","trialsalary":2000,"salary":6000,"idnumber":"341182","address1":"1111","postcode1":"www111","address2":"2222","postcode2":"www222","contactnumber":"17615002988","phonenumber":"110","email":"www666","emergencycontact":"1323654222","c  ontactnumber1":"1323654222","address3":"33333","trialexpired":"2019-05-28 15:03:03","entrydate":"2019-05-28   15:03:03","birthday":"2019-05-28 15:03:03","contractbegindate":"2019-05-28 15:03:03","contractenddate":"2019-05-28 15:03:03"}
func (c *MarketcontractController) AddMarketcontract() {
	var (
		param models.Marketcontract
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error(err)
		util.RetContent.Code = util.FAILED
		util.RetContent.Message = "参数错误"
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	} else {
		beego.Info(param)
	}
	code, msg := models.AddMarketcontract(param)
	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MarketcontractController) DeleteMarketcontract() {
	var (
		param = make(map[string]int64)
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

	id := param["id"]

	beego.Info(id)
	code, msg := models.DeleteMarketcontract(id)
	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
