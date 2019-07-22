package controllers

import (

	//"erpweb/logs"
	"erpweb/models"
	"erpweb/util"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) GetMenus() {
	rets := models.GetMenus()
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}
