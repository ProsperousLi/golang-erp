package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

//暂时不用实现
//维修套餐表
type Matterpackage struct {
	Id   int64  `json:"id" orm:"column(id)"`
	Name string `json:"name" orm:"column(name)"` //套餐名称
}

func GetMatterpackageBypage(pageNum, pageSize int64) []Matterpackage {
	var (
		mas []Matterpackage
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.MATTERPACKAGE_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&mas)
	if err != nil {
		beego.Error(err)
	}
	return mas
}

func GetMatterpackageById(id int64) (ma Matterpackage, err error) {
	ma.Id = id
	err = OSQL.Read(&ma, "id")
	if err != nil {
		beego.Error(err)
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
		beego.Error(err)
		errorCode = util.MATTERPACKAGE_EDIT_FAILED
		return errorCode
	}

	args := edit_Matterpackage(ma)

	num, err2 := OSQL.Update(&ma, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.MATTERPACKAGE_EDIT_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func edit_Matterpackage(param Matterpackage) (args []string) {
	if param.Name != "" {
		args = append(args, "name")
	}
	return args
}

func AddMatterpackage(ma Matterpackage) (errorCode int64) {
	var (
		temp Matterpackage
	)
	errorCode = util.SUCESSFUL
	temp.Id = ma.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		beego.Error("matterpackage have this id=%v", ma.Id)
		errorCode = util.MATTERPACKAGE_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&ma)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.MATTERPACKAGE_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)
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
		beego.Error(err)
		errorCode = util.MATTERPACKAGE_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
