package models

import (
	"erpweb/logs"
	"erpweb/util"
)

type Department struct {
	Id     int64  `json:"id" orm:"column(id)"`
	CompID int64  `json:"compID" orm:"column(compID)"` //公司ID
	Name   string `json:"name" orm:"column(name)"`     //部门名称
}

func GetDepartmentBypage(pageNum, pageSize int64) []Department {
	var (
		departs []Department
	)
	begin := pageSize * pageNum
	logs.FileLogs.Info("begin=%v", begin, ", end =%v", pageSize)
	_, err := OSQL.Raw("select * from "+util.DEPARTMENT_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&departs)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}

	return departs
}

func GetDepartmentById(id int64) (depart Department, err error) {

	depart.Id = id

	err = OSQL.Read(&depart, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
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
		logs.FileLogs.Error("%s", err)
		errorCode = util.DEPART_EDIT_FAILED
		return errorCode
	}

	args := editArgs_depart(depart)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&depart, args...)
		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.DEPART_EDIT_FAILED
			return errorCode
		}
		logs.FileLogs.Info("num=%v", num)
	} else {
		logs.FileLogs.Info("no data update")
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
	logs.FileLogs.Info("args=%v", args)
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
		logs.FileLogs.Error("table have this compID=%v", depart.CompID)
		errorCode = util.DEPART_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&depart)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.DEPART_ADD_FAILED
	}

	logs.FileLogs.Info("num=%v", id)
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
		logs.FileLogs.Error("%v", err)
		errorCode = util.DEPART_DELETE_FAILED
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
