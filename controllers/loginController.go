package controllers

import (
	"erpweb/models"
	"erpweb/util"
)

type LoginController struct {
	BaseController
}

type loginParam struct {
	Username   string
	Password   string
	Vckey      string
	Verifycode string
}

//code: 20000, data: {detailcode: 0, token: ‘’}

//{code: 20000, data: {detailcode: 1, msg: ‘账号密码不正确’, vckey: ‘’,
//verifycode: ‘dsfwekfsldfklsdfkkkk’}}

//{code: 20000, data: {detailcode: 2, msg: ‘验证码错误’,
//vckey: ‘新key’, verifycode: ‘新验证码’}}

//{code: 20000, data: {detailcode: 3, msg:
//‘您的账号已锁定，请联系管理员解除锁定’}}

func (c *LoginController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	//vckey  verifycode
	vckey := c.GetString("password")
	verifycode := c.GetString("password")
	errCode, token := models.Login(username, password, vckey, verifycode)

	util.RetContent.Code = errCode //util.SUCESSFUL
	util.RetContent.Message = token
	util.RetContent.Data = models.AccsMap[token]
	c.Data["json"] = util.RetContent
	c.ServeJSON()
	return
}

func (c *LoginController) Loginout() {
	webToken := c.Ctx.ResponseWriter.Header().Get("x-Token")
	code := models.Loginout(webToken)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
	return
}
