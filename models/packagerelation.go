package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//暂时不用实现
//维修套餐关系表
type Packagerelation struct {
	Id        int64 `json:"id" orm:"column(id)"`
	Packageid int64 `json:"packageid" orm:"column(packageid)"` //套餐id
	Matterid  int64 `json:"matterid" orm:"column(matterid)"`   //物料id(相同套餐物料不重复)
	Num       int64 `json:"num" orm:"column(num)"`             //物料数量
}

func GetPackagerelationBypage(pageNum, pageSize int64) []Packagerelation {
	var (
		pas []Packagerelation
	)
	err := OSQL.Raw("select * from "+util.PACKAGERALATION_TABLE_NAME+" order by id asc limit ?,?",
		pageNum, pageSize).QueryRow(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas
}

func GetPackagerelationById(id int64) (pa Packagerelation, err error) {
	pa.Id = id
	err = OSQL.Read(&pa, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return pa, err
	}
	return pa, nil
}

func EditPackagerelationById(pa Packagerelation) (errorCode int64) {
	var (
		temp Packagerelation
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.LEAVE_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.LEAVE_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddPackagerelation(pa Packagerelation) (errorCode int64) {
	var (
		temp Packagerelation
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", pa.Id)
		errorCode = util.LEAVE_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.LEAVE_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeletePackagerelation(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Packagerelation
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.LEAVE_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
