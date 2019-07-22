package controllers

import (
	"erpweb/logs"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
}

func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	//c.adapterUserInfo()
	c.getInfo()
	//init res struct
	var newRes util.Result
	util.RetContent = newRes
}

//从session里取用户信息
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession("backenduser")
	if a != nil {
		//c.curUser = a.(models.BackendUser)
		c.Data["backenduser"] = a
	}
}

func (c *BaseController) getInfo() {
	method := c.Ctx.Request.Method
	header := c.Ctx.Request.URL
	forms := c.Ctx.Request.Form
	logs.FileLogs.Info("method=%v", method)
	logs.FileLogs.Info("url :%v", header.Path)
	logs.FileLogs.Info("参数 :")
	for k, v := range forms {
		//lg.Info("参数 :", k, v)
		logs.FileLogs.Info("%v=%v", k, v)
	}

}
