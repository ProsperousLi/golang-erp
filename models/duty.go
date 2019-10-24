package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

//暂时不用实现
type Duty struct {
	Id    int64  `json:"id" orm:"column(id)"`
	Name  string `json:"name" orm:"column(name)"`   //岗位名称
	Level int    `json:"level" orm:"column(level)"` //岗位级别
}

func GetDutyBypage(pageNum, pageSize int64) []Duty {
	var (
		dutys []Duty
	)
	begin := pageSize * pageNum
	err := OSQL.Raw("select * from "+util.DUTY_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRow(&dutys)
	if err != nil {
		beego.Error(err)
	}

	return dutys
}

func GetDutyById(id int64) (duty Duty, err error) {
	duty.Id = id
	err = OSQL.Read(&duty, "id")
	if err != nil {
		beego.Error(err)
		return duty, err
	}
	return duty, nil
}

func EditDutyById(duty Duty) (errorCode int64) {
	var (
		temp Duty
	)
	errorCode = util.SUCESSFUL
	temp.Id = duty.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.DUTY_EDIT_FAILED
		return errorCode
	}
	args := editArgs_duty(duty)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&temp, args...)
		if err2 != nil {
			beego.Error(err)
			errorCode = util.DUTY_EDIT_FAILED
			return errorCode
		}
		beego.Info("num=", num)
	} else {
		beego.Info("no data update")
	}
	return errorCode
}

func editArgs_duty(duty Duty) []string {
	var (
		args []string
	)
	if duty.Level != 0 {
		args = append(args, "level")
	}
	if duty.Name != "" {
		args = append(args, "name")
	}

	return args
}

func AddDuty(duty Duty) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Duty
	)
	errorCode = util.SUCESSFUL
	temp.Id = duty.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		beego.Error("duty have this id=", temp.Id)
		errorCode = util.DUTY_ADD_FAILED
		return errorCode
	}
	num, err2 := OSQL.Insert(&temp)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.DUTY_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func DeleteDuty(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Duty
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.DUTY_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
