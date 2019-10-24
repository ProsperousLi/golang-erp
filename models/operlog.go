package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

//操作日志表
type Operlog struct {
	Id       int64  `json:"id" orm:"column(id)"`
	Operator int64  `json:"operator" orm:"column(operator)"` //操作人ID(employee表的id字段)
	Detail   string `json:"detail" orm:"column(detail)"`     //操作详情
	Doat     string `json:"doat" orm:"column(doat)"`         //操作时间
}

func GetOperlogBypage(pageNum, pageSize int64) []Operlog {
	var (
		operlogs []Operlog
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.OPERLOG_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&operlogs)
	if err != nil {
		beego.Error(err)
	}

	return operlogs
}

func GetOperlogByOperator(operator int64) (operlog Operlog, err error) {
	operlog.Operator = operator
	err = OSQL.Read(&operator, "operator")
	if err != nil {
		beego.Error(err)
		return operlog, err
	}
	return operlog, nil
}

func EditOperlogById(operlog Operlog) (errorCode int64) {
	var (
		temp Operlog
	)
	errorCode = util.SUCESSFUL
	temp.Operator = operlog.Operator
	err := OSQL.Read(&temp, "operator")
	if err != nil {
		beego.Error(err)
		errorCode = util.OPERLOG_EDIT_FAILED
		return errorCode
	}

	args := edit_operlog(operlog)

	num, err2 := OSQL.Update(&operlog, args...)
	if err2 != nil {
		beego.Error(err)
		errorCode = util.OPERLOG_EDIT_FAILED
		return errorCode
	}
	beego.Info("num=", num)

	return errorCode
}

func edit_operlog(operlog Operlog) []string {
	var (
		args []string
	)
	if operlog.Detail != "" {
		args = append(args, "detail")
	}
	if operlog.Doat != "" {
		args = append(args, "doat")
	}
	if operlog.Operator != 0 {
		args = append(args, "operator")
	}
	return args
}

func AddOperlog(operlog Operlog) (errorCode int64) {
	var (
		temp Operlog
	)
	errorCode = util.SUCESSFUL
	temp.Operator = operlog.Operator
	err := OSQL.Read(&temp, "operator")
	if err == nil {
		beego.Error("operlog have this operator=", operlog.Operator)
		errorCode = util.OPERLOG_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&operlog)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.OPERLOG_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)

	return errorCode
}

func DeleteOperlog(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Operlog
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.OPERLOG_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
