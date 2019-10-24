package controllers

import (
	"io/ioutil"
	"strings"

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

	actions := c.Ctx.Request.URL.String()

	beego.Info("action=", actions)

	webToken := c.Ctx.Input.Header("x-Token")

	beego.Info("webToken=", webToken)

	if !strings.HasPrefix(actions, "/api/login/") {
		beego.Info("校验token")
		err, code := c.checkToken(webToken)
		if code != util.SUCESSFUL {
			util.RetContent.Code = code
			if err != nil {
				util.RetContent.Message = err.Error()
			}
			c.Data["json"] = util.RetContent
			c.ServeJSON()
			c.StopRun()
		}
	} else {
		beego.Info("login的登录相关接口，不需要校验token")
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
	body, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	beego.Info("method=", method, " url :", header.Path, " 参数 :", "body :", string(body))
	for k, v := range forms {
		beego.Info(k, v)
	}
}

func (c *BaseController) checkToken(webToken string) (error, int64) {
	err, code := models.SSOLogin(webToken)
	return err, code
}
