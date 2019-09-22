package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type AccountController struct {
	BaseController
}

func (c *AccountController) GetAccountList() {
	rets := models.GetAccountsNotPwd()
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *AccountController) GetAccounts() {
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
	rets := models.GetAccountBypage(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *AccountController) GetAccountById() {
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

	userId := param["cardid"]

	logs.FileLogs.Info("%v ---", userId)
	ret, _ := models.GetAccountByUserID(userId)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = ret
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *AccountController) EditAccountById() {
	var (
		param models.Account
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return

	}
	code, msg := models.EditAccountById(param)
	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *AccountController) EditAccountStatusById() {
	//{cardid: “xxx”, status: 1}
	type params struct {
		Cardid string
		Status int8
	}

	var param params

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("%s", err)
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	code := models.EditAccountStatusById(param.Cardid, param.Status)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{"cardid":"123","name":"小李","sex":0,"compID":10,"deptID":1,"dutyID":1,"health":"very good","height":"170","nativeplace":"安徽省","nation":"汉","maritalstatus":"未婚","education":"本科","university":"大连东软","major":"计算  机","qualification":"666","trialsalary":2000,"salary":6000,"idnumber":"341182","address1":"1111","postcode1":"www111","address2":"2222","postcode2":"www222","contactnumber":"17615002988","phonenumber":"110","email":"www666","emergencycontact":"1323654222","c  ontactnumber1":"1323654222","address3":"33333","trialexpired":"2019-05-28 15:03:03","entrydate":"2019-05-28   15:03:03","birthday":"2019-05-28 15:03:03","contractbegindate":"2019-05-28 15:03:03","contractenddate":"2019-05-28 15:03:03"}
func (c *AccountController) AddAccount() {
	var (
		param models.Account
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

	//param.Password =

	code := models.AddAccountment(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *AccountController) DeleteAccount() {
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

	userId := param["cardid"]

	logs.FileLogs.Info("%v ---", userId)
	code := models.DeleteAccount(userId)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{cardid: “xxx”, oldpwd: “xxx”, newpwd: “xxx”}
func (c *AccountController) ModifyPwd() {

	var param models.ModifyPwdStruct

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
		util.RetContent.Code = util.PARAM_FAILED
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	logs.FileLogs.Info("%v ---", param)
	code, msg := models.ModifyPwd(param)
	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//18.重置密码
//{cardid: “xxx”}
func (c *AccountController) ResetAccount(cardid string) {
	//TODO 需要检验操作人的token

	//对人员管理有写权限
	//EditAccountStatusById(cardid,3)
	var account models.Account
	account.Cardid = cardid
	account.Status = 3
	account.Password = util.GETMd5(util.DEFUAL_PWD_PRE + util.DEFUAL_PWD)
	code, msg := models.EditAccountById(account)
	util.RetContent.Code = code
	util.RetContent.Message = msg
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
