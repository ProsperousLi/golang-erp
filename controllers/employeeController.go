package controllers

import (
	"encoding/json"

	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type EmployeeController struct {
	BaseController
}

func (c *EmployeeController) GetEmployees() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
	}
	pageNum := param["pageNum"]
	pageSize := param["pageSize"]
	if pageNum > 0 {
		pageNum = pageNum - 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	emps := models.GetEmployees(pageNum, pageSize)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = emps
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *EmployeeController) GetEmployeeById() {
	var (
		param = make(map[string]int64)
	)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.FileLogs.Error("param is err", string(c.Ctx.Input.RequestBody))
	}

	userId := param["id"]

	//userId = 1
	logs.FileLogs.Info("%v ---", userId)
	emp, code := models.GetEmployeeById(userId)
	util.RetContent.Code = code
	util.RetContent.Data = emp
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *EmployeeController) EditEmployeeById() {
	var (
		emp models.Employee
	)

	logs.FileLogs.Info("param = ", string(c.Ctx.Input.RequestBody))

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &emp)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	code := models.EditEmployeeById(emp)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//{"cardid":"123","name":"小李","sex":0,"compID":10,"deptID":1,"dutyID":1,"health":"very good","height":"170","nativeplace":"安徽省","nation":"汉","maritalstatus":"未婚","education":"本科","university":"大连东软","major":"计算  机","qualification":"666","trialsalary":2000,"salary":6000,"idnumber":"341182","address1":"1111","postcode1":"www111","address2":"2222","postcode2":"www222","contactnumber":"17615002988","phonenumber":"110","email":"www666","emergencycontact":"1323654222","c  ontactnumber1":"1323654222","address3":"33333","trialexpired":"2019-05-28 15:03:03","entrydate":"2019-05-28   15:03:03","birthday":"2019-05-28 15:03:03","contractbegindate":"2019-05-28 15:03:03","contractenddate":"2019-05-28 15:03:03"}
func (c *EmployeeController) AddEmployee() {
	var (
		emp models.Employee
	)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &emp)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	} else {
		logs.FileLogs.Info("%v", emp)
	}
	code := models.AddEmployee(emp)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *EmployeeController) DeleteEmployee() {
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
	code := models.DeleteEmployee(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

//upload File
// func (c *FormController) Post() {
// 	f, h, err := c.GetFile("uploadname")
// 	defer f.Close()
// 	if err != nil {
// 		//fmt.Println("getfile err ", err)
// 	} else {
// 		c.SaveToFile("uploadname", "/www/"+h.Filename)
// 	}
// }

type MainController struct {
	BaseController
}

//curl 127.0.0.1:8080/employee/addPerson -X POST -H "Content-Type:application/json" -d '{"cardid":"123","name":"小李","sex":0,"compID":0,"deptID":1,"dutyID":1,"health":"very good","height":"170","nativeplace":"安徽省","nation":"汉","maritalstatus":"未婚","education":"本科","university":"大连东软","major":"计算机","qualification":"666","trialsalary":2000,"salary":6000,"idnumber":"341182","address1":"1111","postcode1":"www111","address2":"2222","postcode2":"www222","contactnumber":"17615002988","phonenumber":"110","email":"www666","emergencycontact":"1323654222","contactnumber1":"1323654222","address3":"33333","trialexpired":"2019-05-28 15:03:03","entrydate":"2019-05-28 15:03:03","birthday":"2019-05-28 15:03:03","contractbegindate":"2019-05-28 15:03:03","contractenddate":"2019-05-28 15:03:03"}'
func (c *MainController) Index() {
	logs.FileLogs.Info("hello world")
	c.TplName = "ajax/index.html"
	//c.Ctx.Redirect(200, "/ajax/index.html")
}
