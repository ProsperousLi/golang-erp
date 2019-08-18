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
	cardid := c.GetString("cardid")
	password := c.GetString("password")
	err, token := models.Login(cardid, password)
	if err != nil {
		util.RetContent.Code = 20001
		util.RetContent.Message = err.Error()
		c.Data["json"] = util.RetContent
		c.ServeJSON()
		return
	}

	util.RetContent.Code = util.SUCESSFUL
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
