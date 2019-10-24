package controllers

import (
	"encoding/json"

	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type MatterController struct {
	BaseController
}

func (c *MatterController) QueryMatter() {
	mattercode := c.GetString("mattercode")
	name := c.GetString("name")
	rets := models.QueryMatter(mattercode, name)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MatterController) GetMatters() {
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
	pageNum := param["pageNum"]
	pageSize := param["pageSize"]
	if pageNum > 0 {
		pageNum = pageNum - 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	rets := models.GetMatterBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MatterController) GetMatterById() {
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
	ret, _ := models.GetMatterById(id)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MatterController) EditMatterById() {
	var (
		param models.Matter
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		beego.Error(err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code := models.EditMatterById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{"cardid":"123","name":"小李","sex":0,"compID":10,"deptID":1,"dutyID":1,"health":"very good","height":"170","nativeplace":"安徽省","nation":"汉","maritalstatus":"未婚","education":"本科","university":"大连东软","major":"计算  机","qualification":"666","trialsalary":2000,"salary":6000,"idnumber":"341182","address1":"1111","postcode1":"www111","address2":"2222","postcode2":"www222","contactnumber":"17615002988","phonenumber":"110","email":"www666","emergencycontact":"1323654222","c  ontactnumber1":"1323654222","address3":"33333","trialexpired":"2019-05-28 15:03:03","entrydate":"2019-05-28   15:03:03","birthday":"2019-05-28 15:03:03","contractbegindate":"2019-05-28 15:03:03","contractenddate":"2019-05-28 15:03:03"}
func (c *MatterController) AddMatter() {
	var (
		param models.Matter
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
	code := models.AddMatter(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *MatterController) DeleteMatter() {
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

	beego.Info(id)
	code := models.DeleteMatter(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
