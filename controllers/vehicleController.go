package controllers

import (
	"encoding/json"

	"erpweb/models"
	"erpweb/util"

	"github.com/astaxie/beego"
)

type VehicleController struct {
	BaseController
}

func (c *VehicleController) QueryVehicle() {
	custcode := c.GetString("custcode")
	rets := models.GetVehicleByCustcode(custcode)
	util.RetContent.Code = util.SUCESSFUL
	util.RetContent.Data = rets
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *VehicleController) NewVehicle() {
	var (
		param models.Vehicle
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

	//param.Password =

	code := models.AddVehicle(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()
}

func (c *VehicleController) UpdateVehicle() {
	var (
		param models.Vehicle
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

	code := models.EditVehicleById(param)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()

}

func (c *VehicleController) DelVehicle() {
	var (
		param = make(map[string]int64)
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

	id := param["id"]

	code := models.DeleteVehicle(id)
	util.RetContent.Code = code
	c.Data["json"] = util.RetContent
	c.ServeJSON()

}
