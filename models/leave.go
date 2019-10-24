package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

//客户信息表
type Leaves struct {
	Id         int64  `json:"id" orm:"column(id)"`
	Employeeid int64  `json:"employeeid" orm:"column(employeeid)"`
	Leaveat    string `json:"leaveat" orm:"column(leaveat)"` //离职日期
	Reason     string `json:"reason" orm:"column(reason)"`   //离职原因
}

func QueryAllLeave() []Leaves {
	var (
		les []Leaves
	)
	_, err := OSQL.Raw("select * from " + util.LEAVE_TABLE_NAME +
		" order by id asc").QueryRows(&les)
	if err != nil {
		beego.Error(err)
	}
	return les
}

func GetLeaveBypage(pageNum, pageSize int64) []Leaves {
	var (
		les []Leaves
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.LEAVE_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&les)
	if err != nil {
		beego.Error(err)
	}
	return les
}

func GetLeaveById(id int64) (le Leaves, err error) {
	le.Id = id
	err = OSQL.Read(&le, "id")
	if err != nil {
		beego.Error(err)
		return le, err
	}
	return le, nil
}

func EditLeaveById(le Leaves) (errorCode int64) {
	var (
		temp Leaves
	)
	errorCode = util.SUCESSFUL
	temp.Id = le.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.LEAVE_EDIT_FAILED
		return errorCode
	}
	args := editArgs_leave(le)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&le, args...)
		if err2 != nil {
			beego.Error(err2)
			errorCode = util.LEAVE_EDIT_FAILED
			return errorCode
		}
		beego.Info("num=", num)
	} else {
		beego.Info("no data update")
	}
	return errorCode
}

func editArgs_leave(le Leaves) []string {
	var (
		args []string
	)
	if le.Employeeid != 0 {
		args = append(args, "employeeid")
	}
	if le.Leaveat != "" {
		args = append(args, "leaveat")
	}
	if le.Reason != "" {
		args = append(args, "reason")
	}
	beego.Info("args=", args)
	return args
}

func AddLeave(le Leaves) (errorCode int64) {
	var (
		temp Leaves
	)
	errorCode = util.SUCESSFUL
	temp.Id = le.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		beego.Error("leave have this id=", le.Id)
		errorCode = util.LEAVE_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&le)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.LEAVE_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func DeleteLeave(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Leaves
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.LEAVE_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
