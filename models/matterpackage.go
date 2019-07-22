package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//暂时不用实现
//维修套餐表
type Matterpackage struct {
	Id   int64  `json:"id" orm:"column(id)"`
	name string `json:"name" orm:"column(name)"` //套餐名称
}

func GetMatterpackageBypage(pageNum, pageSize int64) []Matterpackage {
	var (
		mas []Matterpackage
	)
	err := OSQL.Raw("select * from "+util.MATTERPACKAGE_TABLE_NAME+" order by id asc limit ?,?",
		pageNum, pageSize).QueryRow(&mas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return mas
}

func GetMatterpackageById(id int64) (ma Matterpackage, err error) {
	ma.Id = id
	err = OSQL.Read(&ma, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ma, err
	}
	return ma, nil
}

func EditMatterpackageById(ma Matterpackage) (errorCode int64) {
	var (
		temp Matterpackage
	)
	errorCode = util.SUCESSFUL
	temp.Id = ma.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.MATTERPACKAGE_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&ma)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.MATTERPACKAGE_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddMatterpackage(ma Matterpackage) (errorCode int64) {
	var (
		temp Matterpackage
	)
	errorCode = util.SUCESSFUL
	temp.Id = ma.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", ma.Id)
		errorCode = util.PACKAGERALATION_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&ma)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.PACKAGERALATION_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteMatterpackage(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Matterpackage
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.PACKAGERALATION_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
