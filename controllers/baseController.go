package controllers

import (
	"erpweb/logs"
	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
}

//TODO 1.检查用户的token是否存在；2.所有的增删改都需要去检查该用户是否有写权限，查询不需要写权限
//权限表最好加载到内存
func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	//c.adapterUserInfo()
	c.getInfo()

	//init res struct
	var newRes util.Result
	util.RetContent = newRes

	var newPageRes util.PageResult
	util.PageResults = newPageRes

	if c.Ctx.Request.Method != "/api/basedata/Login" {
		message, code := c.checkToken()
		if code != util.SUCESSFUL {
			util.RetContent.Code = code
			util.RetContent.Message = message
			c.Data["json"] = util.RetContent
			c.ServeJSON()
			c.StopRun()
		}
	}

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

func (c *BaseController) checkToken() (string, int64) {
	webToken := c.Ctx.ResponseWriter.Header().Get("x-Token")
	err, code := models.SSOLogin(webToken)
	return err.Error(), code
}
