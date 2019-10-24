package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

type Department struct {
	Id     int64  `json:"id" orm:"column(id)"`
	CompID int64  `json:"compID" orm:"column(compID)"` //公司ID
	Name   string `json:"name" orm:"column(name)"`     //部门名称
}

func QueryDept() []Department {
	var (
		departs []Department
	)
	_, err := OSQL.Raw("select * from " + util.DEPARTMENT_TABLE_NAME + " order by id asc ").QueryRows(&departs)
	if err != nil {
		beego.Error(err)
	}

	return departs
}

func GetDepartmentBypage(pageNum, pageSize int64) []Department {
	var (
		departs []Department
	)
	begin := pageSize * pageNum
	beego.Info("begin=", begin, ", end =", pageSize)
	_, err := OSQL.Raw("select * from "+util.DEPARTMENT_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&departs)
	if err != nil {
		beego.Error(err)
	}

	return departs
}

func GetDepartmentById(id int64) (depart Department, err error) {

	depart.Id = id

	err = OSQL.Read(&depart, "id")
	if err != nil {
		beego.Error(err)
		return depart, err
	}

	return depart, nil
}

func EditDepartmentById(depart Department) (errorCode int64) {
	var (
		temp Department
	)
	temp.Id = depart.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.FAILED
		return errorCode
	}

	args := editArgs_depart(depart)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&depart, args...)
		if err2 != nil {
			beego.Error(err2)
			errorCode = util.FAILED
			return errorCode
		}
		beego.Info("num=", num)
	} else {
		beego.Info("no data update")
	}

	return errorCode
}

func editArgs_depart(depart Department) []string {
	// CompID int64  `json:"compID" orm:"column(compID)"` //公司ID
	// Name   string `json:"name" orm:"column(name)"`     //部门名称
	var (
		args []string
	)
	if depart.CompID != 0 {
		args = append(args, "compID")
	}
	if depart.Name != "" {
		args = append(args, "name")
	}
	beego.Info("args=", args)
	return args
}

func AddDepartment(depart Department) (errorCode int64) {
	var (
		temp Department
	)
	temp.CompID = depart.CompID
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "compID")
	if err == nil {
		beego.Error("table have this compID=", depart.CompID)
		errorCode = util.FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&depart)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.FAILED
	}

	beego.Info("num=", id)
	return errorCode
}

func DeleteDepartment(id int64) (errorCode int64) {
	var (
		depart Department
	)
	errorCode = util.SUCESSFUL
	depart.Id = id
	num, err := OSQL.Delete(&depart, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.DEPART_DELETE_FAILED
	}
	beego.Info("num=", num)
	return errorCode
}
