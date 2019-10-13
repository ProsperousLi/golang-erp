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

type LoginResult struct {
	Detailcode int64  `json:"detailcode"`
	Msg        string `json:"msg"`
	Vckey      string `json:"vckey"`
	Verifycode string `json:"verifycode"`
}

func (c *LoginController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	//vckey  verifycode
	vckey := c.GetString("password")
	verifycode := c.GetString("password")
	errCode, token, vckey, verifycode := models.Login(username, password, vckey, verifycode)

	util.RetContent.Code = util.SUCESSFUL
	var loginRes LoginResult
	loginRes.Detailcode = errCode
	if errCode == 1 {
		loginRes.Msg = "账号密码不正确"
	} else if errCode == 2 {
		loginRes.Msg = "验证码错误"
	} else if errCode == 3 {
		loginRes.Msg = "您的账号已锁定，请联系管理员解除锁定"
	}

	loginRes.Vckey = vckey
	loginRes.Verifycode = verifycode

	util.RetContent.Message = token
	util.RetContent.Data = loginRes
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

func (c *LoginController) UserInfo() {
	webToken := c.Ctx.ResponseWriter.Header().Get("x-Token")
	code, userInfo := models.GetUserInfo(webToken)
	util.RetContent.Code = code
	util.RetContent.Data = userInfo
	c.Data["json"] = util.RetContent
	c.ServeJSON()
	return
}

// 返回：{code: 20000, vckey: ‘xxxx’, verifycode: ‘base64’}或者
// {code: 20001, message:”旧验证码信息不正确”}或{code:20001, message: “刷新过于频繁”}
//说明：即使之前的验证码已经过期，仍然返回新的，过期的验证码起码保存1天，
//防止恶意刷验证码,1秒内禁止再次刷新
func (c *LoginController) RefreshVerifyCode() {
	util.RetContent.Code = 20000
	var loginRes LoginResult
	vckey := c.GetString("vckey")
	code, vckey, verifycode := models.RefreshVerifyCode(vckey)
	if code == 1 {
		code = 20001
		loginRes.Msg = "旧验证码信息不正确"
	} else if code == 2 {
		code = 20001
		loginRes.Msg = "刷新过于频繁"
	}

	loginRes.Vckey = vckey
	loginRes.Verifycode = verifycode

	loginRes.Detailcode = code
	util.RetContent.Data = loginRes
	c.Data["json"] = util.RetContent
	c.ServeJSON()
	return
}
