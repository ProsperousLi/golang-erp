package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//车辆信息表
type Vehicle struct {
	Id int64 `json:"id" orm:"column(id)"`
}

func GetVehicleBypage(pageNum, pageSize int64) []Vehicle {
	var (
		pas []Vehicle
	)
	err := OSQL.Raw("select * from "+util.VEHICLE_TABLE_NAME+" order by id asc limit ?,?",
		pageNum, pageSize).QueryRow(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas
}

func GetVehicleById(id int64) (pa Vehicle, err error) {
	pa.Id = id
	err = OSQL.Read(&pa, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return pa, err
	}
	return pa, nil
}

func EditVehicleById(pa Vehicle) (errorCode int64) {
	var (
		temp Vehicle
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.VEHICLE_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.VEHICLE_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddVehicle(pa Vehicle) (errorCode int64) {
	var (
		temp Vehicle
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", pa.Id)
		errorCode = util.VEHICLE_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.VEHICLE_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteVehicle(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Vehicle
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.VEHICLE_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
